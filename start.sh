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

# Modify the app.toml file with the required changes
APP_TOML="$HOME/.junction/config/app.toml"
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.00025amf"/' $APP_TOML
sed -i 's/enable = false/enable = true/' $APP_TOML
sed -i 's/swagger = false/swagger = true/' $APP_TOML
sed -i 's/tcp:\/\/localhost:1317/tcp:\/\/0.0.0.0:1317/' $APP_TOML
echo "App.toml file updated with new gas prices, enabled Swagger, and updated address"

# Modify laddr in config.toml to listen on all interfaces (0.0.0.0)
CONFIG_TOML="$HOME/.junction/config/config.toml"
sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/' $CONFIG_TOML
echo "Config.toml file updated to listen on all interfaces (0.0.0.0) on port 26657"

# Modify node address in client.toml to 0.0.0.0
CLIENT_TOML="$HOME/.junction/config/client.toml"
sed -i 's/node = "tcp:\/\/localhost:26657"/node = "tcp:\/\/0.0.0.0:26657"/' $CLIENT_TOML
echo "Client.toml file updated with new node address (0.0.0.0)"

./build/junctiond genesis validate-genesis

# Run the junctiond node :
./build/junctiond start