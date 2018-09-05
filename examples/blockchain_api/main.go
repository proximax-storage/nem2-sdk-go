package main

import (
	"encoding/json"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
	"golang.org/x/net/context"
)

// Simple Blockchain API request
func main() {

	conf, err := sdk.LoadTestnetConfig("http://catapult.internal.proximax.io:3000")
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
	chainHeightJson, _ := json.Marshal(chainHeight)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(chainHeightJson))

	// Get the chain score
	chainScore, resp, err := client.Blockchain.GetBlockchainScore(context.Background())
	if err != nil {
		panic(err)
	}
	chainScoreJson, _ := json.Marshal(chainScore)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(chainScoreJson))

	// Get the Block Height
	blockHeight, resp, err := client.Blockchain.GetBlockByHeight(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	blockHeightJson, _ := json.Marshal(blockHeight)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(blockHeightJson))

	// Get the Block Transactions
	transactions, resp, err := client.Blockchain.GetBlockTransactions(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	transactionsJson, _ := json.Marshal(transactions)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(transactionsJson))

	// Get the Blockchain Storage Info
	blockchainStorageInfo, resp, err := client.Blockchain.GetBlockchainStorage(context.Background())
	if err != nil {
		panic(err)
	}

	blockchainStorageInfoJson, _ := json.Marshal(blockchainStorageInfo)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(blockchainStorageInfoJson))
}
