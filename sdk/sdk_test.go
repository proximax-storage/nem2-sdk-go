package sdk

const (
	address = "http://10.32.150.136:3000"
)

func setup() (*Client, string) {

	conf, err := LoadTestnetConfig(address)
	if err != nil {
		panic(err)
	}

	return NewClient(nil, conf), address
}