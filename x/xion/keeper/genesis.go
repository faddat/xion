package keeper

import (
	"github.com/burnt-labs/xion/x/xion/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the bank module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	if err := k.SetParams(ctx, *genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the bank module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	rv := types.NewGenesisState(
		k.GetParams(ctx),
	)
	return rv
}