package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitStation = "init_station"

var _ sdk.Msg = &MsgInitStation{}

func NewMsgInitStation(creator string, verificationKey []byte, stationId string, stationInfo string) *MsgInitStation {
	return &MsgInitStation{
		Creator:         creator,
		VerificationKey: verificationKey,
		StationId:       stationId,
		StationInfo:     stationInfo,
	}
}

func (msg *MsgInitStation) Route() string {
	return RouterKey
}

func (msg *MsgInitStation) Type() string {
	return TypeMsgInitStation
}

func (msg *MsgInitStation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitStation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitStation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
