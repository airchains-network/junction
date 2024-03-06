package keeper

import (
	"crypto/sha256"
	"strconv"
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