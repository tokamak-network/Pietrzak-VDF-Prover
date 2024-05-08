package node

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/commit-reveal-recover/crr"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/crrrng"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/util"
	"io/ioutil"
	"log"
	"math/big"
	"sync"
	"time"
)

type Listener struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	ContractABI     abi.ABI
	EventData       []EventInfo
	PrivateKey      *ecdsa.PrivateKey
	CommitStarted   bool
	Mutex           sync.Mutex
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
	StartTime         *big.Int       `json:"startTime"`
	NumOfParticipants *big.Int       `json:"numOfParticipants"`
	Count             *big.Int       `json:"count"`
	Consumer          common.Address `json:"consumer"`
	BStar             []byte         `json:"bStar"`
	CommitsString     []byte         `json:"commitsString"`
	Omega             BigNumber      `json:"omega"`
	Stage             string         `json:"stage"`
	IsCompleted       bool           `json:"isCompleted"`
	IsAllRevealed     bool           `json:"isAllRevealed"`
	T                 *big.Int       `json:"T"`
	NBitLen           *big.Int       `json:"nBitLen"`
	GBitLen           *big.Int       `json:"gBitLen"`
	HBitLen           *big.Int       `json:"hBitLen"`
	NVal              []byte         `json:"nVal"`
	GVal              []byte         `json:"gVal"`
	HVal              []byte         `json:"hVal"`
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

type CommitRevealData struct {
	C                  BigNumber
	A                  BigNumber
	ParticipantAddress common.Address
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

func NewListener(config Config) (*Listener, error) {
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey[2:]) // Strip the "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	abiFilename := "../CRRRNGCoordinator.json"
	contractABI, err := loadContractABI(abiFilename)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(config.ContractAddress)

	return &Listener{
		Client:          client,
		ContractAddress: contractAddress,
		ContractABI:     contractABI,
		PrivateKey:      privateKey,
	}, nil
}

func (l *Listener) SubscribeRandomWordsRequested() {
	var round *big.Int
	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.ContractAddress},
	}

	logs := make(chan types.Log)
	sub, err := l.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	processedRounds := make(map[string]bool)

	util.StartSpinner("🎧 Listening for 'RandomWordsRequested' events...", 5)
	fmt.Println("🎧 Listening for 'RandomWordsRequested' events...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if l.CommitStarted {
				continue // Ignore events if a commit is in progress
			}

			event := EventInfo{}
			l.ContractABI.UnpackIntoInterface(&event, "RandomWordsRequested", vLog.Data)

			round = event.Round
			roundKey := round.String()

			if _, exists := processedRounds[roundKey]; !exists && event.Round != nil && event.Round.Sign() >= 0 && event.Sender.Hex() != "0x0000000000000000000000000000000000000040" {
				fmt.Println("------------------------------------------------------")
				fmt.Printf("🔔 Round: %s, Sender: %s\n", round.String(), event.Sender.Hex())
				l.EventData = append(l.EventData, event)
				processedRounds[roundKey] = true

				l.Mutex.Lock()
				if !l.CommitStarted && l.checkAllEventsReceived() {
					l.CommitStarted = true
					go l.initiateCommitProcess(round)
				}
				l.Mutex.Unlock()
			}
		}
	}
}

func (l *Listener) initiateCommitProcess(round *big.Int) {
	ctx, cancel := context.WithTimeout(context.Background(), 115*time.Second)
	defer cancel()

	// Start commit process
	if err := l.Commit(ctx, round); err != nil {
		log.Printf("Failed to commit: %v", err)
		return
	}

	// Start countdown after commit is successful
	go func() {
		countdown := 120
		bar := pb.StartNew(countdown)
		for i := 0; i < countdown; i++ {
			bar.Increment()
			time.Sleep(1 * time.Second)
		}
		bar.Finish()
		fmt.Println("🕒 Countdown completed. Proceeding to the next step.")

		l.Mutex.Lock()
		l.CommitStarted = false
		l.Mutex.Unlock()

		time.Sleep(10 * time.Second)

		recoveryCtx, recoveryCancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer recoveryCancel()

		util.StartSpinner("Waiting for Recover...", 5)
		l.AutoRecover(recoveryCtx, round)
	}()
}

// Check if all expected events for the round are received
func (l *Listener) checkAllEventsReceived() bool {
	return len(l.EventData) > 0
}

func (l *Listener) Commit(ctx context.Context, round *big.Int) error {
	style := color.New(color.FgHiBlue, color.Bold)
	style.Println("Preparing to commit...")

	chainID, err := l.Client.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch network ID: %v", err)
	}
	//fmt.Println("Chain ID:", chainID)

	auth, err := bind.NewKeyedTransactorWithChainID(l.PrivateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	nonce, err := l.Client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return fmt.Errorf("failed to fetch nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	//fmt.Println("Using Nonce:", auth.Nonce)

	gasPrice, err := l.Client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	//fmt.Println("Gas Price:", auth.GasPrice)

	randomData := make([]byte, 256)
	if _, err := rand.Read(randomData); err != nil {
		return fmt.Errorf("failed to generate random data: %v", err)
	}

	hexData := hex.EncodeToString(randomData)
	byteData, err := hex.DecodeString(hexData)
	if err != nil {
		return fmt.Errorf("failed to decode hex data: %v", err)
	}
	commitData := struct {
		Val    []byte
		Bitlen *big.Int
	}{
		Val:    byteData,
		Bitlen: big.NewInt(2048),
	}

	packedData, err := l.ContractABI.Pack("commit", round, commitData)
	if err != nil {
		return fmt.Errorf("failed to pack data for commit: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 3000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Commit successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())

	return nil
}

func (l *Listener) GetNextRound() (*big.Int, error) {
	config := LoadConfig()
	client, err := ethclient.Dial(config.HttpURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrng.NewCrrrng(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{}
	nextRound, err := instance.GetNextRound(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the next round: %v", err)
	}

	//fmt.Printf("Next Round: %s\n", nextRound.String())
	return nextRound, nil
}

func (l *Listener) GetValuesAtRound(ctx context.Context, round *big.Int) (*ValueAtRound, error) {
	config := LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrng.NewCrrrng(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{Context: ctx}
	rawResult, err := instance.GetValuesAtRound(opts, round)
	if err != nil {
		log.Fatalf("Failed to retrieve values at round: %v", err)
	}

	setupValues, err := GetSetupValues(ctx)
	if err != nil {
		log.Fatalf("Failed to retrieve setup values: %v", err)
	}

	result := &ValueAtRound{
		StartTime:         rawResult.StartTime,
		NumOfParticipants: rawResult.NumOfPariticipants,
		Count:             rawResult.Count,
		Consumer:          rawResult.Consumer,
		BStar:             rawResult.BStar,
		CommitsString:     rawResult.CommitsString,
		Omega:             BigNumber{rawResult.Omega.Val, rawResult.Omega.Bitlen},
		Stage:             GetStage(rawResult.Stage),
		IsCompleted:       rawResult.IsCompleted,
		IsAllRevealed:     rawResult.IsAllRevealed,
		T:                 setupValues.T,
		NBitLen:           setupValues.NBitLen,
		GBitLen:           setupValues.GBitLen,
		HBitLen:           setupValues.HBitLen,
		NVal:              setupValues.NVal,
		GVal:              setupValues.GVal,
		HVal:              setupValues.HVal,
	}

	//jsonResult, err := json.Marshal(result)
	//if err != nil {
	//	log.Fatalf("Failed to marshal result to JSON: %v", err)
	//}

	//fmt.Printf("Values at Round %s: %s\n", round.String(), jsonResult)
	return result, nil
}

func GetSetupValues(ctx context.Context) (*SetupValues, error) {
	config := LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrng.NewCrrrng(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create the contract instance: %v", err)
	}

	opts := &bind.CallOpts{Context: ctx}
	t, nBitLen, gBitLen, hBitLen, nVal, gVal, hVal, err := instance.GetSetUpValues(opts)
	if err != nil {
		log.Fatalf("Failed to retrieve setup values: %v", err)
	}

	result := &SetupValues{
		T:       t,
		NBitLen: nBitLen,
		GBitLen: gBitLen,
		HBitLen: hBitLen,
		NVal:    nVal,
		GVal:    gVal,
		HVal:    hVal,
	}

	//fmt.Printf("Setup Values: %+v\n", result)
	return result, nil
}

func GetCommitRevealValues(ctx context.Context, round *big.Int) (*CommitRevealData, error) {
	config := LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	contractAddress := common.HexToAddress(config.ContractAddress)
	instance, err := crrrng.NewCrrrng(contractAddress, client)
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{Context: ctx}
	index := big.NewInt(0)
	value, err := instance.GetCommitRevealValues(opts, round, index)
	if err != nil {
		return nil, err
	}

	result := &CommitRevealData{
		C: BigNumber{
			Val:    value.C.Val,
			Bitlen: value.C.Bitlen,
		},
		A: BigNumber{
			Val:    value.A.Val,
			Bitlen: value.A.Bitlen,
		},
		ParticipantAddress: value.ParticipantAddress,
	}

	return result, nil
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

func calculateDeltaBytes(delta int) []byte {
	twoPowerOfDelta := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(delta)), nil)
	deltaBytes := twoPowerOfDelta.Bytes()
	targetLength := 32 // 32 bytes, 256 bits

	if len(deltaBytes) < targetLength {
		zeroPadLength := targetLength - len(deltaBytes)
		zeroPad := make([]byte, zeroPadLength)
		deltaBytes = append(zeroPad, deltaBytes...)
	}

	return deltaBytes
}

func (l *Listener) AutoRecover(ctx context.Context, round *big.Int) error {
	style := color.New(color.FgHiBlue, color.Bold)
	style.Println("Preparing to execute recovery process...")

	valuesAtRound, err := l.GetValuesAtRound(ctx, round)
	if err != nil {
		log.Printf("Error retrieving values at round: %v", err)
		return err
	}
	//fmt.Printf("valuesAtRound: %+v\n", valuesAtRound)

	commitData, err := GetCommitRevealValues(ctx, round)
	if err != nil {
		log.Printf("Error retrieving commit-reveal data: %v", err)
		return err
	}
	//fmt.Printf("Commit Data: %+v\n", commitData)

	var commits []*big.Int
	if commitData.C.Val != nil {
		commits = append(commits, commitData.C.ToBigInt())
	}

	n := new(big.Int).SetBytes(valuesAtRound.NVal)
	bStar := new(big.Int).SetBytes(valuesAtRound.BStar)
	T := int(valuesAtRound.T.Int64())

	omegaRecov, proofListRecovery := crr.Recover(n, T, commits, bStar)
	fmt.Printf("[+] Recovered random: %s\n", omegaRecov.String())
	//fmt.Println("Proof List:", proofListRecovery)

	if len(proofListRecovery) == 0 {
		return fmt.Errorf("proofListRecovery is empty")
	}

	x := BigNumber{
		Val:    proofListRecovery[0].X.Bytes(),
		Bitlen: big.NewInt(int64(len(proofListRecovery[0].X.Bytes()) * 8)),
	}
	y := BigNumber{
		Val:    proofListRecovery[0].Y.Bytes(),
		Bitlen: big.NewInt(int64(len(proofListRecovery[0].Y.Bytes()) * 8)),
	}
	v := make([]BigNumber, len(proofListRecovery))
	for i, proof := range proofListRecovery {
		v[i] = BigNumber{
			Val:    proof.V.Bytes(),
			Bitlen: big.NewInt(int64(len(proof.V.Bytes()) * 8)),
		}
	}
	//fmt.Printf("Data being packed: %+v\n", v)

	delta := 9
	twoPowerOfDeltaBytes := calculateDeltaBytes(delta)

	err = l.Recover(ctx, round, v, &x, &y, twoPowerOfDeltaBytes, big.NewInt(int64(delta)))
	if err != nil {
		log.Printf("Failed to execute recovery process: %v", err)
		return err
	}

	return nil
}

func (l *Listener) Recover(ctx context.Context, round *big.Int, v []BigNumber, x *BigNumber, y *BigNumber, twoPowerOfDeltaBytes []byte, delta *big.Int) error {
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

	packedData, err := l.ContractABI.Pack("recover", round, v, x, y, twoPowerOfDeltaBytes, delta)
	if err != nil {
		return fmt.Errorf("failed to pack data for recovery: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 4000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Recover successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())

	return nil
}
