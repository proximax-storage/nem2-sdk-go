// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

const transactionId = "5B55E02EACCB7B00015DB6E1"
const transactionHash = "7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F"

var transaction = &TransferTransaction{
	AbstractTransaction: AbstractTransaction{
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

var (
	aggregateTransactionSerializationCorr = []byte{209, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 65, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 85, 0, 0, 0, 85, 0, 0, 0, 132, 107, 68, 57, 21, 69, 121, 165, 144, 59, 20, 89,
		201, 207, 105, 203, 129, 83, 246, 208, 17, 10, 122, 14, 214, 29, 226, 154,
		228, 129, 11, 242, 3, 144, 84, 65, 144, 80, 185, 131, 126, 250, 180,
		187, 232, 164, 185, 187, 50, 216, 18, 249, 136, 92, 0, 216, 252,
		22, 80, 225, 66, 1, 0, 1, 0, 41, 207, 95, 217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0}

	cosisignatureTransactionSigningCorr = "bf3bc39f2292c028cb0ffa438a9f567a7c4d793d2f8522c8deac74befbcb61af6414adf27b2176d6a24fef612aa6db2f562176a11c46ba6d5e05430042cb5705"

	mosaicDefinitionTransactionSerializationCorr = []byte{156, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 77, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 155, 138, 22, 28, 245, 9, 35, 144, 21, 153, 17, 174, 167, 46, 189, 60, 7, 1, 7, 4, 109, 111, 115, 97, 105, 99, 115, 2, 16, 39, 0, 0, 0, 0, 0, 0}

	mosaicSupplyChangeTransactionSerializationCorr = []byte{137, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 77, 66, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 136, 105, 116, 110, 155, 26, 112, 87, 1, 10, 0, 0, 0, 0, 0, 0, 0}

	transferTransactionSerializationCorr = []byte{165, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		3, 144, 84, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148, 236, 59, 165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172,
		1, 0, 1, 0, 103, 43, 0, 0, 206, 86, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0}

	transferTransactionToAggregateCorr = []byte{85, 0, 0, 0, 154, 73, 54, 100, 6, 172, 169, 82, 184, 139, 173, 245, 241, 233, 190, 108, 228, 150, 129, 65, 3, 90, 96, 190, 80, 50, 115, 234,
		101, 69, 107, 36, 3, 144, 84, 65, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148, 236, 59, 165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172, 1, 0, 1, 0, 103, 43, 0, 0, 206, 86, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0}

	transferTransactionSigningCorr = "A5000000773891AD01DD4CDF6E3A55C186C673E256D7DF9D471846F1943CC3529E4E02B38B9AF3F8D13784645FF5FAAFA94A321B" +
		"94933C673D12DE60E4BC05ABA56F750E1026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF377550390544100000" +
		"00000000000010000000000000090E8FEBD671DD41BEE94EC3BA5831CB608A312C2F203BA84AC01000100672B0000CE56000064000000" +
		"00000000"

	modifyMultisigAccountTransactionSerializationCorr = []byte{189, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		3, 144, 85, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 0, 104, 179, 251, 177, 135, 41, 193, 253, 226, 37, 197, 127, 140, 224, 128, 250, 130, 143, 0, 103, 228, 81, 163, 253, 129, 250, 98, 136, 66, 176, 183, 99, 0, 207, 137, 63, 252, 196, 124, 51, 231, 246, 138, 177, 219, 86, 54, 92, 21, 107, 7, 54, 130, 74, 12, 30, 39, 63, 158, 0, 184, 223, 143, 1, 235}

	registerRootNamespaceTransactionSerializationCorr = []byte{150, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 78, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 16, 39, 0, 0, 0, 0, 0, 0, 126, 233, 179, 184, 175, 223, 83, 64, 12, 110, 101, 119, 110, 97, 109, 101, 115, 112, 97, 99, 101}

	registerSubNamespaceTransactionSerializationCorr = []byte{150, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 144, 78, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 126, 233, 179, 184, 175, 223, 83, 64, 3, 18, 152, 27, 120, 121, 163, 113, 12, 115, 117, 98, 110, 97, 109, 101, 115, 112, 97, 99, 101}

	lockFundsTransactionSerializationCorr = []byte{176, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 144, 76, 65, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 41, 207, 95,
		217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 132,
		152, 179, 141, 137, 193, 220, 138, 68, 142, 165, 130, 73,
		56, 255, 130, 137, 38, 205, 159, 119, 71, 177, 132, 75, 89, 180, 182,
		128, 126, 135, 139}

	secretProofTransactionSigningCorr = "BF000000147827E5FDAB2201ABD3663964B0493166DA7DD18497718F53DF09AAFC55271B57A9E81B4E2F627FD19E9E9B77283D1620FB8E9E32BAC5AC265EB0B43C75B4071026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF3775503904C430000000000000000010000000000000000B778A39A3663719DFC5E48C9D78431B1E45C2AF9DF538782BF199C189DABEAC7680ADA57DCEC8EEE91C4E3BF3BFA9AF6FFDE90CD1D249D1C6121D7B759A001B104009A493664"

	secretProofTransactionToAggregateCorr = []byte{111, 0, 0, 0, 154, 73, 54, 100, 6, 172, 169, 82, 184, 139, 173, 245, 241, 233, 190, 108, 228, 150, 129, 65, 3, 90, 96, 190, 80, 50, 115, 234, 101, 69, 107, 36, 3, 144, 76, 67, 0, 183, 120, 163, 154, 54, 99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249, 223, 83, 135, 130, 191, 25, 156, 24, 157, 171, 234, 199, 104, 10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250, 154, 246, 255, 222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89, 160, 1, 177, 4, 0, 154, 73, 54, 100}

	secretProofTransactionSerializationCorr = []byte{191, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 3, 144, 76, 67, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 183, 120, 163, 154, 54,
		99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249,
		223, 83, 135, 130, 191, 25, 156, 24, 157, 171, 234, 199, 104,
		10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250,
		154, 246, 255, 222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89,
		160, 1, 177, 4, 0, 154, 73, 54, 100}

	secretLockTransactionToAggregateCorr = []byte{154, 0, 0, 0, 154, 73, 54, 100, 6, 172, 169, 82, 184, 139, 173, 245, 241, 233, 190, 108, 228, 150, 129, 65, 3, 90, 96, 190, 80, 50, 115, 234, 101, 69, 107, 36, 3, 144, 76, 66, 41, 207, 95, 217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 183, 120, 163, 154, 54, 99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249, 223, 83, 135, 130, 191, 25, 156, 24, 157, 171, 234, 199, 104, 10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250, 154, 246, 255, 222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89, 160, 1, 177, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148, 236, 59, 165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172}

	secretLockTransactionSerializationCorr = []byte{234, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 3, 144, 76, 66, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 41, 207,
		95, 217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 183, 120,
		163, 154, 54, 99, 113, 157, 252, 94, 72, 201, 215, 132, 49, 177, 228, 92, 42, 249, 223, 83, 135, 130, 191, 25, 156, 24,
		157, 171, 234, 199, 104, 10, 218, 87, 220, 236, 142, 238, 145, 196, 227, 191, 59, 250, 154, 246, 255,
		222, 144, 205, 29, 36, 157, 28, 97, 33, 215, 183, 89, 160, 1, 177, 144, 232, 254, 189, 103, 29, 212, 27, 238, 148,
		236, 59, 165, 131, 28, 182, 8, 163, 18, 194, 242, 3, 186, 132, 172}

	lockFundsTransactionToAggregateCorr = []byte{96, 0, 0, 0, 154, 73, 54, 100, 6, 172, 169, 82, 184, 139, 173, 245, 241, 233, 190, 108, 228, 150, 129, 65, 3, 90, 96, 190, 80, 50, 115, 234, 101, 69, 107, 36, 3, 144, 76, 65, 41, 207, 95, 217, 65, 173, 37, 213, 128, 150, 152, 0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 132, 152, 179, 141, 137, 193, 220, 138, 68, 142, 165, 130, 73, 56, 255, 130, 137, 38, 205, 159, 119, 71, 177, 132, 75, 89, 180, 182, 128, 126, 135, 139}

	secretLockTransactionSigningCorr = "EA0000005A3B75AE172855381353250EA9A1DFEB86E9280C0006B8FD997C2FCECF211C9A260E76CB704A22EAD4648F18E6931381921A4EDC7D309C32275D0147E9BAD3051026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF3775503904C420000000000000000010000000000000029CF5FD941AD25D58096980000000000640000000000000000B778A39A3663719DFC5E48C9D78431B1E45C2AF9DF538782BF199C189DABEAC7680ADA57DCEC8EEE91C4E3BF3BFA9AF6FFDE90CD1D249D1C6121D7B759A001B190E8FEBD671DD41BEE94EC3BA5831CB608A312C2F203BA84AC"

	lockFundsTransactionSigningCorr = "B0000000D079047B87DCEDA0DE68558C1322A453D55D52BDA2778D66C5344BF79EE9E946C731F9ED565E5A854AFC0A1E1476B571940F920F33ADD9BAC245DB46A59794051026D70E1954775749C6811084D6450A3184D977383F0E4282CD47118AF3775503904C410000000000000000010000000000000029CF5FD941AD25D5809698000000000064000000000000008498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"
)

func TestTransactionService_GetTransaction_TransferTransaction(t *testing.T) {
	mockServer.addRouter(&router{
		path:     fmt.Sprintf("/transaction/%s", transactionId),
		respBody: transactionJson,
	})

	cl := mockServer.getTestNetClientUnsafe()

	tx, _, err := cl.Transaction.GetTransaction(context.Background(), transactionId)

	assert.Nilf(t, err, "TransactionService.GetTransaction returned error: %v", err)
	validateStringers(t, transaction, tx)
}

func TestTransactionService_GetTransactions(t *testing.T) {
	mockServer.addRouter(&router{
		path:     "/transaction",
		respBody: "[" + transactionJson + "]",
	})

	cl := mockServer.getTestNetClientUnsafe()

	transactions, resp, err := cl.Transaction.GetTransactions(context.Background(), []string{
		transactionId,
	})

	assert.Nilf(t, err, "TransactionService.GetTransactions returned error: %v", err)
	validateResponse(t, resp)

	for _, tx := range transactions {
		validateStringers(t, transaction, tx)
	}
}

func TestTransactionService_GetTransactionStatus(t *testing.T) {
	mockServer.addRouter(&router{
		path:     "/transaction/7D354E056A10E7ADAC66741D1021B0E79A57998EAD7E17198821141CE87CF63F/status",
		respBody: statusJson,
	})

	cl := mockServer.getTestNetClientUnsafe()

	txStatus, resp, err := cl.Transaction.GetTransactionStatus(context.Background(), transactionHash)

	assert.Nilf(t, err, "TransactionService.GetTransactionStatus returned error: %v", err)
	validateResponse(t, resp)
	validateStringers(t, status, txStatus)
}

func TestTransactionService_GetTransactionStatuses(t *testing.T) {
	mockServer.addRouter(&router{
		path:     "/transaction/statuses",
		respBody: "[" + statusJson + "]",
	})

	cl := mockServer.getTestNetClientUnsafe()

	txStatuses, _, err := cl.Transaction.GetTransactionStatuses(context.Background(), []string{transactionHash})

	assert.Nilf(t, err, "TransactionService.GetTransactionStatuses returned error: %v", err)

	for _, txStatus := range txStatuses {
		validateStringers(t, status, txStatus)
	}
}

func TestAggregateTransactionSerialization(t *testing.T) {
	p, err := NewAccountFromPublicKey("846B4439154579A5903B1459C9CF69CB8153F6D0110A7A0ED61DE29AE4810BF2", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

	ttx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SBILTA367K2LX2FEXG5TFWAS7GEFYAGY7QLFBYKC", MijinTest),
		Mosaics{Xem(10000000)},
		NewPlainMessage(""),
		MijinTest,
	)

	assert.Nilf(t, err, "NewTransferTransaction returned error: %s", err)

	ttx.Signer = p

	atx, err := NewCompleteAggregateTransaction(fakeDeadline, []Transaction{ttx}, MijinTest)

	assert.Nilf(t, err, "NewCompleteAggregateTransaction returned error: %s", err)

	b, err := atx.generateBytes()

	assert.Nilf(t, err, "AggregateTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, aggregateTransactionSerializationCorr, b)
}

func TestAggregateTransactionSigningWithMultipleCosignatures(t *testing.T) {
	p, err := NewAccountFromPublicKey("B694186EE4AB0558CA4AFCFDD43B42114AE71094F5A1FC4A913FE9971CACD21D", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

	ttx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SBILTA367K2LX2FEXG5TFWAS7GEFYAGY7QLFBYKC", MijinTest),
		Mosaics{},
		NewPlainMessage("test-message"),
		MijinTest,
	)

	ttx.Signer = p

	atx, err := NewCompleteAggregateTransaction(fakeDeadline, []Transaction{ttx}, MijinTest)

	assert.Nilf(t, err, "NewCompleteAggregateTransaction returned error: %s", err)

	acc1, err := NewAccountFromPrivateKey("2a2b1f5d366a5dd5dc56c3c757cf4fe6c66e2787087692cf329d7a49a594658b", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPrivateKey returned error: %s", err)

	acc2, err := NewAccountFromPrivateKey("b8afae6f4ad13a1b8aad047b488e0738a437c7389d4ff30c359ac068910c1d59", MijinTest) // TODO from original repo: "bug with private key"

	assert.Nilf(t, err, "NewAccountFromPrivateKey returned error: %s", err)

	stx, err := acc1.SignWithCosignatures(atx, []*Account{acc2})

	assert.Nilf(t, err, "Account.SignWithCosignatures returned error: %s", err)
	assert.Equal(t, "2d010000", stx.Payload[0:8])
	assert.Equal(t, "5100000051000000", stx.Payload[240:256])

	//if !reflect.DeepEqual(stx.Payload[320:474], "039054419050B9837EFAB4BBE8A4B9BB32D812F9885C00D8FC1650E1420D000000746573742D6D65737361676568B3FBB18729C1FDE225C57F8CE080FA828F0067E451A3FD81FA628842B0B763") {
	//	t.Errorf("AggregateTransaction signing returned wrong payload: \n %s", stx.Payload[320:474])
	//} this test is not working in original repo and commented out too
}

func TestCosisignatureTransactionSigning(t *testing.T) {
	rtx := "{\"meta\":{\"hash\":\"671653C94E2254F2A23EFEDB15D67C38332AED1FBD24B063C0A8E675582B6A96\",\"height\":[18160,0],\"id\":\"5A0069D83F17CF0001777E55\",\"index\":0,\"merkleComponentHash\":\"81E5E7AE49998802DABC816EC10158D3A7879702FF29084C2C992CD1289877A7\"},\"transaction\":{\"cosignatures\":[{\"signature\":\"5780C8DF9D46BA2BCF029DCC5D3BF55FE1CB5BE7ABCF30387C4637DDEDFC2152703CA0AD95F21BB9B942F3CC52FCFC2064C7B84CF60D1A9E69195F1943156C07\",\"signer\":\"A5F82EC8EBB341427B6785C8111906CD0DF18838FB11B51CE0E18B5E79DFF630\"}],\"deadline\":[3266625578,11],\"fee\":[0,0],\"signature\":\"939673209A13FF82397578D22CC96EB8516A6760C894D9B7535E3A1E068007B9255CFA9A914C97142A7AE18533E381C846B69D2AE0D60D1DC8A55AD120E2B606\",\"signer\":\"7681ED5023141D9CDCF184E5A7B60B7D466739918ED5DA30F7E71EA7B86EFF2D\",\"transactions\":[{\"meta\":{\"aggregateHash\":\"3D28C804EDD07D5A728E5C5FFEC01AB07AFA5766AE6997B38526D36015A4D006\",\"aggregateId\":\"5A0069D83F17CF0001777E55\",\"height\":[18160,0],\"id\":\"5A0069D83F17CF0001777E56\",\"index\":0},\"transaction\":{\"message\":{\"payload\":\"746573742D6D657373616765\",\"type\":0},\"mosaics\":[{\"amount\":[3863990592,95248],\"id\":[3646934825,3576016193]}],\"recipient\":\"9050B9837EFAB4BBE8A4B9BB32D812F9885C00D8FC1650E142\",\"signer\":\"B4F12E7C9F6946091E2CB8B6D3A12B50D17CCBBF646386EA27CE2946A7423DCF\",\"type\":16724,\"version\":36867}}],\"type\":16705,\"version\":36867}}"
	b := bytes.NewBufferString(rtx)
	tx, err := MapTransaction(b)

	assert.Nilf(t, err, "MapTransaction returned error: %s", err)

	atx := tx.(*AggregateTransaction)

	acc, err := NewAccountFromPrivateKey("26b64cb10f005e5988a36744ca19e20d835ccc7c105aaa5f3b212da593180930", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPrivateKey returned error: %s", err)

	ctx, err := NewCosignatureTransaction(atx)

	assert.Nilf(t, err, "NewCosignatureTransaction returned error: %s", err)

	cstx, err := acc.SignCosignatureTransaction(ctx)

	assert.Nilf(t, err, "Account.SignCosignatureTransaction signing returned error: %s", err)
	assert.Equal(t, cosisignatureTransactionSigningCorr, cstx.Signature)
}

func TestMosaicDefinitionTransactionSerialization(t *testing.T) {
	tx, err := NewMosaicDefinitionTransaction(fakeDeadline, &MosaicId{FullName: "mosaics"}, &NamespaceId{FullName: "sname"}, NewMosaicProperties(true, true, true, 4, big.NewInt(10000)), MijinTest)

	assert.Nilf(t, err, "NewMosaicDefinitionTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "MosaicDefinitionTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, mosaicDefinitionTransactionSerializationCorr, b)
}

func TestMosaicSupplyChangeTransactionSerialization(t *testing.T) {
	id := NewMosaicId(big.NewInt(6300565133566699912))
	tx, err := NewMosaicSupplyChangeTransaction(fakeDeadline, id, Increase, big.NewInt(10), MijinTest)

	assert.Nilf(t, err, "NewMosaicSupplyChangeTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "MosaicSupplyChangeTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, mosaicSupplyChangeTransactionSerializationCorr, b)
}

func TestTransferTransactionSerialization(t *testing.T) {
	tx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		NewPlainMessage(""),
		MijinTest,
	)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "TransferTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, transferTransactionSerializationCorr, b)
}

func TestTransferTransactionToAggregate(t *testing.T) {
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

	tx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		NewPlainMessage(""),
		MijinTest,
	)

	assert.Nilf(t, err, "NewTransferTransaction returned error: %s", err)

	tx.Signer = p

	b, err := toAggregateTransactionBytes(tx)

	assert.Nilf(t, err, "toAggregateTransactionBytes returned error: %s", err)
	assert.Equal(t, transferTransactionToAggregateCorr, b)
}

func TestTransferTransactionSigning(t *testing.T) {
	a, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPrivateKey returned error: %s", err)

	tx, err := NewTransferTransaction(
		fakeDeadline,
		NewAddress("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM", MijinTest),
		Mosaics{{&MosaicId{Id: big.NewInt(95442763262823)}, big.NewInt(100)}},
		NewPlainMessage(""),
		MijinTest,
	)

	assert.Nilf(t, err, "NewTransferTransaction returned error: %s", err)

	stx, err := a.Sign(tx)

	assert.Nilf(t, err, "Account.Sign returned error: %s", err)
	assert.Equal(t, transferTransactionSigningCorr, stx.Payload)
	assert.Equal(t, "350AE56BC97DB805E2098AB2C596FA4C6B37EF974BF24DFD61CD9F77C7687424", stx.Hash.String())
}

func TestModifyMultisigAccountTransactionSerialization(t *testing.T) {
	acc1, err := NewAccountFromPublicKey("68b3fbb18729c1fde225c57f8ce080fa828f0067e451a3fd81fa628842b0b763", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

	acc2, err := NewAccountFromPublicKey("cf893ffcc47c33e7f68ab1db56365c156b0736824a0c1e273f9e00b8df8f01eb", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

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

	assert.Nilf(t, err, "NewModifyMultisigAccountTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "ModifyMultisigAccountTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, modifyMultisigAccountTransactionSerializationCorr, b)
}

func TestRegisterRootNamespaceTransactionSerialization(t *testing.T) {
	tx, err := NewRegisterRootNamespaceTransaction(
		fakeDeadline,
		&NamespaceId{FullName: "newnamespace"},
		big.NewInt(10000),
		MijinTest,
	)

	assert.Nilf(t, err, "NewRegisterRootNamespaceTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "RegisterNamespaceTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, registerRootNamespaceTransactionSerializationCorr, b)
}

func TestRegisterSubNamespaceTransactionSerialization(t *testing.T) {
	tx, err := NewRegisterSubNamespaceTransaction(
		fakeDeadline,
		&NamespaceId{FullName: "subnamespace"},
		&NamespaceId{Id: big.NewInt(4635294387305441662)},
		MijinTest,
	)

	assert.Nilf(t, err, "NewRegisterSubNamespaceTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "RegisterNamespaceTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, registerSubNamespaceTransactionSerializationCorr, b)
}

func TestLockFundsTransactionSerialization(t *testing.T) {
	stx := &SignedTransaction{AggregateBonded, "payload", "8498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"}

	tx, err := NewLockFundsTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), stx, MijinTest)

	assert.Nilf(t, err, "NewLockFundsTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "LockFundsTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, lockFundsTransactionSerializationCorr, b)
}

func TestLockFundsTransactionToAggregate(t *testing.T) {
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

	stx := &SignedTransaction{AggregateBonded, "payload", "8498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"}

	tx, err := NewLockFundsTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), stx, MijinTest)

	assert.Nilf(t, err, "NewLockFundsTransaction returned error: %s", err)

	tx.Signer = p

	b, err := toAggregateTransactionBytes(tx)

	assert.Nilf(t, err, "toAggregateTransactionBytes returned error: %s", err)
	assert.Equal(t, lockFundsTransactionToAggregateCorr, b)
}

func TestLockFundsTransactionSigning(t *testing.T) {
	acc, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPrivateKey returned error: %s", err)

	stx := &SignedTransaction{AggregateBonded, "payload", "8498B38D89C1DC8A448EA5824938FF828926CD9F7747B1844B59B4B6807E878B"}

	tx, err := NewLockFundsTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), stx, MijinTest)

	assert.Nilf(t, err, "NewLockFundsTransaction returned error: %s", err)

	b, err := signTransactionWith(tx, acc)

	assert.Nilf(t, err, "signTransactionWith returned error: %s", err)
	assert.Equal(t, lockFundsTransactionSigningCorr, b.Payload)
	assert.Equal(t, "1F8A695B23F595646D43307DE0C6487AC642520FD31ACC6E6F8163AD2DD98B5A", b.Hash.String())
}

func TestSecretLockTransactionSerialization(t *testing.T) {
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"

	ad, err := NewAddressFromRaw("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM")

	assert.Nilf(t, err, "NewAddressFromRaw returned error: %s", err)

	tx, err := NewSecretLockTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), SHA3_512, s, ad, MijinTest)

	assert.Nilf(t, err, "NewSecretLockTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "SecretLockTransaction.generateBytes returned error: %s", err)
	assert.Equal(t, secretLockTransactionSerializationCorr, b)
}

func TestSecretLockTransactionToAggregate(t *testing.T) {
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

	ad, err := NewAddressFromRaw("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM")

	assert.Nilf(t, err, "NewAddressFromRaw returned error: %s", err)

	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"

	tx, err := NewSecretLockTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), SHA3_512, s, ad, MijinTest)

	assert.Nilf(t, err, "NewSecretLockTransaction returned error: %s", err)

	tx.Signer = p

	b, err := toAggregateTransactionBytes(tx)

	assert.Nilf(t, err, "toAggregateTransactionBytes returned error: %s", err)
	assert.Equal(t, secretLockTransactionToAggregateCorr, b)
}

func TestSecretLockTransactionSigning(t *testing.T) {
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"

	acc, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPrivateKey returned error: %s", err)

	ad, err := NewAddressFromRaw("SDUP5PLHDXKBX3UU5Q52LAY4WYEKGEWC6IB3VBFM")

	assert.Nilf(t, err, "NewAddressFromRaw returned error: %s", err)

	tx, err := NewSecretLockTransaction(fakeDeadline, XemRelative(10), big.NewInt(100), SHA3_512, s, ad, MijinTest)

	assert.Nilf(t, err, "NewSecretLockTransaction returned error: %s", err)

	b, err := acc.Sign(tx)

	assert.Nilf(t, err, "Sign returned error: %s", err)
	assert.Equal(t, secretLockTransactionSigningCorr, b.Payload)
	assert.Equal(t, "B3AF46027909CD24204AF4E7B5B43C3116307D90A1F83A5DE6DBDF1F7759ABC5", b.Hash.String())
}

func TestSecretProofTransactionSerialization(t *testing.T) {
	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	ss := "9a493664"

	tx, err := NewSecretProofTransaction(fakeDeadline, SHA3_512, s, ss, MijinTest)

	assert.Nilf(t, err, "NewSecretProofTransaction returned error: %s", err)

	b, err := tx.generateBytes()

	assert.Nilf(t, err, "generateBytes returned error: %s", err)
	assert.Equal(t, secretProofTransactionSerializationCorr, b)
}

func TestSecretProofTransactionToAggregate(t *testing.T) {
	p, err := NewAccountFromPublicKey("9A49366406ACA952B88BADF5F1E9BE6CE4968141035A60BE503273EA65456B24", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPublicKey returned error: %s", err)

	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	ss := "9a493664"

	tx, err := NewSecretProofTransaction(fakeDeadline, SHA3_512, s, ss, MijinTest)

	assert.Nilf(t, err, "NewSecretProofTransaction returned error: %s", err)

	tx.Signer = p

	b, err := toAggregateTransactionBytes(tx)

	assert.Nilf(t, err, "toAggregateTransactionBytes returned error: %s", err)
	assert.Equal(t, secretProofTransactionToAggregateCorr, b)
}

func TestSecretProofTransactionSigning(t *testing.T) {
	acc, err := NewAccountFromPrivateKey("787225aaff3d2c71f4ffa32d4f19ec4922f3cd869747f267378f81f8e3fcb12d", MijinTest)

	assert.Nilf(t, err, "NewAccountFromPrivateKey returned error: %s", err)

	s := "b778a39a3663719dfc5e48c9d78431b1e45c2af9df538782bf199c189dabeac7680ada57dcec8eee91c4e3bf3bfa9af6ffde90cd1d249d1c6121d7b759a001b1"
	ss := "9a493664"

	tx, err := NewSecretProofTransaction(fakeDeadline, SHA3_512, s, ss, MijinTest)

	assert.Nilf(t, err, "NewSecretProofTransaction returned error: %s", err)

	b, err := signTransactionWith(tx, acc)

	assert.Nilf(t, err, "signTransactionWith returned error: %s", err)
	assert.Equal(t, secretProofTransactionSigningCorr, b.Payload)
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

func TestMapTransaction_MosaicDefinitionTransaction(t *testing.T) {
	txr := `{"transaction":{"signer":"9C2086FE49B7A00578009B705AD719DB7E02A27870C67966AAA40540C136E248","version":43010,"type":16717,"parentId":[2031553063,2912841636],"mosaicId":[1477049789,645988887],"properties":[{"key":0,"value":[2,0]},{"key":1,"value":[0,0]},{"key":2,"value":[5,0]}],"name":"storage_0"}}`

	_, err := MapTransaction(bytes.NewBuffer([]byte(txr)))

	assert.Nilf(t, err, "MapTransaction returned error: %s", err)
}
