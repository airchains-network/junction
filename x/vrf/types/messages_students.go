package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStudents{}

func NewMsgCreateStudents(creator string, name string, details string) *MsgCreateStudents {
	return &MsgCreateStudents{
		Creator: creator,
		Name:    name,
		Details: details,
	}
}

func (msg *MsgCreateStudents) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStudents{}

func NewMsgUpdateStudents(creator string, id uint64, name string, details string) *MsgUpdateStudents {
	return &MsgUpdateStudents{
		Id:      id,
		Creator: creator,
		Name:    name,
		Details: details,
	}
}

func (msg *MsgUpdateStudents) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStudents{}

func NewMsgDeleteStudents(creator string, id uint64) *MsgDeleteStudents {
	return &MsgDeleteStudents{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteStudents) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
