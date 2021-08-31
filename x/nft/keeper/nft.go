package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rizon-world/rizon/x/nft/exported"
	"github.com/rizon-world/rizon/x/nft/types"
)

// GetNFTs returns all NFTs by the specified denom ID
func (k Keeper) GetNFTs(ctx sdk.Context, denom string) (nfts []exported.NFT) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyNFT(denom, ""))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var baseNFT types.BaseNFT
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &baseNFT)
		nfts = append(nfts, baseNFT)
	}

	return nfts
}
