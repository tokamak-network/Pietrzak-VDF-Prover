# VDF Node

```text
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
```


---
# Table of Contents

- [Introduction](#introduction)
- [How to Install](#how-to-install)
- [Getting Started](#how-to-use)
- [Configuration](#configuration)
- [Running the VDF Node](#running-the-vdf-node)

## Introduction

This README file provides instructions on how to install and use the Pietrzak VDF Node. Follow the steps carefully to ensure proper setup and execution of the VDF Node.

## How to Install

1. **Clone the Repository**

    ```bash
    git clone <repository-url>
    ```

2. **Checkout the Specific Branch**

    ```bash
    git checkout vdf-prover
    ```
   If the branch is not up-to-date, pull the latest changes:

    ```bash
    git pull origin vdf-prover
    ```

3. **Install Go**

   Ensure you have Go version 1.22.3 or higher installed. You can download and install Go from [the official Go website](https://golang.org/dl/).

4. **Install pm2 Package**

    ```bash
    npm install pm2 -g
    ```

5. **Configure the JSON Settings**

   Create and configure `config.json` in the `Pietrzak-VDF-Prover/golang/vdf-node` directory:

    ```json
    {
      "RpcURL": "wss://YourWebsocket.url",
      "HttpURL": "https://YourRPC.url",
      "PrivateKey": "0xYourPrivateKey",
      "ContractAddress": "0xYourContractAddress",
      "WalletAddress": "0xYourWalletAddress",
      "SubgraphURL": "https://YourSubgraph.url"
    }
    ```

## Getting Started

1. **Build the Project**

   Navigate to the `Pietrzak-VDF-Prover/golang/vdf-node/cmd` directory and build the project:

    ```bash
    cd Pietrzak-VDF-Prover/golang/vdf-node/cmd
    go build -o /usr/local/bin/vdfnode main_subgraph.go
    ```

2. **Start the VDF Node Using pm2**

    ```bash
    pm2 start /usr/local/bin/vdfnode --name VDFNode
    ```

3. **Check the Logs**

   To view the logs:

    ```bash
    pm2 log VDFNode
    ```

## Configuration

Ensure your `config.json` file contains the correct information as specified:

- **RpcURL**: The WebSocket URL for your RPC.
- **HttpURL**: The HTTP URL for your RPC.
- **PrivateKey**: Your private key in hexadecimal format.
- **ContractAddress**: The address of your contract.
- **WalletAddress**: Your wallet address.
- **SubgraphURL**: The URL for your Subgraph.

## Running the VDF Node

1. Build the Go application in the specified directory.
2. Use pm2 to manage and monitor the VDF Node process.
3. Always check the logs to ensure the node is running correctly and troubleshoot any issues that arise.

Following these steps will help you set up and run the Pietrzak VDF Node successfully. If you encounter any issues, refer to the logs for more information and troubleshooting tips.