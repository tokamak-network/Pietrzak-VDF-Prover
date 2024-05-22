package main

import (
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/node"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/util"
	"log"
	"time"
)

func main() {
	printLogo()
	color.New(color.FgHiCyan, color.Bold).Println("Starting VDF Node...")

	util.StartSpinner("Configuring system...", 5)
	config := node.LoadConfig()
	color.New(color.FgHiGreen, color.Bold).Println("Configuration loaded successfully. Ready to operate.")

	for {
		util.StartSpinner("Initializing listener...", 3)
		listener, err := node.NewListener(config)

		if err != nil {
			log.Fatalf("Listener initialization failed: %v", err)
		}
		color.New(color.FgHiGreen, color.Bold).Println("Listener is now active and ready.")

		// Start a goroutine to handle reconnection
		go handleConnection(listener)

		select {} // Maintain the service running
	}
}

func handleConnection(listener *node.Listener) {
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
		break
	}
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
