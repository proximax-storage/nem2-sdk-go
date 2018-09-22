package main

import (
	"encoding/json"
	"fmt"
	"github.com/proximax-storage/nem2-sdk-go/sdk"
)

// Simple Blockchain API request
func main() {

	conf, err := sdk.LoadTestnetConfig("http://190.216.224.11:3000")
	if err != nil {
		panic(err)
	}

	// Use the default websocket client
	client := sdk.NewClientWs(nil, conf)

	err = client.WsConnect()
	if err != nil {
		panic(err)
	}
	fmt.Println("websocket negotiated uid:", client.Uid)

	//b, err := client.Subscribe.ConfirmedAdded("SCFWMP2M2HP43KJYGOBDVQ3SKX3Q6HFH6HZZ6DNR")
	//for {
	//	data := <-b
	//	fmt.Println(string(data))
	//}

	//c, err := client.Subscribe.UnConfirmedAdded("SCFWMP2M2HP43KJYGOBDVQ3SKX3Q6HFH6HZZ6DNR")
	//for {
	//	data := <-c
	//	fmt.Println(string(data))
	//}

	//d, err := client.Subscribe.UnconfirmedRemoved("SCFWMP2M2HP43KJYGOBDVQ3SKX3Q6HFH6HZZ6DNR")
	//for {
	//	data := <-d
	//	fmt.Println(string(data))
	//}

	//e, err := client.Subscribe.PartialAdded("SCFWMP2M2HP43KJYGOBDVQ3SKX3Q6HFH6HZZ6DNR")
	//for {
	//	data := <-e
	//	fmt.Println(string(data))
	//}

	//f, err := client.Subscribe.PartialRemoved("SCFWMP2M2HP43KJYGOBDVQ3SKX3Q6HFH6HZZ6DNR")
	//for {
	//	data := <-f
	//	fmt.Println(string(data))
	//}

	//// Info If the transaction fails
	//a, err := client.Subscribe.Status("SDXJKWQ5RFTBRQKPJTAB3OIFJFIMLLWMMAUNPXLZ")
	//for {
	//	data := <-a
	//	var datainfo sdk.StatusInfo
	//	err := json.Unmarshal(data, &datainfo)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	r, err := json.MarshalIndent(datainfo, "", "  ")
	//	fmt.Println(string(r))
	//}

	// info once a block is created.
	c, err := client.Subscribe.Block()
	for {
		data := <-c

		var datainfo sdk.BlockInfo
		err := json.Unmarshal(data, &datainfo)
		if err != nil {
			fmt.Println(err)
			return
		}
		r, err := json.MarshalIndent(datainfo, "", "  ")
		fmt.Println(string(r))
	}
}
