#!/bin/bash
set -o errexit -o nounset -o pipefail -x

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

echo "-----------------------"
echo "## Add new CosmWasm contract"
RESP=$(junctiond tx wasm store "$DIR/../../../x/wasm/keeper/testdata/hackatom.wasm" \
  --from validator --gas 1500000 -y --chain-id=testing --node=http://localhost:26657 -b sync -o json --keyring-backend=test)
sleep 6
RESP=$(junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json)
CODE_ID=$(echo "$RESP" | jq -r '.events[]| select(.type=="store_code").attributes[]| select(.key=="code_id").value')
CODE_HASH=$(echo "$RESP" | jq -r '.events[]| select(.type=="store_code").attributes[]| select(.key=="code_checksum").value')
echo "* Code id: $CODE_ID"
echo "* Code checksum: $CODE_HASH"

echo "* Download code"
TMPDIR=$(mktemp -d -t wasmdXXXXXX)
junctiond q wasm code "$CODE_ID" "$TMPDIR/delme.wasm"
rm -f "$TMPDIR/delme.wasm"
echo "-----------------------"
echo "## List code"
junctiond query wasm list-code --node=http://localhost:26657 -o json | jq

echo "-----------------------"
echo "## Create new contract instance"
INIT="{\"verifier\":\"$(junctiond keys show validator -a --keyring-backend=test)\", \"beneficiary\":\"$(junctiond keys show fred -a --keyring-backend=test)\"}"
RESP=$(junctiond tx wasm instantiate "$CODE_ID" "$INIT" --admin="$(junctiond keys show validator -a --keyring-backend=test)" \
  --from validator --amount="100uamf" --label "local0.1.0" \
  --gas 1000000 -y --chain-id=testing -b sync -o json --keyring-backend=test)
sleep 6
junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json | jq

CONTRACT=$(junctiond query wasm list-contract-by-code "$CODE_ID" -o json | jq -r '.contracts[-1]')
echo "* Contract address: $CONTRACT"

echo "## Create new contract instance with predictable address"
RESP=$(junctiond tx wasm instantiate2 "$CODE_ID" "$INIT" $(echo -n "testing" | xxd -ps) \
  --admin="$(junctiond keys show validator -a --keyring-backend=test)" \
  --from validator --amount="100uamf" --label "local0.1.0" \
  --fix-msg \
  --gas 1000000 -y --chain-id=testing -b sync -o json --keyring-backend=test)
sleep 6
junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json | jq


predictedAddress=$(junctiond q wasm build-address "$CODE_HASH" $(junctiond keys show validator -a --keyring-backend=test) $(echo -n "testing" | xxd -ps) "$INIT" | awk '{print $2}')
junctiond q wasm contract "$predictedAddress" -o json | jq

echo "### Query all"
RESP=$(junctiond query wasm contract-state all "$CONTRACT" -o json)
echo "$RESP" | jq
echo "### Query smart"
junctiond query wasm contract-state smart "$CONTRACT" '{"verifier":{}}' -o json | jq
echo "### Query raw"
KEY=$(echo "$RESP" | jq -r ".models[0].key")
junctiond query wasm contract-state raw "$CONTRACT" "$KEY" -o json | jq

echo "-----------------------"
echo "## Execute contract $CONTRACT"
MSG='{"release":{}}'
RESP=$(junctiond tx wasm execute "$CONTRACT" "$MSG" \
  --from validator \
  --gas 1000000 -y --chain-id=testing -b sync -o json --keyring-backend=test)
sleep 6
junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json | jq


echo "-----------------------"
echo "## Set new admin"
echo "### Query old admin: $(junctiond q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
echo "### Update contract"
RESP=$(junctiond tx wasm set-contract-admin "$CONTRACT" "$(junctiond keys show fred -a --keyring-backend=test)" \
  --from validator -y --chain-id=testing -b sync -o json --keyring-backend=test)
sleep 6
junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json | jq

echo "### Query new admin: $(junctiond q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"

echo "-----------------------"
echo "## Migrate contract"
echo "### Upload new code"
RESP=$(junctiond tx wasm store "$DIR/../../../x/wasm/keeper/testdata/burner.wasm" \
  --from validator --gas 1100000 -y --chain-id=testing --node=http://localhost:26657 -b sync -o json --keyring-backend=test)
sleep 6
RESP=$(junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json)
BURNER_CODE_ID=$(echo "$RESP" | jq -r '.events[]| select(.type=="store_code").attributes[]| select(.key=="code_id").value')

echo "### Migrate to code id: $BURNER_CODE_ID"

DEST_ACCOUNT=$(junctiond keys show fred -a --keyring-backend=test)
RESP=$(junctiond tx wasm migrate "$CONTRACT" "$BURNER_CODE_ID" "{\"payout\": \"$DEST_ACCOUNT\"}" --from fred \
  --chain-id=testing -b sync -y -o json --keyring-backend=test)
sleep 6
junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json | jq

echo "### Query destination account: $BURNER_CODE_ID"
junctiond q bank balances "$DEST_ACCOUNT" -o json | jq
echo "### Query contract meta data: $CONTRACT"
junctiond q wasm contract "$CONTRACT" -o json | jq

echo "### Query contract meta history: $CONTRACT"
junctiond q wasm contract-history "$CONTRACT" -o json | jq

echo "-----------------------"
echo "## Clear contract admin"
echo "### Query old admin: $(junctiond q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
echo "### Update contract"
RESP=$(junctiond tx wasm clear-contract-admin "$CONTRACT" \
  --from fred -y --chain-id=testing -b sync -o json --keyring-backend=test)
sleep 6
junctiond q tx $(echo "$RESP"| jq -r '.txhash') -o json | jq
echo "### Query new admin: $(junctiond q wasm contract "$CONTRACT" -o json | jq -r '.contract_info.admin')"
