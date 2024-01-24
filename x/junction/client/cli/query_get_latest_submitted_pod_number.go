package cli

import (
	"strconv"

	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetLatestSubmittedPodNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-latest-submitted-pod-number [station-id]",
		Short: "Query get_latest_submitted_pod_number",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqStationId := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLatestSubmittedPodNumberRequest{

				StationId: reqStationId,
			}

			res, err := queryClient.GetLatestSubmittedPodNumber(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
