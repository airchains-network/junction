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
	NameSpace: this is a unique key for external sequencers like espresso to distinguish rollups
	Address: address through which the data is sent to the sequencer(espresso)
}
*/
type SequencerDetails struct {
	Name      string
	Version   string
	NameSpace string
	Address   string
}

// DADetails represents the Data Availability details of a blockchain component, including its name, type, etc.
/*
{
	Name: mocha4,goldberg,turing....
	Type: celestia,avail,eigen
	Version: git commit version or tag version
	Address: address through which the da blob has been submitted
}
*/
type DADetails struct {
	Name    string
	Type    string
	Version string
	Address string
}

// ProverDetails represents metadata about a prover, including its name and version.
/*
	Name: gevulot,sindri,....
	Version: git commit version or tag version
	Identifier: this can be anything key or address it depends on the proving network if the proving network is a
	blockchain it most probably be an address else it can be a key through which we can get the status of that particular rollup
*/
type ProverDetails struct {
	Name       string
	Version    string
	Identifier string
}

// SchemaDef represents a structure definition
type SchemaDef struct {
	Fields map[string]interface{} // Changed to interface{} to allow nested definitions
}

type EngagementSequencerDetails struct {
	NameSpace string
	Address   string
}
