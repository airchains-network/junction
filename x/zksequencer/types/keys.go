package types

const (
	// ModuleName defines the module name
	ModuleName = "zksequencer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_zksequencer"

	// Version defines the current version the IBC module supports
	Version = "zksequencer-1"

	// PortID is the default port id that module binds to
	PortID = "zksequencer"
)

var (
	ParamsKey = []byte("p_zksequencer")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("zksequencer-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CollegeKey      = "College/value/"
	CollegeCountKey = "College/count/"
)
