package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCollege{}

func NewMsgCreateCollege(creator string, students string, name string, details string) *MsgCreateCollege {
	return &MsgCreateCollege{
		Creator:  creator,
		Students: students,
		Name:     name,
		Details:  details,
	}
}

func (msg *MsgCreateCollege) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCollege{}

func NewMsgUpdateCollege(creator string, id uint64, students string, name string, details string) *MsgUpdateCollege {
	return &MsgUpdateCollege{
		Id:       id,
		Creator:  creator,
		Students: students,
		Name:     name,
		Details:  details,
	}
}

func (msg *MsgUpdateCollege) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCollege{}

func NewMsgDeleteCollege(creator string, id uint64) *MsgDeleteCollege {
	return &MsgDeleteCollege{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteCollege) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
