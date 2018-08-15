package sdk

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var account = &AccountInfo{
	&Address{MIJIN_TEST, "SAONSOGFZZHNEIBRYXHDTDTBR2YSAXKTITRFHG2Y"},
	uint64DTO{1, 0}.toBigInt(),
	"F3824119C9F8B9E81007CAA0EDD44F098458F14503D7C8D7C24F60AF11266E57",
	uint64DTO{0, 0}.toBigInt(),
	uint64DTO{409090909, 0}.toBigInt(),
	uint64DTO{1, 0}.toBigInt(),
	Mosaics{
		&Mosaic{&MosaicId{uint64DTO{3646934825, 3576016193}.toBigInt(), ""}, uint64DTO{3863990592, 95248}.toBigInt()},
	},
}

const accountInfoJson = `
{  
   "meta":{  

   },
   "account":{  
      "address":"901CD938C5CE4ED22031C5CE398E618EB1205D5344E2539B58",
      "addressHeight":[  
         1,
         0
      ],
      "publicKey":"F3824119C9F8B9E81007CAA0EDD44F098458F14503D7C8D7C24F60AF11266E57",
      "publicKeyHeight":[  
         0,
         0
      ],
      "mosaics":[  
         {  
            "id":[  
               3646934825,
               3576016193
            ],
            "amount":[  
               3863990592,
               95248
            ]
         }
      ],
      "importance":[  
         409090909,
         0
      ],
      "importanceHeight":[  
         1,
         0
      ]
   }
}
`

func TestAccountService_GetAccountInfo(t *testing.T) {
	cl, mux, _, teardown, err := setupMockServer()
	if err != nil {
		t.Errorf("Account.GetAccountInfo error setting up mock server: %v", err)
	}
	defer teardown()

	mux.HandleFunc("/account/SAONSOGFZZHNEIBRYXHDTDTBR2YSAXKTITRFHG2Y", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, accountInfoJson)
	})

	acc, _, err := cl.Account.GetAccountInfo(context.Background(), &Address{MIJIN_TEST, "SAONSOGFZZHNEIBRYXHDTDTBR2YSAXKTITRFHG2Y"})
	if err != nil {
		t.Errorf("Account.GetAccountInfo returned error: %s", err)
	}

	if !reflect.DeepEqual(acc, account) {
		t.Errorf("Account.GetAccountInfo returned %s, want %s", acc, account)
	}
}

func TestAccountService_GetAccountsInfo(t *testing.T) {
	cl, mux, _, teardown, err := setupMockServer()
	if err != nil {
		t.Errorf("Account.GetAccountsInfo error setting up mock server: %v", err)
	}
	defer teardown()

	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "["+accountInfoJson+"]")
	})

	acc, _, err := cl.Account.GetAccountsInfo(
		context.Background(),
		[]*Address{{MIJIN_TEST, "SAONSOGFZZHNEIBRYXHDTDTBR2YSAXKTITRFHG2Y"}},
	)

	if err != nil {
		t.Errorf("Account.GetAccountsInfo returned error: %s", err)
	}

	if !reflect.DeepEqual(acc, []*AccountInfo{account}) {
		t.Errorf("Account.GetAccountsInfo returned %s, want %s", acc, account)
	}
}

func TestAccountService_Transactions(t *testing.T) {
	cl, mux, _, teardown, err := setupMockServer()
	if err != nil {
		t.Errorf("Account.Transactions error setting up mock server: %v", err)
	}
	defer teardown()

	mux.HandleFunc("/account/27F6BEF9A7F75E33AE2EB2EBA10EF1D6BEA4D30EBD5E39AF8EE06E96E11AE2A9/transactions", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "["+transactionJson+"]")
	})

	tx, _, err := cl.Account.Transactions(
		context.Background(),
		&PublicAccount{
			&Address{MIJIN_TEST, "SBJ5D7TFIJWPY56JBEX32MUWI5RU6KVKZYITQ2HA"},
			"27F6BEF9A7F75E33AE2EB2EBA10EF1D6BEA4D30EBD5E39AF8EE06E96E11AE2A9",
		},
		&AccountTransactionsOption{},
	)

	if err != nil {
		t.Errorf("Account.Transactions returned error: %s", err)
	}

	want := []Transaction{
		transaction,
	}

	if !reflect.DeepEqual(tx, want) {
		t.Errorf("Account.Transactions returned %s, want %s", tx, want)
	}
}
