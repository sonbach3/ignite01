package capy

import (
	"math/rand"

	"capy/testutil/sample"
	capysimulation "capy/x/capy/simulation"
	"capy/x/capy/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = capysimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateAnimals = "op_weight_msg_animals"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAnimals int = 100

	opWeightMsgUpdateAnimals = "op_weight_msg_animals"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAnimals int = 100

	opWeightMsgDeleteAnimals = "op_weight_msg_animals"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAnimals int = 100

	opWeightMsgCreateAnimalSkill = "op_weight_msg_animal_skill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAnimalSkill int = 100

	opWeightMsgUpdateAnimalSkill = "op_weight_msg_animal_skill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAnimalSkill int = 100

	opWeightMsgDeleteAnimalSkill = "op_weight_msg_animal_skill"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAnimalSkill int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	capyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AnimalsList: []types.Animals{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AnimalsCount: 2,
		AnimalSkillList: []types.AnimalSkill{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AnimalSkillCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&capyGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAnimals int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAnimals, &weightMsgCreateAnimals, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAnimals = defaultWeightMsgCreateAnimals
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAnimals,
		capysimulation.SimulateMsgCreateAnimals(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAnimals int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAnimals, &weightMsgUpdateAnimals, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAnimals = defaultWeightMsgUpdateAnimals
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAnimals,
		capysimulation.SimulateMsgUpdateAnimals(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAnimals int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAnimals, &weightMsgDeleteAnimals, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAnimals = defaultWeightMsgDeleteAnimals
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAnimals,
		capysimulation.SimulateMsgDeleteAnimals(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateAnimalSkill int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAnimalSkill, &weightMsgCreateAnimalSkill, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAnimalSkill = defaultWeightMsgCreateAnimalSkill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAnimalSkill,
		capysimulation.SimulateMsgCreateAnimalSkill(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAnimalSkill int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAnimalSkill, &weightMsgUpdateAnimalSkill, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAnimalSkill = defaultWeightMsgUpdateAnimalSkill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAnimalSkill,
		capysimulation.SimulateMsgUpdateAnimalSkill(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAnimalSkill int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAnimalSkill, &weightMsgDeleteAnimalSkill, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAnimalSkill = defaultWeightMsgDeleteAnimalSkill
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAnimalSkill,
		capysimulation.SimulateMsgDeleteAnimalSkill(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAnimals,
			defaultWeightMsgCreateAnimals,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				capysimulation.SimulateMsgCreateAnimals(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAnimals,
			defaultWeightMsgUpdateAnimals,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				capysimulation.SimulateMsgUpdateAnimals(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAnimals,
			defaultWeightMsgDeleteAnimals,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				capysimulation.SimulateMsgDeleteAnimals(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAnimalSkill,
			defaultWeightMsgCreateAnimalSkill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				capysimulation.SimulateMsgCreateAnimalSkill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAnimalSkill,
			defaultWeightMsgUpdateAnimalSkill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				capysimulation.SimulateMsgUpdateAnimalSkill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAnimalSkill,
			defaultWeightMsgDeleteAnimalSkill,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				capysimulation.SimulateMsgDeleteAnimalSkill(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
