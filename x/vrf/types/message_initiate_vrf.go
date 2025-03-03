package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitiateVrf{}

func NewMsgInitiateVrf(creator string, key uint32, collectionId string, upperBound uint64, serializedRc []byte, proof []byte, originDigest []byte) *MsgInitiateVrf {
	return &MsgInitiateVrf{
		Creator:      creator,
		Key:          key,
		CollectionId: collectionId,
		UpperBound:   upperBound,
		SerializedRc: serializedRc,
		Proof:        proof,
		OriginDigest: originDigest,
	}
}

func (msg *MsgInitiateVrf) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
