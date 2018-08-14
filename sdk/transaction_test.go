package sdk

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const transactionId = "5B55E02EACCB7B00015DB6E1"
const transactionHash = "7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F"

var transaction = &TransferTransaction{
	AbstractTransaction: AbstractTransaction{
		Type:        TRANSFER,
		Version:     uint64(3),
		NetworkType: MIJIN_TEST,
		Signature:   "ADF80CBC864B65A8D94205E9EC6640FA4AE0E3011B27F8A93D93761E454A9853BF0AB1ECB3DF62E1D2D267D3F1913FAB0E2225CE5EA3937790B78FFA1288870C",
		Signer:      &PublicAccount{&Address{MIJIN_TEST, "SBJ5D7TFIJWPY56JBEX32MUWI5RU6KVKZYITQ2HA"}, "27F6BEF9A7F75E33AE2EB2EBA10EF1D6BEA4D30EBD5E39AF8EE06E96E11AE2A9"},
		Fee:         uint64DTO{0, 0}.GetBigInteger(),
		Deadline:    time.Unix(uint64DTO{1094650402, 17}.GetBigInteger().Int64(), int64(time.Millisecond)),
		TransactionInfo: &TransactionInfo{
			Height:              uint64DTO{42, 0}.GetBigInteger(),
			Hash:                "45AC1259DABD7163B2816232773E66FC00342BB8DD5C965D4B784CD575FDFAF1",
			MerkleComponentHash: "45AC1259DABD7163B2816232773E66FC00342BB8DD5C965D4B784CD575FDFAF1",
			Index:               0,
			Id:                  "5B686E97F0C0EA00017B9437",
		},
	},
	Mosaics: Mosaics{
		&Mosaic{&MosaicId{uint64DTO{3646934825, 3576016193}.GetBigInteger(), ""}, uint64DTO{10000000, 0}.GetBigInteger()},
	},
	Address: &Address{MIJIN_TEST, "SBJUINHAC3FKCMVLL2WHBQFPPXYEHOMQY6E2SPVR"},
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
	"confirmed",
	"Success",
	"7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F",
	time.Unix(uint64DTO{1, 0}.GetBigInteger().Int64(), int64(time.Millisecond)),
	uint64DTO{1, 0}.GetBigInteger(),
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
