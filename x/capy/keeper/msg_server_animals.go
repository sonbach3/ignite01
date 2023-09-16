package keeper

import (
	"context"
	"fmt"

	"capy/x/capy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAnimals(goCtx context.Context, msg *types.MsgCreateAnimals) (*types.MsgCreateAnimalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var animals = types.Animals{
		Creator:  msg.Creator,
		Name:     msg.Name,
		Power:    msg.Power,
		Hp:       msg.Hp,
		Location: msg.Location,
	}

	id := k.AppendAnimals(
		ctx,
		animals,
	)

	return &types.MsgCreateAnimalsResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAnimals(goCtx context.Context, msg *types.MsgUpdateAnimals) (*types.MsgUpdateAnimalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var animals = types.Animals{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Name:     msg.Name,
		Power:    msg.Power,
		Hp:       msg.Hp,
		Location: msg.Location,
	}

	// Checks that the element exists
	val, found := k.GetAnimals(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAnimals(ctx, animals)

	return &types.MsgUpdateAnimalsResponse{}, nil
}

func (k msgServer) DeleteAnimals(goCtx context.Context, msg *types.MsgDeleteAnimals) (*types.MsgDeleteAnimalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAnimals(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAnimals(ctx, msg.Id)

	return &types.MsgDeleteAnimalsResponse{}, nil
}
