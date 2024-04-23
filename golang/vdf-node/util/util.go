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
	totalIterations := duration * 4 // 한 바퀴에 4번의 변화
	maxDots := 10                   // 점을 최대 10개까지 추가

	style := color.New(color.FgYellow)   // 스피너 색상 설정
	clearLine := strings.Repeat(" ", 50) // 충분한 공간 확보를 위한 공백 문자

	fmt.Print("\n") // 스피너 시작 전 개행
	for i := 0; i < totalIterations; i++ {
		dots := strings.Repeat(".", i%maxDots+1) // 점을 순환하며 추가
		style.Printf("\r%s %s%s%s", spinner[i%len(spinner)], message, dots, clearLine)
		time.Sleep(100 * time.Millisecond) // 스피너 속도 조절
	}
	style.Printf("\r%s\r", strings.Repeat(" ", len(message)+maxDots+50)) // 마지막 클리어를 확실히 하기 위한 공백 출력
}

func DynamicLoadingIndicator() {
	maxDots := 5
	for {
		for dots := 1; dots <= maxDots; dots++ {
			fmt.Printf("\r%s Listening%s", "🎧", strings.Repeat(".", dots))
			time.Sleep(500 * time.Millisecond) // 0.5초 간격으로 점의 개수를 변경
		}
	}
}
