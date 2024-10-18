package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitStation{}

func NewMsgInitStation(submitter string, stationId string, stationInfo []byte, operators []string) *MsgInitStation {
	return &MsgInitStation{
		Submitter:   submitter,
		StationId:   stationId,
		StationInfo: stationInfo,
		Operators:   operators,
	}
}

func (msg *MsgInitStation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submitter address (%s)", err)
	}
	return nil
}
