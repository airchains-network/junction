package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitRollup{}

func NewMsgInitRollup(creator string, moniker string, chainId string, denomName string, keys []string, supply []uint64, daType string, aclContractAddress string, kmsVerifierAddress string, tfheExecutorAddress string, gatewayContractAddress string, ascContractAddress string, relayerGAddress string, relayerASCAddress string) *MsgInitRollup {
	return &MsgInitRollup{
		Creator:                creator,
		Moniker:                moniker,
		ChainId:                chainId,
		DenomName:              denomName,
		DaType:                 daType,
		Keys:                   keys,
		Supply:                 supply,
		AclContractAddress:     aclContractAddress,
		KmsVerifierAddress:     kmsVerifierAddress,
		TfheExecutorAddress:    tfheExecutorAddress,
		GatewayContractAddress: gatewayContractAddress,
		AscContractAddress:     ascContractAddress,
		RelayerGAddress:        relayerGAddress,
		RelayerASCAddress:      relayerASCAddress,
	}
}

func (msg *MsgInitRollup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
