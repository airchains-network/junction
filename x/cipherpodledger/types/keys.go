package types

const (
	// ModuleName defines the module name
	ModuleName = "cipherpodledger"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cipherpodledger"

	// Version defines the current version the IBC module supports
	Version = "cipherpodledger-1"

	// PortID is the default port id that module binds to
	PortID = "cipherpodledger"
)

var (
	ParamsKey = []byte("p_cipherpodledger")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("cipherpodledger-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
