package cli

import (
	"strconv"

	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetLatestVerifiedPodNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-latest-verified-pod-number [station-id]",
		Short: "Query get_latest_verified_pod_number",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqStationId := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLatestVerifiedPodNumberRequest{

				StationId: reqStationId,
			}

			res, err := queryClient.GetLatestVerifiedPodNumber(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
