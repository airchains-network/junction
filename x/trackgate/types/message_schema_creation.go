package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSchemaCreation{}

func NewMsgSchemaCreation(creator string, extTrackStationId string, version string, schema []byte) *MsgSchemaCreation {
	return &MsgSchemaCreation{
		Creator:           creator,
		ExtTrackStationId: extTrackStationId,
		Version:           version,
		Schema:            schema,
	}
}

func (msg *MsgSchemaCreation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
