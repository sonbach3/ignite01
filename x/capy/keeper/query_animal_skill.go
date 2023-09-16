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

func (k Keeper) AnimalSkillAll(goCtx context.Context, req *types.QueryAllAnimalSkillRequest) (*types.QueryAllAnimalSkillResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var animalSkills []types.AnimalSkill
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	animalSkillStore := prefix.NewStore(store, types.KeyPrefix(types.AnimalSkillKey))

	pageRes, err := query.Paginate(animalSkillStore, req.Pagination, func(key []byte, value []byte) error {
		var animalSkill types.AnimalSkill
		if err := k.cdc.Unmarshal(value, &animalSkill); err != nil {
			return err
		}

		animalSkills = append(animalSkills, animalSkill)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAnimalSkillResponse{AnimalSkill: animalSkills, Pagination: pageRes}, nil
}

func (k Keeper) AnimalSkill(goCtx context.Context, req *types.QueryGetAnimalSkillRequest) (*types.QueryGetAnimalSkillResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	animalSkill, found := k.GetAnimalSkill(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAnimalSkillResponse{AnimalSkill: animalSkill}, nil
}
