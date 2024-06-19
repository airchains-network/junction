#/home/computerkeeda/Desktop/CODES/wasm-integration/attempt-2/junction/build/test-linux-amd64 keys list
#
##/home/computerkeeda/Desktop/CODES/wasm-integration/attempt-2/junction/build/test-linux-amd64 keys add test
#/home/computerkeeda/Desktop/CODES/wasm-integration/attempt-2/junction/build/test-linux-amd64 init junctionXwasmdtesting --chain-id junction --default-denom amf
#/home/computerkeeda/Desktop/CODES/wasm-integration/attempt-2/junction/build/test-linux-amd64 genesis add-genesis-account air1h95ml7ajtadtucck03u6dcq50p2hq6qf7uek2v 100100000000amf
#/home/computerkeeda/Desktop/CODES/wasm-integration/attempt-2/junction/build/test-linux-amd64 genesis gentx test 100000000000amf --chain-id=junction
#/home/computerkeeda/Desktop/CODES/wasm-integration/attempt-2/junction/build/test-linux-amd64 genesis collect-gentxs
#code ~/.junction/
##/home/computerkeeda/Desktop/CODES/wasm-integration/attempt-2/junction/build/test-linux-amd64 start
##
##
## ./build/junctionXwasmd keys list
#
##./build/junctionXwasmd keys add test
##
##./build/junctionXwasmd init junctionXwasmdtesting --chain-id junction --default-denom amf
##
##./build/junctionXwasmd genesis add-genesis-account <key-address> 100100000000amf
##
##./build/junctionXwasmd genesis gentx test 100000000000amf --chain-id=junction
##
##./build/junctionXwasmd genesis collect-gentxs
##
##code ~/.junction/
##
#./build/junctionXwasmd start