package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vrftypes "github.com/airchains-network/junction/x/vrf/types" // Import VRF types
	"github.com/airchains-network/junction/x/zksequencer/types"
)

func (k msgServer) AddStudentViaZksequencer(goCtx context.Context, msg *types.MsgAddStudentViaZksequencer) (*types.MsgAddStudentViaZksequencerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create VRF student using VRF MsgCreateStudents
	vrfMsg := vrftypes.Students{
		Creator: msg.Creator, // Pass the sender from zksequencer
		Name:    msg.Name,
		Details: msg.Details,
	}

	// Call VRF's keeper to create the student
	studentID := k.Keeper.vrfKeeper.AppendStudents(ctx, vrfMsg)

	// Example: Fetch the latest student ID (if needed)
	//studentID := uint64(1) // You might need to query it properly from vrf module

	return &types.MsgAddStudentViaZksequencerResponse{
		Studentid: studentID,
	}, nil
}
