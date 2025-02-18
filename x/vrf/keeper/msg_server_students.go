package keeper

import (
	"context"
	"fmt"

	"github.com/airchains-network/junction/x/vrf/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStudents(goCtx context.Context, msg *types.MsgCreateStudents) (*types.MsgCreateStudentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var students = types.Students{
		Creator: msg.Creator,
		Name:    msg.Name,
		Details: msg.Details,
	}

	id := k.AppendStudents(
		ctx,
		students,
	)

	return &types.MsgCreateStudentsResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateStudents(goCtx context.Context, msg *types.MsgUpdateStudents) (*types.MsgUpdateStudentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var students = types.Students{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Details: msg.Details,
	}

	// Checks that the element exists
	val, found := k.GetStudents(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetStudents(ctx, students)

	return &types.MsgUpdateStudentsResponse{}, nil
}

func (k msgServer) DeleteStudents(goCtx context.Context, msg *types.MsgDeleteStudents) (*types.MsgDeleteStudentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetStudents(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStudents(ctx, msg.Id)

	return &types.MsgDeleteStudentsResponse{}, nil
}
