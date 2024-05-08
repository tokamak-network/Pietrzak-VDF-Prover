package main

import (
	"fmt"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/pkg/util"
	prover "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/vdf"
	"math/big"
	"time"
)

func setup(bits int) (*big.Int, *big.Int) {
	N, err := util.GeneratePrime(bits)
	if err != nil {
		fmt.Println("Error generating prime:", err)
	}
	x := big.NewInt(243)
	x.Mod(x, N) // Ensure x < N
	return N, x
}

func evaluate(N, x *big.Int, T int) ([]prover.Claim, prover.Claim) {
	evalStart := time.Now()
	y := util.CalExp(N, x, T)
	evalExecTime := time.Since(evalStart)
	fmt.Printf("Evaluation Time: %s\n", evalExecTime)
	tHalf := util.CalTHalf(T)
	//v := util.GetExp(expList, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(tHalf)), nil), N)
	v := prover.CalV(N, x, tHalf)
	claim := prover.Claim{
		N: N,
		X: x,
		Y: y,
		T: T,
		V: v,
	}

	proofList := []prover.Claim{claim}

	return proofList, claim
}

func genProof(claim prover.Claim) []prover.Claim {
	start := time.Now()
	proofList := prover.RecHalveProof(claim)
	execTime := time.Since(start)
	fmt.Printf("Execution time for generating and halving proof: %s\n", execTime)
	return proofList
}

func verify(proofList []prover.Claim) bool {
	startVerify := time.Now()
	isVerify := prover.VerifyProof(proofList)
	verifyTime := time.Since(startVerify)

	fmt.Printf("Execution time for verification: %s\n", verifyTime)
	return isVerify
}

func main() {
	// 2^21 ~ 2^25
	// 2097152, 4194304, 8388608, 16777216, 33554432
	T := 2097152

	// lambda
	bits := 2048

	N, x := setup(bits)
	proofList, claim := evaluate(N, x, T)

	proofList = genProof(claim)

	isVerify := verify(proofList)

	if isVerify {
		fmt.Println("Verification succeeded.")
	} else {
		fmt.Println("Verification failed.")
	}
}
