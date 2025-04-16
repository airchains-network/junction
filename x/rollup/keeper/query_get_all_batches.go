package keeper

import (
	"context"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/airchains-network/junction/x/rollup/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAllBatches(goCtx context.Context, req *types.QueryGetAllBatchesRequest) (*types.QueryGetAllBatchesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	// Initialize pagination parameters
	var offset uint64
	var limit uint64
	var order string

	xOffset := req.GetXOffset()
	if xOffset != nil {
		offset = req.GetOffset()
	} else {
		offset = 0
	}

	xLimit := req.GetXLimit()
	if xLimit != nil {
		limit = req.GetLimit()
	} else {
		limit = 10
	}

	xOrder := req.GetXOrder()
	if xOrder != nil {
		order = req.GetOrder()
	} else {
		order = "asc"
	}

	// Get all batches from the store
	batchDbPath, _ := k.GetRollupBatchDbStoreKeys(req.RollupId, 1)
	batchDataStore := prefix.NewStore(storeAdapter, types.KeyPrefix(batchDbPath))

	var batches []types.Batch
	var total uint64

	// Count total items
	iterator := batchDataStore.Iterator(nil, nil)
	for ; iterator.Valid(); iterator.Next() {
		total++
	}
	iterator.Close()

	// Handle pagination bounds
	if offset >= total {
		return &types.QueryGetAllBatchesResponse{
			Batch:  []types.Batch{},
			Total:  total,
			Offset: offset,
			Limit:  limit,
			Order:  order,
		}, nil
	}

	// Adjust limit if it exceeds the total count
	if offset+limit > total {
		limit = total - offset
	}

	// Collect batches with proper ordering
	if order == "asc" {
		var loopError error
		for i := offset; i < offset+limit && i < total; i++ {
			key := strconv.FormatUint(i+1, 10)
			value := batchDataStore.Get([]byte(key))
			if value == nil {
				continue
			}

			var batch types.Batch
			if err := k.cdc.Unmarshal(value, &batch); err != nil {
				loopError = err
				break
			}
			batches = append(batches, batch)
		}
		if loopError != nil {
			return nil, loopError
		}
	} else if order == "desc" {
		var loopError error
		startIndex := total - offset
		endIndex := startIndex - limit
		if endIndex < 0 {
			endIndex = 0
		}
		for i := startIndex; i > endIndex; i-- {
			key := strconv.FormatUint(i, 10)
			value := batchDataStore.Get([]byte(key))
			if value == nil {
				continue
			}

			var batch types.Batch
			if err := k.cdc.Unmarshal(value, &batch); err != nil {
				loopError = err
				break
			}
			batches = append(batches, batch)
		}
		if loopError != nil {
			return nil, loopError
		}
	} else {
		return nil, status.Error(codes.InvalidArgument, "invalid order")
	}

	return &types.QueryGetAllBatchesResponse{
		Batch:  batches,
		Total:  total,
		Offset: offset,
		Limit:  limit,
		Order:  order,
	}, nil
}
