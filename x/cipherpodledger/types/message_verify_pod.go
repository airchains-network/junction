package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifyPod{}

func NewMsgVerifyPod(submittedBy string, stationId string, podNumber uint64, provingNetwork string, zkFheproof []byte) *MsgVerifyPod {
	return &MsgVerifyPod{
		SubmittedBy:    submittedBy,
		StationId:      stationId,
		PodNumber:      podNumber,
		ProvingNetwork: provingNetwork,
		ZkFheproof:     zkFheproof,
	}
}

func (msg *MsgVerifyPod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SubmittedBy)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submittedBy address (%s)", err)
	}
	return nil
}
