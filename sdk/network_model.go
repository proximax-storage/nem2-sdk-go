// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type NetworkType uint8

// NetworkType enums
const (
	MainNet         NetworkType = 104
	TestNet         NetworkType = 152
	Mijin           NetworkType = 96
	MijinTest       NetworkType = 144
	NotSupportedNet NetworkType = 0
)

func (nt NetworkType) String() string {
	return fmt.Sprintf("%d", nt)
}

// Network error
var networkTypeError = errors.New("wrong raw NetworkType value")

func ExtractNetworkType(version uint64) NetworkType {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, version)

	return NetworkType(b[1])
}
