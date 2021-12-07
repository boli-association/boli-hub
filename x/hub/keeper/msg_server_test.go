package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/bolimoney/boli-hub/testutil/keeper"
	"github.com/bolimoney/boli-hub/x/hub/keeper"
	"github.com/bolimoney/boli-hub/x/hub/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HubKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
