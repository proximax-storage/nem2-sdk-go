// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

// structure for Subscribe status
type StatusInfo struct {
	Status string `json:"status"`
	Hash   Hash   `json:"hash"`
}

type SignerInfo struct {
	Signer     string `json:"signer"`
	Signature  string `json:"signature"`
	ParentHash Hash   `json:"parentHash"`
}

type ErrorInfo struct {
	Error error
}

type HashInfo struct {
	Meta struct {
		Hash `json:"hash"`
	} `json:"meta"`
}

// structure for Subscribe PartialRemoved
type PartialRemovedInfo struct {
	HashInfo
}

// structure for Subscribe UnconfirmedRemoved
type UnconfirmedRemoved struct {
	HashInfo
}
