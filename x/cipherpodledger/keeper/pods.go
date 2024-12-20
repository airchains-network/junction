package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/consensys/gnark/frontend"
	"github.com/celestiaorg/celestia-openrpc/types/blob"
)

type Circuit struct {
	Original   [100]frontend.Variable
	Compressed [200]frontend.Variable `gnark:",public"`
}

func (circuit *Circuit) Define(api frontend.API) error {

	symbol := circuit.Compressed[0]
	multiplicity := circuit.Compressed[1]
	multiplicity = api.Sub(multiplicity, frontend.Variable(48))
	y_left := circuit.Compressed

	for i := 0; i < 100; i++ {
		//ensure equality at i-th position
		api.AssertIsEqual(circuit.Original[i], symbol)

		// decrement multiplicity counter
		multiplicity = api.Sub(multiplicity, 1)

		// if counter is at zero, chomp two from compressed list
		for i := 0; i < 198; i++ {
			y_left[i] = api.Select(api.IsZero(multiplicity), y_left[i+2], y_left[i])
		}
		y_left[198] = api.Select(api.IsZero(multiplicity), frontend.Variable(0), y_left[198])
		y_left[199] = api.Select(api.IsZero(multiplicity), frontend.Variable(0), y_left[199])

		asciishift := api.Select(api.IsZero(y_left[1]), y_left[1], api.Sub(y_left[1], frontend.Variable(48)))
		multiplicity = api.Select(api.IsZero(multiplicity), asciishift, multiplicity)

		symbol = y_left[0]
	}

	return nil
}

func (k Keeper) GetPodKeyByte(stationId string, podNumber uint64) (string, []byte) {
	return fmt.Sprintf("%s/%s", types.PodDataStoreKey, stationId), []byte(fmt.Sprintf("%d", podNumber))
}

// DA related pod methods

// Celestia

type CelestiaDABlobSpace struct {
	BlobData []byte
	Commitment []byte
	NamespaceID []byte
	Height uint64
	stationId string
	podRange []uint64
}

type CelestiaPodBundle struct {
	BlobData []byte
	Commitment []byte
	NamespaceID []byte
	Height uint64
}

func (k Keeper) DecodeCelestiaPodBundle(podBundle []byte) (CelestiaPodBundle,*blob.Blob, error) {
	var decodedCelestiaPodBundle CelestiaPodBundle
	err := json.Unmarshal(podBundle, &decodedCelestiaPodBundle)
	if err != nil {
		return CelestiaPodBundle{}, nil, err
	}
	
	var decodedBlob blob.Blob
	err = json.Unmarshal(decodedCelestiaPodBundle.BlobData, &decodedBlob)
	if err != nil {
		return CelestiaPodBundle{}, nil, err
	}

	return decodedCelestiaPodBundle, &decodedBlob, nil
}

// Avail
type AvailDABlobSpace struct {
	BlobData []byte
	Commitment []byte
	stationId string
	podRange []uint64
}

type AvailPodBundle struct {
	BlobData []byte
	Commitment []byte
}

func (k Keeper) DecodeAvailPodBundle(podBundle []byte) (AvailPodBundle, error) {
	var decodedAvailPodBundle AvailPodBundle
	err := json.Unmarshal(podBundle, &decodedAvailPodBundle)
	if err != nil {
		return AvailPodBundle{}, err
	}
	return decodedAvailPodBundle, nil
}

// Eigen
type EigenDABlobSpace struct {
	BlobData []byte
	Commitment []byte
	stationId string
	podRange []uint64
}

type EigenPodBundle struct {
	BlobData []byte
	Commitment []byte
}

func (k Keeper) DecodeEigenPodBundle(podBundle []byte) (EigenPodBundle, error) {
	var decodedEigenPodBundle EigenPodBundle
	err := json.Unmarshal(podBundle, &decodedEigenPodBundle)
	if err != nil {
		return EigenPodBundle{}, err
	}
	return decodedEigenPodBundle, nil
}