package types

const (
	// ModuleName defines the module name
	ModuleName = "rollup"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rollup"

	// Version defines the current version the IBC module supports
	Version = "rollup-1"

	// PortID is the default port id that module binds to
	PortID = "rollup"

	// RollupDataKey is the key to store rollup data with rollup-id as key
	RollupDataKey = "rollup-data-by-rollupid"

	// RollupMonikerKey is the key to store mapping from moniker to rollup-id
	RollupMonikerKey = "rollup-moniker-to-rollupid"

	// RollupCreatorKey is the key to store mapping from creator address to rollup-ids
	LedgerEntryRollupCreatorKey = "ledger-entry-rollup-creator"

	// RollupStakingInfoKey is the key to store mapping from rollup-id to staking info
	RollupStakingInfoKey = "rollup-staking-info"

	// Figures DB
	FiguresDBPath = "Figure/Data/"
)

var (
	ParamsKey = []byte("p_rollup")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("rollup-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
