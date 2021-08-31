package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	QuerySupply			= "supply"
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