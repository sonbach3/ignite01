package keeper

import (
	"context"
	"fmt"

	"capy/x/capy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAnimalSkill(goCtx context.Context, msg *types.MsgCreateAnimalSkill) (*types.MsgCreateAnimalSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var animalSkill = types.AnimalSkill{
		Creator: msg.Creator,
		Skill:   msg.Skill,
	}

	id := k.AppendAnimalSkill(
		ctx,
		animalSkill,
	)

	return &types.MsgCreateAnimalSkillResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAnimalSkill(goCtx context.Context, msg *types.MsgUpdateAnimalSkill) (*types.MsgUpdateAnimalSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var animalSkill = types.AnimalSkill{
		Creator: msg.Creator,
		Id:      msg.Id,
		Skill:   msg.Skill,
	}

	// Checks that the element exists
	val, found := k.GetAnimalSkill(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAnimalSkill(ctx, animalSkill)

	return &types.MsgUpdateAnimalSkillResponse{}, nil
}

func (k msgServer) DeleteAnimalSkill(goCtx context.Context, msg *types.MsgDeleteAnimalSkill) (*types.MsgDeleteAnimalSkillResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAnimalSkill(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAnimalSkill(ctx, msg.Id)

	return &types.MsgDeleteAnimalSkillResponse{}, nil
}
