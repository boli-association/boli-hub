package keeper_test

import (
	"testing"

	testkeeper "github.com/bolimoney/boli-hub/testutil/keeper"
	"github.com/bolimoney/boli-hub/x/hub/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HubKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
