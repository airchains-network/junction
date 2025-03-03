package types

import (
	"testing"

	"github.com/airchains-network/junction/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgInitiateVrf_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInitiateVrf
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInitiateVrf{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInitiateVrf{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
