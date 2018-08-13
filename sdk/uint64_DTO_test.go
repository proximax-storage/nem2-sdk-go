// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import "testing"

func TestNewUint64DTO(t *testing.T) {
	u := NewUint64DTO([]byte("7180930485"), []byte("-888466398"))
	if u.String() != "84b3552d375ffa4b" {
		t.Errorf("fail %d", u.ExtractIntArray())
	}
}
