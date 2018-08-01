// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAddresses_MarshalJSON(t *testing.T) {

	addresses := Addresses{
		Addresses: []*Address{
			&Address{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
			&Address{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
		},
	}
	b, err := json.Marshal(addresses)
	if err != nil {
		t.Fatal(err)
	}

	b1, err := addresses.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, b1) {
		t.Error("not equal standart & self-made marshaling")
	}
	t.Log("standart", string(b))
	t.Log("self-made", string(b1))

	var ad Addresses
	err = json.Unmarshal(b1, &ad)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("standart", ad)
	}
	if !reflect.DeepEqual(ad, addresses) {
		t.Error("not equal unmarshaling")
	}
}
