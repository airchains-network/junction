package keeper

import (
	"fmt"

	"github.com/airchains-network/junction/x/cipherpodledger/types"
	"github.com/consensys/gnark/frontend"
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
