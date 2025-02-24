#!/usr/bin/env bash

# Function to get the module path using `go list`
get_module_path() {
    go list -m -f '{{.Dir}}' "$1" 2>/dev/null
}

# Function to find .proto files in a given directory
find_proto_files() {
    local dir="$1"
    find "$dir" -type f -name "*.proto"
}

# Define Go modules
MODULES=(
    "github.com/cosmos/cosmos-sdk"
    "github.com/cosmos/ibc-go/v8"
)

# Array to store all proto files
PROTO_FILES=()

# Iterate over each module and collect .proto files
for module in "${MODULES[@]}"; do
    echo "Fetching module path for: $module"
    module_path=$(get_module_path "$module")

    if [[ -z "$module_path" ]]; then
        echo "Error: Could not retrieve path for $module"
        continue
    fi

    echo "Module path: $module_path"
    echo "Searching for .proto files in $module_path..."

    # Find and collect .proto files
    while IFS= read -r file; do
        PROTO_FILES+=("$file")
    done < <(find_proto_files "$module_path")

    echo "--------------------------------------"
done

# Check if we found any .proto files
if [[ ${#PROTO_FILES[@]} -eq 0 ]]; then
    echo "No .proto files found. Exiting..."
    exit 1
fi

for proto_file in "${PROTO_FILES[@]}"; do
    buf generate "$proto_file" --template proto/buf.gen.swagger.yml
done

buf generate --template proto/buf.gen.swagger.yml

rm -rf docs/client/messages.swagger.json

go run proto/swagger.go

sleep 1

rm -rf docs/client
