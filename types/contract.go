package types

// define specified type of system contract
const (
	Null               = "Null"
	JustitiaRightToken = "JustitiaRight"
	JustitiaVoting     = "Voting"
	JustitiaWhiteList  = "WhiteList"
	JustitiaMetaData   = "MetaData"
)

type ContractType uint8

const (
	ContractInitialType       ContractType = iota
	ContractJustitiaRight
	ContractVote
	ContractWhiteList
	ContractMetaData
	ConTractMaximumType
)

const MetaDataContractAddress = "0x8be503bcded90ed42eff31f56199399b2b0154ca"