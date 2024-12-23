package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) VerifyPod(goCtx context.Context, msg *types.MsgVerifyPod) (*types.MsgVerifyPodResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// submittedBy := msg.SubmittedBy
	stationId := msg.StationId
	podNumber := msg.PodNumber
	// provingNetwork := msg.ProvingNetwork
	decodedZkFheproof := msg.ZkFheproof

	fhvmMetadataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FhvmKeyStoreKey))
	fhvmMetadataKey := []byte(stationId)
	fhvmMetadataDataBytes := fhvmMetadataDB.Get(fhvmMetadataKey)
	if fhvmMetadataDataBytes == nil {
		return nil, status.Error(codes.NotFound, "stationId is not registered")
	}

	var fhvmMetadata types.FhvmsMeta
	k.cdc.MustUnmarshal(fhvmMetadataDataBytes, &fhvmMetadata)

	// check if the pod number is correct or not
	if fhvmMetadata.LatestPodNumber != podNumber {
		return nil, status.Error(codes.FailedPrecondition, "pod number is not correct")
	}

	// get the pod details
	podStorePath, podStoreKeyByte := k.GetPodKeyByte(stationId, podNumber)
	podStore := prefix.NewStore(storeAdapter, types.KeyPrefix(podStorePath))
	podDataBytes := podStore.Get(podStoreKeyByte)
	if podDataBytes == nil {
		return nil, status.Error(codes.NotFound, "pod not found")
	}

	var podData types.PodData
	k.cdc.MustUnmarshal(podDataBytes, &podData)

	// vk
	verifyingKey := groth16.NewVerifyingKey(ecc.BN254)
	_, err := verifyingKey.ReadFrom(bytes.NewReader(fhvmMetadata.ProvingNetworkVerificationKey))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to read verifying key")
	}

	// proof
	proof := groth16.NewProof(ecc.BN254)
	_, err = proof.ReadFrom(bytes.NewReader(decodedZkFheproof))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to read proof")
	}

	// witness
	var decodedPublicWitness json.RawMessage
	err = json.Unmarshal(podData.ZkFHEPublicWitness, &decodedPublicWitness)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to unmarshal public witness")
	}

	schema, err := frontend.NewSchema(&Circuit{})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create schema")
	}

	publicWitness, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create witness")
	}

	err = publicWitness.FromJSON(schema, decodedPublicWitness)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to set public witness")
	}

	// Verify the proof.
	err = groth16.Verify(proof, verifyingKey, publicWitness)
	if err != nil {
		return nil, status.Error(codes.Internal, "proof verification failed")
	}

	// update the pod data
	updatePodData := types.PodData{
		AscContractAddress: podData.AscContractAddress,
		PodNumber:          podNumber,
		DaBlobId:           podData.DaBlobId,
		SubmittedBy:        podData.SubmittedBy,
		Status:             "verified",
		Timestamp:          podData.Timestamp,
		ProvingNetwork:     podData.ProvingNetwork,
		ZkFHEProof:         decodedZkFheproof,
		ZkFHEPublicWitness: podData.ZkFHEPublicWitness,
		IsProofVerified:    true,
	}

	storingData := k.cdc.MustMarshal(&updatePodData)
	podStore.Set(podStoreKeyByte, storingData)

	// update fhvm metadata
	updateFhvmMetadata := types.FhvmsMeta{
		ChainId:                       fhvmMetadata.ChainId,
		ChainName:                     fhvmMetadata.ChainName,
		Status:                        true,
		ProofType:                     fhvmMetadata.ProofType,
		ProvingNetworkVerificationKey: fhvmMetadata.ProvingNetworkVerificationKey,
		DaProvider:                    fhvmMetadata.DaProvider,
		DaBlobId:                      fhvmMetadata.DaBlobId,
		RelayerGAddress:               fhvmMetadata.RelayerGAddress,
		RelayerAscAddress:             fhvmMetadata.RelayerAscAddress,
		PicContractAddress:            fhvmMetadata.PicContractAddress,
		AclContractAddress:            fhvmMetadata.AclContractAddress,
		TfheExecutorContractAddress:   fhvmMetadata.TfheExecutorContractAddress,
		KmsVerifierContractAddress:    fhvmMetadata.KmsVerifierContractAddress,
		GatewayContractAddress:        fhvmMetadata.GatewayContractAddress,
		AscChildContractAddress:       fhvmMetadata.AscChildContractAddress,
		LatestPodNumber:               fhvmMetadata.LatestPodNumber,
		LatestVerifiedPodNumber:       podNumber,
		FinalityPodNumber:             fhvmMetadata.FinalityPodNumber,
	}

	byteFhvmMetadata := k.cdc.MustMarshal(&updateFhvmMetadata)
	fhvmMetadataDB.Set(fhvmMetadataKey, byteFhvmMetadata)

	// update pod-verified-count
	figureDBStore := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FiguresDBPath))
	podVerifiedCountKey := fmt.Sprintf("zkFHE-pod-verified-count__%s", stationId)
	podNumberString := strconv.FormatUint(podNumber, 10)
	figureDBStore.Set([]byte(podVerifiedCountKey), []byte(podNumberString))

	return &types.MsgVerifyPodResponse{
		Status: true,
	}, nil
}
