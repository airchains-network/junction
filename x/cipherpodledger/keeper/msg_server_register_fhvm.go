package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) RegisterFhvm(goCtx context.Context, msg *types.MsgRegisterFhvm) (*types.MsgRegisterFhvmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	chainId := msg.ChainId
	chainName := msg.ChainName
	proofType := msg.ProofType
	provingNetworkVerificationKey := msg.ProvingNetworkVerificationKey
	daProvider := msg.DaProvider
	daBlobId := msg.DaBlobId
	relayerGaddress := msg.RelayerGaddress
	relayerAscAddress := msg.RelayerAscAddress
	picContractAddress := msg.PicContractAddress
	aclContractAddress := msg.AclContractAddress
	tfheExecutorContractAddress := msg.TfheExecutorContractAddress
	kmsVerifierContractAddress := msg.KmsVerifierContractAddress
	gatewayContractAddress := msg.GatewayContractAddress
	ascChildContractAddress := msg.AscChildContractAddress

	// Check if the chainId is valid or not
	chainIdHash := sha256.Sum256([]byte(ascChildContractAddress))
	if chainId != hex.EncodeToString(chainIdHash[:]) {
		return nil, status.Error(codes.InvalidArgument, "chainId is not valid")
	}

	// FHVM Metadata Store
	fhvmMetadataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FhvmKeyStoreKey))
	fhvmMetadataKey := []byte(chainId)
	fhvmMetadataDataBytes := fhvmMetadataDB.Get(fhvmMetadataKey)
	if fhvmMetadataDataBytes != nil {
		return nil, status.Error(codes.AlreadyExists, "chainId is already registered")
	}

	// ASC Child Contract Registry
	ascChildContractRegistryDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AscChildContractRegistryKey))
	ascChildContractRegistryKey := []byte(ascChildContractAddress)
	ascChildContractRegistryDataBytes := ascChildContractRegistryDB.Get(ascChildContractRegistryKey)
	if ascChildContractRegistryDataBytes != nil {
		return nil, status.Error(codes.AlreadyExists, "ascChildContractAddress is already registered")
	}
	
	// 1. Register the FHVM metadata with the chainId
	var newFhvmMetadata = types.FhvmsMeta{
		ChainId: chainId,
		ChainName: chainName,
		Status: true,
		ProofType: proofType,
		ProvingNetworkVerificationKey: provingNetworkVerificationKey,
		DaProvider: daProvider,
		DaBlobId: daBlobId,
		RelayerGAddress: relayerGaddress,
		RelayerAscAddress: relayerAscAddress,
		PicContractAddress: picContractAddress,
		AclContractAddress: aclContractAddress,
		TfheExecutorContractAddress: tfheExecutorContractAddress,
		KmsVerifierContractAddress: kmsVerifierContractAddress,
		GatewayContractAddress: gatewayContractAddress,
		AscChildContractAddress: ascChildContractAddress,
		LatestVerifiedPodNumber: 0,
		FinalityPodNumber: 0,
	}

	byteFhvmMetadata := k.cdc.MustMarshal(&newFhvmMetadata)
	fhvmMetadataDB.Set(fhvmMetadataKey, byteFhvmMetadata)

	// 2. Register the ascChildContractAddress with the chainId
	ascChildContractRegistryDB.Set(ascChildContractRegistryKey, fhvmMetadataKey)

	return &types.MsgRegisterFhvmResponse{
		Status: true,
	}, nil
}
