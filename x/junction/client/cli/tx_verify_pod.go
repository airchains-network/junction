package cli

import (
	"strconv"

	"github.com/airchains-network/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdVerifyPod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-pod [station-id] [pod-number] [merkle-root-hash] [previous-merkle-root-hash] [zk-proof]",
		Short: "Broadcast message verify_pod",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStationId := args[0]
			argPodNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argMerkleRootHash := args[2]
			argPreviousMerkleRootHash := args[3]
			argZkProof := []byte(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVerifyPod(
				clientCtx.GetFromAddress().String(),
				argStationId,
				argPodNumber,
				argMerkleRootHash,
				argPreviousMerkleRootHash,
				argZkProof,
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
