package nodeSubgraph

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/machinebox/graphql"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/commit-reveal-recover/crr"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/node"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

var roundStatus sync.Map

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

func NewPoFClient(config node.Config) (*PoFClient, error) {
	client, err := ethclient.Dial(config.HttpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey[2:])
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

	return &PoFClient{
		Client:          client,
		ContractAddress: contractAddress,
		ContractABI:     contractABI,
		PrivateKey:      privateKey,
		LeaderRounds:    make(map[*big.Int]common.Address),
		MyAddress:       myAddress,
	}, nil
}

func (l *PoFClient) GetRandomWordRequested() (*RoundResults, error) {
	config := GetConfig()
	client := graphql.NewClient(config.SubgraphURL)

	req := graphql.NewRequest(`
    query MyQuery {
        randomWordsRequesteds(orderBy: blockTimestamp, orderDirection: desc, first: 50) {
            blockTimestamp
            roundInfo {
                commitCount
                validCommitCount
                isRecovered
                isFulfillExecuted
            }
            round
        }
    }`)

	ctx := context.Background()

	var respData struct {
		RandomWordsRequested []RandomWordRequestedStruct `json:"randomWordsRequesteds"`
	}
	if err := client.Run(ctx, req, &respData); err != nil {
		return nil, err
	}

	// Use a map to keep only the latest entry per round
	latestRounds := make(map[string]RandomWordRequestedStruct)
	for _, item := range respData.RandomWordsRequested {
		if existing, ok := latestRounds[item.Round]; ok {
			// Compare blockTimestamps and update if the current item is more recent
			existingTimestamp, _ := strconv.Atoi(existing.BlockTimestamp)
			currentTimestamp, _ := strconv.Atoi(item.BlockTimestamp)
			if currentTimestamp > existingTimestamp {
				latestRounds[item.Round] = item
			}
		} else {
			latestRounds[item.Round] = item
		}
	}

	var rounds []struct {
		RoundInt int
		Data     RandomWordRequestedStruct
	}

	for round, data := range latestRounds {
		roundInt, err := strconv.Atoi(round)
		if err != nil {
			log.Printf("Error converting round to int: %s, %v", round, err)
			continue
		}
		rounds = append(rounds, struct {
			RoundInt int
			Data     RandomWordRequestedStruct
		}{RoundInt: roundInt, Data: data})
	}

	// Filter out items where isFulfillExecuted is true
	var filteredRounds []struct {
		RoundInt int
		Data     RandomWordRequestedStruct
	}

	for _, round := range rounds {
		if !round.Data.RoundInfo.IsFulfillExecuted {
			filteredRounds = append(filteredRounds, round)
		}
	}

	// Sort the filtered rounds by RoundInt
	sort.Slice(filteredRounds, func(i, j int) bool {
		return filteredRounds[i].RoundInt < filteredRounds[j].RoundInt
	})

	// Initialize RoundResults structure
	results := &RoundResults{
		RecoverableRounds:           []string{},
		CommittableRounds:           []string{},
		FulfillableRounds:           []string{},
		ReRequestableRounds:         []string{},
		RecoverDisputeableRounds:    []string{},
		LeadershipDisputeableRounds: []string{},
		CompleteRounds:              []string{},
		RecoveryData:                []RecoveryResult{},
	}

	// Display the filtered rounds
	for _, round := range filteredRounds {
		fmt.Printf("Round: %d, Data: %+v\n", round.RoundInt, round.Data)
	}

	for _, round := range filteredRounds {
		item := round.Data
		reqOne := graphql.NewRequest(`
		query MyQuery($round: String!, $msgSender: String!) {
		  commitCs(where: {round: $round, msgSender: $msgSender}) {
			blockTimestamp
			commitVal
		  }
		}`)

		reqOne.Var("round", item.Round)
		reqOne.Var("msgSender", config.WalletAddress)

		var respOneData struct {
			CommitCs []struct {
				BlockTimestamp string `json:"blockTimestamp"`
				CommitVal      string `json:"commitVal"`
			} `json:"commitCs"`
		}

		if err := client.Run(ctx, reqOne, &respOneData); err != nil {
			fmt.Println("Error running query:", err)
		}

		var myCommitBlockTimestamp time.Time

		for _, data := range respOneData.CommitCs {
			myCommitBlockTimestampInt, err := strconv.ParseInt(data.BlockTimestamp, 10, 64)
			if err != nil {
				log.Printf("Error converting block timestamp to int64: %v", err)
				return nil, err
			}
			myCommitBlockTimestamp = time.Unix(myCommitBlockTimestampInt, 0)
		}

		//commitCount, err := strconv.Atoi(item.RoundInfo.CommitCount)
		//if err != nil {
		//	log.Printf("Error converting ValidCommitCount to int: %v", err)
		//	continue
		//}

		validCommitCount, err := strconv.Atoi(item.RoundInfo.ValidCommitCount)
		if err != nil {
			log.Printf("Error converting ValidCommitCount to int: %v", err)
			continue
		}

		// get recovered data
		recoveredData, err := GetRecoveredData(item.Round)
		var recoverPhaseEndTime time.Time
		var isRecovered bool
		var omega string
		var msgSender string

		if err != nil {
			log.Printf("Error retrieving recovered data for round %s: %v", item.Round, err)
		}

		for _, data := range recoveredData {
			blockTimestamp, err := strconv.ParseInt(data.BlockTimestamp, 10, 64)
			if err != nil {
				log.Printf("Failed to parse block timestamp for round %s: %v", item.Round, err)
				continue
			}

			isRecovered = data.IsRecovered
			omega = data.Omega
			msgSender = data.MsgSender
			blockTime := time.Unix(blockTimestamp, 0)
			recoverPhaseEndTime = blockTime.Add(DisputeDuration * time.Second)
		}

		// get fulfilled data
		fulfillData, err := GetFulfillRandomnessData(item.Round)

		//var isFulfillSuccess bool
		var fulfillSender string

		if err != nil {
			log.Printf("Error retrieving fulfill randomness data for round %s: %v", item.Round, err)
			//isFulfillSuccess = false
		} else {
			for _, data := range fulfillData {
				if data.Success {
					//isFulfillSuccess = true
					fulfillSender = data.MsgSender
					break
				}
			}
		}

		// get committed data
		getCommitData, err := GetCommitData(item.Round)
		if err != nil {
			log.Printf("Error retrieving commit data for round %s: %v", item.Round, err)
		}

		var commitSenders []common.Address
		var isCommitSender bool
		var commitTimeStampStr string

		for _, data := range getCommitData {
			commitSender := common.HexToAddress(data.MsgSender)
			commitSenders = append(commitSenders, commitSender)
			commitTimeStampStr = data.BlockTimestamp
		}

		for _, commitSender := range commitSenders {
			if commitSender == common.HexToAddress(config.WalletAddress) {
				isCommitSender = true
				break
			}
		}

		var isMyAddressLeader bool
		var leaderAddress common.Address
		var recoverData RecoveryResult
		//var loaded bool

		//if validCommitCount > 1 {
		//	recoverData, loaded, err = loadRecoveryDataFromFile(item.Round)
		//	if !loaded {
		//		recoverData, err = l.BeforeRecoverPhase(item.Round)
		//		if err != nil {
		//			log.Printf("Error processing BeforeRecoverPhase for round %s: %v", item.Round, err)
		//			continue
		//		}
		//
		//		err = saveRecoveryDataToFile(recoverData, item.Round)
		//		if err != nil {
		//			log.Printf("Failed to save recovery data to file for round %s: %v", item.Round, err)
		//			continue
		//		}
		//	}
		//
		//	results.RecoveryData = append(results.RecoveryData, recoverData)
		//	isMyAddressLeader, leaderAddress, _ = FindOffChainLeaderAtRound(item.Round, recoverData.OmegaRecov)
		//}

		recoverDataMap := make(map[string]RecoveryResult)

		if validCommitCount >= 2 && !isRecovered {
			recoverData, err = l.BeforeRecoverPhase(item.Round)
			if err != nil {
				log.Printf("Error processing BeforeRecoverPhase for round %s: %v", item.Round, err)
				continue
			}

			if recoverData.OmegaRecov == nil {
				log.Printf("OmegaRecov is nil for round %s", item.Round)
				continue
			}

			recoverDataMap[item.Round] = recoverData
			isMyAddressLeader, leaderAddress, _ = FindOffChainLeaderAtRound(item.Round, recoverData.OmegaRecov)
			results.RecoveryData = append(results.RecoveryData, recoverData)
		} else if validCommitCount >= 2 && isRecovered {
			if strings.ToLower(config.WalletAddress) == msgSender {
				isMyAddressLeader = true
				leaderAddress = common.HexToAddress(msgSender)
			} else {
				isMyAddressLeader = false
				leaderAddress = common.HexToAddress(msgSender)
			}
		}

		var isPreviousRoundRecovered bool
		previousRoundInt, err := strconv.Atoi(item.Round)
		if err != nil {
			log.Printf("Error converting round to int: %v", err)
			continue
		}

		previousRound := strconv.Itoa(previousRoundInt - 1)

		previousRoundData, err := GetRecoveredData(previousRound)
		if err != nil {
			log.Printf("Error retrieving recovered data for previous round %s: %v", previousRound, err)
		} else {
			isPreviousRoundRecovered = false
			for _, data := range previousRoundData {
				if data.IsRecovered {
					isPreviousRoundRecovered = true
					break
				}
			}
		}

		requestBlockTimestampStr := item.BlockTimestamp
		requestBlockTimestampInt, err := strconv.ParseInt(requestBlockTimestampStr, 10, 64)
		if err != nil {
			log.Printf("Error converting block timestamp to int64: %v", err)
			return nil, err
		}
		requestBlockTimestamp := time.Unix(requestBlockTimestampInt, 0)
		//fmt.Println("item Round: ", item.Round, " requestBlockTimestampStr: ", requestBlockTimestampStr)
		//fmt.Println("item Round: ", item.Round, " requestBlockTimestamp: ", requestBlockTimestamp)
		//fmt.Println("myCommitBlockTimestamp: ", myCommitBlockTimestamp)

		//requestBlockTimestampEndTime := requestBlockTimestamp.Add(4 * time.Minute)

		if commitTimeStampStr == "" {
			commitTimeStampStr = "0"
		}

		commitTimeStampInt, err := strconv.ParseInt(commitTimeStampStr, 10, 64)
		if err != nil {
			log.Printf("Error converting commit timestamp to int64: %v", err)
			return nil, err
		}
		commitTimeStampTime := time.Unix(commitTimeStampInt, 0)
		commitPhaseEndTime := commitTimeStampTime.Add(time.Duration(CommitDuration) * time.Second)
		reRequestTime := commitTimeStampTime.Add(120 * time.Second)

		roundStr := item.Round

		if item.Round == "0" {
			isPreviousRoundRecovered = true
		}

		// Commit
		go func() {
			if isPreviousRoundRecovered && !item.RoundInfo.IsRecovered && requestBlockTimestamp.After(myCommitBlockTimestamp) {
				_, reRequestExists := roundStatus.Load(roundStr + ":ReRequested")
				if _, exists := roundStatus.Load(roundStr + ":Committed"); !exists {
					results.CommittableRounds = append(results.CommittableRounds, roundStr)
					roundStatus.Store(roundStr+":Committed", "Processed")

					if reRequestExists {
						roundStatus.Delete(roundStr + ":ReRequested")
					}
				}
			}
		}()

		// Recover
		if !isRecovered && isMyAddressLeader && isCommitSender && commitPhaseEndTime.Before(time.Now()) && !item.RoundInfo.IsRecovered && !item.RoundInfo.IsFulfillExecuted && validCommitCount > 1 {
			if _, exists := roundStatus.Load(roundStr + ":Recovered"); !exists {
				results.RecoverableRounds = append(results.RecoverableRounds, roundStr)
				roundStatus.Store(roundStr+":Recovered", "Processed")
			}
		}

		// Fulfill
		if isMyAddressLeader && isCommitSender && recoverPhaseEndTime.Before(time.Now()) && item.RoundInfo.IsRecovered && !item.RoundInfo.IsFulfillExecuted && validCommitCount > 1 {
			if _, exists := roundStatus.Load(roundStr + ":Fulfilled"); !exists {
				results.FulfillableRounds = append(results.FulfillableRounds, roundStr)
				roundStatus.Store(roundStr+":Fulfilled", "Processed")
			}
		}

		// Re-request
		if isPreviousRoundRecovered && reRequestTime.Before(time.Now()) && !item.RoundInfo.IsRecovered && validCommitCount < 2 && validCommitCount > 0 && commitTimeStampStr != "0" {
			isRoundAlreadyCommittable := false
			for _, committableRound := range results.CommittableRounds {
				if committableRound == roundStr {
					isRoundAlreadyCommittable = true
					break
				}
			}

			if !isRoundAlreadyCommittable {
				_, commitExists := roundStatus.Load(roundStr + ":Committed")
				if _, exists := roundStatus.Load(roundStr + ":ReRequested"); !exists {
					results.ReRequestableRounds = append(results.ReRequestableRounds, roundStr)
					roundStatus.Store(roundStr+":ReRequested", "Processed")

					if commitExists {
						roundStatus.Delete(roundStr + ":Committed")
					}
				}
			}
		}

		// Dispute Recover
		if !isMyAddressLeader && isCommitSender && time.Now().Before(recoverPhaseEndTime) && item.RoundInfo.IsRecovered && !item.RoundInfo.IsFulfillExecuted {
			roundBigInt := new(big.Int)
			roundBigInt.SetString(item.Round, 10)

			if data, exists := recoverDataMap[item.Round]; exists {
				omega = strings.TrimPrefix(omega, "0x")
				omegaBigInt := new(big.Int)
				if _, ok := omegaBigInt.SetString(omega, 16); !ok {
					log.Printf("Failed to parse omega: %s", omega)
				}

				if omegaBigInt.Cmp(data.OmegaRecov) != 0 {
					if _, exists := roundStatus.Load(item.Round + ":DisputeRecovered"); !exists {
						if !containsRound(results.RecoverDisputeableRounds, item.Round) {
							results.RecoverDisputeableRounds = append(results.RecoverDisputeableRounds, item.Round)
							roundStatus.Store(item.Round+":DisputeRecovered", "Processed")

							committedKey := item.Round + ":Committed"
							if _, exists := roundStatus.Load(committedKey); exists {
								roundStatus.Delete(committedKey)
							}
						}
					}
				}
			} else {
				log.Printf("No recovery data found for round %s", item.Round)
			}
		}

		// Dispute Leadership
		if !isMyAddressLeader && isCommitSender && time.Now().Before(recoverPhaseEndTime) && item.RoundInfo.IsRecovered && item.RoundInfo.IsFulfillExecuted {
			fulfillSenderAddress := common.HexToAddress(fulfillSender)

			if fulfillSenderAddress != leaderAddress {
				if _, exists := roundStatus.Load(roundStr + ":DisputeLeadershiped"); !exists {
					if !containsRound(results.LeadershipDisputeableRounds, roundStr) {
						results.LeadershipDisputeableRounds = append(results.LeadershipDisputeableRounds, roundStr)
						roundStatus.Store(roundStr+":DisputeLeadershiped", "Processed")
					}
				}
			}
		}
	}

	fmt.Println("---------------------------------------------------------------------------")
	w := tabwriter.NewWriter(log.Writer(), 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Category\tRounds")
	fmt.Fprintln(w, "RecoverableRounds\t", results.RecoverableRounds)
	fmt.Fprintln(w, "CommittableRounds\t", results.CommittableRounds)
	fmt.Fprintln(w, "FulfillableRounds\t", results.FulfillableRounds)
	fmt.Fprintln(w, "ReRequestableRounds\t", results.ReRequestableRounds)
	fmt.Fprintln(w, "RecoverDisputeableRounds\t", results.RecoverDisputeableRounds)
	fmt.Fprintln(w, "LeadershipDisputeableRounds\t", results.LeadershipDisputeableRounds)
	w.Flush()
	fmt.Println("---------------------------------------------------------------------------")

	return results, nil
}

func containsRound(rounds []string, round string) bool {
	for _, r := range rounds {
		if r == round {
			return true
		}
	}
	return false
}

func (l *PoFClient) ProcessRoundResults() error {
	config := GetConfig()
	isOperator, err := IsOperator(config.WalletAddress)
	if err != nil {
		log.Printf("Error fetching isOperator results: %v", err)
		return err
	}

	if !isOperator {
		ctx := context.Background()
		l.OperatorDeposit(ctx)
	}

	results, err := l.GetRandomWordRequested()
	if err != nil {
		log.Printf("Error fetching round results: %v", err)
		return err
	}

	if len(results.RecoverableRounds) > 0 {
		fmt.Println("Processing Recoverable Rounds...")
		processedRounds := make(map[string]bool)

		for _, roundStr := range results.RecoverableRounds {
			if processedRounds[roundStr] {
				continue
			}

			for _, recoveryData := range results.RecoveryData {
				isMyAddressLeader, _, _ := FindOffChainLeaderAtRound(roundStr, recoveryData.OmegaRecov)
				if isMyAddressLeader {
					round := new(big.Int)
					round, ok := round.SetString(roundStr, 10)
					if !ok {
						log.Printf("Failed to convert round string to big.Int: %s", roundStr)
						continue
					}

					ctx := context.Background()
					l.Recover(ctx, round, recoveryData.Y)

					fmt.Printf("Processing recoverable round: %s\n", roundStr)
					processedRounds[roundStr] = true

					time.NewTicker(5 * time.Second)
					break
				}
			}

			if !processedRounds[roundStr] {
				fmt.Printf("Not recoverable round: %s\n", roundStr)
			}
		}
	}

	if len(results.CommittableRounds) > 0 {
		fmt.Println("Processing Committable Rounds...")
		for _, roundStr := range results.CommittableRounds {
			round := new(big.Int)
			round, ok := round.SetString(roundStr, 10)
			if !ok {
				log.Printf("Failed to convert round string to big.Int: %s", roundStr)
				continue
			}

			ctx := context.Background()
			l.Commit(ctx, round)

			fmt.Printf("Processing committable round: %s\n", roundStr)
		}
	}

	if len(results.FulfillableRounds) > 0 {
		fmt.Println("Processing Fulfillable Rounds...")
		for _, roundStr := range results.FulfillableRounds {
			round := new(big.Int)
			round, ok := round.SetString(roundStr, 10)
			if !ok {
				log.Printf("Failed to convert round string to big.Int: %s", roundStr)
				continue
			}

			// Check if the index is within bounds
			//if i < len(results.RecoveryData) {
			//isMyAddressLeader, _, _ := FindOffChainLeaderAtRound(roundStr, results.RecoveryData[i].OmegaRecov)
			//if isMyAddressLeader {
			ctx := context.Background()
			l.FulfillRandomness(ctx, round)
			//} else {
			//	fmt.Printf("Not fulfillable round: %s\n", round)
			//}
			//} else {
			//	log.Printf("No recovery data available for round: %s", roundStr)
			//}
		}
	}

	if len(results.ReRequestableRounds) > 0 {
		fmt.Println("Processing ReRequestable Rounds...")
		for _, roundStr := range results.ReRequestableRounds {
			round := new(big.Int)
			round, ok := round.SetString(roundStr, 10)
			if !ok {
				log.Printf("Failed to convert round string to big.Int: %s", roundStr)
				continue
			}

			ctx := context.Background()
			l.ReRequestRandomWordAtRound(ctx, round)

			fmt.Printf("Processing re-requestable round: %s\n", round)
		}
	}

	if len(results.RecoverDisputeableRounds) > 0 {
		fmt.Println("Processing Recover Disputeable Rounds...")
		for _, roundStr := range results.RecoverDisputeableRounds {
			recoveredData, err := GetRecoveredData(roundStr)
			if err != nil {
				log.Printf("Error retrieving recovered data for round %s: %v", roundStr, err)
				continue
			}

			round := new(big.Int)
			round, ok := round.SetString(roundStr, 10)
			if !ok {
				log.Printf("Failed to convert round string to big.Int: %s", roundStr)
				continue
			}

			disputeInitiated := false

			for _, data := range recoveredData {
				msgSender := common.HexToAddress(data.MsgSender)
				omega := new(big.Int)
				omega, ok := omega.SetString(data.Omega[2:], 16)
				if !ok {
					log.Printf("Failed to parse omega for round %s: %s", roundStr, data.Omega)
					continue
				}

				fmt.Printf("Recovered Data - MsgSender: %s, Omega: %s\n", msgSender.Hex(), omega.String())

				for _, recoveryData := range results.RecoveryData {
					if recoveryData.OmegaRecov.Cmp(omega) != 0 && !disputeInitiated {
						ctx := context.Background()
						l.DisputeRecover(ctx, round, recoveryData.V, recoveryData.X, recoveryData.Y)
						disputeInitiated = true
					}
				}

				if disputeInitiated {
					fmt.Printf("Processing disputeable round: %s\n", roundStr)
					break
				}
			}

			if !disputeInitiated {
				fmt.Printf("No disputes initiated for round: %s\n", roundStr)
			}
		}
	}

	if len(results.LeadershipDisputeableRounds) > 0 {
		fmt.Println("Processing Leadership Disputeable Rounds...")
		for i, roundStr := range results.LeadershipDisputeableRounds {
			recoveredData, err := GetRecoveredData(roundStr)
			if err != nil {
				log.Printf("Error retrieving recovered data for round %s: %v", roundStr, err)
				continue
			}

			round := new(big.Int)
			round, ok := round.SetString(roundStr, 10)
			if !ok {
				log.Printf("Failed to convert round string to big.Int: %s", roundStr)
				continue
			}

			var msgSender common.Address

			for _, data := range recoveredData {
				msgSender = common.HexToAddress(data.MsgSender)
				fmt.Printf("Recovered Data - MsgSender: %s\n", msgSender.Hex())
			}

			if i < len(results.RecoveryData) {
				isMyAddressLeader, leaderAddress, _ := FindOffChainLeaderAtRound(roundStr, results.RecoveryData[i].OmegaRecov)

				if msgSender != leaderAddress {
					ctx := context.Background()
					if isMyAddressLeader {
						l.DisputeLeadershipAtRound(ctx, round)
						fmt.Printf("MsgSender %s is not the leader for round %s\n", msgSender.Hex(), roundStr)
					}
				}

				fmt.Printf("Processing disputeable round: %s\n", roundStr)
			} else {
				log.Printf("No recovery data available for round: %s", roundStr)
			}
		}
	}

	return nil
}

func FindOffChainLeaderAtRound(round string, OmegaRecov *big.Int) (bool, common.Address, error) {
	config := GetConfig()
	mySender := common.HexToAddress(config.WalletAddress)
	commitDataList, err := GetCommitData(round)
	if err != nil {
		fmt.Printf("Error fetching commit data for round %s: %v\n", round, err)
		return false, common.Address{}, err
	}

	//roundPrefix := fmt.Sprintf("Round %s - ", round) // Creating a prefix string

	var minHash *big.Int
	var leaderAddress common.Address
	var myHash *big.Int

	for _, commit := range commitDataList {
		commitAddress := common.HexToAddress(commit.MsgSender)
		//dataToHash := append([]byte(commit.BlockTimestamp), commitAddress.Bytes()...)
		dataToHash := append(commitAddress.Bytes(), OmegaRecov.Bytes()...)
		currentHash := crypto.Keccak256Hash(dataToHash)
		currentHashInt := new(big.Int).SetBytes(currentHash.Bytes())

		if minHash == nil || currentHashInt.Cmp(minHash) < 0 {
			minHash = currentHashInt
			leaderAddress = commitAddress
		}

		if commitAddress == mySender {
			myHash = currentHashInt
		}
	}

	isMyAddressLeader := myHash != nil && myHash.Cmp(minHash) == 0 && mySender == leaderAddress
	//if isMyAddressLeader {
	//	fmt.Println("---------------------------------------------------------------------------")
	//	color.New(color.FgHiGreen, color.Bold).Printf("%sMy sender's address has the min hash\n", roundPrefix)
	//	color.New(color.FgHiGreen, color.Bold).Printf("%s👑 I am the leader\n", roundPrefix)
	//	fmt.Println("---------------------------------------------------------------------------")
	//} else {
	//	fmt.Println("---------------------------------------------------------------------------")
	//	color.New(color.FgHiRed, color.Bold).Printf("%sMy sender's address does not have the min hash.\n", roundPrefix)
	//	color.New(color.FgHiRed, color.Bold).Printf("%s😢 I am not the leader.\n", roundPrefix)
	//	fmt.Println("---------------------------------------------------------------------------")
	//}

	return isMyAddressLeader, leaderAddress, nil
}

// GetCommitData retrieves commit data for a given round and returns a slice of CommitData and an error
func GetCommitData(round string) ([]CommitData, error) {
	config := GetConfig()
	client := graphql.NewClient(config.SubgraphURL)

	req := graphql.NewRequest(`
    query MyQuery($round: String!) {
        commitCs(where: {round: $round}) {
            round
            msgSender
            blockTimestamp
            commitIndex
			commitVal
            id
        }
    }`)

	// Set the variable for the round
	req.Var("round", round)

	// Define a structure to hold the query response
	var respData struct {
		CommitCs []CommitData `json:"commitCs"`
	}

	ctx := context.Background()
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Printf("Failed to execute query: %v", err)
		return nil, err
	}

	//for _, commit := range respData.CommitCs {
	//	fmt.Printf("Commit Data: Round: %s, Sender: %s, Timestamp: %s, Index: %s, ID: %s\n",
	//		commit.Round, commit.MsgSender, commit.BlockTimestamp, commit.CommitIndex, commit.ID)
	//}

	// Return the list of commit data and no error
	return respData.CommitCs, nil
}

// GetRecoveredData fetches recovered data from a GraphQL endpoint
func GetRecoveredData(round string) ([]RecoveredData, error) {
	config := GetConfig()
	client := graphql.NewClient(config.SubgraphURL)

	req := graphql.NewRequest(`
        query MyQuery($round: String!) {
          recovereds(orderBy: blockTimestamp, orderDirection: asc, where: {round: $round}) {
            round
            blockTimestamp
            id
            msgSender
            omega
            roundInfo {
              isRecovered
            }
          }
        }`)

	req.Var("round", round)

	var respData struct {
		Recovereds []struct {
			Round          string `json:"round"`
			BlockTimestamp string `json:"blockTimestamp"`
			ID             string `json:"id"`
			MsgSender      string `json:"msgSender"`
			Omega          string `json:"omega"`
			RoundInfo      struct {
				IsRecovered bool `json:"isRecovered"`
			} `json:"roundInfo"`
		} `json:"recovereds"`
	}

	ctx := context.Background()
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Printf("Failed to execute query: %v", err)
		return nil, err
	}

	var recoveredData []RecoveredData
	for _, item := range respData.Recovereds {
		recoveredData = append(recoveredData, RecoveredData{
			Round:          item.Round,
			BlockTimestamp: item.BlockTimestamp,
			ID:             item.ID,
			MsgSender:      item.MsgSender,
			Omega:          item.Omega,
			IsRecovered:    item.RoundInfo.IsRecovered,
		})
	}

	return recoveredData, nil
}

func GetFulfillRandomnessData(round string) ([]FulfillRandomnessData, error) {
	config := GetConfig()
	client := graphql.NewClient(config.SubgraphURL)

	req := graphql.NewRequest(`
        query MyQuery($round: String!) {
          fulfillRandomnesses(where: {round: $round}) {
            msgSender
            blockTimestamp
            success
          }
        }`)

	req.Var("round", round)

	var respData struct {
		FulfillRandomnesses []struct {
			MsgSender      string `json:"msgSender"`
			BlockTimestamp string `json:"blockTimestamp"`
			Success        bool   `json:"success"`
		} `json:"fulfillRandomnesses"`
	}

	ctx := context.Background()
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Printf("Failed to execute query: %v", err)
		return nil, err
	}

	var fulfillRandomnessData []FulfillRandomnessData
	for _, item := range respData.FulfillRandomnesses {
		fulfillRandomnessData = append(fulfillRandomnessData, FulfillRandomnessData{
			MsgSender:      item.MsgSender,
			BlockTimestamp: item.BlockTimestamp,
			Success:        item.Success,
		})
	}

	return fulfillRandomnessData, nil
}

func GetSetupValue() SetupValues {
	setupValues := SetupValues{
		T:       big.NewInt(4194304),
		NBitLen: big.NewInt(2047),
		GBitLen: big.NewInt(2046),
		HBitLen: big.NewInt(2044),
		NVal:    common.FromHex("4e502cc741a1a63c4ae0cea62d6eefae5d0395e137075a15b515f0ced5c811334f06272c0f1e85c1bed5445025b039e42d0a949989e2c210c9b68b9af5ada8c0f72fa445ce8f4af9a2e56478c8a6b17a6f1c389445467fe096a4c35262e4b06a6ba67a419bcca5d565e698ead674fca78e5d91fdc18f854b8e43edbca302c5d2d2d47ce49afb7405a4db2e87c98c2fd0718af32c1881e4d6d762f624de2d57663754aedfb02cbcc944812d2f8de4f694c933a1c11ecdbb2e67cf22f410487d598ef3d82190feabf11b5a83a4a058cdda1def94cd244fd30412eb8fa6d467398c21a15af04bf55078d9c73e12e3d0f5939804845b1487fae1fb526fa583e27d71"),
		GVal:    common.FromHex("34bea67f7d10481d71f794f7bf849b91a460b6488fc0def25ff20b19ff63e984e88daef00289931b566f3e25121e8757751e670a04735a78ff255d804caa197aa65da842913a243add64d375e378380e818b330cc9ef2a89753046248e41eff0f87d8ef4f7764e0ed3698b7f87b07805d235627c80e695f3f6095ca6523312a2916456ed011863d5287a33bf603f495071878ebcb06b9303ffa57ac9b5a77121a20fdbe15004010935d65fc39b199692bbadf172ae84a279f63e31997865c133a6cb8ca4e6c29677a46b932c75297347c605b7fe1c292a96d6401f22b4e4ff474e47cfa59ccfef24d99c3777c98bff523f4a587d54ddc395f572bcde1ae93ba1"),
		HVal:    common.FromHex("08d72e28d1cef1b56bc3047d29624445ce203a0c6de5343a5f4873b4017f479e93fc4c3179d4db28dc7e4a6c859469868e50f3347b8736da84cd0995c661b99df90afa21267a8d7588704b9fc249bac3a3087ff1372f8fbfe1f8625c1a42113ebda7fc364a27d8a0c85dab8802f1b3983e867c3b11fedab831b5d6c1d49a906dd5366dd30816c174d6d384295e0229ddb1685eb5c57b9cde512ff50d82bf659eff8b9f3c8d2f0c2737c83eb44463ca23d93e29fa9630c06809b8a6327a29468e19042a7eac025c234be9fe349a19d7b3e5e4acca63f0b4a592b1749a15a1f054689b1809a4b95b27b8513fa1639c98ca9e18113bf36d631944c37459b5575a17"),
	}

	return setupValues
}

func IsOperator(operator string) (bool, error) {
	config := GetConfig()
	client := graphql.NewClient(config.SubgraphURL)

	req := graphql.NewRequest(`
		query MyQuery($operator: String!) {
			operatorNumberChangeds(where: {operator: $operator}) {
				isOperator
			}
		}
	`)

	req.Var("operator", operator)

	var respData struct {
		OperatorNumberChangeds []OperatorNumberChanged `json:"operatorNumberChangeds"`
	}

	ctx := context.Background()
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Printf("Failed to execute query: %v", err)
		return false, err
	}

	for _, record := range respData.OperatorNumberChangeds {
		return record.IsOperator, nil
	}

	return false, nil
}

// BeforeRecoverPhase checks if the local node is the leader by recovering the minimum hash and compares it against its own
func (l *PoFClient) BeforeRecoverPhase(round string) (RecoveryResult, error) {
	setupValues := GetSetupValue()

	// Fetch commit data using the round number
	commitDataList, err := GetCommitData(round)
	if err != nil {
		log.Printf("Error retrieving commit-reveal data: %v", err)
		return RecoveryResult{}, err
	}

	// Process commit data to extract commit values
	var commits []*big.Int
	for _, commitData := range commitDataList {
		if commitData.CommitVal != "" {
			var commitBigInt *big.Int
			var ok bool
			if strings.HasPrefix(commitData.CommitVal, "0x") {
				commitBigInt, ok = new(big.Int).SetString(commitData.CommitVal[2:], 16)
			} else {
				commitBigInt, ok = new(big.Int).SetString(commitData.CommitVal, 10)
			}

			if !ok {
				log.Printf("Failed to convert commit val to big.Int: %s", commitData.CommitVal)
				continue
			}
			commits = append(commits, commitBigInt)
		}
	}

	// Assuming T and NVal are used directly from setupValues for recovery
	omegaRecov, proofListRecovery := crr.Recover(new(big.Int).SetBytes(setupValues.NVal), int(setupValues.T.Int64()), commits)
	if len(proofListRecovery) == 0 {
		return RecoveryResult{}, fmt.Errorf("proofListRecovery is empty")
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

	result := RecoveryResult{
		OmegaRecov: omegaRecov,
		X:          x,
		Y:          y,
		V:          v,
	}

	return result, nil
}

func (l *PoFClient) Recover(ctx context.Context, round *big.Int, y BigNumber) error {
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

	roundStatus.Store(round.String(), "Recovered")

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Recover successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	log.Printf("Recover successful! Tx Hash: %s", signedTx.Hash().Hex())

	return nil
}

func (l *PoFClient) FulfillRandomness(ctx context.Context, round *big.Int) (*types.Transaction, error) {
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

	gasPrice, err := l.Client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice

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

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, l.Client, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		errMsg := fmt.Sprintf("transaction %s reverted", signedTx.Hash().Hex())
		log.Printf("❌ %s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	roundStatus.Store(round.String(), "Fulfilled")

	color.New(color.FgHiGreen, color.Bold).Printf("✅ FulfillRandomness successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	return signedTx, nil
}

func (l *PoFClient) ReRequestRandomWordAtRound(ctx context.Context, round *big.Int) error {
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

	roundStatus.Store(round.String(), "ReRequested")

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Re-request successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	return nil
}

func (l *PoFClient) DisputeRecover(ctx context.Context, round *big.Int, v []BigNumber, x BigNumber, y BigNumber) (*types.Transaction, error) {
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

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, l.Client, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction to be mined: %v", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		errMsg := fmt.Sprintf("transaction %s reverted", signedTx.Hash().Hex())
		log.Printf("❌ %s", errMsg)
		return nil, fmt.Errorf("%s", errMsg)
	}

	roundStatus.Store(round.String(), "DisputeRecovered")

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Dispute recover successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())

	return signedTx, nil
}

func (l *PoFClient) DisputeLeadershipAtRound(ctx context.Context, round *big.Int) error {
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

	roundStatus.Store(round.String(), "DisputeLeadershiped")

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Dispute leadership successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	return nil
}

func (l *PoFClient) Commit(ctx context.Context, round *big.Int) (common.Address, []byte, error) {
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

	roundStatus.Store(round.String(), "Committed")

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Commit successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	fmt.Println("---------------------------------------------------------------------------")

	return auth.From, byteData, nil
}

func (l *PoFClient) OperatorDeposit(ctx context.Context) (common.Address, *types.Transaction, error) {
	style := color.New(color.FgHiBlue, color.Bold)
	style.Println("Preparing to deposit...")

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

	// Set the amount of Ether you want to send in the transaction
	amount := new(big.Int)
	amount.SetString("5000000000000000", 10) // 0.005 ether in wei
	auth.Value = amount                      // Setting the value of the transaction to 0.005 ether

	packedData, err := l.ContractABI.Pack("operatorDeposit")
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to pack data for deposit: %v", err)
	}

	tx := types.NewTransaction(auth.Nonce.Uint64(), l.ContractAddress, amount, 3000000, auth.GasPrice, packedData)
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

	color.New(color.FgHiGreen, color.Bold).Printf("✅  Deposit successful!!\n🔗 Tx Hash: %s\n", signedTx.Hash().Hex())
	fmt.Println("---------------------------------------------------------------------------")

	return auth.From, signedTx, nil // Return the sender address and the transaction
}

func saveRecoveryDataToFile(data RecoveryResult, round string) error {
	// Include the relative path to the data directory
	filename := fmt.Sprintf("../data/recoveryData_%s.json", round)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Ensure the directory exists before writing
	if err := os.MkdirAll("../data", 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(filename, jsonData, 0644)
}

func loadRecoveryDataFromFile(round string) (RecoveryResult, bool, error) {
	filename := fmt.Sprintf("../data/recoveryData_%s.json", round)
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return RecoveryResult{}, false, err
	}
	var data RecoveryResult
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return RecoveryResult{}, false, err
	}
	return data, true, nil
}
