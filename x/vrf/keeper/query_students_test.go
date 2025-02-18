package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/airchains-network/junction/testutil/keeper"
	"github.com/airchains-network/junction/testutil/nullify"
	"github.com/airchains-network/junction/x/vrf/types"
)

func TestStudentsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VrfKeeper(t)
	msgs := createNStudents(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetStudentsRequest
		response *types.QueryGetStudentsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetStudentsRequest{Id: msgs[0].Id},
			response: &types.QueryGetStudentsResponse{Students: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetStudentsRequest{Id: msgs[1].Id},
			response: &types.QueryGetStudentsResponse{Students: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetStudentsRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Students(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestStudentsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VrfKeeper(t)
	msgs := createNStudents(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStudentsRequest {
		return &types.QueryAllStudentsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StudentsAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Students), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Students),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StudentsAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Students), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Students),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StudentsAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Students),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StudentsAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
