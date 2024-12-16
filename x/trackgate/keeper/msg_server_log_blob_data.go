package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
LogBlobData handles the logging of blob data from a DA (Data Availability) relayer.
It processes the incoming blob data and stores it in the state.

Parameters:
  - msg *types.MsgLogBlobData: The message containing the blob data to be logged.

Returns:
  - *types.MsgLogBlobDataResponse: A response indicating the success of the operation.
  - error: An error if any occurred during the processing.

Internal Parameters Extracted from msg.DaData:
  - stationRefID: The reference ID of the station that sent the blob data.
    This typically includes identifiers like namespace (e.g., "eigen") and appID (e.g., "avail").
  - daType: The type of DA used, such as "EigenLayer" or "Avail".
  - height: The block height at which the blob data was created.
  - blob: The actual blob data to be logged.

The function currently extracts these parameters from the message and is expected to perform
further processing and storage operations in the state. These operations are yet to be
implemented (as indicated by `// TODO: Handling the message`).

Example:
Given a msg.DaData with:
  - stationRefID = "eigen/avail"
  - daType = "EigenLayer"
  - height = 12345
  - blob containing some data

This function should log these details in the state.
*/
func (k msgServer) LogBlobData(goCtx context.Context, msg *types.MsgLogBlobData) (*types.MsgLogBlobDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	daData := msg.DaData
	stationRefID := daData.StationRefId
	daType := daData.DaType
	height := daData.Height
	blob := daData.Blob

	return &types.MsgLogBlobDataResponse{}, nil
}
