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

	/**
	 * FhvmKeyStoreKey is the key store key for the FHVM metadata
	 * stored with the `chain ID` as the key.
	 * Key: chainId
	 * Value: FHVM metadata
	 */
	FhvmKeyStoreKey = "fhvm-key-store"

	/**
	 * Store for ASC Child contract address to chainId mapping
	 * Key: ascChildContractAddress
	 * Value: chainId
	 */
	AscChildContractRegistryKey = "asc-child-contract-registry"

	/**
	 * Store for Pod data
	 * Key for this will be created as "pods/{stationId}/ to create this key call func GetPodKeyByte()
	 * Key: podNumber
	 * Value: podData
	 */
	PodDataStoreKey = "pod-data"

	// FiguresDBPath is the path for the figures database
	FiguresDBPath = "Figure/Data/"
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
