package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerifyPod{}

func NewMsgVerifyPod(creator string, stationId string, podNumber uint64, merkleRootHash string, previousMerkleRootHash string, zkProof []byte) *MsgVerifyPod {
	return &MsgVerifyPod{
		Creator:                creator,
		StationId:              stationId,
		PodNumber:              podNumber,
		MerkleRootHash:         merkleRootHash,
		PreviousMerkleRootHash: previousMerkleRootHash,
		ZkProof:                zkProof,
	}
}

func (msg *MsgVerifyPod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
