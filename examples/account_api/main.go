package main

import (
	"context"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

// Simple Blockchain API request
func main() {

	acc, err := sdk.NewAccount("0F3CC33190A49ABB32E7172E348EA927F975F8829107AAA3D6349BB10797D4F6", sdk.MijinTest)
	if err != nil {
		panic(err)
	}

	conf, err := sdk.LoadTestnetConfig("http://190.216.224.11:3000")
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	accI, resp, err := client.Account.GetAccountInfo(context.Background(), acc.Address)

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", accI)

	txs, resp, err := client.Account.Transactions(context.Background(), acc.PublicAccount, nil)

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", txs)
}
