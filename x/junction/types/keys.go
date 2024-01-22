package types

const (
	// ModuleName defines the module name
	ModuleName = "junction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_junction"

	//	our custom keys
	// StationRegistryKeys
	StationRegistryKeys = "Station/StationRegistry/"
	//StationVerificationKeyKeys = "Station/VerificationKey/"
	StationDataKey = "Station/Data/"

	FiguresDBPath = "Figure/Data/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
