package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitStation{}

func NewMsgInitStation(creator string, tracks []string, verificationKey []byte, stationId string, stationInfo string) *MsgInitStation {
	return &MsgInitStation{
		Creator:         creator,
		Tracks:          tracks,
		VerificationKey: verificationKey,
		StationId:       stationId,
		StationInfo:     stationInfo,
	}
}

func (msg *MsgInitStation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
