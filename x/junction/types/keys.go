package types

const (
	// ModuleName defines the module name
	ModuleName = "junction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_junction"
)

var (
	ParamsKey = []byte("p_junction")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
