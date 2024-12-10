#!/bin/bash

# Define the necessary environment variables :
MONIKER="junction-testing"
CHAIN_ID="junction"
DENOM="amf"
KEY_NAME="zaza"
AMOUNT="1500000000000000000000amf"
VALIDATOR_STAKE="1000000000000000amf"
KEY_RING_BACKEND="--keyring-backend test"

# Remove the existing junctiond directory :
rm -rf ~/.junction

# Initialize the junctiond node :
./build/junctiond init "$MONIKER" --default-denom "$DENOM" --chain-id "$CHAIN_ID"

# Generate Keys :
./build/junctiond keys add $KEY_NAME $KEY_RING_BACKEND

# Add the genesis account to the genesis file :
./build/junctiond genesis add-genesis-account "$KEY_NAME" $AMOUNT $KEY_RING_BACKEND

# Stake the validator account in the genesis file :
./build/junctiond genesis gentx "$KEY_NAME" $VALIDATOR_STAKE $KEY_RING_BACKEND --gas-prices="0.0025amf"

# Collect gentx files :
./build/junctiond genesis collect-gentxs

# Modify the genesis file with the required changes
GENESIS_FILE="$HOME/.junction/config/genesis.json"
jq '.app_state.gov.params.max_deposit_period = "600s" |
    .app_state.gov.params.voting_period = "600s" |
    .app_state.gov.params.expedited_voting_period = "300s"' \
    $GENESIS_FILE > "$GENESIS_FILE.tmp" && mv "$GENESIS_FILE.tmp" "$GENESIS_FILE"
echo "Genesis file updated with new voting and deposit periods"

# Detect the operating system
OS=$(uname)

# Function to apply sed command based on the OS
apply_sed() {
    local file="$1"
    local pattern="$2"
    local replacement="$3"

    if [ "$OS" = "Darwin" ]; then
        sed -i '' "s|$pattern|$replacement|" "$file"
    else
        sed -i "s|$pattern|$replacement|" "$file"
    fi
}
# Paths to configuration files
APP_TOML="$HOME/.junction/config/app.toml"
CONFIG_TOML="$HOME/.junction/config/config.toml"
CLIENT_TOML="$HOME/.junction/config/client.toml"

# Modify the app.toml file with the required changes
apply_sed "$APP_TOML" 'minimum-gas-prices = ""' 'minimum-gas-prices = "0.00025amf"'
apply_sed "$APP_TOML" 'enable = false' 'enable = true'
apply_sed "$APP_TOML" 'swagger = false' 'swagger = true'
apply_sed "$APP_TOML" 'tcp://localhost:1317' 'tcp://0.0.0.0:1317'
echo "App.toml file updated with new gas prices, enabled Swagger, and updated address"

# Modify laddr in config.toml to listen on all interfaces (0.0.0.0)
apply_sed "$CONFIG_TOML" 'laddr = "tcp://127.0.0.1:26657"' 'laddr = "tcp://0.0.0.0:26657"'
echo "Config.toml file updated to listen on all interfaces (0.0.0.0) on port 26657"

# Modify node address in client.toml to 0.0.0.0
apply_sed "$CLIENT_TOML" 'node = "tcp://localhost:26657"' 'node = "tcp://0.0.0.0:26657"'
echo "Client.toml file updated with new node address (0.0.0.0)"

# Validate genesis file
./build/junctiond genesis validate-genesis

# Run the junctiond node
./build/junctiond start