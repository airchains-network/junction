package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgVerifyPod = "verify_pod"

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

func (msg *MsgVerifyPod) Route() string {
	return RouterKey
}

func (msg *MsgVerifyPod) Type() string {
	return TypeMsgVerifyPod
}

func (msg *MsgVerifyPod) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVerifyPod) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVerifyPod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		//return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
		invalidAddressErrorMessage := errors.New("invalid creator address")
		return invalidAddressErrorMessage
	}
	return nil
}
