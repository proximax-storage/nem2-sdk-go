// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
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
	MAIN_NET                  NetworkType = 104
	TEST_NET                  NetworkType = 152
	MIJIN                     NetworkType = 96
	MIJIN_TEST                NetworkType = 144
	NOT_SUPPORTED_NETWORKTYPE NetworkType = 0
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
