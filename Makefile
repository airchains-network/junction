# Makefile for building and linting Go project
DOCKER := $(shell which docker)
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
GO_VERSION=1.23
COMMIT := $(shell git log -1 --format='%H')

LEDGER_ENABLED ?= true

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))
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
	@shasum -a 256 $(BUILD_DIR)/$(BINARY_NAME) >> $(CHECKSUM_FILE)

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

build-static-linux-amd64: go.sum $(BUILD_DIR)/
	mkdir -p $(BUILD_DIR)
	$(DOCKER) buildx create --name junctionbuilder || true
	$(DOCKER) buildx use junctionbuilder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
		--build-arg BUILD_TAGS=$(build_tags_comma_sep),muslc \
		--platform linux/amd64 \
		-t junction-amd64 \
		--load \
		-f Dockerfile.builder .
	$(DOCKER) rm -f junctionbinary || true
	$(DOCKER) create -ti --name junctionbinary junction-amd64
	$(DOCKER) cp junctionbinary:/bin/junctiond $(BUILD_DIR)/junctiond-linux-amd64
	$(DOCKER) rm -f junctionbinary

.PHONY: default build install test clean lint print-system build-all build-static-linux-amd64
