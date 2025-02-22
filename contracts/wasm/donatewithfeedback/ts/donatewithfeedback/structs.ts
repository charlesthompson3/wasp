// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib"

export class Donation {
    amount   : i64 = 0;          // amount donated
    donator  : wasmlib.ScAgentID = new wasmlib.ScAgentID(); // who donated
    error    : string = "";      // error to be reported to donator if anything goes wrong
    feedback : string = "";      // the feedback for the person donated to
    timestamp: i64 = 0;          // when the donation took place

    static fromBytes(bytes: u8[]): Donation {
        let decode = new wasmlib.BytesDecoder(bytes);
        let data = new Donation();
        data.amount = decode.int64();
        data.donator = decode.agentID();
        data.error = decode.string();
        data.feedback = decode.string();
        data.timestamp = decode.int64();
        decode.close();
        return data;
    }

    bytes(): u8[] {
        return new wasmlib.BytesEncoder().
            int64(this.amount).
            agentID(this.donator).
            string(this.error).
            string(this.feedback).
            int64(this.timestamp).
            data();
    }
}

export class ImmutableDonation {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    value(): Donation {
        return Donation.fromBytes(wasmlib.getBytes(this.objID, this.keyID,wasmlib. TYPE_BYTES));
    }
}

export class MutableDonation {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    setValue(value: Donation): void {
        wasmlib.setBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES, value.bytes());
    }

    value(): Donation {
        return Donation.fromBytes(wasmlib.getBytes(this.objID, this.keyID,wasmlib. TYPE_BYTES));
    }
}
