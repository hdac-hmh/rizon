package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rizon-world/rizon/x/nft/types"
)

func (k Keeper) deleteOwner(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyOwner(owner, denomID, tokenID))
}

// TODO:
// 1. Question?
// owner를 key, tokenID를 value로 했을때, 하나의 tokenID에 owner가 여러명이 될 수 있는 구조가 될 수 있어 보임.
// tokenID를 key, owner address를 value로 하지 않는 이유는?
func (k Keeper) setOwner(ctx sdk.Context,
	denomID, tokenID string,
	owner sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)

	bz := types.MustMarshalTokenID(k.cdc, tokenID)
	store.Set(types.KeyOwner(owner, denomID, tokenID), bz)
}

func (k Keeper) swapOwner(ctx sdk.Context, denomID, tokenID string, srcOwner, dstOwner sdk.AccAddress) {

	// delete old owner key
	k.deleteOwner(ctx, denomID, tokenID, srcOwner)

	// set new owner key
	k.setOwner(ctx, denomID, tokenID, dstOwner)
}
