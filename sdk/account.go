package sdk

type Address struct {
	Address     string
	NetworkType NetworkType
}

//TODO: this marshal return one string - change in future
func (ref *Address) MarshalJSON() (buf []byte, err error) {
	return append([]byte(`"`+ref.Address), byte(ref.NetworkType), '"'), nil
}

type Addresses []Address

func (ref Addresses) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"addresses": [`)
	for i, address := range ref {
		b, _ := address.MarshalJSON()
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, b...)
	}

	buf = append(buf, ']', '}')
	return
}

type PublicAccount struct {
	Address   Address
	PublicKey string
}
