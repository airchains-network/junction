package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitPod{}

func NewMsgSubmitPod(submittedBy string, ascContractAddress string, podNumber uint64, stationId string, daBlobId string, timestamp int32, provingNetwork string, zkFheproof []byte, zkFhepublicWitness []byte) *MsgSubmitPod {
	return &MsgSubmitPod{
		SubmittedBy:        submittedBy,
		AscChildContractAddress: ascContractAddress,
		PodNumber:          podNumber,
		StationId:          stationId,
		DaBlobId:           daBlobId,
		Timestamp:          timestamp,
		ProvingNetwork:     provingNetwork,
		ZkFheproof:         zkFheproof,
		ZkFhepublicWitness: zkFhepublicWitness,
	}
}

func (msg *MsgSubmitPod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SubmittedBy)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submittedBy address (%s)", err)
	}
	return nil
}
