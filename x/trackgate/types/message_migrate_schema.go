package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMigrateSchema{}

func NewMsgMigrateSchema(operator string, extTrackStationId string, newSchemaKey string) *MsgMigrateSchema {
	return &MsgMigrateSchema{
		Operator:          operator,
		ExtTrackStationId: extTrackStationId,
		NewSchemaKey:      newSchemaKey,
	}
}

func (msg *MsgMigrateSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}
	return nil
}
