package sdk

type Address struct {
	Address     string
	NetworkType NetworkType
}

func (ref *Address) MarshalJSON() (buf []byte, err error) {
	return append([]byte(ref.Address), byte(ref.NetworkType)), nil
}

type PublicAccount struct {
	Address   Address
	PublicKey string
}
