// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"time"
)

const (
	baseUrl     = "http://localhost:3000"
	networkType = sdk.MijinTest
	privateKey  = "0F3CC33190A49ABB32E7172E348EA927F975F8829107AAA3D6349BB10797D4F6"
)

// WebSockets make possible receiving notifications when a transaction or event occurs in the blockchain.
// The notification is received in real time without having to poll the API waiting for a reply.
func main() {

	conf, err := sdk.NewConfig(baseUrl,networkType)
	if err != nil {
		panic(err)
	}

	acc, err := sdk.NewAccountFromPrivateKey(privateKey, networkType)

	ws, err := sdk.NewConnectWs(baseUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("websocket negotiated uid:", ws.Uid)

	// The UnconfirmedAdded channel notifies when a transaction related to an
	// address is in unconfirmed state and waiting to be included in a block.
	// The message contains the transaction.
	chUnconfirmedAdded, _ := ws.Subscribe.UnconfirmedAdded(acc.Address.Address)
	go func() {
		for {
			data := <-chUnconfirmedAdded.Ch
			fmt.Printf("UnconfirmedAdded Tx Hash: %v \n", data.GetAbstractTransaction().Hash)
			chUnconfirmedAdded.Unsubscribe()
		}
	}()
	//
	//// The confirmedAdded channel notifies when a transaction related to an
	//// address is included in a block. The message contains the transaction.
	chConfirmedAdded, _ := ws.Subscribe.ConfirmedAdded(acc.Address.Address)
	go func() {
		for {
			data := <-chConfirmedAdded.Ch
			fmt.Printf("ConfirmedAdded Tx Hash: %v \n", data.GetAbstractTransaction().Hash)
			chConfirmedAdded.Unsubscribe()
			fmt.Println("Successful transfer!")
		}
	}()

	//The status channel notifies when a transaction related to an address rises an error.
	//The message contains the error message and the transaction hash.
	chStatus, _ := ws.Subscribe.Status(acc.Address.Address)

	go func() {
		for {
			data := <-chStatus.Ch
			chStatus.Unsubscribe()
			fmt.Printf("Hash: %v \n", data.Hash)
			panic(fmt.Sprint("Status: ", data.Status))
		}
	}()

	time.Sleep(time.Second * 5)
	// Use the default http client
	client := sdk.NewClient(nil, conf)

	ttx, err := sdk.NewTransferTransaction(
		sdk.NewDeadline(time.Hour*1),
		sdk.NewAddress("SBILTA367K2LX2FEXG5TFWAS7GEFYAGY7QLFBYKC", networkType),
		sdk.Mosaics{sdk.Xem(10000000)},
		sdk.NewPlainMessage(""),
		networkType,
	)

	stx, err := acc.Sign(ttx)
	if err != nil {
		panic(fmt.Errorf("TransaferTransaction signing returned error: %s", err))
	}

	// Get the chain height
	restTx, resp, err := client.Transaction.Announce(context.Background(), stx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", restTx)
	fmt.Printf("Response Status Code == %d\n\n", resp.StatusCode)
	fmt.Printf("Hash: \t\t%v\n", stx.Hash)
	fmt.Printf("Signer: \t%X\n\n", acc.KeyPair.PublicKey.Raw)

	// The block channel notifies for every new block.
	// The message contains the block information.
	chBlock, _ := ws.Subscribe.Block()

	for {
		data := <-chBlock.Ch
		fmt.Printf("Block received with height: %v \n", data.Height)
	}
}
