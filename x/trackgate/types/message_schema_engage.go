package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSchemaEngage{}

func NewMsgSchemaEngage(operator string, extTrackStationId string, schemaObject []byte, acknowledgementHash string, podNumber uint64, sequencerDetails []byte) *MsgSchemaEngage {
	return &MsgSchemaEngage{
		Operator:            operator,
		ExtTrackStationId:   extTrackStationId,
		SchemaObject:        schemaObject,
		AcknowledgementHash: acknowledgementHash,
		PodNumber:           podNumber,
		SequencerDetails:    sequencerDetails,
	}
}

func (msg *MsgSchemaEngage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}
	return nil
}
