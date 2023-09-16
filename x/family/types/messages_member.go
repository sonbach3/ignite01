package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMember = "create_member"
	TypeMsgUpdateMember = "update_member"
	TypeMsgDeleteMember = "delete_member"
)

var _ sdk.Msg = &MsgCreateMember{}

func NewMsgCreateMember(creator string, amount sdk.Coin, name string) *MsgCreateMember {
	return &MsgCreateMember{
		Creator: creator,
		Amount:  amount,
		Name:    name,
	}
}

func (msg *MsgCreateMember) Route() string {
	return RouterKey
}

func (msg *MsgCreateMember) Type() string {
	return TypeMsgCreateMember
}

func (msg *MsgCreateMember) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMember) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMember) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMember{}

func NewMsgUpdateMember(creator string, id uint64, amount sdk.Coin, name string) *MsgUpdateMember {
	return &MsgUpdateMember{
		Id:      id,
		Creator: creator,
		Amount:  amount,
		Name:    name,
	}
}

func (msg *MsgUpdateMember) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMember) Type() string {
	return TypeMsgUpdateMember
}

func (msg *MsgUpdateMember) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMember) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMember) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMember{}

func NewMsgDeleteMember(creator string, id uint64) *MsgDeleteMember {
	return &MsgDeleteMember{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMember) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMember) Type() string {
	return TypeMsgDeleteMember
}

func (msg *MsgDeleteMember) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMember) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMember) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
