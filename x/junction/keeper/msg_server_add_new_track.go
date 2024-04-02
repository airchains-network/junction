package keeper

import (
	"context"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AddNewTrack(goCtx context.Context, msg *types.MsgAddNewTrack) (*types.MsgAddNewTrackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var creator = msg.Creator
	var stationId = msg.StationId
	var newTrackAddress = msg.NewTrackAddress
	var newTrackVotingPower = msg.NewTrackVotingPower
	var signatures = msg.Signatures
	var publicKeys = msg.PublicKeys
	var votes = msg.Votes

	station, err := k.getStationById(ctx, stationId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "station not found")
	}
	tracks := station.Tracks
	var found bool
	for _, track := range tracks {
		if track == creator {
			found = true
			break
		}
	}
	if !found {
		return nil, status.Error(codes.NotFound, "sender not found in tracks")
	}
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	var voteBytes [][]byte
	for i := 0; i < len(votes); i++ {
		// voteBytes[i] = []byte{boolToByte(votes[i])}
		tempVoteByteValue := []byte{boolToByte(votes[i])}
		voteBytes = append(voteBytes, tempVoteByteValue)
	}
	votesBool, signers, success, err := VerifyVrfDisputeSignatures(tracks, signatures, publicKeys, voteBytes)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	if !success {
		return nil, status.Error(codes.InvalidArgument, "")
	}
	voteResult := countDisputeVotes(votesBool, signers)
	if voteResult.Result {
		station.Tracks = append(station.Tracks, newTrackAddress)
		station.VotingPower = append(station.VotingPower, newTrackVotingPower)

		stationDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationDataKey))
		byteStation := k.cdc.MustMarshal(&station)
		byteStationId := []byte(station.Id)
		stationDataDB.Set(byteStationId, byteStation)

		return &types.MsgAddNewTrackResponse{
			Success: true,
		}, nil
	} else {
		return nil, status.Error(codes.InvalidArgument, "more than 50% of votes are not in agreement")
	}
}
