package crr

import (
	"crypto/rand"
	"fmt"
	crrUtil "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/commit-reveal-recover/pkg/util"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/pkg/util"
	prover "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-prover/vdf"
	"math/big"
)

func Setup(N *big.Int, x *big.Int, T int) (*big.Int, []prover.Claim) {
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
	setupProofList := prover.RecHalveProof(claim)
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

func Recover(N *big.Int, T int, c []*big.Int, bStar *big.Int) (*big.Int, []prover.Claim) {
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

	fmt.Println("Revealed Random: ", recov, "\n")

	omegaRecov, expListRecov := util.CalExp(N, recov, T)
	fmt.Println("omegaRecov: ", omegaRecov, "\n")
	fmt.Println("expListRecov: ", expListRecov, "\n")
	claim := crrUtil.ConstructClaim(N, recov, omegaRecov, T)

	fmt.Println("claim: ", claim, "\n")

	proofListRecovery := prover.RecHalveProof(claim)

	return omegaRecov, proofListRecovery
}
