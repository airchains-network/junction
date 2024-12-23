package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterFhvm{}

func NewMsgRegisterFhvm(creator string, chainId string, chainName string, proofType string, provingNetworkVerificationKey []byte, daProvider string, daBlobId string, relayerGaddress string, relayerAscAddress string, picContractAddress string, aclContractAddress string, tfheExecutorContractAddress string, kmsVerifierContractAddress string, gatewayContractAddress string, ascChildContractAddress string) *MsgRegisterFhvm {
	return &MsgRegisterFhvm{
		Creator:                       creator,
		ChainId:                       chainId,
		ChainName:                     chainName,
		ProofType:                     proofType,
		ProvingNetworkVerificationKey: provingNetworkVerificationKey,
		DaProvider:                    daProvider,
		DaBlobId:                      daBlobId,
		RelayerGaddress:               relayerGaddress,
		RelayerAscAddress:             relayerAscAddress,
		PicContractAddress:            picContractAddress,
		AclContractAddress:            aclContractAddress,
		TfheExecutorContractAddress:   tfheExecutorContractAddress,
		KmsVerifierContractAddress:    kmsVerifierContractAddress,
		GatewayContractAddress:        gatewayContractAddress,
		AscChildContractAddress:       ascChildContractAddress,
	}
}

func (msg *MsgRegisterFhvm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
