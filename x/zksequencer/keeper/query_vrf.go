package keeper

import (
	vrftypes "github.com/airchains-network/junction/x/vrf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetStudentFromVRF Query Student from vrf
func (k Keeper) GetStudentFromVRF(ctx sdk.Context, Id uint64) (*vrftypes.Students, error) {
	req := &vrftypes.QueryGetStudentsRequest{Id: Id}

	// Use VRF Keeper to query student data
	vrfResp, err := k.vrfKeeper.Students(ctx, req)
	if err != nil {
		return nil, err
	}

	return &vrfResp.Students, nil
}
