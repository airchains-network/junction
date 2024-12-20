package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/cipherpodledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) IntegrityCheck(goCtx context.Context, msg *types.MsgIntegrityCheck) (*types.MsgIntegrityCheckResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgIntegrityCheckResponse{}, nil
}
