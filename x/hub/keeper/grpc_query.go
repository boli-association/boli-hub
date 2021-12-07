package keeper

import (
	"github.com/bolimoney/boli-hub/x/hub/types"
)

var _ types.QueryServer = Keeper{}
