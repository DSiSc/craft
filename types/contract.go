package types

// define specified type of system contract
const (
	Null                   = "Null"
	JustitiaRightToken     = "JustitiaRight"
	JustitiaVoting         = "Voting"
	JustitiaWhiteList      = "WhiteList"
	JustitiaMetaData       = "MetaData"
	JustitiaCrossFundsPool = "CrossFundsPool"
	DposBftVotingContract  = "DposPbft"
	MspContract            = "MspContract"
)

type ContractType int

const (
	InitialContractType ContractType = iota
	JustitiaRightContractType
	VoteContractType
	WhiteListContractType
	MetaDataContractType
	CrossFundsPoolContractType
	MaximumContractType
)

const (
	MinimunNodesForDpos                 = uint64(4)
	MetaDataContractAddress             = "8be503bcded90ed42eff31f56199399b2b0154ca"
	JustiitaContractDefaultAddress      = "bd770416a3345f91e4b34576cb804a576fa48eb1"
	VotingContractDefaultAddress        = "5a443704dd4b594b382c22a083e2bd3090a6fef3"
	WhiteListContractTypeDefaultAddress = "47e9fbef8c83a1714f1951f142132e6e90f5fa5d"
	CrossFundsPoolDefaultAddress        = "47c5e40890bce4a473a49d7501808b9633f29782"
	DposBftVotingContractAddress        = "3fc3a28f299d1fd21312ddc2c362e92b97a5a831"
)
