package keeper

import (
	"context"

	"capy/x/capy/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AnimalsAll(goCtx context.Context, req *types.QueryAllAnimalsRequest) (*types.QueryAllAnimalsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var animalss []types.Animals
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	animalsStore := prefix.NewStore(store, types.KeyPrefix(types.AnimalsKey))

	pageRes, err := query.Paginate(animalsStore, req.Pagination, func(key []byte, value []byte) error {
		var animals types.Animals
		if err := k.cdc.Unmarshal(value, &animals); err != nil {
			return err
		}

		animalss = append(animalss, animals)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAnimalsResponse{Animals: animalss, Pagination: pageRes}, nil
}

func (k Keeper) Animals(goCtx context.Context, req *types.QueryGetAnimalsRequest) (*types.QueryGetAnimalsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	animals, found := k.GetAnimals(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAnimalsResponse{Animals: animals}, nil
}
