package node

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config := LoadConfig()

	if config.RpcURL == "" {
		t.Error("Failed to load RpcURL from config.json")
	}
	if config.PrivateKey == "" {
		t.Error("Failed to load PrivateKey from config.json")
	}

	fmt.Println("Loaded configuration:")
	fmt.Printf("RPC URL: %s\n", config.RpcURL)
	fmt.Printf("Private Key: %s\n", config.PrivateKey)
	fmt.Printf("Contract Address: %s\n", config.ContractAddress)
}
