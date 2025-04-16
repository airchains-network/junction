package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitProver{}

func NewMsgInitProver(creator string, rollupId string, proverVerificationKey []byte, proverType string, proverEndpoint string) *MsgInitProver {
	return &MsgInitProver{
		Creator:               creator,
		RollupId:              rollupId,
		ProverVerificationKey: proverVerificationKey,
		ProverType:            proverType,
		ProverEndpoint:        proverEndpoint,
	}
}

func (msg *MsgInitProver) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
