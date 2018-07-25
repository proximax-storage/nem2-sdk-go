package sdk

type Address struct {
	Address string
	NetworkType NetworkType
}

type PublicAccount struct {
	Address Address
	PublicKey string
}