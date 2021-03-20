// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package coretypes

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/hashing"
	"golang.org/x/xerrors"
	"io"
)

// ChainID represents the global identifier of the chain
// It is wrapped AliasAddress, an address without a private key behind
type ChainID struct {
	*ledgerstate.AliasAddress
}

var NilChainID = ChainID{}

// NewChainIDFromAddress temporary
func NewChainIDFromAddress(addr ledgerstate.Address) (*ChainID, error) {
	alias, ok := addr.(*ledgerstate.AliasAddress)
	if !ok {
		return nil, xerrors.New("chain id must be an alias address")
	}
	return &ChainID{alias}, nil
}

// NewChainIDFromBase58 constructor decodes base58 string to the ChainID
func NewChainIDFromBase58(b58 string) (*ChainID, error) {
	alias, err := ledgerstate.AliasAddressFromBase58EncodedString(b58)
	if err != nil {
		return nil, err
	}
	return &ChainID{alias}, nil
}

// NewChainIDFromBytes reconstructs a ChainID from its binary representation.
func NewChainIDFromBytes(data []byte) (*ChainID, error) {
	alias, _, err := ledgerstate.AliasAddressFromBytes(data)
	if err != nil {
		return nil, err
	}
	return &ChainID{alias}, nil
}

// NewRandomChainID creates a random chain ID.
func NewRandomChainID(seed ...[]byte) *ChainID {
	var h hashing.HashValue
	if len(seed) > 0 {
		h = hashing.HashData(seed[0])
	} else {
		h = hashing.RandomHash(nil)
	}
	return &ChainID{ledgerstate.NewAliasAddress(h[:])}
}

func (chid *ChainID) Equals(chid1 *ChainID) bool {
	return chid.AliasAddress.Equals(chid1.AliasAddress)
}

func (chid *ChainID) Clone() (ret ChainID) {
	ret.AliasAddress = chid.AliasAddress.Clone().(*ledgerstate.AliasAddress)
	return
}

func (chid *ChainID) Base58() string {
	return chid.AliasAddress.Base58()
}

// String human readable form (base58 encoding)
func (chid *ChainID) String() string {
	return "$/" + chid.Base58()
}

// AsAddress Temporary
func (chid *ChainID) AsAddress() ledgerstate.Address {
	return chid.AliasAddress
}

func (chid *ChainID) Read(r io.Reader) error {
	var buf [ledgerstate.AddressLength]byte
	if n, err := r.Read(buf[:]); err != nil || n != ledgerstate.AddressLength {
		return xerrors.Errorf("error while parsing address (err=%v)", err)
	}
	alias, _, err := ledgerstate.AliasAddressFromBytes(buf[:])
	if err != nil {
		return err
	}
	chid.AliasAddress = alias
	return nil
}

func (chid *ChainID) Write(w io.Writer) error {
	_, err := w.Write(chid.AliasAddress.Bytes())
	return err
}
