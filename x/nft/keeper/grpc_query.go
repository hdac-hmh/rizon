package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rizon-world/rizon/x/nft/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Supply queries the total supply of a given denom or owner
func (k Keeper) Supply(c context.Context, request *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var supply uint64
	switch {
	case len(request.Owner) == 0 && len(request.DenomId) > 0:
		supply = k.GetTotalSupply(ctx, request.DenomId)
	default:
		owner, err := sdk.AccAddressFromBech32(request.Owner)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid owner address %s", request.Owner)
		}
		supply = k.GetTotalSupplyOfOwner(ctx, request.DenomId, owner)
	}
	return &types.QuerySupplyResponse{Amount: supply}, nil
}

// Owner queries the NFTs of the specified owner
func (k Keeper) Owner(c context.Context, request *types.QueryOwnerRequest) (*types.QueryOwnerResponse, error) {
	return &types.QueryOwnerResponse{}, nil
}

// Collection queries the NFTs of the specified denom
func (k Keeper) Collection(c context.Context, request *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	return &types.QueryCollectionResponse{}, nil
}

// Denom queries the definition of a given denom
func (k Keeper) Denom(c context.Context, request *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	return &types.QueryDenomResponse{}, nil
}

// Denoms queries all the denoms
func (k Keeper) Denoms(c context.Context, request *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	return &types.QueryDenomsResponse{}, nil
}

// NFT queries the NFT for the given denom and token ID
func (k Keeper) NFT(c context.Context, request *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	return &types.QueryNFTResponse{}, nil
}
