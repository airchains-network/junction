#!/bin/bash
# Description: Initialize the node

CHAINID=junction
KEY=node

../build/junction-test tx wasm store ./light_client3.wasm --from $KEY   --chain-id $CHAINID --gas auto --gas-adjustment 1.2 --fees 2000amf --keyring-backend test   -y









