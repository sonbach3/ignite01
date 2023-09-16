package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "capy/testutil/keeper"
	"capy/testutil/nullify"
	"capy/x/capy/types"
)

func TestAnimalsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAnimals(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAnimalsRequest
		response *types.QueryGetAnimalsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAnimalsRequest{Id: msgs[0].Id},
			response: &types.QueryGetAnimalsResponse{Animals: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetAnimalsRequest{Id: msgs[1].Id},
			response: &types.QueryGetAnimalsResponse{Animals: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetAnimalsRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Animals(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAnimalsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CapyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAnimals(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAnimalsRequest {
		return &types.QueryAllAnimalsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AnimalsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Animals), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Animals),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AnimalsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Animals), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Animals),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AnimalsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Animals),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AnimalsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
