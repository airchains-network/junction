package keeper

import (
	"encoding/json"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/ComputerKeeda/junction/x/junction/types"
	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) initStationHelper(ctx sdk.Context, station types.Stations, creator string) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	var vk bls12381.VerifyingKey
	err := json.Unmarshal(station.VerificationKey, &vk)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid verification key")
	}

	//	database of list of stations under each track member
	stationRegistry := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationRegistryKeys))

	// check if the user is sending the unique id or not
	stationDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationDataKey))
	uniqueStationIDCheck := stationDataDB.Get([]byte(station.Id))
	if uniqueStationIDCheck != nil {
		return status.Error(codes.InvalidArgument, "station id already exists")
	}

	byteStation := k.cdc.MustMarshal(&station)
	byteStationId := []byte(station.Id)
	tracksBytes, tbe := json.Marshal(station.Tracks)
	if tbe != nil {
		return status.Error(codes.InvalidArgument, "invalid tracks")
	}
	stationRegistry.Set(byteStationId, tracksBytes)
	stationDataDB.Set(byteStationId, byteStation)

	figuresDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
	stationCountByte := figuresDB.Get([]byte("station-count"))

	if stationCountByte == nil {
		figuresDB.Set([]byte("station-count"), []byte("1"))
	} else {
		stationCount := string(stationCountByte)

		stationCountUint64, err := strconv.ParseUint(stationCount, 10, 64)
		if err != nil {
			return sdkerrors.ErrInvalidRequest
		}
		stationCountUint64++

		stationCount = strconv.FormatUint(stationCountUint64, 10)
		figuresDB.Set([]byte("station-count"), []byte(stationCount))
	}
	return nil
}

func (k Keeper) getStationById(ctx sdk.Context, stationId string) (types.Stations, error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	stationDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationDataKey))
	stationByte := stationDataDB.Get([]byte(stationId))
	var station types.Stations
	if stationByte == nil {
		return station, sdkerrors.ErrInvalidRequest
	}
	k.cdc.MustUnmarshal(stationByte, &station)
	return station, nil
}

func findKeyByValue(store *prefix.Store, targetValue string) (key []byte, found bool) {
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		if string(iterator.Value()) == targetValue {
			return iterator.Key(), true
		}
	}
	// Return nil if no match is found
	return nil, false
}
func (k Keeper) GetStationsIdByAddressHelper(ctx sdk.Context, trackMemberAddress string) (stationIds []string, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationRegistryKeys))

	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		if string(iterator.Value()) == trackMemberAddress {
			stationIds = append(stationIds, string(iterator.Key()))
		}
	}

	if len(stationIds) == 0 {
		return stationIds, false
	}
	return stationIds, true
}
