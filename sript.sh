#!/bin/bash
# Description: Initialize the node

# Ensure jq is installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# Variables
KEY=node
CHAINID=junction
MONIKER=noooblien
KEYRING="test"
DENOM="amf"
TOTAL_SUPPLY="10000000000000000"
VALIDATOR_STAKE="10000000000"

# Clean build directory
rm -rf ./build/
make build

# Remove the existing data
rm -rf ~/.junction

# Initialize the node
./build/junction-test init $MONIKER --chain-id $CHAINID --default-denom $DENOM

# Prepare your account
./build/junction-test keys add $KEY --keyring-backend $KEYRING

# Add the account to genesis
./build/junction-test genesis add-genesis-account $KEY ${TOTAL_SUPPLY}${DENOM} --keyring-backend $KEYRING

# Create a gentx
./build/junction-test genesis gentx $KEY ${VALIDATOR_STAKE}${DENOM} \
    --keyring-backend $KEYRING \
    --chain-id $CHAINID \
    --moniker=$MONIKER \
    --commission-rate="0.10" \
    --commission-max-rate="0.20" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation=${VALIDATOR_STAKE}

# Collect the gentxs
./build/junction-test genesis collect-gentxs

# Validate the genesis file
./build/junction-test genesis validate-genesis

# Update staking parameters in genesis
jq '.app_state.staking.params.bond_denom = "amf"' ~/.junction/config/genesis.json > temp.json && mv temp.json ~/.junction/config/genesis.json

# Update crisis parameters in genesis
jq '.app_state.crisis.constant_fee.denom = "amf"' ~/.junction/config/genesis.json > temp.json && mv temp.json ~/.junction/config/genesis.json

# Update gov parameters in genesis
jq '.app_state.gov.params.min_deposit[0].denom = "amf"' ~/.junction/config/genesis.json > temp.json && mv temp.json ~/.junction/config/genesis.json
jq '.app_state.gov.params.expedited_min_deposit[0].denom = "amf"' ~/.junction/config/genesis.json > temp.json && mv temp.json ~/.junction/config/genesis.json

# Update mint parameters in genesis
jq '.app_state.mint.params.mint_denom = "amf"' ~/.junction/config/genesis.json > temp.json && mv temp.json ~/.junction/config/genesis.json

# Verify the genesis file
echo "Verifying genesis file..."
cat ~/.junction/config/genesis.json | jq '.app_state.staking.validators'
cat ~/.junction/config/genesis.json | jq '.app_state.genutil.gen_txs'

# Start the node
./build/junction-test start --api.enable --minimum-gas-prices 0.00025${DENOM}