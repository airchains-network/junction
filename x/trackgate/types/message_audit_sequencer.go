package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAuditSequencer{}

func NewMsgAuditSequencer(verifier string, sequencerChecks []ExtSequencerCheck) *MsgAuditSequencer {
	return &MsgAuditSequencer{
		Verifier:        verifier,
		SequencerChecks: sequencerChecks,
	}
}

func (msg *MsgAuditSequencer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Verifier)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid verifier address (%s)", err)
	}
	return nil
}
