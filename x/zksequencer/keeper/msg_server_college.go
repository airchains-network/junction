package keeper

import (
	"context"
	"fmt"

	"github.com/airchains-network/junction/x/zksequencer/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCollege(goCtx context.Context, msg *types.MsgCreateCollege) (*types.MsgCreateCollegeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var college = types.College{
		Creator:  msg.Creator,
		Students: msg.Students,
		Name:     msg.Name,
		Details:  msg.Details,
	}

	id := k.AppendCollege(
		ctx,
		college,
	)

	return &types.MsgCreateCollegeResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateCollege(goCtx context.Context, msg *types.MsgUpdateCollege) (*types.MsgUpdateCollegeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var college = types.College{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Students: msg.Students,
		Name:     msg.Name,
		Details:  msg.Details,
	}

	// Checks that the element exists
	val, found := k.GetCollege(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCollege(ctx, college)

	return &types.MsgUpdateCollegeResponse{}, nil
}

func (k msgServer) DeleteCollege(goCtx context.Context, msg *types.MsgDeleteCollege) (*types.MsgDeleteCollegeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetCollege(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCollege(ctx, msg.Id)

	return &types.MsgDeleteCollegeResponse{}, nil
}
