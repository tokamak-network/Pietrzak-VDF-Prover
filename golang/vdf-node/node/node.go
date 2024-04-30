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
	Val    []byte
	Bitlen *big.Int
}

type ValueAtRound struct {
	StartTime         *big.Int
	NumOfParticipants *big.Int
	Count             *big.Int
	Consumer          common.Address
	BStar             []byte
	CommitsString     []byte
	Omega             BigNumber
	Stage             uint8
	IsCompleted       bool
	IsAllRevealed     bool
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

			if _, exists := processedRounds[roundKey]; !exists && event.Round != nil && event.Round.Sign() >= 0 {
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
		l.CommitStarted = false // Reset the flag to allow new events to be processed
		l.Mutex.Unlock()
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

	round, err = l.GetNextRound()
	if err != nil {
		log.Printf("round %s: %v", round.String(), err)
	}
	fmt.Printf("round", round)

	//result, err := l.GetValuesAtRound(ctx, round)
	//fmt.Println("result: ", result)

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

	fmt.Printf("Next Round: %s\n", nextRound.String())
	return nextRound, nil
}

func (l *Listener) GetValuesAtRound(ctx context.Context, round *big.Int) (*ValueAtRound, error) {
	config := LoadConfig()
	client, err := ethclient.Dial(config.RpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
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

	result := &ValueAtRound{
		StartTime:     rawResult.StartTime,
		Count:         rawResult.Count,
		Consumer:      rawResult.Consumer,
		BStar:         rawResult.BStar,
		CommitsString: rawResult.CommitsString,
		Omega:         BigNumber{rawResult.Omega.Val, rawResult.Omega.Bitlen}, // 예제, BigNumber 구조체 필요
		Stage:         rawResult.Stage,
		IsCompleted:   rawResult.IsCompleted,
		IsAllRevealed: rawResult.IsAllRevealed,
	}

	fmt.Printf("Values at Round %s: %+v\n", round.String(), result)
	return result, nil
}

func (l *Listener) Recover(round *big.Int, v []*big.Int, x *big.Int, y *big.Int, bigNumTwoPowerOfDelta []byte, delta *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	fmt.Println("Preparing to execute recovery process...")

	// Fetch network ID and create an authorized transactor
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

	// Packing the data for the 'recover' function call
	packedData, err := l.ContractABI.Pack("recover", round, v, x, y, bigNumTwoPowerOfDelta, delta)
	if err != nil {
		return fmt.Errorf("failed to pack data for recovery: %v", err)
	}

	// Create and send the transaction
	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, nil, 4000000, auth.GasPrice, packedData)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), l.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign the transaction: %v", err)
	}

	if err := l.Client.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("failed to send the signed transaction: %v", err)
	}

	fmt.Printf("Recovery process initiated! Transaction Hash: %s\n", signedTx.Hash().Hex())

	return nil
}
