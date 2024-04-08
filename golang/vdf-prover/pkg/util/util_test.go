package util

import (
	"math/big"
	"testing"
)

func TestCalExp(t *testing.T) {
	T := 32
	N := big.NewInt(161603)
	g := big.NewInt(64)
	g.Mod(g, N)

	result, expList := CalExp(N, g, T)

	for i, val := range expList {
		t.Logf("expList[%d]: %s\n", i, val.String())
	}

	t.Logf("Final result: %s\n", result.String())

	if len(expList) != T+1 {
		t.Errorf("Expected expList of length %d, got %d", T+1, len(expList))
	}
}
