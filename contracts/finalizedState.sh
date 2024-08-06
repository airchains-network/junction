#!/bin/bash
# Description: Submit New Finalized State

#CONTRACT_ADDRESS=wasmstation1pvrwmjuusn9wh34j7y520g8gumuy9xtl3gvprlljfdpwju3x7ucs553dqf
#
#VIEW_NUM=2
#BLOCK_HEIGHT=100
#BLOCK_COMM_ROOT="1234567890"
#FEE_LEDGER_COMM="2345678901"
#STAKE_TABLE_BLS_KEY_COMM="3456789012"
#STAKE_TABLE_SCHNORR_KEY_COMM="4567890123"
#STAKE_TABLE_AMOUNT_COMM="5678901234"
#THRESHOLD="200"
#
#../build/wasmstationd tx wasm execute $CONTRACT_ADDRESS '{
#  "NewFinalizedState": {
#    "new_state": {
#      "view_num": 2,
#      "block_height": 100,
#      "block_comm_root": {
#        "value": "1234567890"
#      },
#      "fee_ledger_comm": {
#        "value": "2345678901"
#      },
#      "stake_table_bls_key_comm": {
#        "value": "3456789012"
#      },
#      "stake_table_schnorr_key_comm": {
#        "value": "4567890123"
#      },
#      "stake_table_amount_comm": {
#        "value": "5678901234"
#      },
#      "threshold": "200"
#    },
#    "proof": "dummy_proof_string"
#  }
#}' --from node --keyring-backend "test" --gas auto --gas-adjustment 1.2 --fees 5000stake --chain-id station-1


#!/bin/bash
# Description: Submit New Finalized State

#!/bin/bash
# Description: Submit New Finalized State (Simplified)

#!/bin/bash
# Description: Simplified Submit New Finalized State

#!/bin/bash
# Description: Submit New Finalized State

#CONTRACT_ADDRESS=wasmstation1pvrwmjuusn9wh34j7y520g8gumuy9xtl3gvprlljfdpwju3x7ucs553dqf

CONTRACT_ADDRESS=wasmstation1gurgpv8savnfw66lckwzn4zk7fp394lpe667dhu7aw48u40lj6jsad5qmf


#PROOF="{ \\\"wire0\\\": { \\\"x\\\": \\\"0\\\", \\\"y\\\": \\\"1\\\" }, \\\"wire1\\\": { \\\"x\\\": \\\"2\\\", \\\"y\\\": \\\"3\\\" }, \\\"wire2\\\": { \\\"x\\\": \\\"4\\\", \\\"y\\\": \\\"5\\\" }, \\\"wire3\\\": { \\\"x\\\": \\\"6\\\", \\\"y\\\": \\\"7\\\" }, \\\"wire4\\\": { \\\"x\\\": \\\"8\\\", \\\"y\\\": \\\"9\\\" }, \\\"prod_perm\\\": { \\\"x\\\": \\\"10\\\", \\\"y\\\": \\\"11\\\" }, \\\"split0\\\": { \\\"x\\\": \\\"12\\\", \\\"y\\\": \\\"13\\\" }, \\\"split1\\\": { \\\"x\\\": \\\"14\\\", \\\"y\\\": \\\"15\\\" }, \\\"split2\\\": { \\\"x\\\": \\\"16\\\", \\\"y\\\": \\\"17\\\" }, \\\"split3\\\": { \\\"x\\\": \\\"18\\\", \\\"y\\\": \\\"19\\\" }, \\\"split4\\\": { \\\"x\\\": \\\"20\\\", \\\"y\\\": \\\"21\\\" }, \\\"zeta\\\": { \\\"x\\\": \\\"22\\\", \\\"y\\\": \\\"23\\\" }, \\\"zeta_omega\\\": { \\\"x\\\": \\\"24\\\", \\\"y\\\": \\\"25\\\" }, \\\"wire_eval0\\\": { \\\"value\\\": \\\"26\\\" }, \\\"wire_eval1\\\": { \\\"value\\\": \\\"27\\\" }, \\\"wire_eval2\\\": { \\\"value\\\": \\\"28\\\" }, \\\"wire_eval3\\\": { \\\"value\\\": \\\"29\\\" }, \\\"wire_eval4\\\": { \\\"value\\\": \\\"30\\\" }, \\\"sigma_eval0\\\": { \\\"value\\\": \\\"31\\\" }, \\\"sigma_eval1\\\": { \\\"value\\\": \\\"32\\\" }, \\\"sigma_eval2\\\": { \\\"value\\\": \\\"33\\\" }, \\\"sigma_eval3\\\": { \\\"value\\\": \\\"34\\\" }, \\\"prod_perm_zeta_omega_eval\\\": { \\\"value\\\": \\\"35\\\" } }"

#!/bin/bash
# Description: Submit New Finalized State


VIEW_NUM=4
BLOCK_HEIGHT=40
BLOCK_COMM_ROOT="1234567890"
FEE_LEDGER_COMM="2345678901"
STAKE_TABLE_BLS_KEY_COMM="3456789012"
STAKE_TABLE_SCHNORR_KEY_COMM="4567890123"
STAKE_TABLE_AMOUNT_COMM="5678901234"
THRESHOLD=100

PROOF=$(cat <<EOF
{
      "wire0": { "x": "2166381329", "y": "4388970787" },
      "wire1": { "x": "1725591613", "y": "3661609453" },
      "wire2": { "x": "1460205438", "y": "3873421494" },
      "wire3": { "x": "1665976827", "y": "2286691652" },
      "wire4": { "x": "3411989198", "y": "3655282556" },
      "prod_perm": { "x": "1455902546", "y": "2571162080" },
      "split0": { "x": "3468037418", "y": "3922838781" },
      "split1": { "x": "2335782196", "y": "3881474989" },
      "split2": { "x": "1583955451", "y": "2568964436" },
      "split3": { "x": "1771385625", "y": "2309006736" },
      "split4": { "x": "1802720890", "y": "3904911082" },
      "zeta": { "x": "2013144571", "y": "2923796178" },
      "zeta_omega": { "x": "1966944859", "y": "2783630353" },
  "wire_eval0": "26",
  "wire_eval1": "27",
  "wire_eval2": "28",
  "wire_eval3": "29",
  "wire_eval4": "30",
  "sigma_eval0": "31",
  "sigma_eval1": "32",
  "sigma_eval2": "33",
  "sigma_eval3": "34",
  "prod_perm_zeta_omega_eval": "35"
}
EOF
)

TX_MSG=$(cat <<EOF
{
  "NewFinalizedState": {
    "new_state": {
      "view_num": $VIEW_NUM,
      "block_height": $BLOCK_HEIGHT,
      "block_comm_root": "$BLOCK_COMM_ROOT" ,
      "fee_ledger_comm": "$FEE_LEDGER_COMM" ,
      "stake_table_bls_key_comm": "$STAKE_TABLE_BLS_KEY_COMM" ,
      "stake_table_schnorr_key_comm": "$STAKE_TABLE_SCHNORR_KEY_COMM" ,
      "stake_table_amount_comm": "$STAKE_TABLE_AMOUNT_COMM" ,
      "threshold": "$THRESHOLD"
    },
    "proof": $PROOF
  }
}
EOF
)

echo "$TX_MSG"

../build/wasmstationd tx wasm execute $CONTRACT_ADDRESS "$TX_MSG" \
  --from node \
  --keyring-backend test \
  --gas auto \
  --gas-adjustment 1.2 \
  --fees 5000stake \
  --chain-id station-1 \
  -y
