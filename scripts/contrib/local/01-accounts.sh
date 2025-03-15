#!/bin/bash
set -o errexit -o nounset -o pipefail

BASE_ACCOUNT=$(junctiond keys show validator -a --keyring-backend=test)
junctiond q auth account "$BASE_ACCOUNT" -o json | jq

echo "## Add new account"
junctiond keys add fred --keyring-backend=test

echo "## Check balance"
NEW_ACCOUNT=$(junctiond keys show fred -a --keyring-backend=test)
junctiond q bank balances "$NEW_ACCOUNT" -o json || true

echo "## Transfer tokens"
junctiond tx bank send validator "$NEW_ACCOUNT" 1uamf --gas 1000000 -y --chain-id=testing --node=http://localhost:26657 -b sync -o json --keyring-backend=test | jq

echo "## Check balance again"
junctiond q bank balances "$NEW_ACCOUNT" -o json | jq
