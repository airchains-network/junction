package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitiateVrf{}

func NewMsgInitiateVrf(creator string, podNumber uint64, stationId string, occupancy uint64, creatorsVrfKey string, extraArg []byte) *MsgInitiateVrf {
	return &MsgInitiateVrf{
		Creator:        creator,
		PodNumber:      podNumber,
		StationId:      stationId,
		Occupancy:      occupancy,
		CreatorsVrfKey: creatorsVrfKey,
		ExtraArg:       extraArg,
	}
}

func (msg *MsgInitiateVrf) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
