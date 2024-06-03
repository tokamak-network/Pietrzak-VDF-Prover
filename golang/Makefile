# Define the default target
.PHONY: run config install-go install-go-dependency install-dependency

# Define the target to run the main.go file
run:
	@cd vdf-node/cmd && go run main.go

# Define the target to edit the config.json file
config:
	@cd vdf-node && vi config.json

install-go:
	@echo "Checking if Go is installed and working..."
	@if go version; then \
		echo "Go is already installed, skipping installation..."; \
	else \
		echo "Installing Go..."; \
		wget -P ~/ https://go.dev/dl/go1.22.3.linux-amd64.tar.gz; \
		sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf ~/go1.22.3.linux-amd64.tar.gz; \
		export PATH=$$PATH:/usr/local/go/bin; \
		echo "New version of Go installed. Checking version..."; \
		go version; \
	fi

install-go-dependency:
	@echo "Running go mod tidy"
	@go mod tidy

install-dependency: install-go install-go-dependency
	@echo "All dependencies have been installed."