package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/junction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveTrack(goCtx context.Context, msg *types.MsgRemoveTrack) (*types.MsgRemoveTrackResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var creator = msg.Creator
	var stationId = msg.StationId
	var trackAddress = msg.TrackAddress
	var signatures = msg.Signatures
	var votes = msg.Votes
	var publicKeys = msg.PublicKeys

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

	// formatted track slice
	index := -1
	// Find the index of the value to remove
	for i, v := range tracks {
		if v == trackAddress {
			index = i
			break
		}
	}

	formattedTracks := tracks
	formattedTracksVotingPower := station.VotingPower
	if index != -1 {
		// Remove the element at index
		formattedTracks = append(formattedTracks[:index], formattedTracks[index+1:]...)
		formattedTracksVotingPower = append(formattedTracksVotingPower[:index], formattedTracksVotingPower[index+1:]...)
	} else {
		return nil, status.Error(codes.NotFound, "track not found")
	}

	votesBool, signers, success, err := VerifyVrfDisputeSignatures(formattedTracks, signatures, publicKeys, voteBytes)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	if !success {
		return nil, status.Error(codes.InvalidArgument, "")
	}
	voteResult := countDisputeVotes(votesBool, signers)

	if voteResult.Result {
		updatedStation := types.Stations{
			Tracks:               formattedTracks,
			VotingPower:          formattedTracksVotingPower,
			LatestPod:            station.LatestPod,
			LatestMerkleRootHash: station.LatestMerkleRootHash,
			VerificationKey:      station.VerificationKey,
			StationInfo:          station.StationInfo,
			Id:                   station.Id,
			Creator:              station.Creator,
			Spsp:                 station.Spsp,
			TrackType:            station.TrackType,
			DaType:               station.DaType,
			Prover:               station.Prover,
		}

		stationDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StationDataKey))
		byteStation := k.cdc.MustMarshal(&updatedStation)
		byteStationId := []byte(station.Id)
		stationDataDB.Set(byteStationId, byteStation)

		return &types.MsgRemoveTrackResponse{
			Success: true,
		}, nil
	} else {
		return nil, status.Error(codes.InvalidArgument, "more than 50% of votes are not in agreement")
	}
}
