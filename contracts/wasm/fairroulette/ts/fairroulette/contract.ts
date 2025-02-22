// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib"
import * as sc from "./index";

export class ForcePayoutCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncForcePayout);
}

export class ForcePayoutContext {
    state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState();
}

export class ForceResetCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncForceReset);
}

export class ForceResetContext {
    state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState();
}

export class PayWinnersCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPayWinners);
}

export class PayWinnersContext {
    state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState();
}

export class PlaceBetCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPlaceBet);
    params: sc.MutablePlaceBetParams = new sc.MutablePlaceBetParams();
}

export class PlaceBetContext {
    params: sc.ImmutablePlaceBetParams = new sc.ImmutablePlaceBetParams();
    state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState();
}

export class PlayPeriodCall {
    func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPlayPeriod);
    params: sc.MutablePlayPeriodParams = new sc.MutablePlayPeriodParams();
}

export class PlayPeriodContext {
    params: sc.ImmutablePlayPeriodParams = new sc.ImmutablePlayPeriodParams();
    state: sc.MutableFairRouletteState = new sc.MutableFairRouletteState();
}

export class LastWinningNumberCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewLastWinningNumber);
    results: sc.ImmutableLastWinningNumberResults = new sc.ImmutableLastWinningNumberResults();
}

export class LastWinningNumberContext {
    results: sc.MutableLastWinningNumberResults = new sc.MutableLastWinningNumberResults();
    state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState();
}

export class RoundNumberCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewRoundNumber);
    results: sc.ImmutableRoundNumberResults = new sc.ImmutableRoundNumberResults();
}

export class RoundNumberContext {
    results: sc.MutableRoundNumberResults = new sc.MutableRoundNumberResults();
    state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState();
}

export class RoundStartedAtCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewRoundStartedAt);
    results: sc.ImmutableRoundStartedAtResults = new sc.ImmutableRoundStartedAtResults();
}

export class RoundStartedAtContext {
    results: sc.MutableRoundStartedAtResults = new sc.MutableRoundStartedAtResults();
    state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState();
}

export class RoundStatusCall {
    func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewRoundStatus);
    results: sc.ImmutableRoundStatusResults = new sc.ImmutableRoundStatusResults();
}

export class RoundStatusContext {
    results: sc.MutableRoundStatusResults = new sc.MutableRoundStatusResults();
    state: sc.ImmutableFairRouletteState = new sc.ImmutableFairRouletteState();
}

export class ScFuncs {

    static forcePayout(ctx: wasmlib.ScFuncCallContext): ForcePayoutCall {
        let f = new ForcePayoutCall();
        return f;
    }

    static forceReset(ctx: wasmlib.ScFuncCallContext): ForceResetCall {
        let f = new ForceResetCall();
        return f;
    }

    static payWinners(ctx: wasmlib.ScFuncCallContext): PayWinnersCall {
        let f = new PayWinnersCall();
        return f;
    }

    static placeBet(ctx: wasmlib.ScFuncCallContext): PlaceBetCall {
        let f = new PlaceBetCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static playPeriod(ctx: wasmlib.ScFuncCallContext): PlayPeriodCall {
        let f = new PlayPeriodCall();
        f.func.setPtrs(f.params, null);
        return f;
    }

    static lastWinningNumber(ctx: wasmlib.ScViewCallContext): LastWinningNumberCall {
        let f = new LastWinningNumberCall();
        f.func.setPtrs(null, f.results);
        return f;
    }

    static roundNumber(ctx: wasmlib.ScViewCallContext): RoundNumberCall {
        let f = new RoundNumberCall();
        f.func.setPtrs(null, f.results);
        return f;
    }

    static roundStartedAt(ctx: wasmlib.ScViewCallContext): RoundStartedAtCall {
        let f = new RoundStartedAtCall();
        f.func.setPtrs(null, f.results);
        return f;
    }

    static roundStatus(ctx: wasmlib.ScViewCallContext): RoundStatusCall {
        let f = new RoundStatusCall();
        f.func.setPtrs(null, f.results);
        return f;
    }
}
