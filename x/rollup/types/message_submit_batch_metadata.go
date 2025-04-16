package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitBatchMetadata{}

func NewMsgSubmitBatchMetadata(creator string, batchNo uint64, rollupId string, daName string, daCommitment string, daHash string, daPointer string, daNamespace string) *MsgSubmitBatchMetadata {
	return &MsgSubmitBatchMetadata{
		Creator:      creator,
		BatchNo:      batchNo,
		RollupId:     rollupId,
		DaName:       daName,
		DaCommitment: daCommitment,
		DaHash:       daHash,
		DaPointer:    daPointer,
		DaNamespace:  daNamespace,
	}
}

func (msg *MsgSubmitBatchMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
