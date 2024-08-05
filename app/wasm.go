package app

import (
	wasmkeeper "github.com/airchains-network/junction/x/wasm/keeper"
)

// Deprecated: Use BuiltInCapabilities from github.com/airchains-network/junction/x/wasm/keeper
func AllCapabilities() []string {
	return wasmkeeper.BuiltInCapabilities()
}
