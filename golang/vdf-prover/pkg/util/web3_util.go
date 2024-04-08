package util

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

func PadHex(hexStr string) string {
	n := (64 - (len(hexStr) % 64)) % 64

	return strings.Repeat("0", n) + hexStr
}

func HashEth(hexStrings ...string) *big.Int {
	var input string
	for _, hexStr := range hexStrings {
		paddedHex := PadHex(hexStr)
		input += paddedHex
	}

	inputBytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("hex.DecodeString error:", err)
		return big.NewInt(0)
	}

	hashBytes := crypto.Keccak256(inputBytes)
	result := new(big.Int).SetBytes(hashBytes)

	return result
}

func HashEth128(strings ...string) *big.Int {
	hashBigInt := HashEth(strings...)

	return new(big.Int).Rsh(hashBigInt, 128)
}
