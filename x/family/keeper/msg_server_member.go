package keeper

import (
	"context"
	"fmt"

	"capy/x/family/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMember(goCtx context.Context, msg *types.MsgCreateMember) (*types.MsgCreateMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var member = types.Member{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Name:    msg.Name,
	}

	id := k.AppendMember(
		ctx,
		member,
	)

	return &types.MsgCreateMemberResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMember(goCtx context.Context, msg *types.MsgUpdateMember) (*types.MsgUpdateMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var member = types.Member{
		Creator: msg.Creator,
		Id:      msg.Id,
		Amount:  msg.Amount,
		Name:    msg.Name,
	}

	// Checks that the element exists
	val, found := k.GetMember(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMember(ctx, member)

	return &types.MsgUpdateMemberResponse{}, nil
}

func (k msgServer) DeleteMember(goCtx context.Context, msg *types.MsgDeleteMember) (*types.MsgDeleteMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMember(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMember(ctx, msg.Id)

	return &types.MsgDeleteMemberResponse{}, nil
}
