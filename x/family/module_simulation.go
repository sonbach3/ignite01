package family

import (
	"math/rand"

	"capy/testutil/sample"
	familysimulation "capy/x/family/simulation"
	"capy/x/family/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = familysimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateMember = "op_weight_msg_member"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMember int = 100

	opWeightMsgUpdateMember = "op_weight_msg_member"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMember int = 100

	opWeightMsgDeleteMember = "op_weight_msg_member"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMember int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	familyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MemberList: []types.Member{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MemberCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&familyGenesis)
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

	var weightMsgCreateMember int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMember, &weightMsgCreateMember, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMember = defaultWeightMsgCreateMember
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMember,
		familysimulation.SimulateMsgCreateMember(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMember int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMember, &weightMsgUpdateMember, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMember = defaultWeightMsgUpdateMember
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMember,
		familysimulation.SimulateMsgUpdateMember(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMember int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMember, &weightMsgDeleteMember, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMember = defaultWeightMsgDeleteMember
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMember,
		familysimulation.SimulateMsgDeleteMember(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMember,
			defaultWeightMsgCreateMember,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				familysimulation.SimulateMsgCreateMember(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMember,
			defaultWeightMsgUpdateMember,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				familysimulation.SimulateMsgUpdateMember(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteMember,
			defaultWeightMsgDeleteMember,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				familysimulation.SimulateMsgDeleteMember(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
