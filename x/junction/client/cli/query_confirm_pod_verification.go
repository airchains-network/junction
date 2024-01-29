package cli

import (
	"strconv"

	"github.com/airchains-network/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdConfirmPodVerification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-pod-verification [station-id] [pod-number] [merkle-root-hash] [previous-merkle-root-hash] [zk-proof]",
		Short: "Query confirm_pod_verification",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqStationId := args[0]
			reqPodNumber, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			reqMerkleRootHash := args[2]
			reqPreviousMerkleRootHash := args[3]
			reqZkProof := []byte(args[4])

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryConfirmPodVerificationRequest{

				StationId:              reqStationId,
				PodNumber:              reqPodNumber,
				MerkleRootHash:         reqMerkleRootHash,
				PreviousMerkleRootHash: reqPreviousMerkleRootHash,
				ZkProof:                reqZkProof,
			}

			res, err := queryClient.ConfirmPodVerification(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
