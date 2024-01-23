package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitPod = "submit_pod"

var _ sdk.Msg = &MsgSubmitPod{}

func NewMsgSubmitPod(creator string, stationId string, podNumber uint64, merkleRootHash string, previousMerkleRootHash string, publicWitness []byte, timestamp uint64) *MsgSubmitPod {
	return &MsgSubmitPod{
		Creator:                creator,
		StationId:              stationId,
		PodNumber:              podNumber,
		MerkleRootHash:         merkleRootHash,
		PreviousMerkleRootHash: previousMerkleRootHash,
		PublicWitness:          publicWitness,
		Timestamp:              timestamp,
	}
}

func (msg *MsgSubmitPod) Route() string {
	return RouterKey
}

func (msg *MsgSubmitPod) Type() string {
	return TypeMsgSubmitPod
}

func (msg *MsgSubmitPod) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitPod) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitPod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
