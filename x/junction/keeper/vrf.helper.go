package keeper

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"cosmossdk.io/store/prefix"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

type ExtraArg struct {
	SerializedRc []byte `json:"serializedRc"`
	Proof        []byte `json:"proof"`
	VrfOutput    []byte `json:"vrfOutput"`
}

func GetVRFKeyByte(stationId string, podNumber uint64) (string, []byte) {
	podStoreKey := "vrf/" + stationId
	podNumberString := strconv.FormatUint(podNumber, 10)
	podStoreKeyByte := []byte(podNumberString)
	return podStoreKey, podStoreKeyByte
}

func GetVRFDisputeKeyByte(stationId string, podNumber uint64) (string, []byte) {
	disputeStoreKey := "vrfDispute/" + stationId
	podNumberString := strconv.FormatUint(podNumber, 10)
	disputeStoreKeyByte := []byte(podNumberString)
	return disputeStoreKey, disputeStoreKeyByte
}

// Function to generate a deterministic random number from a proof
func GenerateDeterministicRandomNumber(proof []byte) ([]byte, error) {
	// Apply a SHA-256 hash to the proof to generate the random number
	hash := sha256.Sum256(proof)
	return hash[:], nil
}

// LoadHexPublicKey loads a public key from a hexadecimal string
func LoadHexPublicKey(hexPublicKey string) (kyber.Point, error) {
	// Decode the hexadecimal string to a byte slice
	publicKeyBytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		return nil, fmt.Errorf("error decoding public key: %w", err)
	}

	// Initialize the Kyber suite for the Edwards25519 curve
	suite := edwards25519.NewBlakeSHA256Ed25519()

	// Convert the byte slice into a Kyber point representing the public key
	publicKey := suite.Point()
	if err := publicKey.UnmarshalBinary(publicKeyBytes); err != nil {
		return nil, fmt.Errorf("error unmarshalling public key: %w", err)
	}

	return publicKey, nil
}

func VerifyVRFProof(hexPublicKey string, serializedRC []byte, proof []byte, blockNum int64, vrfOutput []byte) (bool, error) {
	suite := edwards25519.NewBlakeSHA256Ed25519()

	publicKey, err := LoadHexPublicKey(hexPublicKey)
	if err != nil {
		return false, fmt.Errorf("error loading public key: %w", err)
	}

	// Deserialize R and s from the proof
	pointSize := suite.Point().MarshalSize()
	R, s := suite.Point(), suite.Scalar()
	R.UnmarshalBinary(proof[:pointSize])
	s.SetBytes(proof[pointSize:])

	// Recompute e = H(R||data) from the proof and data
	hash := sha256.New()
	rBytes, _ := R.MarshalBinary()
	hash.Write(rBytes)
	hash.Write(serializedRC)
	e := suite.Scalar().SetBytes(hash.Sum(nil))

	// Verify the equation R == g^s * y^-e
	gs := suite.Point().Mul(s, nil) // g^s, correct usage

	// Correct calculation for y^e where 'y' is publicKey (a point) and 'e' is a scalar.
	ye := suite.Point().Mul(e, publicKey)

	yeInv := suite.Point().Neg(ye)            // -y^e, correct usage
	expectedR := suite.Point().Add(gs, yeInv) // g^s * y^-e, correct combination

	if !R.Equal(expectedR) {
		return false, fmt.Errorf("invalid VRF proof")
	}

	// Verify the VRF output matches the hash of R and data
	vrfHash := sha256.New()
	vrfHash.Write(rBytes)
	vrfHash.Write(serializedRC)
	expectedVrfOutput := vrfHash.Sum(nil)
	if !bytes.Equal(vrfOutput, expectedVrfOutput) {
		return false, fmt.Errorf("invalid VRF output")
	}

	return true, nil
}

func (k Keeper) GetVrfDisputeHelper(ctx sdk.Context, stationId string, podNumber uint64) (vrfDispute types.VrfDisputeResult, sdkErr error) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	vrfDisputeStoreKey, vrfDisputeStoreKeyByte := GetVRFKeyByte(stationId, podNumber)
	vrfDisputeStore := prefix.NewStore(storeAdapter, types.KeyPrefix(vrfDisputeStoreKey))
	vrfDisputeDetailsByte := vrfDisputeStore.Get(vrfDisputeStoreKeyByte)

	if vrfDisputeDetailsByte == nil {
		return vrfDispute, sdkerrors.ErrKeyNotFound
	}

	var vrfDetails types.VrfDisputeResult
	k.cdc.MustUnmarshal(vrfDisputeDetailsByte, &vrfDetails)

	return vrfDetails, nil
}

// signature verification
func VerifyVrfDisputeSignatures(tracks []string, signatures [][]byte, pubKeys [][]byte, message [][]byte) (votes []bool, signers []string, success bool, err error) {

	pubKeysLength := len(pubKeys)
	signaturesLength := len(signatures)
	messageLength := len(message)
	tracksLength := len(tracks)

	if pubKeysLength == 0 {
		return nil, nil, false, fmt.Errorf("no public keys")
	}
	if tracksLength == 0 {
		return nil, nil, false, fmt.Errorf("no tracks")
	}
	if signaturesLength == 0 {
		return nil, nil, false, fmt.Errorf("no signatures")
	}
	if messageLength == 0 {
		return nil, nil, false, fmt.Errorf("no message")
	}

	// make sure the length of the public keys, signatures, and message are the same
	if pubKeysLength != signaturesLength || pubKeysLength != messageLength || signaturesLength != messageLength {
		return nil, nil, false, fmt.Errorf("invalid input length")
	}

	pubKeysMap := make(map[string]bool)
	for _, key := range pubKeys {
		if pubKeysMap[string(key)] {
			return nil, nil, false, fmt.Errorf("duplicate public keys")
		}
		pubKeysMap[string(key)] = true
	}

	for i := 0; i < pubKeysLength; i++ {

		publicKeyBytes := pubKeys[i] // byte{/* your public key bytes here */}

		// Initialize a Protobuf codec for Amino
		interfaceRegistry := codecTypes.NewInterfaceRegistry()
		cryptocodec.RegisterInterfaces(interfaceRegistry)
		marshaler := codec.NewProtoCodec(interfaceRegistry)

		// Attempt to unmarshal the public key bytes into a PubKey interface
		var pubKey cryptotypes.PubKey
		err := marshaler.UnmarshalInterface(publicKeyBytes, &pubKey)
		if err != nil {
			return nil, nil, false, fmt.Errorf("error unmarshalling public key: %w", err)
		}

		// check if pubKey.Address() exists in stationsAddresses
		isExists := false
		for _, address := range tracks {
			if address == pubKey.Address().String() {
				isExists = true
				break
			}
		}
		if !isExists {
			return nil, nil, false, fmt.Errorf("public key address does not exist in stationsAddresses: %s", pubKey.Address().String())
		}

		verificationResult := pubKey.VerifySignature(message[i], signatures[i])
		if !verificationResult {
			return nil, nil, false, fmt.Errorf("signature verification failed")
		}

		vote := byteToBool(message[i][0])

		votes = append(votes, vote)
		signers = append(signers, pubKey.Address().String())
	}

	return votes, signers, true, nil
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func byteToBool(b byte) bool {
	return b != 0
}

// count votes in the dispute
func countDisputeVotes(votes []bool, addr []string) (voteResults types.VrfDisputeResult) {

	totalVotes := len(votes)
	consentVote := 0
	dissentVote := 0
	agreementPercentage := 0.0
	result := false
	voteMsg := ""

	for _, vote := range votes {
		if vote {
			consentVote++
		} else {
			dissentVote++
		}
	}

	agreementPercentage = (float64(consentVote) / float64(totalVotes)) * 100
	agreementPercentage = float64(int(agreementPercentage*10000)) / 10000

	if consentVote > (66*totalVotes)/100 {
		result = true
	}

	if result {
		voteMsg = "VRN is invalid and the dispute result is inclined in favor of the verifier"
	} else {
		voteMsg = "VRN is valid and the dispute result is inclined in favor of the creator"
	}

	voteResults = types.VrfDisputeResult{
		Votes:               votes,
		AddressList:         addr,
		ConsentVote:         uint64(consentVote),
		DissentVote:         uint64(dissentVote),
		AgreementPercentage: float32(agreementPercentage),
		Result:              result,
		Message:             voteMsg,
	}

	return voteResults
}
