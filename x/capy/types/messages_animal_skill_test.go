package types

import (
	"testing"

	"capy/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateAnimalSkill_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAnimalSkill
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAnimalSkill{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAnimalSkill{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateAnimalSkill_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAnimalSkill
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAnimalSkill{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAnimalSkill{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteAnimalSkill_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAnimalSkill
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAnimalSkill{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAnimalSkill{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
