package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterCollection{}

func NewMsgRegisterCollection(collectionOwner string, collectionName string, collectionId string, members []string, offset uint32) *MsgRegisterCollection {
	return &MsgRegisterCollection{
		CollectionOwner: collectionOwner,
		CollectionName:  collectionName,
		CollectionId:    collectionId,
		Members:         members,
		Offset:          offset,
	}
}

func (msg *MsgRegisterCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.CollectionOwner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid collectionOwner address (%s)", err)
	}
	return nil
}
