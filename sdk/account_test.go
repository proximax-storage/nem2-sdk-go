// Copyright 2018 ProximaX Limited. All rights reserved. // Use of this source code is governed by the Apache 2.0 // license that can be found in the LICENSE file.
package sdk

import (
	"context"
	"fmt"
	"github.com/proximax-storage/proximax-utils-go/mock"
	"github.com/proximax-storage/proximax-utils-go/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	account = &AccountInfo{
		&Address{MijinTest, "SAONSOGFZZHNEIBRYXHDTDTBR2YSAXKTITRFHG2Y"},
		uint64DTO{1, 0}.toBigInt(),
		"F3824119C9F8B9E81007CAA0EDD44F098458F14503D7C8D7C24F60AF11266E57",
		uint64DTO{0, 0}.toBigInt(),
		uint64DTO{409090909, 0}.toBigInt(),
		uint64DTO{1, 0}.toBigInt(),
		Mosaics{
			&Mosaic{&MosaicId{uint64DTO{3646934825, 3576016193}.toBigInt(), ""}, uint64DTO{3863990592, 95248}.toBigInt()},
		},
	}

	accountClient = mockServer.getTestNetClientUnsafe().Account
)

const (
	accountInfoJson = `{  
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
)

var (
	nemTestAddress1 = "SAONSOGFZZHNEIBRYXHDTDTBR2YSAXKTITRFHG2Y"
	nemTestAddress2 = "SBJ5D7TFIJWPY56JBEX32MUWI5RU6KVKZYITQ2HA"
	publicKey1      = "27F6BEF9A7F75E33AE2EB2EBA10EF1D6BEA4D30EBD5E39AF8EE06E96E11AE2A9"
)

func TestAccountService_GetAccountInfo(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf("/account/%s", nemTestAddress1),
		RespBody: accountInfoJson,
	})

	acc, resp, err := accountClient.GetAccountInfo(context.Background(), &Address{MijinTest, nemTestAddress1})

	assert.Nilf(t, err, "AccountService.GetAccountInfo returned error: %s", err)

	if tests.IsOkResponse(t, resp) {
		tests.ValidateStringers(t, account, acc)
	}
}

func TestAccountService_GetAccountsInfo(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     "/account",
		RespBody: "[" + accountInfoJson + "]",
	})

	accounts, resp, err := accountClient.GetAccountsInfo(
		context.Background(),
		[]*Address{{MijinTest, nemTestAddress1}},
	)

	assert.Nilf(t, err, "AccountService.GetAccountsInfo returned error: %s", err)

	if tests.IsOkResponse(t, resp) {
		for _, acc := range accounts {
			tests.ValidateStringers(t, account, acc)
		}
	}
}

func TestAccountService_Transactions(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf("/account/%s/transactions", publicKey1),
		RespBody: "[" + transactionJson + "]",
	})

	transactions, resp, err := accountClient.Transactions(
		context.Background(),
		&PublicAccount{
			&Address{MijinTest, nemTestAddress2},
			publicKey1,
		},
		&AccountTransactionsOption{},
	)

	assert.Nilf(t, err, "AccountService.Transactions returned error: %s", err)

	if tests.IsOkResponse(t, resp) {
		for _, tx := range transactions {
			tests.ValidateStringers(t, transaction, tx)
		}
	}
}
