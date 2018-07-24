# nem2-sdk-go

nem2-sdk-go is a Golang client library for [Catapult API](http://docs.nem.ninja/#/) 

## Usage ##

```go
import "github.com/proximax/nem2-go-sdk/sdk"
```

Create a Catapult network configuration 

Using the *Testnet* network
```go
conf, err := sdk.LoadTestnetConfig("http://catapult.internal.proximax.io:3000")
```
Or using the *Mainnet* network
```go
conf, err := sdk.LoadMainnetConfig("http://catapult.internal.proximax.io:3000")
```

Construct a new Catapult client
```go
client := sdk.NewClient(nil, conf)
```

Using the client to call a method from a Service API

```go
// Get the chain height
chainHeight, resp, err := client.Blockchain.GetChainHeight(context.Background())
```

## Context ##

A [Context](https://golang.org/pkg/context/) type is the first argument in any service method for specifying
deadlines, cancelation signals, and other request-scoped values
```go
// Get the chain height
chainHeight, resp, err := client.Blockchain.GetChainHeight(context.Background())
```

## Examples ##

Examples are in the `examples` folder