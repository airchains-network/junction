package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddNewTrack{}

func NewMsgAddNewTrack(creator string, stationId string, newTrackAddress string, newTrackVotingPower string, signatures []string, votes []string, publicKeys []string) *MsgAddNewTrack {
	return &MsgAddNewTrack{
		Creator:             creator,
		StationId:           stationId,
		NewTrackAddress:     newTrackAddress,
		NewTrackVotingPower: newTrackVotingPower,
		Signatures:          signatures,
		Votes:               votes,
		PublicKeys:          publicKeys,
	}
}

func (msg *MsgAddNewTrack) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
