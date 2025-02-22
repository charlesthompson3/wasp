// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblocklog

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ImmutableGetBlockInfoParams struct {
	id int32
}

func (s ImmutableGetBlockInfoParams) BlockIndex() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, ParamBlockIndex.KeyID())
}

type MutableGetBlockInfoParams struct {
	id int32
}

func (s MutableGetBlockInfoParams) BlockIndex() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, ParamBlockIndex.KeyID())
}

type ImmutableGetEventsForBlockParams struct {
	id int32
}

func (s ImmutableGetEventsForBlockParams) BlockIndex() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, ParamBlockIndex.KeyID())
}

type MutableGetEventsForBlockParams struct {
	id int32
}

func (s MutableGetEventsForBlockParams) BlockIndex() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, ParamBlockIndex.KeyID())
}

type ImmutableGetEventsForContractParams struct {
	id int32
}

func (s ImmutableGetEventsForContractParams) ContractHname() wasmlib.ScImmutableHname {
	return wasmlib.NewScImmutableHname(s.id, ParamContractHname.KeyID())
}

func (s ImmutableGetEventsForContractParams) FromBlock() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, ParamFromBlock.KeyID())
}

func (s ImmutableGetEventsForContractParams) ToBlock() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, ParamToBlock.KeyID())
}

type MutableGetEventsForContractParams struct {
	id int32
}

func (s MutableGetEventsForContractParams) ContractHname() wasmlib.ScMutableHname {
	return wasmlib.NewScMutableHname(s.id, ParamContractHname.KeyID())
}

func (s MutableGetEventsForContractParams) FromBlock() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, ParamFromBlock.KeyID())
}

func (s MutableGetEventsForContractParams) ToBlock() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, ParamToBlock.KeyID())
}

type ImmutableGetEventsForRequestParams struct {
	id int32
}

func (s ImmutableGetEventsForRequestParams) RequestID() wasmlib.ScImmutableRequestID {
	return wasmlib.NewScImmutableRequestID(s.id, ParamRequestID.KeyID())
}

type MutableGetEventsForRequestParams struct {
	id int32
}

func (s MutableGetEventsForRequestParams) RequestID() wasmlib.ScMutableRequestID {
	return wasmlib.NewScMutableRequestID(s.id, ParamRequestID.KeyID())
}

type ImmutableGetRequestIDsForBlockParams struct {
	id int32
}

func (s ImmutableGetRequestIDsForBlockParams) BlockIndex() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, ParamBlockIndex.KeyID())
}

type MutableGetRequestIDsForBlockParams struct {
	id int32
}

func (s MutableGetRequestIDsForBlockParams) BlockIndex() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, ParamBlockIndex.KeyID())
}

type ImmutableGetRequestReceiptParams struct {
	id int32
}

func (s ImmutableGetRequestReceiptParams) RequestID() wasmlib.ScImmutableRequestID {
	return wasmlib.NewScImmutableRequestID(s.id, ParamRequestID.KeyID())
}

type MutableGetRequestReceiptParams struct {
	id int32
}

func (s MutableGetRequestReceiptParams) RequestID() wasmlib.ScMutableRequestID {
	return wasmlib.NewScMutableRequestID(s.id, ParamRequestID.KeyID())
}

type ImmutableGetRequestReceiptsForBlockParams struct {
	id int32
}

func (s ImmutableGetRequestReceiptsForBlockParams) BlockIndex() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, ParamBlockIndex.KeyID())
}

type MutableGetRequestReceiptsForBlockParams struct {
	id int32
}

func (s MutableGetRequestReceiptsForBlockParams) BlockIndex() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, ParamBlockIndex.KeyID())
}

type ImmutableIsRequestProcessedParams struct {
	id int32
}

func (s ImmutableIsRequestProcessedParams) RequestID() wasmlib.ScImmutableRequestID {
	return wasmlib.NewScImmutableRequestID(s.id, ParamRequestID.KeyID())
}

type MutableIsRequestProcessedParams struct {
	id int32
}

func (s MutableIsRequestProcessedParams) RequestID() wasmlib.ScMutableRequestID {
	return wasmlib.NewScMutableRequestID(s.id, ParamRequestID.KeyID())
}
