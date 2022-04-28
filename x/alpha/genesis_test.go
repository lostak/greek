package alpha_test

import (
	"testing"

	keepertest "github.com/lostak/greek/testutil/keeper"
	"github.com/lostak/greek/testutil/nullify"
	"github.com/lostak/greek/x/alpha"
	"github.com/lostak/greek/x/alpha/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AlphaKeeper(t)
	alpha.InitGenesis(ctx, *k, genesisState)
	got := alpha.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
