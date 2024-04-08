package main

import (
	"fmt"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/pkg/util"
	prover "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/vdf"
	"math/big"
	"time"
)

type Claim struct {
	N *big.Int
	X *big.Int
	Y *big.Int
	T int
	V *big.Int
}

func main() {

	// 2^21 ~ 2^25
	// 2097152, 4194304, 8388608, 16777216, 33554432
	T := 2097152
	bits := 512 // 소수를 생성할 비트 수
	N, err := util.GeneratePrime(bits)
	if err != nil {
		fmt.Println("Error generating prime:", err)
		return
	}
	x := big.NewInt(243)
	x.Mod(x, N)

	evalStart := time.Now()
	y, expList := util.CalExp(N, x, T)

	tHalf := util.CalTHalf(T)
	v := util.GetExp(expList, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(tHalf)), nil), N)

	claim := prover.Claim{
		N: N,
		X: x,
		Y: y,
		T: T,
		V: v,
	}

	proofList := []prover.Claim{claim}
	evalExecTime := time.Since(evalStart)
	fmt.Printf("Evaluation Time: %s\n", evalExecTime)
	start := time.Now()
	proofList = prover.RecHalveProof(claim)

	//for i, proof := range proofList {
	//	fmt.Printf("proofList[%d]: (N: %s, X: %s, Y: %s, T: %d, V: %s)\n", i, proof.N, proof.X, proof.Y, proof.T, proof.V)
	//}

	execTime := time.Since(start)
	fmt.Printf("Execution time for generating and halving proof: %s\n", execTime)

	startVerify := time.Now()

	isVerify := prover.VerifyProof(proofList)
	if isVerify {
		fmt.Println("Verification succeeded.")
	} else {
		fmt.Println("Verification failed.")
	}

	verifyTime := time.Since(startVerify)
	fmt.Printf("Execution time for verification: %s\n", verifyTime)
}

func joinStrings(strs []string, sep string) string {
	var result string
	for i, s := range strs {
		if i > 0 {
			result += sep + ""
		}
		result += s
	}
	return result
}
