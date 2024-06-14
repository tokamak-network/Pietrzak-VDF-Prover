package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"math/big"
	"strings"
	"time"
)

type EventInfo struct {
	Round  *big.Int
	Sender common.Address
}

func StartSpinner(message string, duration int) {
	spinner := []string{"|", "/", "-", "\\"}
	totalIterations := duration * 4
	maxDots := 10

	style := color.New(color.FgYellow)
	clearLine := strings.Repeat(" ", 50)

	fmt.Print("\n")
	for i := 0; i < totalIterations; i++ {
		dots := strings.Repeat(".", i%maxDots+1)
		style.Printf("\r%s %s%s%s", spinner[i%len(spinner)], message, dots, clearLine)
		time.Sleep(100 * time.Millisecond)
	}
	style.Printf("\r%s\r", strings.Repeat(" ", len(message)+maxDots+50))
}

func DynamicLoadingIndicator() {
	maxDots := 5
	for {
		for dots := 1; dots <= maxDots; dots++ {
			fmt.Printf("\r%s Listening%s", "🎧", strings.Repeat(".", dots))
			time.Sleep(500 * time.Millisecond)
		}
	}
}
