package sdk

import (
	"errors"
	"strconv"
	"sync"
)

type Address struct {
	Address     string `json:"addres"`
	NetworkType NetworkType
}

func createFromPublicKey(publicKey string, networkType NetworkType) Address {
	return Address{generateEncoded(uint8(networkType), publicKey),
		networkType,
	}

}
func generateEncoded(version uint8, publicKey string) string {

	// step 1: sha3 hash of the public key
	return ""
}

//TODO: this marshal return one string - change in future
func (ref *Address) MarshalJSON() (buf []byte, err error) {
	return append([]byte(`"`+ref.Address), '"'), nil
}

type Addresses struct {
	list []*Address
	lock sync.RWMutex
}

func (ref *Addresses) AddAddress(address *Address) {
	ref.lock.Lock()
	defer ref.lock.Unlock()

	ref.list = append(ref.list, address)
}
func (ref *Addresses) GetAddress(i int) (*Address, error) {

	if (i >= 0) && (i < len(ref.list)) {
		ref.lock.RLock()
		defer ref.lock.RUnlock()
		return ref.list[i], nil
	}

	return nil, errors.New("index out of range - " + strconv.Itoa(i))

}
func (ref *Addresses) Marshal1JSON() (buf []byte, err error) {
	buf = []byte(`{"addresses":[`)
	for i, address := range ref.list {
		b, _ := address.MarshalJSON()
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, b...)
	}

	buf = append(buf, ']', '}')
	return
}
func (ref *Addresses) UnmarshalJSON(buf []byte) error {
	return nil
}

type PublicAccount struct {
	Address   Address
	PublicKey string
}

func NewPublicAccount(publicKey string, networkType NetworkType) *PublicAccount { /* public  */
	ref := &PublicAccount{
		createFromPublicKey(publicKey, networkType),
		publicKey,
	}
	return ref
}
