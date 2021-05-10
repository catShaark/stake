package keeper

import (
	"github.com/romelukaku/stake/x/stake/types"
)

var _ types.QueryServer = Keeper{}
