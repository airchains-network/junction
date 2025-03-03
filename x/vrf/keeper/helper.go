package keeper

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	sdktypes "cosmossdk.io/store/types"
	"github.com/airchains-network/junction/x/vrf/types"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

// GenerateDeterministicRandomNumber generates a deterministic random number from a proof
func GenerateDeterministicRandomNumber(proof []byte) ([]byte, error) {
	// Apply an SHA-256 hash to the proof to generate the random number
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

func VerifyVRFProof(hexPublicKey string, serializedRC []byte, proof []byte, vrfOutput []byte) (bool, error) {
	suite := edwards25519.NewBlakeSHA256Ed25519()

	publicKey, err := LoadHexPublicKey(hexPublicKey)
	if err != nil {
		return false, fmt.Errorf("error loading public key: %w", err)
	}

	// Deserialize R and s from the proof
	pointSize := suite.Point().MarshalSize()
	R, s := suite.Point(), suite.Scalar()
	if err := R.UnmarshalBinary(proof[:pointSize]); err != nil {
		return false, fmt.Errorf("error unmarshalling R: %w", err)
	}
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

// cleanupOldVrfRecords removes the oldest VRF records if the total count exceeds 100,
// and updates the collection's offset accordingly.
func (k msgServer) cleanupOldVrfRecords(vrfStore sdktypes.KVStore, collection *types.Collection, collectionKey []byte, collectionStore sdktypes.KVStore, latestKey uint32) {
	// Calculate the total number of records using the sequential nature of keys.
	recordCount := latestKey - collection.Offset + 1
	if recordCount <= 100 {
		return
	}

	// Calculate the new offset to retain only the latest 100 records.
	newOffset := latestKey - 100 + 1

	// Delete records that fall before the new offset.
	for key := collection.Offset; key < newOffset; key++ {
		keyBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(keyBytes, key)
		vrfStore.Delete(keyBytes)
	}

	// Update the collection's offset.
	collection.Offset = newOffset
	updatedCollectionBytes := k.cdc.MustMarshal(collection)
	collectionStore.Set(collectionKey, updatedCollectionBytes)

}

// isCollectionMember checks if the given owner is a member of the collection.
func isCollectionMember(members []string, owner string) bool {
	for _, member := range members {
		if member == owner {
			return true
		}
	}
	return false
}

func VerifyVrfDisputeSignatures(members []string, signatures [][]byte, pubKeys [][]byte, votesArr []bool, collection types.Collection) (votes []bool, signers []string, success bool, err error) {

	pubKeysLength := len(pubKeys)
	signaturesLength := len(signatures)
	messageLength := len(votesArr)
	membersLength := len(members)

	if pubKeysLength == 0 {
		return nil, nil, false, fmt.Errorf("no public keys")
	}
	if membersLength == 0 {
		return nil, nil, false, fmt.Errorf("no members")
	}
	if signaturesLength == 0 {
		return nil, nil, false, fmt.Errorf("no signatures")
	}
	if messageLength == 0 {
		return nil, nil, false, fmt.Errorf("no votesArr")
	}

	// make sure the length of the public keys, signatures, and votesArr are the same
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

		pubKey := secp256k1.PubKey(publicKeyBytes)
		var message string
		if votesArr[i] {
			//	if votesArr[i]  is true we are going to put message value as collection's dispute vote yes hash
			message = collection.DisputeVoteYesHash
		} else {
			//	if votesArr[i]  is false we are going to put message value as collection's dispute vote no hash
			message = collection.DisputeVoteNoHash
		}
		msgHash := sha256.Sum256([]byte(message))
		verificationResult := pubKey.VerifySignature(msgHash[:], signatures[i])
		if !verificationResult {
			return nil, nil, false, fmt.Errorf("signature verification failed")
		}

		votes = append(votes, votesArr[i])
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
		Members:             addr,
		ConsentVote:         uint32(consentVote),
		DissentVote:         uint32(dissentVote),
		AgreementPercentage: float32(agreementPercentage),
		Result:              result,
		Message:             voteMsg,
	}

	return voteResults
}
