// Copyright 2017 Author: Ruslan Bikchentaev. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestAddresses_MarshalJSON(t *testing.T) {

	addresses := Addresses{
		list: []*Address{
			&Address{Address: "SDRDGFTDLLCB67D4HPGIMIHPNSRYRJRT7DOBGWZY"},
			&Address{Address: "SBCPGZ3S2SCC3YHBBTYDCUZV4ZZEPHM2KGCP4QXX"},
		},
	}
	b, err := json.Marshal(addresses)
	if err != nil {
		t.Fatal(err)
	}

	b1, err := addresses.Marshal1JSON()
	if err != nil {
		t.Fatal(err)
	}

	addresses.Marshal1JSON()

	if !bytes.Equal(b, b1) {
		t.Error("not equal")
	}
	t.Log("standart", string(b))
	t.Log("self-made", string(b1))

	ad := &Addresses{}
	err = json.Unmarshal(b1, ad)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(ad)
	}
}
