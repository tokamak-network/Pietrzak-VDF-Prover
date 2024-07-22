package main

import (
	"fmt"
	"math/big"
	"sync"
)

func isPrime(n *big.Int, k int) bool {
	return n.ProbablyPrime(k)
}

func findPrimes(start, end *big.Int, wg *sync.WaitGroup, primesChan chan<- *big.Int) {
	defer wg.Done()
	prime := new(big.Int).Set(start)
	if prime.Bit(0) == 0 {
		prime.Add(prime, big.NewInt(1))
	}

	for prime.Cmp(end) < 0 {
		if isPrime(prime, 20) {
			primesChan <- new(big.Int).Set(prime)
		}
		prime.Add(prime, big.NewInt(2))
	}
}

func main() {
	start := new(big.Int).Exp(big.NewInt(2), big.NewInt(255), nil)
	end := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))

	numWorkers := 10
	var wg sync.WaitGroup
	primesChan := make(chan *big.Int, 100)
	step := new(big.Int).Div(new(big.Int).Sub(end, start), big.NewInt(int64(numWorkers)))

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startRange := new(big.Int).Add(start, new(big.Int).Mul(step, big.NewInt(int64(i))))
		endRange := new(big.Int).Add(startRange, step)
		if i == numWorkers-1 {
			endRange = end
		}
		go findPrimes(startRange, endRange, &wg, primesChan)
	}

	go func() {
		wg.Wait()
		close(primesChan)
	}()

	for prime := range primesChan {
		fmt.Println("Found prime:", prime)
	}
}
