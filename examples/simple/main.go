package main

import (
	"encoding/json"
	"fmt"
	"github.com/proximax/nem2-go-sdk/sdk"
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
	chainHeight, resp, err := client.Blockchain.GetChainHeight(context.Background())
	if err != nil {
		panic(err)
	}
	chainHeightJson, _ := json.Marshal(chainHeight)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(chainHeightJson))

	// Get the chain score
	chainScore, resp, err := client.Blockchain.GetChainScore(context.Background())
	if err != nil {
		panic(err)
	}
	chainScoreJson, _ := json.Marshal(chainScore)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(chainScoreJson))

	// Get the Block Height
	blockHeight, resp, err := client.Blockchain.GetBlockHeight(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	blockHeightJson, _ := json.Marshal(blockHeight)
	fmt.Printf("Response Status Code == %d\n", resp.StatusCode)
	fmt.Printf("%s\n\n", string(blockHeightJson))
}
