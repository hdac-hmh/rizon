package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/nft/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      codec.Marshaler
}

// NewKeeper creates a new instance of the NFT Keeper
func NewKeeper(cdc codec.Marshaler, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,

		cdc: cdc,
	}
}

// IssueDenom issues a denom according to the given params
func (k Keeper) IssueDenom(ctx sdk.Context,
	id, name, schema, symbol string,
	creator sdk.AccAddress,
	mintRestricted, updateRestricted bool,
) error {
	return k.SetDenom(ctx, types.NewDenom(id, name, schema, symbol, creator, mintRestricted, updateRestricted))
}
