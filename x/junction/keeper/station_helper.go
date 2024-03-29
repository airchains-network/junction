package keeper

import (
	"encoding/json"
	"fmt"
	"strconv"

	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/airchains-network/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/* trunk-ignore(golangci-lint/staticcheck) */
func (k Keeper) initStationHelper(ctx sdk.Context, station types.Stations, creator string) error {
	var vk bls12381.VerifyingKey
	err := json.Unmarshal(station.VerificationKey, &vk)
	if err != nil {
		return sdkerrors.ErrInvalidRequest
	}

	//	database of list of stations under each creator
	stationRegistry := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationRegistryKeys))
	// creator is limited to create only one station at this time (this will be changed in future testnet development)
	checkStationExist := stationRegistry.Get([]byte(creator))
	if checkStationExist != nil {
		errorMsg := fmt.Sprintf("station already exist for %s", creator)
		return status.Error(codes.InvalidArgument, errorMsg)
	}

	// check if the user is sending the unique id or not
	stationDataDB := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationDataKey))
	uniqueStationIDCheck := stationDataDB.Get([]byte(station.Id))
	if uniqueStationIDCheck != nil {
		return sdkerrors.ErrConflict
	}

	byteStation := k.cdc.MustMarshal(&station)
	byteStationId := []byte(station.Id)
	creatorByte := []byte(creator)
	stationRegistry.Set(creatorByte, byteStationId)
	stationDataDB.Set(byteStationId, byteStation)

	figuresDB := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FiguresDBPath))
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
	stationDataDB := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationDataKey))
	stationByte := stationDataDB.Get([]byte(stationId))
	var station types.Stations
	if stationByte == nil {
		return station, sdkerrors.ErrInvalidRequest
	}
	k.cdc.MustUnmarshal(stationByte, &station)
	return station, nil
}

func (k Keeper) GetStationIdByAddressHelper(ctx sdk.Context, address string) (stationId string, found bool) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationRegistryKeys))
	stationIdByte := store.Get([]byte(address))
	if stationIdByte == nil {
		return "nil", false
	}

	return string(stationIdByte), true
}
