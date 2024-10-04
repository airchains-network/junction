package types

const (
	// ModuleName defines the module name
	ModuleName = "trackgate"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_trackgate"

	ExtTrackStationsDataStoreKey = "ext_track_stations_data"

	ExtTrackSchemaStoreKey            = "ext_track_schema"
	ExtTrackVersionFinderStoreKey     = "ext_track_version_finder"
	ExtTrackSchemaEngagementsStoreKey = "ext_track_schema_engagements"
	TrackGateFigureStoreKey           = "trackgate_figur"
)

var (
	ParamsKey = []byte("p_trackgate")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
