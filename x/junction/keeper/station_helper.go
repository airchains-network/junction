package keeper

import (
	"encoding/json"
	"github.com/ComputerKeeda/junction/x/junction/types"
	bls12381 "github.com/airchains-network/gnark/backend/groth16/bls12-381"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

/* trunk-ignore(golangci-lint/staticcheck) */
func (k Keeper) initStationHelper(ctx sdk.Context, station types.Stations, creator string) *sdkerrors.Error {
	var vk bls12381.VerifyingKey
	err := json.Unmarshal(station.VerificationKey, &vk)
	if err != nil {
		return sdkerrors.ErrInvalidRequest
	}

	//	database of list of stations under each creator
	stationRegistry := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationRegistryKeys))
	byteStationId := []byte(station.Id)
	// creator is limited to create only one station at this time (this will be changed in future testnet development)
	checkStationExist := stationRegistry.Get([]byte(creator))
	if checkStationExist != nil {
		return sdkerrors.ErrInvalidAddress
	}

	stationRegistry.Set([]byte(creator), []byte(byteStationId))
	//stationVerificationKeyDB := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationVerificationKeyKeys))

	//newVerificationKey := types.StationVerificationKey{
	//	StationId:       station.Id,
	//	VerificationKey: station.VerificationKey,
	//}
	//stationVerificationKeyDB.Set([]byte(station.Id), station.VerificationKey)

	stationDataDB := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StationDataKey))
	b := k.cdc.MustMarshal(&station)
	stationDataDB.Set([]byte(station.Id), b)

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
