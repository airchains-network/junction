#!/bin/bash
set -o errexit -o nounset -o pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

sleep 1
echo "## Submit a CosmWasm gov proposal"
RESP=$(junctiond tx wasm submit-proposal store-instantiate "$DIR/../../../x/wasm/keeper/testdata/reflect_2_0.wasm" \
  '{}' --label="testing" \
  --title "testing" --summary "Testing" --deposit "1000000000uamf" \
  --admin $(junctiond keys show -a validator --keyring-backend=test) \
  --amount 123uamf \
  --keyring-backend=test \
  --from validator --gas auto --gas-adjustment=1.5 -y  --chain-id=testing --node=http://localhost:26657 -b sync -o json)
echo $RESP
sleep 6
junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json | jq

