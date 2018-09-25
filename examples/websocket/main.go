package main

import (
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

// Simple Blockchain API request
func main() {

	conf, err := sdk.LoadTestnetConfig("http://190.216.224.11:3000")
	if err != nil {
		panic(err)
	}

	client, err := sdk.NewConnectWs(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("websocket negotiated uid:", client.Uid)

	a, _ := client.Subscribe.Status("SCFWMP2M2HP43KJYGOBDVQ3SKX3Q6HFH6HZZ6DNR")

	go func() {
		for {
			fmt.Printf("Status: %s \n", <-a.ChIn)
			//a.Unsubscribe()
		}
	}()

	// info once a block is created.
	c, _ := client.Subscribe.Block()

	for {
		fmt.Printf("Block: %s \n", <-c.ChIn)
	}
}
