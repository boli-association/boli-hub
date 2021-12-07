package hub_test

import (
	"testing"

	keepertest "github.com/bolimoney/boli-hub/testutil/keeper"
	"github.com/bolimoney/boli-hub/testutil/nullify"
	"github.com/bolimoney/boli-hub/x/hub"
	"github.com/bolimoney/boli-hub/x/hub/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HubKeeper(t)
	hub.InitGenesis(ctx, *k, genesisState)
	got := hub.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
