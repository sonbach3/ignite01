package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAnimals = "create_animals"
	TypeMsgUpdateAnimals = "update_animals"
	TypeMsgDeleteAnimals = "delete_animals"
)

var _ sdk.Msg = &MsgCreateAnimals{}

func NewMsgCreateAnimals(creator string, name string, power uint64, hp uint64, location string) *MsgCreateAnimals {
	return &MsgCreateAnimals{
		Creator:  creator,
		Name:     name,
		Power:    power,
		Hp:       hp,
		Location: location,
	}
}

func (msg *MsgCreateAnimals) Route() string {
	return RouterKey
}

func (msg *MsgCreateAnimals) Type() string {
	return TypeMsgCreateAnimals
}

func (msg *MsgCreateAnimals) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAnimals) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAnimals) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAnimals{}

func NewMsgUpdateAnimals(creator string, id uint64, name string, power uint64, hp uint64, location string) *MsgUpdateAnimals {
	return &MsgUpdateAnimals{
		Id:       id,
		Creator:  creator,
		Name:     name,
		Power:    power,
		Hp:       hp,
		Location: location,
	}
}

func (msg *MsgUpdateAnimals) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAnimals) Type() string {
	return TypeMsgUpdateAnimals
}

func (msg *MsgUpdateAnimals) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAnimals) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAnimals) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAnimals{}

func NewMsgDeleteAnimals(creator string, id uint64) *MsgDeleteAnimals {
	return &MsgDeleteAnimals{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAnimals) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAnimals) Type() string {
	return TypeMsgDeleteAnimals
}

func (msg *MsgDeleteAnimals) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAnimals) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAnimals) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
