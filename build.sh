#!/bin/bash

# Define the binary name
BINARY="junctiond"
BUILD_DIR="build"

# Define version information
VERSION="v0.5.3"
COMMIT=$(git rev-parse --short HEAD)
BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')

# Detect the platform
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)

# Clearing the build directory
echo "Deleting BUILD_DIR this may take more than 5sec..."
sudo rm -rf $BUILD_DIR
#sleep 5  # Delay for 5 seconds

# Create a build directory
mkdir -p $BUILD_DIR

# Set output binary name
OUTPUT="$BUILD_DIR/$BINARY-$GOOS-$GOARCH/junctiond"

# Build the binary with version information
echo "Building $BINARY for $GOOS/$GOARCH..."
go build -ldflags "-X 'main.Version=$VERSION' -X 'main.Commit=$COMMIT' -X 'main.BuildDate=$BUILD_DATE'" -o $OUTPUT ./cmd/junctiond/main.go

if [ $? -ne 0 ]; then
    echo "Failed to build $BINARY for $GOOS/$GOARCH"
    exit 1
fi

echo "Build completed successfully! Binary is located at $OUTPUT"
echo "Moved binary to the usr/local/bin"
sudo mv ./build/$BINARY-$GOOS-$GOARCH/junctiond /usr/local/bin/