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
	FiguresDBPath  = "Figure/Data/"
	// PodStoreKey this is the DB path of the pod store where all the data is stored in key value pair and the key is in this kind of format `pods__{stationId}__{podNumber}` and value is the byte of pods
	//PodStoreKey = "Pod/Data"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
