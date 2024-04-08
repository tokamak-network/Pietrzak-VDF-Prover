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

//// evalV computes the value of v based on the input parameters.
//func CalV(N, x *big.Int, T *big.Int) *big.Int {
//	i := 1
//	sumPart := big.NewInt(0)
//	productPart := big.NewInt(1)
//
//	two := big.NewInt(2)
//
//	// 2^i
//	exp1 := new(big.Int).Exp(two, big.NewInt(int64(i)), nil)
//
//	// 2^(i-1)
//	exp2 := new(big.Int).Exp(two, big.NewInt(int64(i-1)), nil)
//
//	for j := 0; j < 1<<(i-1); j++ {
//
//		}
//
//		// T / 2^i + j * T / 2^(i-1)
//		temp := new(big.Int).Set(T)
//		temp.Div(temp, exp1) // T / 2^i
//		temp2 := new(big.Int).Mul(big.NewInt(int64(j)), T)
//		temp2.Div(temp2, exp2) // j * T / 2^(i-1)
//		temp.Add(temp, temp2)  // Add them together
//
//		// 2 to the power of the above result
//		tmp := new(big.Int).Exp(two, temp, nil)
//
//		// productPart * 2^...
//		tmp.Mul(productPart, tmp)
//
//		// Add to sumPart
//		sumPart.Add(sumPart, tmp)
//	}
//
//	// Perform the final operation using the calculated sum_part
//	v := new(big.Int).Exp(x, sumPart, N)
//
//	return v
//}
