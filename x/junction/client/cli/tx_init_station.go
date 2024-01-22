package cli

import (
	"strconv"

	"github.com/airchains-network/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInitStation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-station [verification-key] [station-id] [station-info]",
		Short: "Broadcast message init_station",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVerificationKey := args[0]
			argStationId := args[1]
			argStationInfo := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitStation(
				clientCtx.GetFromAddress().String(),
				argVerificationKey,
				argStationId,
				argStationInfo,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
