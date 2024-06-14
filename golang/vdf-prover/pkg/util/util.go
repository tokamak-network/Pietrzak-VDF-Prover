package util

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/sha3"
	"math/big"
	"time"
)

func ModHash(strings []string, n int64) *big.Int {
	input := ""

	for _, s := range strings {
		input += s
	}

	hasher := sha3.New256()
	hasher.Write([]byte(input))
	hashed := hasher.Sum(nil)
	hashInt := big.NewInt(0).SetBytes(hashed)
	nBig := big.NewInt(n)
	r := big.NewInt(0).Mod(hashInt, nBig)

	return r
}

func CalExp(N, x *big.Int, T int) *big.Int {
	startTime := time.Now()

	// Use Lsh directly on a new big.Int initialized to 1
	expT := new(big.Int).Lsh(big.NewInt(1), uint(T))

	// Now calculate x^(2^T) mod N
	result := new(big.Int).Exp(x, expT, N)

	fmt.Printf("CalExp run time: %s\n", time.Since(startTime))
	return result
}

func GetExp(expList []*big.Int, exp, N *big.Int) *big.Int {
	startTime := time.Now()

	res := big.NewInt(1)
	i := 0
	bigExp := new(big.Int).Set(exp)

	for bigExp.Sign() > 0 {
		if bigExp.Bit(0) == 1 {
			res.Mul(res, expList[i])
			res.Mod(res, N)
		}
		bigExp.Rsh(bigExp, 1)
		i++
	}

	fmt.Printf("GetExp run time: %s\n", time.Since(startTime))
	return res
}

func CalTHalf(T int) int {
	var tHalf int
	if T%2 == 0 {
		tHalf = T / 2
	} else {
		tHalf = (T + 1) / 2
	}

	return tHalf
}

func GeneratePrime(bits int) (*big.Int, error) {

	prime, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	return prime, nil
}
