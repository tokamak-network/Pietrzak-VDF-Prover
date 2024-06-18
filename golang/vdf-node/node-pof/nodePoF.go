package nodePoF

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/commit-reveal-recover/crr"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/crrrngpof"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/node"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
)

const (
	CommitDuration  = 120    // Commit duration in seconds
	DisputeDuration = 200    // Dispute duration in seconds
	ContextTimeout  = 600000 // Context timeout duration in seconds
)

type PoFListenerInterface interface {
	SubscribeRandomWordsRequested() error
	CheckRoundCondition() error
	GetNextRound() (*big.Int, error)
}

type PoFListener struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	ContractABI     abi.ABI
	EventData       []EventInfo
	PrivateKey      *ecdsa.PrivateKey
	Mutex           sync.Mutex
	WaitGroup       sync.WaitGroup
	LeaderRounds    map[*big.Int]common.Address
	MyAddress       common.Address
	RoundData       map[string]*big.Int
	StartTimes      map[string]time.Time
}

type EventInfo struct {
	Round  *big.Int
	Sender common.Address
}

type BigNumber struct {
	Val    []byte   `json:"val"`
	Bitlen *big.Int `json:"bitlen"`
}

type ValueAtRound struct {
	StartTime     *big.Int       `json:"startTime"`
	RequestedTime *big.Int       `json:"requestedTime"`
	CommitCounts  *big.Int       `json:"commitCounts"`
	Consumer      common.Address `json:"consumer"`
	CommitsString []byte         `json:"commitsString"`
	Omega         BigNumber      `json:"omega"`
	Stage         string         `json:"stage"`
	IsCompleted   bool           `json:"isCompleted"`
	IsVerified    bool           `json:"isVerified"`
	T             *big.Int       `json:"T"`
	NBitLen       *big.Int       `json:"nBitLen"`
	GBitLen       *big.Int       `json:"gBitLen"`
	HBitLen       *big.Int       `json:"hBitLen"`
	NVal          []byte         `json:"nVal"`
	GVal          []byte         `json:"gVal"`
	HVal          []byte         `json:"hVal"`
}

type SetupValues struct {
	T       *big.Int
	NBitLen *big.Int
	GBitLen *big.Int
	HBitLen *big.Int
	NVal    []byte
	GVal    []byte
	HVal    []byte
	Stage   string
}

type CommitValue struct {
	Commit          BigNumber      `json:"commit"`
	OperatorAddress common.Address `json:"operatorAddress"`
}

func loadContractABI(filename string) (abi.ABI, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to read ABI file: %v", err)
	}

	var abiObject struct {
		Abi []interface{} `json:"abi"`
	}
	if err := json.Unmarshal(fileContent, &abiObject); err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse ABI JSON: %v", err)
	}

	abiBytes, err := json.Marshal(abiObject.Abi)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to re-marshal ABI: %v", err)
	}

	contractAbi, err := abi.JSON(bytes.NewReader(abiBytes))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse contract ABI: %v", err)
	}
	return contractAbi, nil
}

func NewPoFListener(config node.Config) (*PoFListener, error) {
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey[2:]) // Strip the "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	abiFilename := "../CRRNGCoordinatorPoF.json"
	contractABI, err := loadContractABI(abiFilename)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	myAddress := common.HexToAddress(config.WalletAddress)

	return &PoFListener{
		Client:          client,
		ContractAddress: contractAddress,
		ContractABI:     contractABI,
		PrivateKey:      privateKey,
		LeaderRounds:    make(map[*big.Int]common.Address),
		MyAddress:       myAddress,
		RoundData:       make(map[string]*big.Int),
		StartTimes:      make(map[string]time.Time),
	}, nil
}

func (l *PoFListener) CheckRoundCondition() error {
	color.New(color.FgHiRed, color.Bold).Printf("🚨 Checking previous rounds...\n")

	nextRound, err := l.GetNextRound()
	if err != nil {
		log.Fatalf("Error retrieving next round: %v", err)
		return nil
	}

	currentRound := new(big.Int).Sub(nextRound, big.NewInt(1))
	if currentRound.Cmp(big.NewInt(0)) < 0 {
		color.New(color.FgHiGreen, color.Bold).Printf("✅ Check Complete!! \n")
		color.New(color.FgHiYellow, color.Bold).Printf("🚫 No rounds have started yet.\n")
		return nil
	}
	//log.Printf("Current round number is: %s", currentRound.String())

	lastRecoveredRound, err := l.GetLastRecoveredRound()
	if err != nil {
		log.Fatalf("Error retrieving last recovered round: %v", err)
		return nil
	}
	//log.Printf("Last recovered round number is: %s", lastRecoveredRound.String())

	lastFulfilledRound, err := l.GetLastFulfilledRound()
	if err != nil {
		log.Fatalf("Error retrieving last fulfilled round: %v", err)
		return nil
	}
	//log.Printf("Last fulfilled round number is: %s", lastFulfilledRound.String())

	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout*time.Minute)
	defer cancel()

	for checkRound := new(big.Int).Set(lastRecoveredRound); checkRound.Cmp(currentRound) <= 0; checkRound.Add(checkRound, big.NewInt(1)) {
		color.New(color.FgHiYellow, color.Bold).Printf("Current checking round: %s\n", checkRound)

		if lastRecoveredRound.Cmp(big.NewInt(0)) == 0 {
			operators, err := l.GetCommittedOperatorsAtRound(checkRound)
			if err != nil {
				log.Printf("Error retrieving operators at round %s: %v", checkRound.String(), err)
				return err
			}
			if operators == nil {
				log.Printf("No operators committed at round %s", checkRound.String())
				return nil
			}

			valueAtRound, err := l.GetValuesAtRound(ctx, lastRecoveredRound)
			if err != nil {
				log.Printf("Error retrieving values at round 0: %v", err)
			}
			l.initiateCommitProcess(lastRecoveredRound)

			startTimeInSeconds := valueAtRound.StartTime.Int64()
			startTime := time.Unix(startTimeInSeconds, 0)
			commitDeadline := startTime.Add(time.Second * time.Duration(CommitDuration))

			if time.Now().After(commitDeadline) && !valueAtRound.IsCompleted {
				color.New(color.FgHiRed, color.Bold).Printf("🚨 Round %s is not fully recovered, initiating recovery process.\n", checkRound)
				l.ReRequestRandomWordAtRound(ctx, lastRecoveredRound)
				l.initiateCommitProcess(lastRecoveredRound)
			} else if !time.Now().After(commitDeadline) && !valueAtRound.IsCompleted {
				l.initiateCommitProcess(lastRecoveredRound)
			}

			return nil
		}

		if lastRecoveredRound.Cmp(checkRound) == 0 {
			fmt.Println("lastRecoveredRound: ", checkRound)
			if lastFulfilledRound.Cmp(currentRound) == 0 {
				log.Printf("Last fulfilled round %s matches current round %s, and matches last recovered round %s", lastFulfilledRound.String(), currentRound.String(), lastRecoveredRound.String())
			} else {
				// If they do not match, attempt to fulfill the randomness for the check round
				signedTx, err := l.FulfillRandomness(ctx, checkRound)
				if err != nil {
					log.Printf("Failed to fulfill randomness for round %s: %v", checkRound.String(), err)
					return err
				}
				log.Printf("FulfillRandomness successful! Tx Hash: %s", signedTx.Hash().Hex())
			}
		} else {
			valueAtRound, err := l.GetValuesAtRound(ctx, checkRound)
			if err != nil {
				log.Printf("Error retrieving values at round %s: %v", checkRound.String(), err)
				continue
			}

			startTimeInSeconds := valueAtRound.StartTime.Int64()
			startTime := time.Unix(startTimeInSeconds, 0)
			commitDeadline := startTime.Add(time.Second * time.Duration(CommitDuration))

			if time.Now().After(commitDeadline) && !valueAtRound.IsCompleted {
				l.ReRequestRandomWordAtRound(ctx, lastRecoveredRound)
				l.initiateCommitProcess(lastRecoveredRound)
			} else if !time.Now().After(commitDeadline) && !valueAtRound.IsCompleted {
				l.initiateCommitProcess(lastRecoveredRound)
			} else {
				l.Recover(ctx, checkRound, valueAtRound.Omega)
			}
		}
	}

	return nil
}

func (l *PoFListener) SubscribeRandomWordsRequested() error {
	var round *big.Int
	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.ContractAddress},
	}

	logs := make(chan types.Log)
	sub, err := l.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}

	processedRounds := make(map[string]bool)
	l.RoundData = make(map[string]*big.Int)

	//util.StartSpinner("🎧 Listening for 'RandomWordsRequested' events...", 5)
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("🎧 Listening for 'RandomWordsRequested' events...")

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			event := EventInfo{}
			l.ContractABI.UnpackIntoInterface(&event, "RandomWordsRequested", vLog.Data)

			round = event.Round
			roundKey := round.String()

			if _, exists := processedRounds[roundKey]; !exists && event.Round != nil && event.Round.Sign() >= 0 && event.Sender.Hex() != "0x0000000000000000000000000000000000000040" && event.Sender.Hex() != "0x0000000000000000000000000000000000000080" {
				fmt.Println("---------------------------------------------------------------------------")
				fmt.Printf("🔔 Round: %s, Sender: %s\n", round.String(), event.Sender.Hex())
				l.EventData = append(l.EventData, event)
				processedRounds[roundKey] = true

				l.Mutex.Lock()
				l.RoundData[roundKey] = round
				l.Mutex.Unlock()

				roundCtx, cancel := context.WithTimeout(context.Background(), ContextTimeout*time.Second)
				defer cancel()

				// Get values for the current round
				valuesAtRound, err := l.GetValuesAtRound(roundCtx, round)
				if err != nil {
					log.Printf("Error retrieving values at round: %v", err)
					return err
				}

				// Check if the current round is completed
				isCurrentRoundCompleted := valuesAtRound.IsCompleted
				if !isCurrentRoundCompleted {
					zero := big.NewInt(0)
					if round.Cmp(zero) > 0 {
						// If the current round is not completed and this is not the first round, check the previous round
						prevRound := new(big.Int).Sub(round, big.NewInt(1))

						prevRoundCtx, prevCancel := context.WithTimeout(context.Background(), ContextTimeout*time.Second)
						defer prevCancel()

						prevValuesAtRound, err := l.GetValuesAtRound(prevRoundCtx, prevRound)
						if err != nil {
							log.Printf("Error retrieving values at previous round: %v", err)
							return err
						}

						isPrevRoundCompleted := prevValuesAtRound.IsCompleted
						if isPrevRoundCompleted {
							// If the previous round is completed, initiate commit process for the current round
							l.initiateCommitProcess(round)
						} else {
							color.New(color.FgHiRed, color.Bold).Printf("🚨 Previous round %s is not completed, cannot initiate commit process for current round %s\n", prevRound.String(), round.String())
							fmt.Println("---------------------------------------------------------------------------")
						}
					} else {
						// If this is the first round, just initiate the commit process
						l.initiateCommitProcess(round)
					}
				} else {
					log.Printf("Current round %s is already completed", round.String())
				}
			}
		}
	}
}

func (l *PoFListener) initiateCommitProcess(round *big.Int) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout*time.Second)
	defer cancel()

	// Start commit process
	sender, _, err := l.Commit(ctx, round)
	if err != nil {
		log.Printf("Failed to commit: %v", err)
		return
	}

	valuesAtRound, err := l.GetValuesAtRound(ctx, round)
	if err != nil {
		log.Printf("Error retrieving values at round: %v", err)
		return
	}

	go func(round *big.Int, startTime *big.Int) {
		// Convert blockchain start time to time.Time
		startTimeUnix := time.Unix(startTime.Int64(), 0)

		// Calculate elapsed time
		elapsed := time.Since(startTimeUnix)
		remaining := CommitDuration - int(elapsed.Seconds())
		if remaining < 0 {
			remaining = 0
		}

		for i := remaining; i > 0; i-- {
			if i%5 == 0 {
				fmt.Printf("⏳ [Commit Phase] Round %s - Countdown: %d seconds remaining\n", round.String(), i)
			}
			time.Sleep(1 * time.Second)
		}
		fmt.Printf("🕒 Round %s Countdown completed. Proceeding to the next step.\n", round.String())

		time.Sleep(10 * time.Second)

		recoveryCtx, recoveryCancel := context.WithTimeout(context.Background(), ContextTimeout*time.Second)
		defer recoveryCancel()

		//util.StartSpinner("Waiting for Auto Recover...", 5)
		l.AutoRecover(recoveryCtx, round, sender)
	}(round, valuesAtRound.StartTime)
}

func (l *PoFListener) Commit(ctx context.Context, round *big.Int) (common.Address, []byte, error) {
	style := color.New(color.FgHiBlue, color.Bold)
	style.Println("Preparing to commit...")

	chainID, err := l.Client.NetworkID(ctx)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, chainID)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	nonce, err := l.Client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	gasPrice, err := l.Client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice

	randomData := make([]byte, 32)
	if _, err := rand.Read(randomData); err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to generate random data: %v", err)
	}

	hexData := hex.EncodeToString(randomData)
	byteData, err := hex.DecodeString(hexData)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to decode hex data: %v", err)
	}

	commitData := struct {
		Val    []byte
		Bitlen *big.Int
	}{
		Val:    byteData,
		Bitlen: big.NewInt(int64(len(byteData) * 8)), // Assuming byteData is directly the value committed
	}

	packedData, err := l.ContractABI.Pack("commit", round, commitData)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to pack data for commit: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 3000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	receipt, err := bind.WaitMined(ctx, l.Client, signedTx)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		errMsg := fmt.Sprintf("transaction %s reverted", signedTx.Hash().Hex())
		log.Printf("❌ %s", errMsg)
		fmt.Println("---------------------------------------------------------------------------")
		return common.Address{}, nil, fmt.Errorf("%s", errMsg)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Commit successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	fmt.Println("---------------------------------------------------------------------------")

	return auth.From, byteData, nil // Return the sender address and the committed value
}

func (l *PoFListener) GetNextRound() (*big.Int, error) {
	config := node.LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrngpof.NewCrrrngpof(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{}
	nextRound, err := instance.GetNextRound(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the next round: %v", err)
	}

	return nextRound, nil
}

func (l *PoFListener) GetCommittedOperatorsAtRound(round *big.Int) ([]common.Address, error) {
	config := node.LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrngpof.NewCrrrngpof(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
		return nil, err
	}

	opts := &bind.CallOpts{}
	operators, err := instance.GetCommittedOperatorsAtRound(opts, round)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the committed operators at round %s: %v", round.String(), err)
	}

	if len(operators) == 0 {
		return nil, nil
	}

	return operators, nil
}

func (l *PoFListener) GetLastFulfilledRound() (*big.Int, error) {
	config := node.LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrngpof.NewCrrrngpof(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{}
	lastFulfilledRound, err := instance.LastFulfilledRound(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the last fulfilled round: %v", err)
	}

	return lastFulfilledRound, nil
}

func (l *PoFListener) GetLastRecoveredRound() (*big.Int, error) {
	config := node.LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrngpof.NewCrrrngpof(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{}
	lastRecoveredRound, err := instance.LastRecoveredRound(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the last recovered round: %v", err)
	}

	return lastRecoveredRound, nil
}

func (l *PoFListener) GetValuesAtRound(ctx context.Context, round *big.Int) (*ValueAtRound, error) {
	config := node.LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrngpof.NewCrrrngpof(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{Context: ctx}
	rawResult, err := instance.GetValuesAtRound(opts, round)
	if err != nil {
		log.Fatalf("Failed to retrieve values at round: %v", err)
	}

	//setupValues, err := GetSetupValues(ctx)
	if err != nil {
		log.Fatalf("Failed to retrieve setup values: %v", err)
	}

	result := &ValueAtRound{
		StartTime: rawResult.StartTime,
		//RequestedTime: rawResult.RequestedTime,
		CommitCounts:  rawResult.CommitCounts,
		Consumer:      rawResult.Consumer,
		CommitsString: rawResult.CommitsString,
		Omega:         BigNumber{Val: rawResult.Omega.Val, Bitlen: rawResult.Omega.Bitlen},
		Stage:         GetStage(rawResult.Stage),
		IsCompleted:   rawResult.IsCompleted,
		IsVerified:    rawResult.IsVerified,
		T:             big.NewInt(4194304),
		NBitLen:       big.NewInt(2047),
		GBitLen:       big.NewInt(2046),
		HBitLen:       big.NewInt(2044),
		NVal:          common.FromHex("4e502cc741a1a63c4ae0cea62d6eefae5d0395e137075a15b515f0ced5c811334f06272c0f1e85c1bed5445025b039e42d0a949989e2c210c9b68b9af5ada8c0f72fa445ce8f4af9a2e56478c8a6b17a6f1c389445467fe096a4c35262e4b06a6ba67a419bcca5d565e698ead674fca78e5d91fdc18f854b8e43edbca302c5d2d2d47ce49afb7405a4db2e87c98c2fd0718af32c1881e4d6d762f624de2d57663754aedfb02cbcc944812d2f8de4f694c933a1c11ecdbb2e67cf22f410487d598ef3d82190feabf11b5a83a4a058cdda1def94cd244fd30412eb8fa6d467398c21a15af04bf55078d9c73e12e3d0f5939804845b1487fae1fb526fa583e27d71"),
		GVal:          common.FromHex("34bea67f7d10481d71f794f7bf849b91a460b6488fc0def25ff20b19ff63e984e88daef00289931b566f3e25121e8757751e670a04735a78ff255d804caa197aa65da842913a243add64d375e378380e818b330cc9ef2a89753046248e41eff0f87d8ef4f7764e0ed3698b7f87b07805d235627c80e695f3f6095ca6523312a2916456ed011863d5287a33bf603f495071878ebcb06b9303ffa57ac9b5a77121a20fdbe15004010935d65fc39b199692bbadf172ae84a279f63e31997865c133a6cb8ca4e6c29677a46b932c75297347c605b7fe1c292a96d6401f22b4e4ff474e47cfa59ccfef24d99c3777c98bff523f4a587d54ddc395f572bcde1ae93ba1"),
		HVal:          common.FromHex("08d72e28d1cef1b56bc3047d29624445ce203a0c6de5343a5f4873b4017f479e93fc4c3179d4db28dc7e4a6c859469868e50f3347b8736da84cd0995c661b99df90afa21267a8d7588704b9fc249bac3a3087ff1372f8fbfe1f8625c1a42113ebda7fc364a27d8a0c85dab8802f1b3983e867c3b11fedab831b5d6c1d49a906dd5366dd30816c174d6d384295e0229ddb1685eb5c57b9cde512ff50d82bf659eff8b9f3c8d2f0c2737c83eb44463ca23d93e29fa9630c06809b8a6327a29468e19042a7eac025c234be9fe349a19d7b3e5e4acca63f0b4a592b1749a15a1f054689b1809a4b95b27b8513fa1639c98ca9e18113bf36d631944c37459b5575a17"),
	}

	//fmt.Printf("GET Values: %+v\n", result)
	return result, nil
}

func GetCommitValue(ctx context.Context, round *big.Int, totalCommits int64) ([]*CommitValue, error) {
	config := node.LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrngpof.NewCrrrngpof(contractAddress, client)
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{Context: ctx}
	commitValues := make([]*CommitValue, 0, totalCommits)

	for i := int64(0); i < totalCommits; i++ {
		index := big.NewInt(i)
		value, err := instance.GetCommitValue(opts, round, index)
		if err != nil {
			return nil, err
		}

		commitValue := &CommitValue{
			Commit: BigNumber{
				Val:    value.Commit.Val,
				Bitlen: value.Commit.Bitlen,
			},
			OperatorAddress: value.OperatorAddress,
		}
		commitValues = append(commitValues, commitValue)
	}

	//fmt.Println("commitValue: ", commitValues)
	return commitValues, nil
}

func GetStage(stageValue uint8) string {
	stages := []string{"Finished", "Commit", "Reveal"}
	if int(stageValue) < len(stages) {
		return stages[stageValue]
	}
	return "Unknown"
}

func (b BigNumber) ToBigInt() *big.Int {
	return new(big.Int).SetBytes(b.Val)
}

func (l *PoFListener) AutoRecover(ctx context.Context, round *big.Int, mySender common.Address) error {
	//style := color.New(color.FgHiBlue, color.Bold)
	//style.Println("Preparing to execute recovery process...")

	valuesAtRound, err := l.GetValuesAtRound(ctx, round)
	if err != nil {
		log.Printf("Error retrieving values at round: %v", err)
		return err
	}

	commitDataList, err := GetCommitValue(ctx, round, valuesAtRound.CommitCounts.Int64())
	if err != nil {
		log.Printf("Error retrieving commit-reveal data: %v", err)
		return err
	}

	isMyHashMin, minSender, err := l.FindMinHashAndCompare(ctx, round, mySender)
	if err != nil {
		log.Printf("Error in FindMinHashAndCompare: %v", err)
		return err
	}

	var commits []*big.Int
	for _, commitData := range commitDataList {
		if commitData.Commit.Val != nil {
			commits = append(commits, commitData.Commit.ToBigInt())
		}
	}

	n := new(big.Int).SetBytes(valuesAtRound.NVal)
	T := int(valuesAtRound.T.Int64())

	omegaRecov, proofListRecovery := crr.Recover(n, T, commits)

	if len(proofListRecovery) == 0 {
		return fmt.Errorf("proofListRecovery is empty")
	}

	x := BigNumber{
		Val:    proofListRecovery[0].X.Bytes(),
		Bitlen: big.NewInt(int64(proofListRecovery[0].X.BitLen())),
	}

	y := BigNumber{
		Val:    proofListRecovery[0].Y.Bytes(),
		Bitlen: big.NewInt(int64(proofListRecovery[0].Y.BitLen())),
	}

	v := make([]BigNumber, len(proofListRecovery))
	for i, proof := range proofListRecovery {
		v[i] = BigNumber{
			Val:    proof.V.Bytes(),
			Bitlen: big.NewInt(int64(proof.V.BitLen())),
		}
	}

	if isMyHashMin {
		fmt.Println("---------------------------------------------------------------------------")
		color.New(color.FgHiGreen, color.Bold).Println("My hash is the minimum hash for this round. I am the leader, proceeding with the leader responsibilities.")
		l.LeaderRounds[round] = mySender

		valuesAtRound, err := l.GetValuesAtRound(ctx, round)
		if err != nil {
			log.Printf("Error retrieving values at round: %v", err)
			return err
		}

		currentTime := time.Now()
		startTime := time.Unix(valuesAtRound.StartTime.Int64(), 0)
		commitDuration := time.Duration(CommitDuration) * time.Second
		recoverTime := startTime.Add(commitDuration)

		if currentTime.Before(recoverTime) {
			time.Sleep(recoverTime.Sub(currentTime))
		}

		if time.Now().After(recoverTime) {
			time.Sleep(CommitDuration * time.Second)
			err = l.Recover(ctx, round, y)
			if err != nil {
				log.Printf("Failed to execute recovery process: %v", err)
				delete(l.LeaderRounds, round)
				return err
			}
		}
	} else {
		fmt.Println("---------------------------------------------------------------------------")
		color.New(color.FgHiGreen, color.Bold).Println("Leader Address: ", minSender)

		l.LeaderRounds[round] = minSender
		l.SubscribeRecovered(ctx, omegaRecov.String(), v, x, y, minSender)
		log.Printf("My hash is not the minimum hash for this round. Minimum hash is held by: %s", minSender.Hex())
	}

	return nil
}

func (l *PoFListener) FindMinHashAndCompare(ctx context.Context, round *big.Int, mySender common.Address) (bool, common.Address, error) {
	roundPrefix := fmt.Sprintf("Round %s - ", round.String()) // Creating a prefix string

	valuesAtRound, err := l.GetValuesAtRound(ctx, round)
	if err != nil {
		log.Printf("%sError fetching values at round: %v", roundPrefix, err)
		return false, common.Address{}, err
	}

	commitDataList, err := GetCommitValue(ctx, round, valuesAtRound.CommitCounts.Int64())
	if err != nil {
		log.Printf("%sError fetching commit values: %v", roundPrefix, err)
		return false, common.Address{}, err
	}

	var minHash *big.Int
	var minSender common.Address
	var myHash *big.Int

	for _, commit := range commitDataList {
		if commit == nil {
			continue
		}
		currentHash := crypto.Keccak256Hash(append(valuesAtRound.Omega.Val, commit.OperatorAddress.Bytes()...))
		currentHashInt := new(big.Int).SetBytes(currentHash.Bytes())

		if minHash == nil || currentHashInt.Cmp(minHash) < 0 {
			minHash = currentHashInt
			minSender = commit.OperatorAddress
			//log.Printf("%sNew min hash %s found from address %s", roundPrefix, minHash.String(), minSender.Hex())
		}

		if commit.OperatorAddress == mySender {
			myHash = currentHashInt
			//log.Printf("%sMy sender's hash %s from address %s", roundPrefix, myHash.String(), mySender.Hex())
		}
	}

	isMyHashMin := myHash != nil && myHash.Cmp(minHash) == 0 && mySender == minSender
	if isMyHashMin {
		fmt.Println("---------------------------------------------------------------------------")
		color.New(color.FgHiGreen, color.Bold).Printf("%sMy sender's address has the min hash\n", roundPrefix)
		color.New(color.FgHiGreen, color.Bold).Printf("%s👑 I am the leader\n", roundPrefix)
		fmt.Println("---------------------------------------------------------------------------")
	} else {
		fmt.Println("---------------------------------------------------------------------------")
		color.New(color.FgHiRed, color.Bold).Printf("%sMy sender's address does not have the min hash.\n", roundPrefix)
		color.New(color.FgHiRed, color.Bold).Printf("%s😢 I am not the leader.\n", roundPrefix)
		fmt.Println("---------------------------------------------------------------------------")
		//log.Printf("%sMin hash %s from address %s", roundPrefix, minHash.String(), minSender.Hex())
	}

	return isMyHashMin, minSender, nil
}

func (l *PoFListener) SubscribeRecovered(ctx context.Context, expectedOmega string, v []BigNumber, x BigNumber, y BigNumber, minSender common.Address) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.ContractAddress},
		Topics: [][]common.Hash{
			{
				l.ContractABI.Events["Recovered"].ID,
			},
		},
	}

	logs := make(chan types.Log)
	sub, err := l.Client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}

	//util.StartSpinner("🎧 Listening for 'Recovered' events...", 1)
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("🎧 Listening for 'Recovered' events...")

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			event := struct {
				Round     *big.Int
				Omega     []byte
				Recoverer common.Address
			}{}
			err := l.ContractABI.UnpackIntoInterface(&event, "Recovered", vLog.Data)
			if err != nil {
				log.Printf("Error unpacking Recovered event: %v", err)
				continue
			}

			omegaRecovered := new(big.Int).SetBytes(event.Omega).String()
			fmt.Printf("🔔 Recovered Event - Round: %s, Omega: %s\n", event.Round.String(), omegaRecovered)

			go func() {
				l.Mutex.Lock()
				delete(l.RoundData, event.Round.String())
				fmt.Println("---------------------------------------------------------------------------")
				fmt.Println("Pending rounds after deletion:")
				for r, v := range l.RoundData {
					fmt.Printf("Round: %s, Value: %s\n", r, v.String())
				}

				var minRound *big.Int
				for _, r := range l.RoundData {
					if minRound == nil || r.Cmp(minRound) < 0 {
						minRound = r
					}
				}
				l.Mutex.Unlock()

				if minRound != nil {
					// Get the start time for the minimum round
					fmt.Printf("Proceeding with the next round: %s\n", minRound.String())
					l.initiateCommitProcess(minRound)
				} else {
					color.New(color.FgHiRed, color.Bold).Printf("No pending rounds to process.\n")
				}
			}()

			_, leader, err := l.GetDisputeEndTimeAndLeaderAtRound(ctx, event.Round)
			if err != nil {
				log.Printf("Error retrieving end time and leader during dispute phase: %v", err)
				return err
			}

			if omegaRecovered != expectedOmega {
				color.New(color.FgHiRed, color.Bold).Printf("\nRecovered omega value does not match expected value during dispute phase: %s\n", omegaRecovered)

				signedTx, err := l.DisputeRecover(ctx, event.Round, v, x, y)
				if err != nil {
					log.Printf("Error initiating DisputeRecover during dispute phase: %v", err)
					return err
				}

				receipt, err := bind.WaitMined(ctx, l.Client, signedTx)
				if err != nil {
					log.Printf("Error waiting for transaction to be mined during dispute phase: %v", err)
					return err
				}

				if receipt.Status == 0 {
					log.Printf("Transaction reverted during dispute phase: %s", signedTx.Hash().Hex())
				} else {
					log.Printf("Transaction successful during dispute phase: %s", signedTx.Hash().Hex())
				}
			} else {
				color.New(color.FgHiBlue, color.Bold).Println("\nRecovered omega value matches the expected value")
				fmt.Println("🔔 Checking if the sender is the leader...")
				color.New(color.FgHiGreen, color.Bold).Println("l.LeaderRounds[round]: ", minSender)
				fmt.Println("")
			}

			l.SubscribeFulfillRandomness(ctx, event.Round, leader, minSender)
		}
	}
}

func (l *PoFListener) SubscribeFulfillRandomness(ctx context.Context, round *big.Int, leader common.Address, minSender common.Address) error {
	fulfillRandomnessTopic := []common.Hash{crypto.Keccak256Hash([]byte("FulfillRandomness(uint256,uint256,bool,address)"))}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.ContractAddress},
		Topics:    [][]common.Hash{fulfillRandomnessTopic},
	}

	logs := make(chan types.Log)
	sub, err := l.Client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("🎧 Listening for 'FulfillRandomness' events...")

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			event := struct {
				Round       *big.Int
				HashedOmega *big.Int
				Success     bool
				Leader      common.Address
			}{}
			err := l.ContractABI.UnpackIntoInterface(&event, "FulfillRandomness", vLog.Data)
			if err != nil {
				if strings.Contains(err.Error(), "abi: improperly encoded boolean value") {
					continue
				}
				log.Printf("Error unpacking FulfillRandomness event: %v", err)
				continue
			}

			fmt.Printf("🔔 FulfillRandomness Event - Round: %s, Hashed Omega: %s, Success: %t\n", event.Round.String(), event.HashedOmega.String(), event.Success)
			fmt.Println("---------------------------------------------------------------------------")

			if leader == minSender {
				color.New(color.FgHiBlue, color.Bold).Println("The leader for round: ", event.Round.String(), " is ", minSender)
				color.New(color.FgHiBlue, color.Bold).Println("No dispute in leadership for round: ", event.Round.String(), " is needed.")
			} else {
				color.New(color.FgHiRed, color.Bold).Printf("The sender of this event is not the leader for round %s. The actual leader is %s.\n", event.Round.String(), l.LeaderRounds[event.Round].Hex())
				fmt.Println("---------------------------------------------------------------------------")
				l.DisputeLeadershipAtRound(ctx, round)
			}
		}
	}
}

func (l *PoFListener) Recover(ctx context.Context, round *big.Int, y BigNumber) error {
	fmt.Println("---------------------------------------------------------------------------")
	color.New(color.FgHiRed, color.Bold).Println("Recover proceeding...")

	chainID, err := l.Client.NetworkID(ctx)
	if err != nil {
		log.Printf("Failed to fetch network ID: %v", err)
		return fmt.Errorf("failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, chainID)
	if err != nil {
		log.Printf("Failed to create authorized transactor: %v", err)
		return fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	nonce, err := l.Client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		log.Printf("Failed to fetch nonce: %v", err)
		return fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice, err = l.Client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// In PoF mode, the node does not send v and x values.
	packedData, err := l.ContractABI.Pack("recover", round, y)
	if err != nil {
		log.Printf("Failed to pack data for recovery: %v", err)
		return fmt.Errorf("failed to pack data for recovery: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 6000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		log.Printf("Failed to sign the transaction: %v", err)
		return fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		log.Printf("Failed to send the signed transaction: %v", err)
		return fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	log.Printf("Recover transaction sent! Tx Hash: %s", signedTx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, l.Client, signedTx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		errMsg := fmt.Sprintf("transaction %s reverted", signedTx.Hash().Hex())
		log.Printf("❌ %s", errMsg)
		return fmt.Errorf("%s", errMsg)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Recover successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	log.Printf("Recover successful! Tx Hash: %s", signedTx.Hash().Hex())

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	delete(l.RoundData, round.String())
	fmt.Println("---------------------------------------------------------------------------")
	color.New(color.FgHiRed, color.Bold).Printf("Pending rounds after deletion: ")
	for r, v := range l.RoundData {
		fmt.Printf("Round: %s, Value: %s\n", r, v.String())
	}

	var minRound *big.Int
	for _, r := range l.RoundData {
		if minRound == nil || r.Cmp(minRound) < 0 {
			minRound = r
		}
	}

	if minRound != nil {
		// Get the start time for the minimum round
		fmt.Printf("Proceeding with the next round: %s\n", minRound.String())
		fmt.Println("---------------------------------------------------------------------------")
		l.initiateCommitProcess(minRound)
	} else {
		fmt.Printf("No pending rounds to process.\n")
	}

	fmt.Printf("Proceeding to dispute process for round: %s.\n", round.String())

	l.handleDisputePhase(round)
	return nil
}

func (l *PoFListener) handleDisputePhase(round *big.Int) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout*time.Second)
		defer cancel()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		secondsLeft := DisputeDuration
		for {
			select {
			case <-ticker.C:
				secondsLeft--
				if secondsLeft%5 == 0 {
					fmt.Printf("⏳ [Dispute Phase] Round %s - Countdown: %d seconds remaining\n", round.String(), secondsLeft)
				}

				if secondsLeft == 0 {
					fmt.Println()
					log.Printf("Dispute phase completed for round %s. Proceeding with FulfillRandomness.", round.String())
					go l.proceedWithFulfillRandomness(round)
					return
				}
			case <-ctx.Done():
				log.Printf("Context canceled during dispute phase for round %s", round.String())
				return
			}
		}
	}()
}

func (l *PoFListener) proceedWithFulfillRandomness(round *big.Int) {
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout*time.Second)
	defer cancel()

	//_, leader, err := l.GetDisputeEndTimeAndLeaderAtRound(ctx, round)

	signedTx, err := l.FulfillRandomness(ctx, round)
	if err != nil {
		log.Printf("Error in FulfillRandomness: %v", err)
		return
	}

	receipt, err := bind.WaitMined(ctx, l.Client, signedTx)
	if err != nil {
		log.Printf("Error waiting for transaction to be mined: %v", err)
		return
	}

	if receipt.Status == 0 {
		log.Printf("Transaction reverted: %s", signedTx.Hash().Hex())
	} else {
		log.Printf("Transaction successful: %s", signedTx.Hash().Hex())
	}
}

func (l *PoFListener) DisputeLeadershipAtRound(ctx context.Context, round *big.Int) error {
	chainID, err := l.Client.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	nonce, err := l.Client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice, err = l.Client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Call the 'disputeLeadershipAtRound' function from the smart contract
	packedData, err := l.ContractABI.Pack("disputeLeadershipAtRound", round)
	if err != nil {
		return fmt.Errorf("failed to pack data for disputeLeadershipAtRound: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 6000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Dispute successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())

	return nil
}

func (l *PoFListener) DisputeRecover(ctx context.Context, round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
	chainID, err := l.Client.NetworkID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	nonce, err := l.Client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice, err = l.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	packedData, err := l.ContractABI.Pack("disputeRecover", round, v, x, y)
	if err != nil {
		return nil, fmt.Errorf("failed to pack data for dispute recover: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 6000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		return nil, fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Dispute Recover successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())

	return signedTx, nil
}

func (l *PoFListener) FulfillRandomness(ctx context.Context, round *big.Int) (*types.Transaction, error) {
	log.Println("Starting FulfillRandomness process")

	chainID, err := l.Client.NetworkID(ctx)
	if err != nil {
		log.Printf("Failed to fetch network ID: %v", err)
		return nil, fmt.Errorf("failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, chainID)
	if err != nil {
		log.Printf("Failed to create authorized transactor: %v", err)
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	nonce, err := l.Client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		log.Printf("Failed to fetch nonce: %v", err)
		return nil, fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice, err = l.Client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	packedData, err := l.ContractABI.Pack("fulfillRandomness", round)
	if err != nil {
		log.Printf("Failed to pack data for fulfillRandomness: %v", err)
		return nil, fmt.Errorf("failed to pack data for fulfillRandomness: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 6000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		log.Printf("Failed to sign the transaction: %v", err)
		return nil, fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		log.Printf("Failed to send the signed transaction: %v", err)
		return nil, fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  FulfillRandomness successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	//log.Printf("FulfillRandomness successful! Tx Hash: %s", signedTx.Hash().Hex())

	return signedTx, nil
}

func (l *PoFListener) GetDisputeEndTimeAndLeaderAtRound(ctx context.Context, round *big.Int) (uint64, common.Address, error) {
	config := node.LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Use an existing client connection; assuming l.Client is already initialized and connected
	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrngpof.NewCrrrngpof(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{
		Context: ctx, // Using context to handle request cancellation and timeouts
	}

	// Using the generated Go bindings to call a method directly
	endTimeBigInt, leader, err := instance.GetDisputeEndTimeAndLeaderAtRound(opts, round)
	if err != nil {
		return 0, common.Address{}, fmt.Errorf("failed to retrieve end time and leader at round %v: %v", round, err)
	}

	// Converting big.Int to uint64 for endTime
	if !endTimeBigInt.IsUint64() {
		return 0, leader, fmt.Errorf("endTime is not a uint64")
	}

	return endTimeBigInt.Uint64(), leader, nil
}

func (l *PoFListener) ReRequestRandomWordAtRound(ctx context.Context, round *big.Int) error {
	style := color.New(color.FgHiBlue, color.Bold)
	style.Println("Preparing to re-request random word at round...")

	chainID, err := l.Client.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	nonce, err := l.Client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	gasPrice, err := l.Client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice

	// Pack the transaction data for calling the smart contract function
	packedData, err := l.ContractABI.Pack("reRequestRandomWordAtRound", round)
	if err != nil {
		return fmt.Errorf("failed to pack data for reRequestRandomWordAtRound: %v", err)
	}

	// Create and sign the transaction
	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 3000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign the transaction: %v", err)
	}

	// Send the transaction
	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, l.Client, signedTx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		errMsg := fmt.Sprintf("transaction %s reverted", signedTx.Hash().Hex())
		log.Printf("❌ %s", errMsg)
		return fmt.Errorf("%s", errMsg)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Re-request successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	return nil
}
