package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAnimalSkill = "create_animal_skill"
	TypeMsgUpdateAnimalSkill = "update_animal_skill"
	TypeMsgDeleteAnimalSkill = "delete_animal_skill"
)

var _ sdk.Msg = &MsgCreateAnimalSkill{}

func NewMsgCreateAnimalSkill(creator string, skill []uint64) *MsgCreateAnimalSkill {
	return &MsgCreateAnimalSkill{
		Creator: creator,
		Skill:   skill,
	}
}

func (msg *MsgCreateAnimalSkill) Route() string {
	return RouterKey
}

func (msg *MsgCreateAnimalSkill) Type() string {
	return TypeMsgCreateAnimalSkill
}

func (msg *MsgCreateAnimalSkill) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAnimalSkill) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAnimalSkill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAnimalSkill{}

func NewMsgUpdateAnimalSkill(creator string, id uint64, skill []uint64) *MsgUpdateAnimalSkill {
	return &MsgUpdateAnimalSkill{
		Id:      id,
		Creator: creator,
		Skill:   skill,
	}
}

func (msg *MsgUpdateAnimalSkill) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAnimalSkill) Type() string {
	return TypeMsgUpdateAnimalSkill
}

func (msg *MsgUpdateAnimalSkill) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAnimalSkill) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAnimalSkill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAnimalSkill{}

func NewMsgDeleteAnimalSkill(creator string, id uint64) *MsgDeleteAnimalSkill {
	return &MsgDeleteAnimalSkill{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAnimalSkill) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAnimalSkill) Type() string {
	return TypeMsgDeleteAnimalSkill
}

func (msg *MsgDeleteAnimalSkill) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAnimalSkill) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAnimalSkill) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
