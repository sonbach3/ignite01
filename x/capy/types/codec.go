package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAnimals{}, "capy/CreateAnimals", nil)
	cdc.RegisterConcrete(&MsgUpdateAnimals{}, "capy/UpdateAnimals", nil)
	cdc.RegisterConcrete(&MsgDeleteAnimals{}, "capy/DeleteAnimals", nil)
	cdc.RegisterConcrete(&MsgCreateAnimalSkill{}, "capy/CreateAnimalSkill", nil)
	cdc.RegisterConcrete(&MsgUpdateAnimalSkill{}, "capy/UpdateAnimalSkill", nil)
	cdc.RegisterConcrete(&MsgDeleteAnimalSkill{}, "capy/DeleteAnimalSkill", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAnimals{},
		&MsgUpdateAnimals{},
		&MsgDeleteAnimals{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAnimalSkill{},
		&MsgUpdateAnimalSkill{},
		&MsgDeleteAnimalSkill{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
