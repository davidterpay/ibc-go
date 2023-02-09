package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	host "github.com/davidterpay/ibc-go/modules/core/24-host"
	"github.com/davidterpay/ibc-go/modules/core/simulation"
	"github.com/davidterpay/ibc-go/modules/core/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: math.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	// Remark: the current RandomizedGenState function
	// is actually not random as it does not utilize concretely the random value r.
	// This tests will pass for any value of r.
	simulation.RandomizedGenState(&simState)

	var ibcGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[host.ModuleName], &ibcGenesis)

	require.NotNil(t, ibcGenesis.ClientGenesis)
	require.NotNil(t, ibcGenesis.ConnectionGenesis)
	require.NotNil(t, ibcGenesis.ChannelGenesis)
}
