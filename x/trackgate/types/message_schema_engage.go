package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSchemaEngage{}

func NewMsgSchemaEngage(operator string, extTrackStationId string, schemaKey string, schemaObject []byte, stateRoot string) *MsgSchemaEngage {
	return &MsgSchemaEngage{
		Operator:          operator,
		ExtTrackStationId: extTrackStationId,
		SchemaKey:         schemaKey,
		SchemaObject:      schemaObject,
		StateRoot:         stateRoot,
	}
}

func (msg *MsgSchemaEngage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}
	return nil
}
