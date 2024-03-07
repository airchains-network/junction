package keeper

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

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
