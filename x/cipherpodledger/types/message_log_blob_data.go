package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLogBlobData{}

func NewMsgLogBlobData(creator string, stationId string, podBundle []byte, podRange []uint64) *MsgLogBlobData {
	return &MsgLogBlobData{
		Creator:   creator,
		StationId: stationId,
		PodBundle: podBundle,
		PodRange:  podRange,
	}
}

func (msg *MsgLogBlobData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
