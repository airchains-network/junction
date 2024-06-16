#!/bin/bash

# Define the binary name
BINARY="junctiond"
BUILD_DIR="build"

# Detect the platform
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

# Create a build directory
mkdir -p $BUILD_DIR

# Set output binary name
OUTPUT="$BUILD_DIR/$BINARY-$GOOS-$GOARCH"

# Build the binary
echo "Building $BINARY for $GOOS/$GOARCH..."
go build -o $OUTPUT ./cmd/junctiond/main.go

if [ $? -ne 0 ]; then
    echo "Failed to build $BINARY for $GOOS/$GOARCH"
    exit 1
fi

echo "Build completed successfully! Binary is located at $OUTPUT"