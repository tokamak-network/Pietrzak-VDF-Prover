package main

import (
	"fmt"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/commit-reveal-recover/crr"
	prover "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/vdf"
	"math/big"
)

func main() {
	// 2^21 ~ 2^25
	// 2097152, 4194304, 8388608, 16777216, 33554432
	//T := 2097152
	//bits := 512
	//N, err := util.GeneratePrime(bits)
	//if err != nil {
	//	fmt.Println("Error generating prime:", err)
	//	return
	//}
	T := 128
	N := big.NewInt(401)
	x := big.NewInt(243)

	// Test XGenerator
	//x := crrUtil.XGenerator(N)
	//fmt.Println("Generated x:", x)

	// Test Setup Phase
	y, setupProofList := crr.Setup(N, x, T)

	fmt.Println("Output (y):", y.String())
	fmt.Println("setupProofList:", setupProofList)

	verified := prover.VerifyProof(setupProofList)
	if verified {
		fmt.Println("Verification successful!")
	} else {
		fmt.Println("Verification failed.")
	}

	// Test Commit Phase
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

	fmt.Println("\nbStar (Commitments Hash):", bStar)

	// Test Reveal Phase
	omega := crr.Reveal(N, y, a, c, bStar)
	fmt.Println("\nFinal Omega: ", omega)

	// Test Recover Phase
	omegaRecov, proofListRecovery := crr.Recover(N, T, c, bStar)
	fmt.Println("\nomegaRecov: ", omegaRecov)
	fmt.Println("\nproofListRecovery: ", proofListRecovery)

	recovVerified := prover.VerifyProof(proofListRecovery)
	if recovVerified {
		fmt.Println("RECOVERY Verification successful!")
	} else {
		fmt.Println("RECOVERY Verification failed.")
	}

}
