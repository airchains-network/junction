package types

import (
	"testing"

	"github.com/airchains-network/junction/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterCollection_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterCollection
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRegisterCollection{
				CollectionOwner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRegisterCollection{
				CollectionOwner: sample.AccAddress(),
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
