package sdk

import (
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

func TestTransferTransactionSerialization(t *testing.T) {
	want := []byte{byte(165), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		3, byte(144), 84, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, byte(144), byte(232), byte(254), byte(189), byte(103), byte(29), byte(212), byte(27), byte(238), byte(148), byte(236), byte(59), byte(165), byte(131), byte(28), byte(182), byte(8), byte(163), byte(18), byte(194), byte(242), byte(3), byte(186), byte(132), byte(172),
		1, 0, 1, 0, 103, 43, 0, 0, byte(206), 86, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0}
	tx := &TransferTransaction{
		Recipient: NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics:   Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		Message:   &Message{0, ""},
		abstractTransaction: abstractTransaction{
			Deadline:    &Deadline{time.Unix(0, 1459468801*int64(time.Millisecond))},
			NetworkType: MijinTest,
			Version:     3,
			Type:        Transfer,
		},
	}
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
	p, err := NewPublicAccount("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)
	tx := &TransferTransaction{
		Recipient: NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics:   Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		Message:   &Message{0, ""},
		abstractTransaction: abstractTransaction{
			Deadline:    &Deadline{time.Unix(0, 1459468801*int64(time.Millisecond))},
			NetworkType: MijinTest,
			Version:     3,
			Type:        Transfer,
			Signer:      p,
		},
	}
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
	a, err := NewAccount("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)
	tx := &TransferTransaction{
		Recipient: NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics:   Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		Message:   &Message{0, ""},
		abstractTransaction: abstractTransaction{
			Deadline:    &Deadline{time.Unix(0, 1459468801*int64(time.Millisecond))},
			NetworkType: MijinTest,
			Version:     3,
			Type:        Transfer,
		},
	}
	stx, err := SignTransaction(tx, a)
	if err != nil {
		t.Errorf("TransaferTransaction signing returned error: %s", err)
	}

	if !reflect.DeepEqual(stx.Payload, want) {
		t.Errorf("TransaferTransaction signing returned wrong payload: \n %s, want: \n %s", a, want)
	}

	if !reflect.DeepEqual(stx.Hash, want1) {
		t.Errorf("TransaferTransaction signing returned wrong hash: \n %s, want: \n %s", a, want)
	}
}
