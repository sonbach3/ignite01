package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"capy/x/capy/types"
)

func TestAnimalSkillMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateAnimalSkill(ctx, &types.MsgCreateAnimalSkill{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestAnimalSkillMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateAnimalSkill
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAnimalSkill{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAnimalSkill{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAnimalSkill{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateAnimalSkill(ctx, &types.MsgCreateAnimalSkill{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateAnimalSkill(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAnimalSkillMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteAnimalSkill
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAnimalSkill{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAnimalSkill{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteAnimalSkill{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateAnimalSkill(ctx, &types.MsgCreateAnimalSkill{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAnimalSkill(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
