// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.
package test

import (
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"golang.org/x/net/context"
	"math/big"
	"testing"
)

const (
	iter     = 1000
	pageSize = 32
)

func TestMosaicService_GetMosaicsFromNamespaceExt(t *testing.T) {
	cfg, _ := sdk.NewConfig("http://190.216.224.11:3000", sdk.MijinTest)
	ctx := context.TODO()

	serv := sdk.NewClient(nil, cfg)
	h, err := serv.Blockchain.GetBlockchainHeight(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for i := uint64(1); i < h.Uint64() && i <= iter; i++ {

		h := big.NewInt(int64(i))
		trans, err := serv.Blockchain.GetBlockTransactions(ctx, h)
		if err != nil {
			t.Fatal(err)
			continue
		}

		if len(trans) == 0 {
			t.Logf("%d block, empty transactiona", h)
		}

		for j, val := range trans {
			if val == nil {
				t.Logf("empty trans #%d", j)
				continue
			}
			//t.Log(val.String())
			switch val.GetAbstractTransaction().Type {
			case sdk.MosaicDefinition:
				tran := val.(*sdk.MosaicDefinitionTransaction)

				if tran.NamespaceId == nil {
					t.Logf("empty nsId or MosaicId")
					t.Log(tran)
					continue
				}
				mscInfoArr, err := serv.Mosaic.GetMosaicsFromNamespaceUpToMosaic(ctx, tran.NamespaceId, tran.MosaicId, pageSize)
				if err != nil {
					t.Error(err)
				}

				for _, mscInfo := range mscInfoArr {
					t.Logf("%+v", mscInfo)
				}
			case sdk.MosaicSupplyChange:
				tran := val.(*sdk.MosaicSupplyChangeTransaction)

				if tran.MosaicId == nil {
					t.Logf("empty MosaicId")
					t.Log(tran)
					continue
				}
				mscInfo, err := serv.Mosaic.GetMosaic(ctx, tran.MosaicId)
				if err != nil {
					t.Error(err)
				}

				t.Logf("%+v", mscInfo)
			case sdk.Transfer:
				tran := val.(*sdk.TransferTransaction)
				if tran.Mosaics == nil {
					t.Logf("empty Mosaics")
					t.Log(tran)
					continue
				}
				mosaicIDs := make([]*sdk.MosaicId, len(tran.Mosaics))
				for _, val := range tran.Mosaics {
					mosaicIDs = append(mosaicIDs, val.MosaicId)
				}
				mscInfoArr, err := serv.Mosaic.GetMosaicNames(ctx, mosaicIDs)
				if err != nil {
					t.Error(err)
				}

				for _, mscInfo := range mscInfoArr {
					t.Logf("%+v", mscInfo)
				}
			case sdk.RegisterNamespace:
				tran := val.(*sdk.RegisterNamespaceTransaction)
				nsInfo, err := serv.Namespace.GetNamespace(ctx, tran.NamespaceId)
				if err != nil {
					t.Error(err)
				}

				t.Logf("%#v", nsInfo)
			default:
				t.Log(val)
			}
		}

	}
}
