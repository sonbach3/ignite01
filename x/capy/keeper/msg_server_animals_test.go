package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"capy/x/capy/types"
)

func TestAnimalsMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAnimals(ctx, &types.MsgCreateAnimals{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAnimalsMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateAnimals
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAnimals{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAnimals{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAnimals{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateAnimals(ctx, &types.MsgCreateAnimals{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAnimals(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAnimalsMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteAnimals
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAnimals{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAnimals{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAnimals{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateAnimals(ctx, &types.MsgCreateAnimals{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAnimals(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
