package sdk

import (
	"context"
	"reflect"
	"testing"
)

const transactionId = "5B55E02EACCB7B00015DB6E1"
const transactionHash = "7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F"

var transaction = &TransferTransaction{
	AbstractTransaction: AbstractTransaction{
		Type:        TRANSFER,
		Version:     uint64(3),
		NetworkType: MIJIN_TEST,
		Signature:   "A036C1F27D1DE649BA783AE7A984AFA2CAFC2E4888E76009BF5B6146468E898F391A0EE7FFF8B65507FD5245C0967510133453C015B37DADED561F4380707507",
		Signer:      "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
		Fee:         []uint64{0, 0},
		Deadline:    []uint64{1, 0},
		TransactionInfo: TransactionInfo{
			Height:              []uint64{1, 0},
			Hash:                "7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F",
			MerkleComponentHash: "7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F",
			Index:               7,
			Id:                  "5B55E02EACCB7B00015DB6E1",
		},
	},
	Mosaics: MosaicsDTO{{
		MosaicId: []uint64{3646934825, 3576016193},
		Amount:   []uint64{3863990592, 95248},
	}},
	Address: "903AE571BCD5C85B746F48B4506C8B322AB8B186F48D3B0BC5",
}

var status = &TransactionStatus{
	"confirmed",
	"Success",
	"7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F",
	[]uint64{1, 0},
	[]uint64{1, 0},
}

func TestTransactionService_GetTransaction_TransferTransaction(t *testing.T) {
	client, _ := setup()

	tx, _, err := client.Transaction.GetTransaction(context.Background(), transactionId)
	if err != nil {
		t.Errorf("Transaction.GetTransaction returned error: %s", err)
	} else {

		want := transaction
		if !reflect.DeepEqual(tx, want) {
			t.Errorf("Transaction.GetTransaction returned %s, want %s", tx, want)
		}
	}
}

func TestTransactionService_GetTransactions(t *testing.T) {
	client, _ := setup()

	tx, _, err := client.Transaction.GetTransactions(context.Background(), []string{
		transactionId,
	})

	if err != nil {
		t.Errorf("Transaction.GetTransaction returned error: %v", err)
	}

	want := []Transaction{
		transaction,
	}
	if !reflect.DeepEqual(tx, want) {
		t.Errorf("Transaction.GetTransaction returned %s, want %s", tx, want)
	}
}

func TestTransactionService_GetTransactionStatus(t *testing.T) {
	client, _ := setup()

	tx, _, err := client.Transaction.GetTransactionStatus(context.Background(), transactionHash)
	if err != nil {
		t.Errorf("Transaction.GetTransaction returned error: %s", err)
	}

	if !reflect.DeepEqual(tx, status) {
		t.Errorf("Transaction.GetTransaction returned %s, want %s", tx, status)
	}
}

func TestTransactionService_GetTransactionStatuses(t *testing.T) {
	client, _ := setup()

	tx, _, err := client.Transaction.GetTransactionStatuses(context.Background(), []string{transactionHash})
	if err != nil {
		t.Errorf("Transaction.GetTransaction returned error: %s", err)
	}

	want := []*TransactionStatus{status}
	if !reflect.DeepEqual(tx, want) {
		t.Errorf("Transaction.GetTransaction returned %s, want %s", tx, want)
	}
}
