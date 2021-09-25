package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/rust/testcore"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/vm/wasmlib"

	"github.com/iotaledger/goshimmer/packages/ledgerstate/utxoutil"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestMainCallsFromFullEP(t *testing.T) { run2(t, testMainCallsFromFullEP) }
func testMainCallsFromFullEP(t *testing.T, w bool) {
	ctx := setupTest(t, w, true)
	user := ctx.ContractCreator()

	f := testcore.ScFuncs.CheckContextFromFullEP(ctx.Sign(user))
	chainID := ctx.Convertor.ScChainID(ctx.Chain.ChainID)
	f.Params.ChainID().SetValue(chainID)
	f.Params.AgentID().SetValue(wasmlib.NewScAgentID(chainID, testcore.HScName))
	f.Params.Caller().SetValue(user.ScAgentID())
	f.Params.ChainOwnerID().SetValue(ctx.Originator().ScAgentID())
	f.Params.ContractCreator().SetValue(user.ScAgentID())
	f.Func.TransferIotas(1).Post()
	require.NoError(t, ctx.Err)
}

func TestMainCallsFromViewEP(t *testing.T) { run2(t, testMainCallsFromViewEP) }
func testMainCallsFromViewEP(t *testing.T, w bool) {
	ctx := setupTest(t, w, true)
	user := ctx.ContractCreator()

	f := testcore.ScFuncs.CheckContextFromViewEP(ctx)
	chainID := ctx.Convertor.ScChainID(ctx.Chain.ChainID)
	f.Params.ChainID().SetValue(chainID)
	f.Params.AgentID().SetValue(wasmlib.NewScAgentID(chainID, testcore.HScName))
	f.Params.ChainOwnerID().SetValue(ctx.Originator().ScAgentID())
	f.Params.ContractCreator().SetValue(user.ScAgentID())
	f.Func.Call()
	require.NoError(t, ctx.Err)
}

func TestMintedSupplyOk(t *testing.T) { run2(t, testMintedSupplyOk) }
func testMintedSupplyOk(t *testing.T, w bool) {
	//ctx := setupTest(t, w, true)
	//user := ctx.ContractCreator()
	//
	//f := testcore.ScFuncs.GetMintedSupply(ctx.Sign(user))
	//chainID := ctx.Convertor.ScChainID(ctx.Chain.ChainID)
	//f.Params.ChainID().SetValue(chainID)
	//f.Params.AgentID().SetValue(wasmlib.NewScAgentID(chainID, testcore.HScName))
	//f.Params.ChainOwnerID().SetValue(ctx.Originator().ScAgentID())
	//f.Params.ContractCreator().SetValue(user.ScAgentID())
	//f.Func.Call()
	//require.NoError(t, ctx.Err)

	_, chain := setupChain(t, nil)

	user, userAddress, _ := setupDeployer(t, chain)
	setupTestSandboxSC(t, chain, user, w)

	newSupply := uint64(42)
	req := solo.NewCallParams(ScName, sbtestsc.FuncGetMintedSupply.Name).
		WithIotas(1).
		WithMint(userAddress, newSupply)
	tx, ret, err := chain.PostRequestSyncTx(req, user)
	require.NoError(t, err)

	mintedAmounts := colored.BalancesFromL1Map(utxoutil.GetMintedAmounts(tx))
	t.Logf("minting request tx: %s", tx.ID().Base58())

	require.Len(t, mintedAmounts, 1)
	var col colored.Color
	for col1 := range mintedAmounts {
		col = col1
		break
	}
	t.Logf("Minted: amount = %d color = %s", newSupply, col.Base58())

	extraIota := uint64(0)
	if w {
		extraIota = 1
	}
	chain.Env.AssertAddressIotas(userAddress, solo.Saldo-3-extraIota-newSupply)
	chain.Env.AssertAddressBalance(userAddress, col, newSupply)

	colorBack, ok, err := codec.DecodeColor(ret.MustGet(sbtestsc.VarMintedColor))
	require.NoError(t, err)
	require.True(t, ok)
	t.Logf("color back: %s", colorBack.Base58())
	require.EqualValues(t, col, colorBack)
	supplyBack, ok, err := codec.DecodeUint64(ret.MustGet(sbtestsc.VarMintedSupply))
	require.NoError(t, err)
	require.True(t, ok)
	t.Logf("supply back: %d", supplyBack)
	require.EqualValues(t, int(newSupply), int(supplyBack))
}
