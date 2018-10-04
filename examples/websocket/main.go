package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"time"
)

// WebSockets make possible receiving notifications when a transaction or event occurs in the blockchain.
// The notification is received in real time without having to poll the API waiting for a reply.
func main() {

	host := "http://http://catapult.internal.proximax.io:3000"

	conf, err := sdk.LoadTestnetConfig(host)
	if err != nil {
		panic(err)
	}

	p, err := sdk.NewAccount("0F3CC33190A49ABB32E7172E348EA927F975F8829107AAA3D6349BB10797D4F6", sdk.MijinTest)

	ws, err := sdk.NewConnectWs(host)
	if err != nil {
		panic(err)
	}

	fmt.Println("websocket negotiated uid:", ws.Uid)

	// The UnconfirmedAdded channel notifies when a transaction related to an
	// address is in unconfirmed state and waiting to be included in a block.
	// The message contains the transaction.
	a, _ := ws.Subscribe.UnconfirmedAdded(p.Address.Address)
	go func() {
		for {
			data := <-a.ChIn
			ch := data.(sdk.Transaction)
			fmt.Printf("UnconfirmedAdded Tx Hash: %v \n", ch.GetAbstractTransaction().Hash)
			a.Unsubscribe()
		}
	}()

	// The confirmedAdded channel notifies when a transaction related to an
	// address is included in a block. The message contains the transaction.
	b, _ := ws.Subscribe.ConfirmedAdded(p.Address.Address)
	go func() {
		for {
			data := <-b.ChIn
			ch := data.(sdk.Transaction)
			fmt.Printf("ConfirmedAdded Tx Hash: %v \n", ch.GetAbstractTransaction().Hash)
			b.Unsubscribe()
			fmt.Println("Successful transfer!")
		}
	}()

	//The status channel notifies when a transaction related to an address rises an error.
	//The message contains the error message and the transaction hash.
	c, _ := ws.Subscribe.Status(p.Address.Address)

	go func() {
		for {
			data := <-c.ChIn
			ch := data.(sdk.StatusInfo)
			fmt.Printf("Status: %v \n", ch.Status)
			fmt.Printf("Hash: %v \n", ch.Hash)
		}
	}()

	time.Sleep(time.Second * 5)
	// Use the default http client
	client := sdk.NewClient(nil, conf)

	ttx, err := sdk.NewTransferTransaction(
		sdk.NewDeadline(time.Hour*1),
		sdk.NewAddress("SBILTA367K2LX2FEXG5TFWAS7GEFYAGY7QLFBYKC", sdk.MijinTest),
		sdk.Mosaics{sdk.Xem(10000000)},
		sdk.NewPlainMessage(""),
		sdk.MijinTest,
	)

	stx, err := p.Sign(ttx)
	if err != nil {
		panic(fmt.Errorf("TransaferTransaction signing returned error: %s", err))
	}

	// Get the chain height
	restTx, resp, err := client.Transaction.Announce(context.Background(), stx)
	if err != nil {
		panic(err)
	}
	txJson, _ := json.Marshal(restTx)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Println("Transaction Hash:", stx.Hash)
	fmt.Printf("%s\n\n", string(txJson))

	// The block channel notifies for every new block.
	// The message contains the block information.
	d, _ := ws.Subscribe.Block()

	for {
		data := <-d.ChIn
		ch := data.(*sdk.BlockInfo)
		fmt.Printf("Block received with height: %v \n", ch.Height)
	}
}
