package keeper_test

import (
	"testing"

	testkeeper "github.com/lostak/greek/testutil/keeper"
	"github.com/lostak/greek/x/alpha/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AlphaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
