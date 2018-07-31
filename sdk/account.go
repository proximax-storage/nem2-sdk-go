package sdk

import (
	"encoding/base32"
	"errors"
	"github.com/json-iterator/go"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"strconv"
	"sync"
	"unsafe"
)

type Address struct {
	Address     string `json:"address"`
	networkType NetworkType
}

func createFromPublicKey(publicKey string, networkType NetworkType) (*Address, error) {
	add, err := generateEncoded(byte(uint8(networkType)), publicKey)
	if err != nil {
		return nil, err
	}
	return &Address{add,
		networkType,
	}, nil

}

const NUM_CHECKSUM_BYTES = 4

func generateEncoded(version byte, publicKey string) (string, error) {

	// step 1: sha3 hash of the public key

	sha3PublicKeyHash := sha3.New256()
	_, err := sha3PublicKeyHash.Write([]byte(publicKey))
	if err == nil {

		// step 2: ripemd160 hash of (1)
		ripemd160StepOneHash := ripemd160.New()
		_, err = ripemd160StepOneHash.Write(sha3PublicKeyHash.Sum(nil))
		if err == nil {

			// step 3: add version byte in front of (2)
			versionPrefixedRipemd160Hash := ripemd160StepOneHash.Sum([]byte{version})

			// step 4: get the checksum of (3)
			sha3StepThreeHash := sha3.New256()
			_, err = sha3StepThreeHash.Write(versionPrefixedRipemd160Hash)
			if err == nil {

				stepThreeChecksum := sha3StepThreeHash.Sum(nil)

				// step 5: concatenate (3) and (4)
				concatStepThreeAndStepSix := append(versionPrefixedRipemd160Hash, stepThreeChecksum[:NUM_CHECKSUM_BYTES]...)

				// step 6: base32 encode (5)
				return base32.HexEncoding.EncodeToString(concatStepThreeAndStepSix), nil
			}
		}
	}
	return "", err
}

//TODO: this marshal return one string - change in future
func (ref *Address) MarshalJSON() (buf []byte, err error) {
	return append([]byte(`"`+ref.Address), '"'), nil
}

type Addresses struct {
	Addresses []*Address
	lock      sync.RWMutex
}

func (ref *Addresses) AddAddress(address *Address) {
	ref.lock.Lock()
	defer ref.lock.Unlock()

	ref.Addresses = append(ref.Addresses, address)
}
func (ref *Addresses) GetAddress(i int) (*Address, error) {

	if (i >= 0) && (i < len(ref.Addresses)) {
		ref.lock.RLock()
		defer ref.lock.RUnlock()
		return ref.Addresses[i], nil
	}

	return nil, errors.New("index out of range - " + strconv.Itoa(i))

}
func (ref *Addresses) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"addresses":[`)
	for i, address := range ref.Addresses {
		b, _ := address.MarshalJSON()
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, b...)
	}

	buf = append(buf, ']', '}')
	return
}
func AddressesIsEmpty(ptr unsafe.Pointer) bool {
	return len((*Addresses)(ptr).Addresses) == 0
}
func AddressesEncode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	buf, err := (*Addresses)(ptr).MarshalJSON()
	if err == nil {
		stream.Write(buf)
	}

}

type PublicAccount struct {
	Address   *Address
	PublicKey string
}

func NewPublicAccount(publicKey string, networkType NetworkType) (*PublicAccount, error) {
	add, err := createFromPublicKey(publicKey, networkType)
	if err != nil {
		return nil, err
	}
	ref := &PublicAccount{
		add,
		publicKey,
	}
	return ref, nil
}

func init() {
	jsoniter.RegisterTypeEncoderFunc("sdk.Addresses", AddressesEncode, AddressesIsEmpty)
}
