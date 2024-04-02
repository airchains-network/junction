package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveTrack{}

func NewMsgRemoveTrack(creator string, stationId string, trackAddress string, signatures [][]byte, votes []bool, publicKeys [][]byte) *MsgRemoveTrack {
	return &MsgRemoveTrack{
		Creator:      creator,
		StationId:    stationId,
		TrackAddress: trackAddress,
		Signatures:   signatures,
		Votes:        votes,
		PublicKeys:   publicKeys,
	}
}

func (msg *MsgRemoveTrack) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
