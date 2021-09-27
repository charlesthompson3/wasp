// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmsolo

import (
	"flag"
	"testing"
	"time"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/goshimmer/packages/ledgerstate/utxoutil"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/iscp/colored"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/util"
	"github.com/iotaledger/wasp/packages/vm/core/root"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
	"github.com/iotaledger/wasp/packages/vm/wasmlib"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/stretchr/testify/require"
)

const (
	Debug      = true
	StackTrace = false
	TraceHost  = false
)

type SoloContext struct {
	Chain     *solo.Chain
	Convertor SoloConvertor
	creator   *SoloAgent
	Err       error
	keyPair   *ed25519.KeyPair
	mint      uint64
	offLedger bool
	scName    string
	Tx        *ledgerstate.Transaction
	wasmHost  wasmhost.WasmHost
}

var (
	_        wasmlib.ScFuncCallContext = &SoloContext{}
	_        wasmlib.ScViewCallContext = &SoloContext{}
	SoloHost wasmlib.ScHost
	GoDebug  = flag.Bool("godebug", false, "debug go smart contract code")
)

// NewSoloContext can be used to create a SoloContext associated with a smart contract with
// minimal information and will verify successful creation before returning ctx.
// It will start a default chain "chain1" before initializing the smart contract.
// It takes the scName and onLoad() function for the contract and optionally
// an init.Func initialized with the parameters to pass to the contract's init() function.
// Unless you want to use a different chain than the default "chain1" this will be your
// function of choice to set up a smart contract for your tests
func NewSoloContext(t *testing.T, scName string, onLoad func(), init ...*wasmlib.ScInitFunc) *SoloContext {
	ctx := NewSoloContextForChain(t, nil, nil, scName, onLoad, init...)
	require.NoError(t, ctx.Err)
	return ctx
}

// NewSoloContextForChain can be used to create a SoloContext associated with a smart contract on a particular chain.
// When chain is nil the function will start a default chain "chain1" before initializing the smart contract.
// It takes the scName and onLoad() function for the contract and optionally
// an init.Func initialized with the parameters to pass to the contract's init() function.
// You can check for any error that occurred by checking the ctx.Err member.
func NewSoloContextForChain(t *testing.T, chain *solo.Chain, creator *SoloAgent, scName string, onLoad func(), init ...*wasmlib.ScInitFunc) *SoloContext {
	if chain == nil {
		chain = StartChain(t, "chain1")
	}

	ctx := &SoloContext{scName: scName, Chain: chain, creator: creator}
	var keyPair *ed25519.KeyPair
	if creator != nil {
		keyPair = creator.Pair
	}
	ctx.Err = deploy(chain, keyPair, scName, onLoad, init...)
	if ctx.Err != nil {
		return ctx
	}
	return ctx.init(onLoad)
}

func NewSoloContextForCore(t *testing.T, chain *solo.Chain, scName string, onLoad func()) *SoloContext {
	if chain == nil {
		chain = StartChain(t, "chain1")
	}

	ctx := &SoloContext{scName: scName, Chain: chain}
	return ctx.init(onLoad)
}

// TODO can we make upload work through off-ledger request instead?
// that way we can get rid of all the extra token code when checking balances

func deploy(chain *solo.Chain, keyPair *ed25519.KeyPair, scName string, onLoad func(), init ...*wasmlib.ScInitFunc) error {
	var params []interface{}
	if len(init) != 0 {
		params = init[0].Params()
	}

	if SoloHost == nil {
		retDict, err := chain.CallView(root.Contract.Name, root.FuncFindContract.Name,
			root.ParamHname, iscp.Hn(scName),
		)
		if err != nil {
			return err
		}
		retBin, err := retDict.Get(root.ParamContractFound)
		if err != nil {
			return err
		}
		if len(retBin) == 1 && retBin[0] == 0xff {
			// a contract with that name already exists: probably native code
			return nil
		}
	}

	SoloHost = nil
	if *GoDebug {
		wasmproc.GoWasmVM = wasmhost.NewWasmGoVM(scName, onLoad)
		hprog, err := chain.UploadWasm(keyPair, []byte("go:"+scName))
		if err != nil {
			return err
		}
		return chain.DeployContract(keyPair, scName, hprog, params...)
	}

	// wasmproc.GoWasmVM = NewWasmTimeJavaVM()
	// wasmproc.GoWasmVM = NewWartVM()
	// wasmproc.GoWasmVM = NewWasmerVM()
	wasmFile := scName + "_bg.wasm"
	exists, _ := util.ExistsFilePath("../pkg/" + wasmFile)
	if exists {
		wasmFile = "../pkg/" + wasmFile
	}
	return chain.DeployWasmContract(keyPair, scName, wasmFile, params...)
	//hprog, err := chain.UploadWasmFromFile(keyPair, wasmFile)
	//if err != nil {
	//	return err
	//}
	//return chain.DeployContract(keyPair, scName, hprog, params...)
}

// StartChain starts a new chain named chainName.
func StartChain(t *testing.T, chainName string) *solo.Chain {
	wasmhost.HostTracing = TraceHost
	// wasmhost.ExtendedHostTracing = TraceHost
	env := solo.New(t, Debug, StackTrace)
	return env.NewChain(nil, chainName)
}

// AdvanceClockBy is used to forward the internal clock by the provided step duration.
func (ctx *SoloContext) AdvanceClockBy(step time.Duration) {
	ctx.Chain.Env.AdvanceClockBy(step)
}

// Balance returns the account balance of the specified agent on the chain associated with ctx.
// To return the account balance of the contract associated with ctx use nil as agent.
// The optional color parameter can be used to retrieve the balance for the specific color.
// When color is omitted, wasmlib.IOTA is assumed.
func (ctx *SoloContext) Balance(agent *SoloAgent, color ...wasmlib.ScColor) int64 {
	var account *iscp.AgentID
	if agent == nil {
		account = iscp.NewAgentID(ctx.Chain.ChainID.AsAddress(), iscp.Hn(ctx.scName))
	} else {
		account = iscp.NewAgentID(agent.address, 0)
	}
	balances := ctx.Chain.GetAccountBalance(account)
	switch len(color) {
	case 0:
		return int64(balances.Get(colored.IOTA))
	case 1:
		col, err := colored.ColorFromBytes(color[0].Bytes())
		require.NoError(ctx.Chain.Env.T, err)
		return int64(balances.Get(col))
	default:
		require.Fail(ctx.Chain.Env.T, "too many color arguments")
		return 0
	}
}

// ContractExists checks to see if the contract named scName exists in the chain associated with ctx.
func (ctx *SoloContext) ContractExists(scName string) error {
	_, err := ctx.Chain.FindContract(scName)
	return err
}

func (ctx *SoloContext) Host() wasmlib.ScHost {
	return nil
}

// init further initializes the SoloContext.
func (ctx *SoloContext) init(onLoad func()) *SoloContext {
	ctx.wasmHost.Init(ctx.Chain.Log)
	ctx.wasmHost.TrackObject(wasmproc.NewNullObject(&ctx.wasmHost.KvStoreHost))
	ctx.wasmHost.TrackObject(NewSoloScContext(ctx))
	if SoloHost == nil {
		SoloHost = wasmlib.ConnectHost(&ctx.wasmHost)
	}
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
	onLoad()
	return ctx
}

// InitFuncCallContext is a function that is required to use SoloContext as an ScFuncCallContext
func (ctx *SoloContext) InitFuncCallContext() {
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
}

// InitViewCallContext is a function that is required to use SoloContext as an ScViewCallContext
func (ctx *SoloContext) InitViewCallContext() {
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
}

// NewSoloAgent creates a new SoloAgent with solo.Saldo tokens in its address
func (ctx *SoloContext) NewSoloAgent() *SoloAgent {
	return NewSoloAgent(ctx.Chain.Env)
}

// TODO add context methods that are identical to func context methods
// which means that Balances, ContractCreator and Minted need to change

// ContractCreator returns a SoloAgent representing the contract creator
func (ctx *SoloContext) ContractCreator() *SoloAgent {
	if ctx.creator != nil {
		return ctx.creator
	}
	return ctx.Originator()
}

// Minted returns the color and amount of newly minted tokens
func (ctx *SoloContext) Minted() (wasmlib.ScColor, uint64) {
	t := ctx.Chain.Env.T
	t.Logf("minting request tx: %s", ctx.Tx.ID().Base58())
	mintedAmounts := colored.BalancesFromL1Map(utxoutil.GetMintedAmounts(ctx.Tx))
	require.Len(t, mintedAmounts, 1)
	var mintedColor wasmlib.ScColor
	var mintedAmount uint64
	for c := range mintedAmounts {
		mintedColor = ctx.Convertor.ScColor(c)
		mintedAmount = mintedAmounts[c]
		break
	}
	t.Logf("Minted: amount = %d color = %s", mintedAmount, mintedColor.String())
	return mintedColor, mintedAmount
}

// OffLedger tells SoloContext to Post() the next request off-ledger
func (ctx *SoloContext) OffLedger(agent *SoloAgent) wasmlib.ScFuncCallContext {
	ctx.offLedger = true
	ctx.keyPair = agent.Pair
	return ctx
}

// Originator returns a SoloAgent representing the chain originator
func (ctx *SoloContext) Originator() *SoloAgent {
	c := ctx.Chain
	return &SoloAgent{Env: c.Env, Pair: c.OriginatorKeyPair, address: c.OriginatorAddress}
}

// Sign is used to force a different agent for signing a Post() request
func (ctx *SoloContext) Sign(agent *SoloAgent, mint ...uint64) wasmlib.ScFuncCallContext {
	ctx.keyPair = agent.Pair
	if len(mint) != 0 {
		ctx.mint = mint[0]
	}
	return ctx
}

// Transfer creates a new ScTransfers proxy
func (ctx *SoloContext) Transfer() wasmlib.ScTransfers {
	return wasmlib.NewScTransfers()
}

// WaitForPendingRequests waits for expectedRequests pending requests to be processed.
// The function will wait for maxWait (default 5 seconds) duration before giving up with a timeout.
// The function returns the false in case of a timeout.
func (ctx *SoloContext) WaitForPendingRequests(expectedRequests int, maxWait ...time.Duration) bool {
	_ = wasmlib.ConnectHost(SoloHost)
	info := ctx.Chain.MempoolInfo()
	result := ctx.Chain.WaitForRequestsThrough(expectedRequests+info.OutPoolCounter, maxWait...)
	_ = wasmlib.ConnectHost(&ctx.wasmHost)
	return result
}
