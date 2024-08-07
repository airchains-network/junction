# Makefile for building and linting Go project

# Basic project settings
BINARY_NAME=junctiond
BUILD_DIR=./build
SOURCE_DIR=./cmd/junctiond

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

# Build the project
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
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

# Lint the project
lint:
	@echo "Running linter..."
	$(GOLINT) run $(LINT_OPTIONS)

.PHONY: default build install test clean lint
