package main

import (
	"flag"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/node"
	nodePoF "github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/node-pof"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/util"
	"log"
	"os"
	"time"
)

var timeoutTimer *time.Timer // 타임아웃 타이머 전역 변수

func main() {
	printLogo()

	pofMode := flag.Bool("pof", false, "Activate Proof of Flow mode")
	testMode := flag.Bool("test", false, "Activate Test mode")
	flag.Parse()

	if *pofMode {
		color.New(color.FgHiCyan, color.Bold).Println("Starting VDF Node in PoF mode...")
	}

	if *testMode {
		color.New(color.FgHiCyan, color.Bold).Println("Starting VDF Node in Test mode...")
	}

	util.StartSpinner("Configuring system...", 5)

	config := node.LoadConfig()
	color.New(color.FgHiGreen, color.Bold).Println("Configuration loaded successfully. Ready to operate.")

	var listener node.ListenerInterface
	var err error

	if *pofMode {
		listener, err = nodePoF.NewPoFListener(config)
	} else {
		log.Fatal("No mode selected, shutting down.")
	}

	if err != nil {
		log.Fatalf("Listener initialization failed: %v", err)
	}

	if listener == nil {
		log.Fatal("Listener is not initialized.")
	}

	color.New(color.FgHiGreen, color.Bold).Println("Listener is now active and ready.")
	go handleConnection(listener)

	resetTimeout(5 * time.Minute)

	select {}
}

func handleConnection(listener nodePoF.PoFListenerInterface) {
	if err := listener.CheckRoundCondition(); err != nil {
		log.Printf("Error during CheckRoundCondition: %v", err)
	}

	for {
		err := listener.SubscribeRandomWordsRequested()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
				color.New(color.FgHiRed, color.Bold).Println("WebSocket connection lost. Reconnecting in 5 seconds...")
				time.Sleep(5 * time.Second)
				continue
			} else {
				log.Fatalf("Listener failed: %v", err)
			}
		}
		resetTimeout(5 * time.Minute)
	}
}

func resetTimeout(duration time.Duration) {
	if timeoutTimer != nil {
		timeoutTimer.Stop()
	}
	timeoutTimer = time.AfterFunc(duration, func() {
		log.Println("No requests for 5 minutes, shutting down the application...")
		os.Exit(0)
	})
}

func printLogo() {
	color.Cyan(`
                  _____          ___                  ___           ___          _____          ___     
      ___        /  /::\        /  /\                /__/\         /  /\        /  /::\        /  /\    
     /__/\      /  /:/\:\      /  /:/_               \  \:\       /  /::\      /  /:/\:\      /  /:/_   
     \  \:\    /  /:/  \:\    /  /:/ /\               \  \:\     /  /:/\:\    /  /:/  \:\    /  /:/ /\  
      \  \:\  /__/:/ \__\:|  /  /:/ /:/           _____\__\:\   /  /:/  \:\  /__/:/ \__\:|  /  /:/ /:/_ 
  ___  \__\:\ \  \:\ /  /:/ /__/:/ /:/           /__/::::::::\ /__/:/ \__\:\ \  \:\ /  /:/ /__/:/ /:/ /\
 /__/\ |  |:|  \  \:\  /:/  \  \:\/:/            \  \:\~~\~~\/ \  \:\ /  /:/  \  \:\  /:/  \  \:\/:/ /:/
 \  \:\|  |:|   \  \:\/:/    \  \::/              \  \:\  ~~~   \  \:\  /:/    \  \:\/:/    \  \::/ /:/ 
  \  \:\__|:|    \  \::/      \  \:\               \  \:\        \  \:\/:/      \  \::/      \  \:\/:/  
   \__\::::/      \__\/        \  \:\               \  \:\        \  \::/        \__\/        \  \::/   
                                \__\/                \__\/         \__\/                       \__\/    
	`)
}
