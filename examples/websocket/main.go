// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"sync"
	"time"
)

const (
	baseUrl          = "http://localhost:3000"
	networkType      = sdk.MijinTest
	senderPrivateKey = "C6DA6B91EB951E6A22908FC8BC7ADD1FD5AE0B73864745DB240474532C34EFD8"
	recipientAddress = "SBILTA367K2LX2FEXG5TFWAS7GEFYAGY7QLFBYKC"
)

// WebSockets make possible receiving notifications when a transaction or event occurs in the blockchain.
// The notification is received in real time without having to poll the API waiting for a reply.
func main() {
	var wg sync.WaitGroup

	wg.Add(4)

	conf, err := sdk.NewConfig(baseUrl, networkType)
	if err != nil {
		panic(err)
	}

	// Sender account
	accSender, err := sdk.NewAccountFromPrivateKey(senderPrivateKey, networkType)
	if err != nil {
		panic(err)
	}

	// Recipient account
	accRecipient, err := sdk.NewAddressFromRaw(recipientAddress)
	if err != nil {
		panic(err)
	}

	// timeout in milliseconds
	// 60000 ms = 60 seconds
	// 0 = without timeout
	ws, err := sdk.NewConnectWs(baseUrl, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("websocket negotiated uid:", ws.Uid)

	// The block channel notifies for every new block.
	// The message contains the block information.
	chBlock, err := ws.Subscribe.Block()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			data := <-chBlock.Ch
			fmt.Printf("Block received with height: %v \n\n", data.Height)
		}
	}()

	// The UnconfirmedAdded channel notifies when a transaction related to an
	// address is in unconfirmed state and waiting to be included in a block.
	// The message contains the transaction.
	unconfirmedTxSender, err := ws.Subscribe.UnconfirmedAdded(accSender.Address)
	if err != nil {
		panic(err)
	}
	go func() {
		defer wg.Done()
		for {
			data := <-unconfirmedTxSender.Ch
			fmt.Printf("[01]Sender \t\t UnconfirmedTxn Hash: %v \n\n", data.GetAbstractTransaction().Hash)
			unconfirmedTxSender.Unsubscribe()
			break
		}
	}()

	unconfirmedTxRecipient, err := ws.Subscribe.UnconfirmedAdded(accRecipient)
	if err != nil {
		panic(err)
	}
	go func() {
		defer wg.Done()
		for {
			data := <-unconfirmedTxRecipient.Ch
			fmt.Printf("[02]Recipient \t UnconfirmedTxn Hash: %v \n", data.GetAbstractTransaction().Hash)
			unconfirmedTxRecipient.Unsubscribe()
			break
		}
	}()

	//// The confirmedAdded channel notifies when a transaction related to an
	//// address is included in a block. The message contains the transaction.
	confirmedTxSender, err := ws.Subscribe.ConfirmedAdded(accSender.Address)
	if err != nil {
		panic(err)
	}
	go func() {
		defer wg.Done()
		for {
			data := <-confirmedTxSender.Ch
			fmt.Printf("[01]Sender \t\t ConfirmedTxn Hash: %v \n", data.GetAbstractTransaction().Hash)
			fmt.Printf("[01]Sender \t\t Height Txn: %v \n", data.GetAbstractTransaction().Height)

			confirmedTxSender.Unsubscribe()
			fmt.Printf("[01]Successful \t transfer! \n\n")
			break
		}
	}()

	confirmedTxRecipient, err := ws.Subscribe.ConfirmedAdded(accRecipient)
	if err != nil {
		panic(err)
	}
	go func() {
		defer wg.Done()
		for {
			data := <-confirmedTxRecipient.Ch
			fmt.Printf("[02]Recipient \t ConfirmedTxn Hash: %v \n", data.GetAbstractTransaction().Hash)
			fmt.Printf("[02]Recipient \t Height Txn: %v \n", data.GetAbstractTransaction().Height)
			confirmedTxRecipient.Unsubscribe()
			fmt.Printf("[02]Successful \t transfer! \n\n")
			break
		}
	}()

	//The status channel notifies when a transaction related to an address rises an error.
	//The message contains the error message and the transaction hash.
	statusSender, err := ws.Subscribe.Status(accSender.Address)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			data := <-statusSender.Ch
			statusSender.Unsubscribe()
			fmt.Printf("Hash: %v \n", data.Hash)
			panic(fmt.Sprint("Sender Status: ", data.Status))
		}
	}()

	statusRecipient, err := ws.Subscribe.Status(accRecipient)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			data := <-statusRecipient.Ch
			statusRecipient.Unsubscribe()
			fmt.Printf("Hash: %v \n", data.Hash)
			panic(fmt.Sprint("Recipient Status: ", data.Status))
		}
	}()

	errChannelSender, err := ws.Subscribe.Error(accSender.Address)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			data := <-errChannelSender.Ch
			errChannelSender.Unsubscribe()
			panic(fmt.Sprint("Sender ChannelError: ", data.Error))
		}
	}()

	errChannelRecipient, err := ws.Subscribe.Error(accSender.Address)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			data := <-errChannelRecipient.Ch
			errChannelRecipient.Unsubscribe()
			panic(fmt.Sprint("Recipient ChannelError: ", data.Error))
		}
	}()

	time.Sleep(time.Second * 5)

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	ttx, err := sdk.NewTransferTransaction(
		sdk.NewDeadline(time.Hour*1),
		sdk.NewAddress(accRecipient.Address, networkType),
		[]*sdk.Mosaic{sdk.Xem(10)},
		sdk.NewPlainMessage(""),
		networkType,
	)

	// Sign transaction
	stx, err := accSender.Sign(ttx)
	if err != nil {
		panic(fmt.Errorf("TransaferTransaction signing returned error: %s", err))
	}

	// Announce transaction
	restTx, err := client.Transaction.Announce(context.Background(), stx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", restTx)
	fmt.Printf("Hash: \t\t%v\n", stx.Hash)
	fmt.Printf("Signer: \t%X\n\n", accSender.KeyPair.PublicKey.Raw)

	wg.Wait()
}
