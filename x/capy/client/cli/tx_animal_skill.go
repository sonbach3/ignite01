package cli

import (
	"strconv"

	"capy/x/capy/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"strings"
)

func CmdCreateAnimalSkill() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-animal-skill [skill]",
		Short: "Create a new animalSkill",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCastSkill := strings.Split(args[0], listSeparator)
			argSkill := make([]uint64, len(argCastSkill))
			for i, arg := range argCastSkill {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argSkill[i] = value
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAnimalSkill(clientCtx.GetFromAddress().String(), argSkill)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAnimalSkill() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-animal-skill [id] [skill]",
		Short: "Update a animalSkill",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argCastSkill := strings.Split(args[1], listSeparator)
			argSkill := make([]uint64, len(argCastSkill))
			for i, arg := range argCastSkill {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argSkill[i] = value
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAnimalSkill(clientCtx.GetFromAddress().String(), id, argSkill)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAnimalSkill() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-animal-skill [id]",
		Short: "Delete a animalSkill by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAnimalSkill(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
