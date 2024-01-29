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

func CmdSubmitPod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-pod [station-id] [pod-number] [merkle-root-hash] [previous-merkle-root-hash] [public-witness] [timestamp]",
		Short: "Broadcast message submit_pod",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStationId := args[0]
			argPodNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argMerkleRootHash := args[2]
			argPreviousMerkleRootHash := args[3]
			argPublicWitness := []byte(args[4])
			//argTimestamp := args[5]
			var timestampValue uint64
			timestampValue, err = strconv.ParseUint(args[5], 10, 64)
			if err != nil {
				//timestampValue = uint64(time.Now().Unix())
				return err
			}
			argTimestamp := timestampValue

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitPod(
				clientCtx.GetFromAddress().String(),
				argStationId,
				argPodNumber,
				argMerkleRootHash,
				argPreviousMerkleRootHash,
				argPublicWitness,
				argTimestamp,
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
