package simulation

import (
	"math/rand"

	"github.com/airchains-network/junction/x/zksequencer/keeper"
	"github.com/airchains-network/junction/x/zksequencer/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddStudentViaZksequencer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddStudentViaZksequencer{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddStudentViaZksequencer simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "AddStudentViaZksequencer simulation not implemented"), nil, nil
	}
}
