#!/bin/bash
set -o errexit -o nounset -o pipefail

PASSWORD=${PASSWORD:-1234567890}
STAKE=${STAKE_TOKEN:-uamf}
FEE=${FEE_TOKEN:-uamf}
CHAIN_ID=${CHAIN_ID:-testing}
MONIKER=${MONIKER:-junction}

junctiond init --chain-id "$CHAIN_ID" "$MONIKER"
# staking/governance token is hardcoded in config, change this
## OSX requires: -i.
sed -i. "s/\"stake\"/\"$STAKE\"/" "$HOME"/.junctiond/config/genesis.json
if ! junctiond keys show validator --keyring-backend=test; then
  (
    echo "$PASSWORD"
    echo "$PASSWORD"
  ) | junctiond keys add validator --keyring-backend=test
fi
# hardcode the validator account for this instance
echo "$PASSWORD" | junctiond genesis add-genesis-account validator "10000000000000000$STAKE" --keyring-backend=test
# (optionally) add a few more genesis accounts
for addr in "$@"; do
  echo "$addr"
  junctiond genesis add-genesis-account "$addr" "10000000000000$STAKE" --keyring-backend=test
done
# submit a genesis validator tx
## Workraround for https://github.com/cosmos/cosmos-sdk/issues/8251
(
  echo "$PASSWORD"
  echo "$PASSWORD"
  echo "$PASSWORD"
) | junctiond genesis gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID" --amount="250000000$STAKE" --keyring-backend=test
## should be:
# (echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | junctiond gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID"
junctiond genesis collect-gentxs
