#TODO
# Define the binary name
BINARY_NAME=auth

# Define the local bin directory for project-specific tools
LOCAL_BIN_DIR := $(CURDIR)/bin

# Path to the goose binary within the local bin directory
GOOSE_BIN := $(LOCAL_BIN_DIR)/goose

install-tools: ## Installs Go tools required by the project into a local bin directory.
		@echo "${YELLOW}Installing project tools to $(LOCAL_BIN_DIR)...${RESET}"
		@mkdir -p $(LOCAL_BIN_DIR)
		@echo "${GREEN}Tools installed.${RESET}"

# Build the project
build:
		go build -o bin/${BINARY_NAME} ./cmd/api

run-%:
		go run $*/$(MODULE)/main.go

# Clean build artifacts
clean:
		go clean
		rm -f ${BINARY_NAME}

# Install dependencies
deps:
		go mod download

# Lint code (requires golangci-lint)
lint:
		golangci-lint run

.PHONY: install-tools check-goose-installed api run clean deps lint% 