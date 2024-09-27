package types

// StationInfoDetails represents the detailed information of a station including its name, type, and associated components.
/*
	Name: moniker of the station
	Type: evm,svm,cosmwasm etc...
	Operators: List of node operators or validator identities
*/
type StationInfoDetails struct {
	StationName      string
	Type             string
	FheEnabled       bool
	Operators        []string // List of node operators or validator identities
	SequencerDetails SequencerDetails
	DADetails        DADetails
	ProverDetails    ProverDetails
}

// SequencerDetails represents metadata about a sequencer, including its name and version.
/*
{
	Name: espresso,...
	Version: git commit version or tag version
}
*/
type SequencerDetails struct {
	Name    string
	Version string
}

// DADetails represents the Data Availability details of a blockchain component, including its name, type, etc.
/*
{
	Name: mocha4,goldberg,turing....
	Type: celestia,avail,eigen
	Version: git commit version or tag version
}
*/
type DADetails struct {
	Name    string
	Type    string
	Version string
}

// ProverDetails represents metadata about a prover, including its name and version.
/*
	Name: gevulot,sindri,....
	Version: git commit version or tag version
*/
type ProverDetails struct {
	Name    string
	Version string
}
