#!/bin/bash
CHAINID=station-1
KEY=node




../build/junction-test tx wasm instantiate 1 '{
  "genesis_state": {
    "view_num": 1,
    "block_height": 0,
    "block_comm_root": "1234567890",
    "fee_ledger_comm": "2345678901",
    "stake_table_bls_key_comm": "3456789012",
    "stake_table_schnorr_key_comm": "4567890123",
    "stake_table_amount_comm": "5678901234",
    "threshold": "100"
  },
  "blocks_per_epoch": 1000000000
}' --from node --keyring-backend "test" --label "light client contract" --admin air1d4rep4g7ekw8mn5wvym9hkglxssxk90xwdenng --gas auto --gas-adjustment 1.2 --fees 5000amf --chain-id junction

