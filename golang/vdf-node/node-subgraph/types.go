package nodeSubgraph

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/node"
	"math/big"
	"sync"
)

const (
	CommitDuration  = 120 // Commit duration in seconds
	DisputeDuration = 180 // Dispute duration in seconds
	ContextTimeout  = 600000
)

type BigNumber struct {
	Val    []byte   `json:"val"`
	Bitlen *big.Int `json:"bitlen"`
}

type RecoveredData struct {
	Round          string `json:"round"`
	BlockTimestamp string `json:"blockTimestamp"`
	ID             string `json:"id"`
	MsgSender      string `json:"msgSender"`
	Omega          string `json:"omega"`
	IsRecovered    bool   `json:"isRecovered"`
}

type PoFClient struct {
	Client          *ethclient.Client
	ContractAddress common.Address
	ContractABI     abi.ABI
	PrivateKey      *ecdsa.PrivateKey
	Mutex           sync.Mutex
	WaitGroup       sync.WaitGroup
	LeaderRounds    map[*big.Int]common.Address
	MyAddress       common.Address
}

func GetConfig() node.Config {
	config := node.LoadConfig()
	return config
}

type RandomWordRequestedStruct struct {
	BlockTimestamp string `json:"blockTimestamp"`
	RoundInfo      struct {
		CommitCount       string `json:"commitCount"`
		ValidCommitCount  string `json:"validCommitCount"`
		IsRecovered       bool   `json:"isRecovered"`
		IsFulfillExecuted bool   `json:"isFulfillExecuted"`
	} `json:"roundInfo"`
	Round string `json:"round"`
}

type CommitData struct {
	Round          string `json:"round"`
	MsgSender      string `json:"msgSender"`
	BlockTimestamp string `json:"blockTimestamp"`
	CommitIndex    string `json:"commitIndex"`
	CommitVal      string `json:"commitVal"`
	ID             string `json:"id"`
}

type RoundResults struct {
	RecoverableRounds           []string
	CommittableRounds           []string
	FulfillableRounds           []string
	ReRequestableRounds         []string
	RecoverDisputeableRounds    []string
	LeadershipDisputeableRounds []string
	CompleteRounds              []string
}

type SetupValues struct {
	T       *big.Int
	NBitLen *big.Int
	GBitLen *big.Int
	HBitLen *big.Int
	NVal    []byte `json:"nVal"`
	GVal    []byte `json:"gVal"`
	HVal    []byte `json:"hVal"`
}

type OperatorNumberChanged struct {
	IsOperator bool `json:"isOperator"`
}

type RecoveryResult struct {
	OmegaRecov *big.Int
	X          BigNumber
	Y          BigNumber
	V          []BigNumber
}
