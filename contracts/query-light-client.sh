


CONTRACT_ADDRESS=wasmstation1gurgpv8savnfw66lckwzn4zk7fp394lpe667dhu7aw48u40lj6jsad5qmf
#echo "Querying Config..."
../build/wasmstationd  query wasm contract-state smart $CONTRACT_ADDRESS '{"GetConfig": {}}'

echo "Querying Finalized State..."
#../build/wasmstationd  query wasm contract-state smart $CONTRACT_ADDRESS '{"GetFinalizedState": {}}'

#echo "Querying Genesis State..."
#../build/wasmstationd  query wasm contract-state smart $CONTRACT_ADDRESS '{"GetGenesisState": {}}'
##
#echo "Querying Version..."
#../build/wasmstationd  query wasm contract-state smart $CONTRACT_ADDRESS '{"GetVersion": {}}'