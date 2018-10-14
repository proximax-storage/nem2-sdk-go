// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const transactionId = "5B55E02EACCB7B00015DB6E1"
const transactionHash = "7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F"

var transaction = &TransferTransaction{
	abstractTransaction: abstractTransaction{
		Type:        Transfer,
		Version:     uint64(3),
		NetworkType: MijinTest,
		Signature:   "ADF80CBC864B65A8D94205E9EC6640FA4AE0E3011B27F8A93D93761E454A9853BF0AB1ECB3DF62E1D2D267D3F1913FAB0E2225CE5EA3937790B78FFA1288870C",
		Signer:      &PublicAccount{&Address{MijinTest, "SBJ5D7TFIJWPY56JBEX32MUWI5RU6KVKZYITQ2HA"}, "27F6BEF9A7F75E33AE2EB2EBA10EF1D6BEA4D30EBD5E39AF8EE06E96E11AE2A9"},
		Fee:         uint64DTO{0, 0}.toBigInt(),
		Deadline:    &Deadline{time.Unix(0, uint64DTO{1094650402, 17}.toBigInt().Int64()*int64(time.Millisecond))},
		TransactionInfo: &TransactionInfo{
			Height:              uint64DTO{42, 0}.toBigInt(),
			Hash:                "45AC1259DABD7163B2816232773E66FC00342BB8DD5C965D4B784CD575FDFAF1",
			MerkleComponentHash: "45AC1259DABD7163B2816232773E66FC00342BB8DD5C965D4B784CD575FDFAF1",
			Index:               0,
			Id:                  "5B686E97F0C0EA00017B9437",
		},
	},
	Mosaics: Mosaics{
		&Mosaic{&MosaicId{uint64DTO{3646934825, 3576016193}.toBigInt(), ""}, uint64DTO{10000000, 0}.toBigInt()},
	},
	Recipient: &Address{MijinTest, "SBJUINHAC3FKCMVLL2WHBQFPPXYEHOMQY6E2SPVR"},
	Message:   &Message{Type: 0, Payload: ""},
}

var fakeDeadline = &Deadline{time.Unix(1459468800, 1000000)}

const transactionJson = `
{
   "meta":{
      "height":[42, 0],
      "hash":"45AC1259DABD7163B2816232773E66FC00342BB8DD5C965D4B784CD575FDFAF1",
      "merkleComponentHash":"45AC1259DABD7163B2816232773E66FC00342BB8DD5C965D4B784CD575FDFAF1",
      "index":0,
      "id":"5B686E97F0C0EA00017B9437"
   },
   "transaction":{
      "signature":"ADF80CBC864B65A8D94205E9EC6640FA4AE0E3011B27F8A93D93761E454A9853BF0AB1ECB3DF62E1D2D267D3F1913FAB0E2225CE5EA3937790B78FFA1288870C",
      "signer":"27F6BEF9A7F75E33AE2EB2EBA10EF1D6BEA4D30EBD5E39AF8EE06E96E11AE2A9",
      "version":36867,
      "type":16724,
      "fee":[
         0,
         0
      ],
      "deadline":[
         1094650402,
         17
      ],
      "recipient":"90534434E016CAA132AB5EAC70C0AF7DF043B990C789A93EB1",
      "message":{
         "type":0,
         "payload":""
      },
      "mosaics":[
         {
            "id":[
               3646934825,
               3576016193
            ],
            "amount":[
               10000000,
               0
            ]
         }
      ]
   }
}
`

var status = &TransactionStatus{
	&Deadline{time.Unix(uint64DTO{1, 0}.toBigInt().Int64(), int64(time.Millisecond))},
	"confirmed",
	"Success",
	"7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F",
	uint64DTO{1, 0}.toBigInt(),
}

const statusJson = `{
	"group": "confirmed",
	"status": "Success",
	"hash": "7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F",
	"deadline": [1,0],
	"height": [1, 0]
}`

func TestTransactionService_GetTransaction_TransferTransaction(t *testing.T) {
	cl, mux, _, teardown, err := setupMockServer()
	if err != nil {
		t.Errorf("Transaction.GetTransaction error setting up mock server: %v", err)
	}
	defer teardown()

	mux.HandleFunc("/transaction/5B55E02EACCB7B00015DB6E1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, transactionJson)
	})

	tx, _, err := cl.Transaction.GetTransaction(context.Background(), transactionId)
	if err != nil {
		t.Errorf("Transaction.GetTransaction returned error: %s", err)
	}

	if !reflect.DeepEqual(tx, transaction) {
		t.Errorf("Transaction.GetTransaction returned %s, want %s", tx, transaction)
	}
}

func TestTransactionService_GetTransactions(t *testing.T) {
	cl, mux, _, teardown, err := setupMockServer()
	if err != nil {
		t.Errorf("Transaction.GetTransactions error setting up mock server: %v", err)
	}
	defer teardown()

	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "["+transactionJson+"]")
	})

	tx, _, err := cl.Transaction.GetTransactions(context.Background(), []string{
		transactionId,
	})

	if err != nil {
		t.Errorf("Transaction.GetTransactions returned error: %v", err)
	}

	want := []Transaction{
		transaction,
	}

	if !reflect.DeepEqual(tx, want) {
		t.Errorf("Transaction.GetTransactions returned %s, want %s", tx, want)
	}
}

func TestTransactionService_GetTransactionStatus(t *testing.T) {
	cl, mux, _, teardown, err := setupMockServer()
	if err != nil {
		t.Errorf("Transaction.GetTransactionStatus error setting up mock server: %v", err)
	}
	defer teardown()

	mux.HandleFunc("/transaction/7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, statusJson)
	})

	tx, resp, err := cl.Transaction.GetTransactionStatus(context.Background(), transactionHash)
	if err != nil {
		t.Errorf("Transaction.GetTransactionStatus returned error: %s, resp status: %s", err, resp.Status)
	}

	if !reflect.DeepEqual(tx, status) {
		t.Errorf("Transaction.GetTransactionStatus returned %s, want %s, resp status: %s", tx, status, resp.Status)
	}
}

func TestTransactionService_GetTransactionStatuses(t *testing.T) {
	cl, mux, _, teardown, err := setupMockServer()
	if err != nil {
		t.Errorf("Transaction.GetTransactionStatuses error setting up mock server: %v", err)
	}
	defer teardown()

	mux.HandleFunc("/transaction/statuses", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "["+statusJson+"]")
	})

	tx, _, err := cl.Transaction.GetTransactionStatuses(context.Background(), []string{transactionHash})
	if err != nil {
		t.Errorf("Transaction.GetTransactionStatuses returned error: %s", err)
	}

	want := []*TransactionStatus{status}
	if !reflect.DeepEqual(tx, want) {
		t.Errorf("Transaction.GetTransactionStatuses returned %s, want %s", tx, want)
	}
}

func TestAggregateTransactionSerialization(t *testing.T) {
	want := []byte{209, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 65, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 85, 0, 0, 0, 85, 0, 0, 0, 132, 107, 68, 57, 21, 69, 121, 165, 144, 59, 20, 89,
		201, 207, 105, 203, 129, 83, 246, 208, 17, 10, 122, 14, 214, 29, 226, 154,
		228, 129, 11, 242, 3, 144, 84, 65, 144, 80, 185, 131, 126, 250, 180,
		187, 232, 164, 185, 187, 50, 216, 18, 249, 136, 92, 0, 216, 252,
		22, 80, 225, 66, 1, 0, 1, 0, 41, 207, 95, 217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0}
	p, err := NewAccountFromPublicKey("846B4439154579A5903B1459C9CF69CB8153F6D0110A7A0ED61DE29AE4810BF2", MijinTest)
	ttx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SBILTA367K2LX2FEXG5TFWAS7GEFYAGY7QLFBYKC", MijinTest),
		Mosaics{Xem(10000000)},
		NewPlainMessage(""),
		MijinTest,
	)

	ttx.Signer = p

	atx, err := NewCompleteAggregateTransaction(fakeDeadline, []Transaction{ttx}, MijinTest)

	b, err := atx.generateBytes()
	if err != nil {
		t.Errorf("AggregateTransaction generateBytes() returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("AggregateTransaction generateBytes() returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestAggregateTransactionSigningWithMultipleCosignatures(t *testing.T) {
	p, err := NewAccountFromPublicKey("B694186EE4AB0558CA4AFCFDD43B42114AE71094F5A1FC4A913FE9971CACD21D", MijinTest)
	ttx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SBILTA367K2LX2FEXG5TFWAS7GEFYAGY7QLFBYKC", MijinTest),
		Mosaics{},
		NewPlainMessage("test-message"),
		MijinTest,
	)

	ttx.Signer = p

	atx, err := NewCompleteAggregateTransaction(fakeDeadline, []Transaction{ttx}, MijinTest)

	acc1, err := NewAccountFromPrivateKey("2a2b1f5d366a5dd5dc56c3c757cf4fe6c66e2787087692cf329d7a49a594658b", MijinTest)
	acc2, err := NewAccountFromPrivateKey("b8afae6f4ad13a1b8aad047b488e0738a437c7389d4ff30c359ac068910c1d59", MijinTest) // TODO from original repo: "bug with private key"

	stx, err := acc1.SignWithCosignatures(atx, []*Account{acc2})

	if err != nil {
		t.Errorf("AggregaterTransaction signing returned error: %s", err)
	}

	if !reflect.DeepEqual(stx.Payload[0:8], "2d010000") {
		t.Errorf("AggregateTransaction signing returned wrong payload: \n %s", stx.Payload[0:8])
	}

	if !reflect.DeepEqual(stx.Payload[240:256], "5100000051000000") {
		t.Errorf("AggregateTransaction signing returned wrong payload: \n %s", stx.Payload[240:256])
	}
	//if !reflect.DeepEqual(stx.Payload[320:474], "039054419050B9837EFAB4BBE8A4B9BB32D812F9885C00D8FC1650E1420D000000746573742D6D65737361676568B3FBB18729C1FDE225C57F8CE080FA828F0067E451A3FD81FA628842B0B763") {
	//	t.Errorf("AggregateTransaction signing returned wrong payload: \n %s", stx.Payload[320:474])
	//} this test is not working in original repo and commented out too
}

func TestCosisignatureTransactionSigning(t *testing.T) {
	want1 := "bf3bc39f2292c028cb0ffa438a9f567a7c4d793d2f8522c8deac74befbcb61af6414adf27b2176d6a24fef612aa6db2f562176a11c46ba6d5e05430042cb5705"
	rtx := "{\"meta\":{\"hash\":\"671653C94E2254F2A23EFEDB15D67C38332AED1FBD24B063C0A8E675582B6A96\",\"height\":[18160,0],\"id\":\"5A0069D83F17CF0001777E55\",\"index\":0,\"merkleComponentHash\":\"81E5E7AE49998802DABC816EC10158D3A7879702FF29084C2C992CD1289877A7\"},\"transaction\":{\"cosignatures\":[{\"signature\":\"5780C8DF9D46BA2BCF029DCC5D3BF55FE1CB5BE7ABCF30387C4637DDEDFC2152703CA0AD95F21BB9B942F3CC52FCFC2064C7B84CF60D1A9E69195F1943156C07\",\"signer\":\"A5F82EC8EBB341427B6785C8111906CD0DF18838FB11B51CE0E18B5E79DFF630\"}],\"deadline\":[3266625578,11],\"fee\":[0,0],\"signature\":\"939673209A13FF82397578D22CC96EB8516A6760C894D9B7535E3A1E068007B9255CFA9A914C97142A7AE18533E381C846B69D2AE0D60D1DC8A55AD120E2B606\",\"signer\":\"7681ED5023141D9CDCF184E5A7B60B7D466739918ED5DA30F7E71EA7B86EFF2D\",\"transactions\":[{\"meta\":{\"aggregateHash\":\"3D28C804EDD07D5A728E5C5FFEC01AB07AFA5766AE6997B38526D36015A4D006\",\"aggregateId\":\"5A0069D83F17CF0001777E55\",\"height\":[18160,0],\"id\":\"5A0069D83F17CF0001777E56\",\"index\":0},\"transaction\":{\"message\":{\"payload\":\"746573742D6D657373616765\",\"type\":0},\"mosaics\":[{\"amount\":[3863990592,95248],\"id\":[3646934825,3576016193]}],\"recipient\":\"9050B9837EFAB4BBE8A4B9BB32D812F9885C00D8FC1650E142\",\"signer\":\"B4F12E7C9F6946091E2CB8B6D3A12B50D17CCBBF646386EA27CE2946A7423DCF\",\"type\":16724,\"version\":36867}}],\"type\":16705,\"version\":36867}}"
	b := bytes.NewBufferString(rtx)
	tx, err := MapTransaction(b)
	atx := tx.(*AggregateTransaction)
	acc, err := NewAccountFromPrivateKey("26b64cb10f005e5988a36744ca19e20d835ccc7c105aaa5f3b212da593180930", MijinTest)
	ctx, err := NewCosignatureTransaction(atx)
	cstx, err := acc.SignCosignatureTransaction(ctx)

	if err != nil {
		t.Errorf("CosignatureTransaction signing returned error: %s", err)
	}

	if !reflect.DeepEqual(cstx.Signature, want1) {
		t.Errorf("CosignatureTransaction signing returned wrong signature: \n %s", cstx.Signature)
	}

}

func TestMosaicDefinitionTransactionSerialization(t *testing.T) {
	want := []byte{156, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 77, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 155, 138, 22, 28, 245, 9, 35, 144, 21, 153, 17, 174, 167, 46, 189, 60, 7, 1, 7, 4, 109, 111, 115, 97, 105, 99, 115, 2, 16, 39, 0, 0, 0, 0, 0, 0}

	tx, err := NewMosaicDefinitionTransaction(fakeDeadline, &MosaicId{FullName: "mosaics"}, &NamespaceId{FullName: "sname"}, NewMosaicProperties(true, true, true, 4, big.NewInt(10000)), MijinTest)

	b, err := tx.generateBytes()

	if err != nil {
		t.Errorf("MosaicDefinitionTransaction serialization returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("MosaicDefinitionTransaction serialization returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestMosaicSupplyChangeTransactionSerialization(t *testing.T) {
	want := []byte{137, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 77, 66, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 136, 105, 116, 110, 155, 26, 112, 87, 1, 10, 0, 0, 0, 0, 0, 0, 0}

	id := NewMosaicId(big.NewInt(6300565133566699912))
	tx, err := NewMosaicSupplyChangeTransaction(fakeDeadline, id, Increase, big.NewInt(10), MijinTest)

	b, err := tx.generateBytes()

	if err != nil {
		t.Errorf("MosaicSupplyChangeTransaction serialization returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("MosaicSupplyChangeTransaction serialization returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestTransferTransactionSerialization(t *testing.T) {
	want := []byte{165, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		3, 144, 84, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148, 236, 59, 165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172,
		1, 0, 1, 0, 103, 43, 0, 0, 206, 86, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0}
	tx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		NewPlainMessage(""),
		MijinTest,
	)
	b, err := tx.generateBytes()
	if err != nil {
		t.Errorf("TransaferTransaction generateBytes() returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("TransaferTransaction generateBytes() returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestTransferTransactionToAggregate(t *testing.T) {
	want1 := []int8{85, 0, 0, 0, -102, 73, 54, 100, 6, -84, -87, 82, -72, -117, -83, -11, -15, -23, -66, 108, -28, -106, -127,
		65, 3, 90, 96, -66, 80, 50, 115, -22, 101, 69, 107, 36, 3, -112, 84, 65, -112, -24, -2, -67, 103, 29, -44, 27, -18, -108, -20, 59,
		-91, -125, 28, -74, 8, -93, 18, -62, -14, 3, -70, -124, -84, 1, 0, 1, 0, 103, 43, 0, 0, -50, 86, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0}
	want := make([]byte, len(want1))
	for i, w := range want1 {
		want[i] = byte(w)
	}
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)
	tx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		NewPlainMessage(""),
		MijinTest,
	)
	tx.Signer = p
	b, err := toAggregateTransactionBytes(tx)
	if err != nil {
		t.Errorf("TransaferTransaction toAggregate() returned error: %s", err)
	}

	if !reflect.DeepEqual(b, []uint8(want)) {
		t.Errorf("TransaferTransaction toAggregate() returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestTransferTransactionSigning(t *testing.T) {
	want := "A5000000773891AD01DD4CDF6E3A55C186C673E256D7DF9D471846F1943CC3529E4E02B38B9AF3F8D13784645FF5FAAFA94A321B" +
		"94933C673D12DE60E4BC05ABA56F750E1026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF377550390544100000" +
		"00000000000010000000000000090E8FEBD671DD41BEE94EC3BA5831CB608A312C2F203BA84AC01000100672B0000CE56000064000000" +
		"00000000"
	want1 := "350AE56BC97DB805E2098AB2C596FA4C6B37EF974BF24DFD61CD9F77C7687424"
	a, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)

	tx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		NewPlainMessage(""),
		MijinTest,
	)
	stx, err := a.Sign(tx)
	if err != nil {
		t.Errorf("TransaferTransaction signing returned error: %s", err)
	}

	if !reflect.DeepEqual(stx.Payload, want) {
		t.Errorf("TransaferTransaction signing returned wrong payload: \n %s, want: \n %s", stx.Payload, want)
	}

	if !reflect.DeepEqual(stx.Hash, want1) {
		t.Errorf("TransaferTransaction signing returned wrong hash: \n %s, want: \n %s", stx.Hash, want1)
	}
}

func TestModifyMultisigAccountTransactionSerialization(t *testing.T) {
	want := []byte{189, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		3, 144, 85, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 0, 104, 179, 251, 177, 135, 41, 193, 253, 226, 37, 197, 127, 140, 224, 128, 250, 130, 143, 0, 103, 228, 81, 163, 253, 129, 250, 98, 136, 66, 176, 183, 99, 0, 207, 137, 63, 252, 196, 124, 51, 231, 246, 138, 177, 219, 86, 54, 92, 21, 107, 7, 54, 130, 74, 12, 30, 39, 63, 158, 0, 184, 223, 143, 1, 235}

	acc1, err := NewAccountFromPublicKey("68b3fbb18729c1fde225c57f8ce080fa828f0067e451a3fd81fa628842b0b763", MijinTest)
	acc2, err := NewAccountFromPublicKey("cf893ffcc47c33e7f68ab1db56365c156b0736824a0c1e273f9e00b8df8f01eb", MijinTest)
	tx, err := NewModifyMultisigAccountTransaction(
		fakeDeadline,
		2,
		1,
		[]*MultisigCosignatoryModification{
			{
				Add,
				acc1,
			},
			{
				Add,
				acc2,
			},
		},
		MijinTest,
	)

	b, err := tx.generateBytes()

	if err != nil {
		t.Errorf("MosaicDefinitionTransaction serialization returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("MosaicDefinitionTransaction serialization returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestRegisterRootNamespaceTransactionSerialization(t *testing.T) {
	want := []byte{150, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 78, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 16, 39, 0, 0, 0, 0, 0, 0, 126, 233, 179, 184, 175, 223, 83, 64, 12, 110, 101, 119, 110, 97, 109, 101, 115, 112, 97, 99, 101}
	tx, err := NewRegisterRootNamespaceTransaction(
		fakeDeadline,
		&NamespaceId{FullName: "newnamespace"},
		big.NewInt(10000),
		MijinTest,
	)
	b, err := tx.generateBytes()
	if err != nil {
		t.Errorf("RegisterRootNamespaceTransaction generateBytes() returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("RegisterRootNamespaceTransaction generateBytes() returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestRegisterSubNamespaceTransactionSerialization(t *testing.T) {
	want := []byte{150, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 78, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 126, 233, 179, 184, 175, 223, 83, 64, 3, 18, 152, 27, 120, 121, 163, 113, 12, 115, 117, 98, 110, 97, 109, 101, 115, 112, 97, 99, 101}

	tx, err := NewRegisterSubNamespaceTransaction(
		fakeDeadline,
		&NamespaceId{FullName: "subnamespace"},
		&NamespaceId{Id: big.NewInt(4635294387305441662)},
		MijinTest,
	)
	b, err := tx.generateBytes()
	if err != nil {
		t.Errorf("RegisterSubNamespaceTransaction generateBytes() returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("RegisterSubNamespaceTransaction generateBytes() returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestLockFundsTransactionSerialization(t *testing.T) {
	want := []byte{176, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 144, 76, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 41, 207, 95,
		217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 132,
		152, 179, 141, 137, 193, 220, 138, 68, 142, 165, 130, 73,
		56, 255, 130, 137, 38, 205, 159, 119, 71, 177, 132, 75, 89, 180, 182,
		128, 126, 135, 139}

	stx := &SignedTransaction{AggregateBonded, "payload", "8498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"}
	tx, err := NewLockFundsTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), stx, MijinTest)
	b, err := tx.generateBytes()
	if err != nil {
		t.Errorf("LockFundsTransaction generateBytes() returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("LockFundsTransaction generateBytes() returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestLockFundsTransactionToAggregate(t *testing.T) {
	want1 := []int16{96, 0, 0, 0, -102, 73, 54, 100, 6, -84, -87, 82, -72, -117, -83, -11, -15, -23, -66, 108, -28, -106, -127,
		65, 3, 90, 96, -66, 80, 50, 115, -22, 101, 69, 107, 36, 3, 144, 76, 65, 41, 207, 95,
		217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 132,
		152, 179, 141, 137, 193, 220, 138, 68, 142, 165, 130, 73,
		56, 255, 130, 137, 38, 205, 159, 119, 71, 177, 132, 75, 89, 180, 182,
		128, 126, 135, 139}
	want := make([]byte, len(want1))
	for i, w := range want1 {
		want[i] = byte(w)
	}
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)
	stx := &SignedTransaction{AggregateBonded, "payload", "8498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"}
	tx, err := NewLockFundsTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), stx, MijinTest)
	tx.Signer = p
	b, err := toAggregateTransactionBytes(tx)
	if err != nil {
		t.Errorf("LockFundsTransaction toAggregate returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("LockFundsTransaction toAggregate returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestLockFundsTransactionSigning(t *testing.T) {
	want := "B0000000D079047B87DCEDA0DE68558C1322A453D55D52BDA2778D66C5344BF79EE9E946C731F9ED565E5A854AFC0A1E1476B571940F920F33ADD9BAC245DB46A59794051026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF3775503904C410000000000000000010000000000000029CF5FD941AD25D5809698000000000064000000000000008498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"

	acc, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)

	stx := &SignedTransaction{AggregateBonded, "payload", "8498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"}
	tx, err := NewLockFundsTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), stx, MijinTest)
	b, err := signTransactionWith(tx, acc)
	if err != nil {
		t.Errorf("LockFundsTransaction signing returned error: %s", err)
	}

	if !reflect.DeepEqual(b.Payload, want) {
		t.Errorf("LockFundsTransaction signing returned wrong result: \n %+v, want: \n %+v", b.Payload, want)
	}

	if !reflect.DeepEqual(b.Hash, "1F8A695B23F595646D43307DE0C6487AC642520FD31ACC6E6F8163AD2DD98B5A") {
		t.Errorf("LockFundsTransaction signing returned wrong result: \n %+v, want: \n %+v", b.Hash, "1F8A695B23F595646D43307DE0C6487AC642520FD31ACC6E6F8163AD2DD98B5A")
	}
}

func TestSecretLockTransactionSerialization(t *testing.T) {
	want := []byte{234, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 3, 144, 76, 66, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 41, 207,
		95, 217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 183, 120,
		163, 154, 54, 99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249, 223, 83, 135, 130, 191, 25, 156, 24,
		157, 171, 234, 199, 104, 10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250, 154, 246, 255,
		222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89, 160, 1, 177, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148,
		236, 59, 165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172}
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	ad, err := NewAddressFromRaw("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM")
	tx, err := NewSecretLockTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), SHA3_512, s, ad, MijinTest)
	b, err := tx.generateBytes()
	if err != nil {
		t.Errorf("SecretLockTransaction generateBytes() returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("SecretLockTransaction generateBytes() returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestSecretLockTransactionToAggregate(t *testing.T) {
	want1 := []int16{-102, 0, 0, 0, -102, 73, 54, 100, 6, -84, -87, 82, -72, -117, -83, -11, -15, -23, -66, 108, -28, -106, -127,
		65, 3, 90, 96, -66, 80, 50, 115, -22, 101, 69, 107, 36, 3, 144, 76, 66, 41, 207,
		95, 217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 183, 120,
		163, 154, 54, 99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249, 223, 83, 135, 130, 191, 25, 156, 24,
		157, 171, 234, 199, 104, 10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250, 154, 246, 255,
		222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89, 160, 1, 177, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148,
		236, 59, 165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172}
	want := make([]byte, len(want1))
	for i, w := range want1 {
		want[i] = byte(w)
	}
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)
	ad, err := NewAddressFromRaw("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM")
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	tx, err := NewSecretLockTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), SHA3_512, s, ad, MijinTest)
	tx.Signer = p
	b, err := toAggregateTransactionBytes(tx)
	if err != nil {
		t.Errorf("SecretLockTransaction toAggregate returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("SecretLockTransaction toAggregate returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestSecretLockTransactionSigning(t *testing.T) {
	want := "EA0000005A3B75AE172855381353250EA9A1DFEB86E9280C0006B8FD997C2FCECF211C9A260E76CB704A22EAD4648F18E6931381921A4EDC7D309C32275D0147E9BAD3051026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF3775503904C420000000000000000010000000000000029CF5FD941AD25D58096980000000000640000000000000000B778A39A3663719DFC5E48C9D78431B1E45C2AF9DF538782BF199C189DABEAC7680ADA57DCEC8EEE91C4E3BF3BFA9AF6FFDE90CD1D249D1C6121D7B759A001B190E8FEBD671DD41BEE94EC3BA5831CB608A312C2F203BA84AC"
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	acc, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)
	ad, err := NewAddressFromRaw("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM")
	tx, err := NewSecretLockTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), SHA3_512, s, ad, MijinTest)
	b, err := acc.Sign(tx)
	if err != nil {
		t.Errorf("SecretLockTransaction signing returned error: %s", err)
	}

	if !reflect.DeepEqual(b.Payload, want) {
		t.Errorf("SecretLockTransaction signing returned wrong result: \n %+v, want: \n %+v", b.Payload, want)
	}
	if !reflect.DeepEqual(b.Hash, "B3AF46027909CD24204AF4E7B5B43C3116307D90A1F83A5DE6DBDF1F7759ABC5") {
		t.Errorf("SecretLockTransaction signing returned wrong hash: \n %+v, want: \n %+v", b.Hash, "B3AF46027909CD24204AF4E7B5B43C3116307D90A1F83A5DE6DBDF1F7759ABC5")
	}
}

func TestSecretProofTransactionSerialization(t *testing.T) {
	want := []byte{191, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 3, 144, 76, 67, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 183, 120, 163, 154, 54,
		99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249,
		223, 83, 135, 130, 191, 25, 156, 24, 157, 171, 234, 199, 104,
		10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250,
		154, 246, 255, 222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89,
		160, 1, 177, 4, 0, 154, 73, 54, 100}
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	ss := "9a493664"

	tx, err := NewSecretProofTransaction(fakeDeadline, SHA3_512, s, ss, MijinTest)

	b, err := tx.generateBytes()
	if err != nil {
		t.Errorf("SecretProofTransaction serialization returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("ecretProofTransaction serialization returned wrong result: \n %+v, want: \n %+v", b, want)
	}

}

func TestSecretProofTransactionToAggregate(t *testing.T) {
	want1 := []int16{111, 0, 0, 0, -102, 73, 54, 100, 6, -84, -87, 82, -72, -117, -83, -11, -15, -23, -66, 108, -28, -106, -127,
		65, 3, 90, 96, -66, 80, 50, 115, -22, 101, 69, 107, 36, 3, 144, 76, 67, 0, 183, 120, 163, 154, 54,
		99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249,
		223, 83, 135, 130, 191, 25, 156, 24, 157, 171, 234, 199, 104,
		10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250,
		154, 246, 255, 222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89,
		160, 1, 177, 4, 0, 154, 73, 54, 100}
	want := make([]byte, len(want1))
	for i, w := range want1 {
		want[i] = byte(w)
	}
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	ss := "9a493664"
	tx, err := NewSecretProofTransaction(fakeDeadline, SHA3_512, s, ss, MijinTest)
	tx.Signer = p
	b, err := toAggregateTransactionBytes(tx)
	if err != nil {
		t.Errorf("SecretProofTransaction toAggregate returned error: %s", err)
	}

	if !reflect.DeepEqual(b, want) {
		t.Errorf("SecretProofTransaction toAggregate returned wrong result: \n %+v, want: \n %+v", b, want)
	}
}

func TestSecretProofTransactionSigning(t *testing.T) {
	want := "BF000000147827E5FDAB2201ABD3663964B0493166DA7DD18497718F53DF09AAFC55271B57A9E81B4E2F627FD19E9E9B77283D1620FB8E9E32BAC5AC265EB0B43C75B4071026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF3775503904C430000000000000000010000000000000000B778A39A3663719DFC5E48C9D78431B1E45C2AF9DF538782BF199C189DABEAC7680ADA57DCEC8EEE91C4E3BF3BFA9AF6FFDE90CD1D249D1C6121D7B759A001B104009A493664"
	acc, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	ss := "9a493664"

	tx, err := NewSecretProofTransaction(fakeDeadline, SHA3_512, s, ss, MijinTest)

	b, err := signTransactionWith(tx, acc)
	if err != nil {
		t.Errorf("SecretProofTransaction signing returned error: %s", err)
	}

	if !reflect.DeepEqual(b.Payload, want) {
		t.Errorf("SecretProofTransaction signing returned wrong result: \n %+v, want: \n %+v", b.Payload, want)
	}
}

func TestDeadline(t *testing.T) {
	if !time.Now().Before(NewDeadline(time.Hour * 2).Time) {
		t.Error("now is before deadline localtime")
	}
	if !time.Now().Add(time.Hour * 2).Add(-time.Second).Before(NewDeadline(time.Hour * 2).Time) {
		t.Error("now plus 2 hours is before deadline localtime")
	}
	if !time.Now().Add(time.Hour * 2).Add(time.Second * 2).After(NewDeadline(time.Hour * 2).Time) {
		t.Error("now plus 2 hours and 2 seconds is after deadline localtime")
	}
}
