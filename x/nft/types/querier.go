package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QuerySupply			= "supply"
	QueryNFT            = "nft"
)

// QuerySupplyParams defines the params for queries:
type QuerySupplyParams struct {
	Denom string
	Owner sdk.AccAddress
}

// NewQuerySupplyParams creates a new instance of QuerySupplyParams
func NewQuerySupplyParams(denom string, owner sdk.AccAddress) QuerySupplyParams {
	return QuerySupplyParams{
		Denom: denom,
		Owner: owner,
	}
}

// QueryNFTParams params for query 'custom/nfts/nft'
type QueryNFTParams struct {
	Denom   string
	TokenID string
}

// NewQueryNFTParams creates a new instance of QueryNFTParams
func NewQueryNFTParams(denom, id string) QueryNFTParams {
	return QueryNFTParams{
		Denom:   denom,
		TokenID: id,
	}
}
