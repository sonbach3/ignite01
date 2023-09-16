package keeper

import (
	"context"

	"capy/x/family/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MemberAll(goCtx context.Context, req *types.QueryAllMemberRequest) (*types.QueryAllMemberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var members []types.Member
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	memberStore := prefix.NewStore(store, types.KeyPrefix(types.MemberKey))

	pageRes, err := query.Paginate(memberStore, req.Pagination, func(key []byte, value []byte) error {
		var member types.Member
		if err := k.cdc.Unmarshal(value, &member); err != nil {
			return err
		}

		members = append(members, member)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMemberResponse{Member: members, Pagination: pageRes}, nil
}

func (k Keeper) Member(goCtx context.Context, req *types.QueryGetMemberRequest) (*types.QueryGetMemberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	member, found := k.GetMember(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMemberResponse{Member: member}, nil
}
