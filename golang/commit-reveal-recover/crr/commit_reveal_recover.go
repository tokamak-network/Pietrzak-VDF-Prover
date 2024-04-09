package crr

import (
	"crypto/rand"
	"fmt"
	crrUtil "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/commit-reveal-recover/pkg/util"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/pkg/util"
	prover "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/vdf"
	"math/big"
	"time"
)

func Setup(N *big.Int, x *big.Int, T int) (*big.Int, []prover.Claim) {
	setupEvalTime := time.Now()
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
	setupEvalDuration := time.Since(setupEvalTime)
	fmt.Println("Setup Eval Time: ", setupEvalDuration)

	setupProofTime := time.Now()
	setupProofList := prover.RecHalveProof(claim)
	setupProofDuration := time.Since(setupProofTime)
	fmt.Println("Setup Proof Time: ", setupProofDuration)

	return y, setupProofList
}

func Commit(N *big.Int, x *big.Int, member int) ([]*big.Int, []*big.Int, *big.Int) {
	a := make([]*big.Int, member)
	c := make([]*big.Int, member)

	fmt.Println("[+] Number of participants: ", member, "\n")

	for i := range a {
		a[i], _ = rand.Int(rand.Reader, N)
		fmt.Printf("a_%d is generated as %s\n", i, a[i].String())
	}

	for i := range c {
		c[i] = new(big.Int).Exp(x, a[i], N)
		fmt.Printf("c_%d is generated as %s\n", i, c[i].String())
	}

	cStrs := make([]string, len(c))
	for i, ci := range c {
		cStrs[i] = ci.String()
	}

	bStar := util.HashEth128(cStrs...)

	return a, c, bStar
}

func Reveal(N *big.Int, y *big.Int, a []*big.Int, c []*big.Int, bStar *big.Int) *big.Int {
	omega := big.NewInt(1)

	for i := range a {
		cHex := fmt.Sprintf("%x", c[i])
		bStarHex := fmt.Sprintf("%x", bStar)

		hashValue := util.HashEth(cHex, bStarHex)

		temp := new(big.Int).Exp(y, hashValue, N)
		omega.Mul(omega, temp).Mod(omega, N)
	}

	return omega
}

func calExpRecov(N, x *big.Int, T int) *big.Int {
	expList := make([]*big.Int, T+1)
	expList[0] = new(big.Int).Set(x)
	result := new(big.Int).Set(x)

	for i := 1; i <= T; i++ {
		result.Mul(result, result).Mod(result, N)
		expList[i] = new(big.Int).Set(result)
	}

	return result
}

func Recover(N *big.Int, T int, c []*big.Int, bStar *big.Int) (*big.Int, []prover.Claim) {
	RecovEvalTime := time.Now()
	var cHexStrings []string
	for _, ci := range c {
		ciHex := fmt.Sprintf("%x", ci)
		cHexStrings = append(cHexStrings, ciHex)
	}

	if bStar == nil {
		bStar = util.HashEth(cHexStrings...)
	}

	recov := big.NewInt(1)
	for _, ciHex := range cHexStrings {
		bStarHex := fmt.Sprintf("%x", bStar)
		hashValue := util.HashEth(ciHex, bStarHex)

		ci, _ := new(big.Int).SetString(ciHex, 16)

		temp := new(big.Int).Exp(ci, hashValue, N)
		recov.Mul(recov, temp).Mod(recov, N)
	}

	omegaRecov := calExpRecov(N, recov, T)
	claim := crrUtil.ConstructClaim(N, recov, omegaRecov, T)
	RecovEvalDuration := time.Since(RecovEvalTime)
	fmt.Println("Recover Eval Time: ", RecovEvalDuration)

	RecovProofTime := time.Now()
	proofListRecovery := prover.RecHalveProof(claim)
	RecovProofDuration := time.Since(RecovProofTime)
	fmt.Println("Recover Proof Time: ", RecovProofDuration)

	return omegaRecov, proofListRecovery
}
