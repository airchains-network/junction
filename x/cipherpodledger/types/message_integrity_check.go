package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIntegrityCheck{}

func NewMsgIntegrityCheck(creator string, stationId string, blobRef string, daProvider string) *MsgIntegrityCheck {
	return &MsgIntegrityCheck{
		Creator:    creator,
		StationId:  stationId,
		BlobRef:    blobRef,
		DaProvider: daProvider,
	}
}

func (msg *MsgIntegrityCheck) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
