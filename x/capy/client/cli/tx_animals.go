package cli

import (
	"strconv"

	"capy/x/capy/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateAnimals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-animals [name] [power] [hp] [location]",
		Short: "Create a new animals",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argPower, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argHp, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argLocation := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAnimals(clientCtx.GetFromAddress().String(), argName, argPower, argHp, argLocation)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAnimals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-animals [id] [name] [power] [hp] [location]",
		Short: "Update a animals",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argName := args[1]

			argPower, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			argHp, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argLocation := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAnimals(clientCtx.GetFromAddress().String(), id, argName, argPower, argHp, argLocation)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAnimals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-animals [id]",
		Short: "Delete a animals by id",
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

			msg := types.NewMsgDeleteAnimals(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
