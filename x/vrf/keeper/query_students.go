package keeper

import (
	"context"

	"github.com/airchains-network/junction/x/vrf/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StudentsAll(ctx context.Context, req *types.QueryAllStudentsRequest) (*types.QueryAllStudentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var studentss []types.Students

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	studentsStore := prefix.NewStore(store, types.KeyPrefix(types.StudentsKey))

	pageRes, err := query.Paginate(studentsStore, req.Pagination, func(key []byte, value []byte) error {
		var students types.Students
		if err := k.cdc.Unmarshal(value, &students); err != nil {
			return err
		}

		studentss = append(studentss, students)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStudentsResponse{Students: studentss, Pagination: pageRes}, nil
}

func (k Keeper) Students(ctx context.Context, req *types.QueryGetStudentsRequest) (*types.QueryGetStudentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	students, found := k.GetStudents(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetStudentsResponse{Students: students}, nil
}
