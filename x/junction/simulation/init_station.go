package simulation

import (
	"math/rand"

	"github.com/airchains-network/junction/x/junction/keeper"
	"github.com/airchains-network/junction/x/junction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgInitStation(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgInitStation{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the InitStation simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "InitStation simulation not implemented"), nil, nil
	}
}
