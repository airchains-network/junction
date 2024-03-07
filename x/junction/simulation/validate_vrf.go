package simulation

import (
	"math/rand"

	"github.com/ComputerKeeda/junction/x/junction/keeper"
	"github.com/ComputerKeeda/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgValidateVrf(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgValidateVrf{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ValidateVrf simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "ValidateVrf simulation not implemented"), nil, nil
	}
}
