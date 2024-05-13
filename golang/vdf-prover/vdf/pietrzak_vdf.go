package vdf

import (
	"fmt"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/pkg/util"
	"math/big"
)

type Claim struct {
	N *big.Int
	X *big.Int
	Y *big.Int
	T int
	V *big.Int
}

func CalV(N, x *big.Int, T int) *big.Int {
	two := big.NewInt(2)
	exp1 := big.NewInt(int64(T))

	// 2^T
	exp2 := new(big.Int).Exp(two, exp1, nil)

	// v ← x^(2^T)
	v := new(big.Int).Exp(x, exp2, N)

	return v
}

func RecHalveProof(claim Claim) []Claim {
	proofList := []Claim{claim}
	for claim.T > 1 {
		claim = HalveProof(claim)
		proofList = append(proofList, claim)
	}

	return proofList
}

func RecHalveProofWithDelta(claim Claim) []Claim {
	proofList := []Claim{claim}
	delta := 9
	deltaPower := 1 << delta
	for claim.T > deltaPower*2 {
		claim = HalveProof(claim)
		proofList = append(proofList, claim)
	}

	return proofList
}

func HalveProof(claim Claim) Claim {
	xStr, yStr, vStr := fmt.Sprintf("%x", claim.X), fmt.Sprintf("%x", claim.Y), fmt.Sprintf("%x", claim.V)

	// T ← T/2
	tHalf := util.CalTHalf(claim.T)

	// ri ← H(xi + yi + µi)
	r := util.HashEth128(xStr, yStr, vStr)

	// x(i+1) ← xi^ri * µi mod N
	xPrime := new(big.Int).Exp(claim.X, r, claim.N)
	xPrime.Mul(xPrime, claim.V).Mod(xPrime, claim.N)

	// y(i+1) ← µi^ri * yi mod N
	yPrime := new(big.Int).Exp(claim.V, r, claim.N)
	yPrime.Mul(yPrime, claim.Y).Mod(yPrime, claim.N)

	// µi ← xi^2^(T)
	vPrime := CalV(claim.N, xPrime, tHalf/2)

	claim.X = xPrime
	claim.Y = yPrime
	claim.T = tHalf
	claim.V = vPrime

	return claim
}

func VerifyProof(proofList []Claim) bool {
	// µ >= N → return False
	for _, claim := range proofList {
		if claim.V.Cmp(claim.N) >= 0 {
			fmt.Println("Verification failed: V is greater than or equal to N.")
			return false
		}
	}

	// y(t+1) = x(t+1)^2 return True
	N := proofList[0].N
	x := new(big.Int).Set(proofList[0].X)
	y := new(big.Int).Set(proofList[0].Y)
	v := new(big.Int).Set(proofList[0].V)
	T := proofList[0].T

	for i := 0; i < len(proofList); i++ {
		xStr, yStr, vStr := fmt.Sprintf("%x", x), fmt.Sprintf("%x", y), fmt.Sprintf("%x", v)

		if T > 1 {
			// ri ← H(xi + yi + µi)
			r := util.HashEth128(xStr, yStr, vStr)

			tHalf := util.CalTHalf(T)
			T = tHalf

			// x(i+1) ← xi^ri * µi mod N
			xPrime := new(big.Int).Exp(x, r, N)
			xPrime.Mul(xPrime, v).Mod(xPrime, N)

			// y(i+1) ← µi^ri * yi mod N
			yPrime := new(big.Int).Exp(v, r, N)
			yPrime.Mul(yPrime, y).Mod(yPrime, N)

			vPrime := CalV(N, xPrime, T/2)

			x, y, v = xPrime, yPrime, vPrime
		}
	}

	xSquared := new(big.Int).Mul(x, x)
	xSquared.Mod(xSquared, N)

	// y(t+1) = x(t+1)^2
	if xSquared.Cmp(y) != 0 {
		fmt.Println("Verification failed: y(t+1) does not equal x(t+1)^2.")
		return false
	}

	return true
}
