package lselysium

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlin-network/estake-native/v2/x/lselysium/keeper"
	"github.com/merlin-network/estake-native/v2/x/lselysium/types"
)

// BeginBlocker updates liquid validator set changes for the current block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	k.UpdateLiquidValidatorSet(ctx)
}
