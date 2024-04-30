package main

import (
	"github.com/fatih/color"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/node"
	"github.com/tokamak-network/Pietrzak-VDF-Prover/golang/vdf-node/util"
	"log"
)

func main() {
	printLogo()
	color.New(color.FgHiCyan, color.Bold).Println("Starting VDF Node...")

	util.StartSpinner("Configuring system...", 5)
	config := node.LoadConfig()
	color.New(color.FgHiGreen, color.Bold).Println("Configuration loaded successfully. Ready to operate.")

	util.StartSpinner("Initializing listener...", 3)
	listener, err := node.NewListener(config)
	if err != nil {
		log.Fatalf("Listener initialization failed: %v", err)
	}
	color.New(color.FgHiGreen, color.Bold).Println("Listener is now active and ready.")

	listener.SubscribeRandomWordsRequested()
	select {} // Maintain the service running

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
