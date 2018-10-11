// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package utils

// Gets the i'th bit of a byte array.
//     *
//     * @param h The byte array.
//     * @param i The bit index.
//     * @return The value of the i'th bit in h
func GetBit(h []byte, i uint) int {

	sByte := int(h[i>>3])
	return (sByte >> (i & 7)) & 1
}

func GetBitToBool(h []byte, i uint) bool {
	return BoolFromInt[GetBit(h, i)]
}
