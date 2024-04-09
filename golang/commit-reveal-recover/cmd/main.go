package main

import (
	"fmt"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/commit-reveal-recover/crr"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/pkg/util"
	prover "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/vdf"
	"math/big"
	"time"
)

func main() {
	// 2^21 ~ 2^25
	// 2097152, 4194304, 8388608, 16777216, 33554432
	T := 2097152
	bits := 2048
	N, err := util.GeneratePrime(bits)
	if err != nil {
		fmt.Println("Error generating prime:", err)
		return
	}
	//T := 128
	//N := big.NewInt(401)
	x := big.NewInt(243)

	//Test XGenerator
	//x := crrUtil.XGenerator(N)

	// Test Setup Phase
	start2 := time.Now()
	y, setupProofList := crr.Setup(N, x, T)
	duration2 := time.Since(start2)

	fmt.Println("Output (y):", y.String())
	fmt.Println("setupProofList:", setupProofList)
	fmt.Println("Setup Phase: ", duration2)

	verified := prover.VerifyProof(setupProofList)
	if verified {
		fmt.Println("Verification successful!")
	} else {
		fmt.Println("Verification failed.")
	}

	// Test Commit Phase
	start3 := time.Now()
	member := 5
	a, c, bStar := crr.Commit(N, x, member)
	fmt.Println("\na values (Secrets):")
	for i, val := range a {
		fmt.Printf("a_%d: %s\n", i+1, val.String())
	}

	fmt.Println("\nc values (Commitments):")
	for i, val := range c {
		fmt.Printf("c_%d: %s\n", i+1, val.String())
	}
	duration3 := time.Since(start3)
	fmt.Println("Commit Phase: ", duration3)
	fmt.Println("\nbStar (Commitments Hash):", bStar)

	// Test Reveal Phase
	start4 := time.Now()
	omega := crr.Reveal(N, y, a, c, bStar)
	duration4 := time.Since(start4)
	fmt.Println("Reveal Phase: ", duration4)
	fmt.Println("\nFinal Omega: ", omega)

	// Test Recover Phase
	start5 := time.Now()
	omegaRecov, proofListRecovery := crr.Recover(N, T, c, bStar)
	fmt.Println("\nomegaRecov: ", omegaRecov)
	//fmt.Println("\nproofListRecovery: ", proofListRecovery)

	recovVerified := prover.VerifyProof(proofListRecovery)
	if recovVerified {
		fmt.Println("RECOVERY Verification successful!")
	} else {
		fmt.Println("RECOVERY Verification failed.")
	}
	duration5 := time.Since(start5)
	fmt.Println("Recover Phase: ", duration5)

}
