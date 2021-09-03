package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rizon-world/rizon/x/nft/types"
)

func (k Keeper) setOwner(ctx sdk.Context,
	denomID, tokenID string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	bz := types.MustMarshalTokenID(k.cdc, tokenID)
	store.Set(types.KeyOwner(owner, denomID, tokenID), bz)
}
