package node

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	RpcURL          string `json:"RpcURL"`
	HttpURL         string `json:"HttpURL"`
	PrivateKey      string `json:"PrivateKey"`
	ContractAddress string `json:"ContractAddress"`
	WalletAddress   string `json:"WalletAddress"`
}

func LoadConfig() Config {
	var config Config
	configFile, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err)
	}
	return config
}

func LoadConfigPoF() Config {
	var config Config
	configFile, err := ioutil.ReadFile("../config-pof.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err)
	}
	return config
}

func LoadConfigTest() Config {
	var config Config
	configFile, err := ioutil.ReadFile("../config-test.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err)
	}
	return config
}
