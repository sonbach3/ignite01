package keeper

import (
	"capy/x/family/types"
)

var _ types.QueryServer = Keeper{}
