package simulation

import (
	"math/rand"

	"github.com/airchains-network/junction/x/trackgate/keeper"
	"github.com/airchains-network/junction/x/trackgate/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgIntegrityCheck(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgIntegrityCheck{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the IntegrityCheck simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "IntegrityCheck simulation not implemented"), nil, nil
	}
}
