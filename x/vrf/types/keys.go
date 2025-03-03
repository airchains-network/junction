package types

const (
	// ModuleName defines the module name
	ModuleName = "vrf"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	VrfDisputeStoreKey = "vrfDispute"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vrf"

	// Version defines the current version the IBC module supports
	Version = "vrf-1"

	// PortID is the default port id that module binds to
	PortID = "vrf"
)

var (
	ParamsKey = []byte("p_vrf")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("vrf-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
