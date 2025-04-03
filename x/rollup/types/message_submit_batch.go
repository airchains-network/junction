package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitBatch{}

func NewMsgSubmitBatch(creator string, rollupId string, batchNo uint64, merkleRootHash string, previousMerkleRootHash string, zkProof []byte, publicWitness []byte) *MsgSubmitBatch {
	return &MsgSubmitBatch{
		Creator:                creator,
		RollupId:               rollupId,
		BatchNo:                batchNo,
		MerkleRootHash:         merkleRootHash,
		PreviousMerkleRootHash: previousMerkleRootHash,
		ZkProof:                zkProof,
		PublicWitness:          publicWitness,
	}
}

func (msg *MsgSubmitBatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
