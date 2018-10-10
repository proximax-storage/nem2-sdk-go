// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

type SignerInfo struct {
	Signer     string `json:"signer"`
	Signature  string `json:"signature"`
	ParentHash string `json:"parentHash"`
}
