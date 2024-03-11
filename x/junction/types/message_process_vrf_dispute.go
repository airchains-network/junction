package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProcessVrfDispute{}

func NewMsgProcessVrfDispute(creator string, podNumber uint64, stationId string, signatures [][]byte, votes []bool, publicKeys [][]byte) *MsgProcessVrfDispute {
	return &MsgProcessVrfDispute{
		Creator:    creator,
		PodNumber:  podNumber,
		StationId:  stationId,
		Signatures: signatures,
		Votes:      votes,
		PublicKeys: publicKeys,
	}
}

func (msg *MsgProcessVrfDispute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
