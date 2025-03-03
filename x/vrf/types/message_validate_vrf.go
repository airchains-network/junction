package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgValidateVrf{}

func NewMsgValidateVrf(creator string, key uint32, collectionId string, serializedRc []byte) *MsgValidateVrf {
	return &MsgValidateVrf{
		Creator:      creator,
		Key:          key,
		CollectionId: collectionId,
		SerializedRc: serializedRc,
	}
}

func (msg *MsgValidateVrf) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
