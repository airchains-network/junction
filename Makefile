# Makefile for building and linting Go project

# Basic project settings
BINARY_NAME=junctiond
BUILD_DIR=./build
SOURCE_DIR=./cmd/junctiond
CHECKSUM_FILE=$(BUILD_DIR)/checksums.txt

# Go commands
GO_BUILD=go build
GO_INSTALL=go install
GO_CLEAN=go clean
GO_TEST=go test

# Versioning
VERSION ?= $(shell git describe --tags `git rev-list --tags --max-count=1`)
COMMIT := $(shell git log -1 --format='%H')
TMVERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')

# Build tags
build_tags = netgo
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

# Linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=junction \
          -X github.com/cosmos/cosmos-sdk/version.AppName=$(BINARY_NAME) \
          -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
          -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
          -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags)" \
          -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TMVERSION)

# Linter settings
GOLINT=golangci-lint
LINT_OPTIONS=--enable-all
ifeq ($(VERSION),1.0.0)
  LINT_OPTIONS+=--disable=errcheck
endif

# Default target
default: build

# Build the project for the current architecture
build: go.sum
	@echo "Building $(BINARY_NAME) binary..."
	$(GO_BUILD) -tags "$(build_tags)" -ldflags '$(ldflags)' -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

# Install the binary
install:
	@echo "Installing $(BINARY_NAME) binary..."
	$(GO_INSTALL) -tags "$(build_tags)" -ldflags '$(ldflags)' $(SOURCE_DIR)

# Run tests
test:
	@echo "Running tests..."
	$(GO_TEST) ./...

# Clean build artifacts
clean:
	@echo "Cleaning up old versions..."
	$(GO_CLEAN)
	rm -f $(BUILD_DIR)/$(BINARY_NAME) $(CHECKSUM_FILE)

# Lint the project
lint:
	@echo "Running linter..."
	$(GOLINT) run $(LINT_OPTIONS)

# Print system information
print-system:
	@echo "System architecture: $(shell uname -s)/$(shell uname -m)"

# Build the project for different architectures and generate checksums
build-all: go.sum
	@echo "Building $(BINARY_NAME) for different architectures..."
	rm -f $(CHECKSUM_FILE)
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -tags "$(build_tags)" -ldflags '$(ldflags)' -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(SOURCE_DIR)
	@shasum -a 256 $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 >> $(CHECKSUM_FILE)
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -tags "$(build_tags)" -ldflags '$(ldflags)' -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(SOURCE_DIR)
	@shasum -a 256 $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 >> $(CHECKSUM_FILE)
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -tags "$(build_tags)" -ldflags '$(ldflags)' -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(SOURCE_DIR)
	@shasum -a 256 $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 >> $(CHECKSUM_FILE)
	GOOS=darwin GOARCH=arm64 $(GO_BUILD) -tags "$(build_tags)" -ldflags '$(ldflags)' -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(SOURCE_DIR)
	@shasum -a 256 $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 >> $(CHECKSUM_FILE)
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -tags "$(build_tags)" -ldflags '$(ldflags)' -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(SOURCE_DIR)
	@shasum -a 256 $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe >> $(CHECKSUM_FILE)

.PHONY: default build install test clean lint print-system build-all
