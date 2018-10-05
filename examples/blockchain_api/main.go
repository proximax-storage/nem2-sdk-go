package main

import (
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"golang.org/x/net/context"
	"math/big"
)

const baseUrl = "http://catapult.internal.proximax.io:3000"

// Simple Blockchain API request
func main() {

	conf, err := sdk.LoadTestnetConfig(baseUrl)
	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Get the chain height
	chainHeight, resp, err := client.Blockchain.GetBlockchainHeight(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", chainHeight)

	// Get the chain score
	chainScore, resp, err := client.Blockchain.GetBlockchainScore(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", chainScore)

	// Get the Block by height
	blockHeight, resp, err := client.Blockchain.GetBlockByHeight(context.Background(), big.NewInt(9999))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%v\n\n", blockHeight)

	// Get the Block Transactions
	transactions, resp, err := client.Blockchain.GetBlockTransactions(context.Background(), big.NewInt(1))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", transactions)

	// Get the Blockchain Storage Info
	blockchainStorageInfo, resp, err := client.Blockchain.GetBlockchainStorage(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%v\n\n", blockchainStorageInfo)
}
