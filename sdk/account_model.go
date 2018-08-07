package sdk

import (
	"encoding/base32"
	"encoding/hex"
	"github.com/kataras/iris/core/errors"
	"github.com/proximax-storage/nem2-sdk-go/crypto"
	"math/big"
	"strconv"
	"strings"
	"sync"
)

type Account struct {
	*PublicAccount
	*crypto.KeyPair
}

type PublicAccount struct {
	Address   *Address
	PublicKey string
}

type AccountInfo struct {
	Address          *Address
	AddressHeight    *big.Int
	PublicKey        string
	PublicKeyHeight  *big.Int
	Importance       *big.Int
	ImportanceHeight *big.Int
	Mosaics          []*Mosaic
}

type accountInfoDTO struct {
	Account struct {
		Address          string       `json:"address"`
		AddressHeight    uint64DTO    `json:"addressHeight"`
		PublicKey        string       `json:"publicKey"`
		PublicKeyHeight  uint64DTO    `json:"publicKeyHeight"`
		Importance       uint64DTO    `json:"importance"`
		ImportanceHeight uint64DTO    `json:"importanceHeight"`
		Mosaics          []*mosaicDTO `json:"mosaics"`
	} `json:"account"`
}

func (dto *accountInfoDTO) toStruct() (*AccountInfo, error) {
	var err error
	ms := make(Mosaics, len(dto.Account.Mosaics))
	for i, m := range dto.Account.Mosaics {
		ms[i], err = m.toStruct()
	}
	if err != nil {
		return nil, err
	}

	add, err := NewAddressFromEncoded(dto.Account.Address)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{
		add,
		dto.Account.AddressHeight.toStruct(),
		dto.Account.PublicKey,
		dto.Account.PublicKeyHeight.toStruct(),
		dto.Account.Importance.toStruct(),
		dto.Account.ImportanceHeight.toStruct(),
		ms,
	}, nil
}

type Address struct {
	Type    NetworkType
	Address string
}

func (ad *Address) Pretty() string {
	res := ""
	for i := 0; i < 6; i++ {
		res += ad.Address[i*6:i*6+6] + "-"
	}
	res += ad.Address[len(ad.Address)-4:]
	return res
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
func (ref *Addresses) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"addresses":[`)
	for i, address := range ref.list {
		b := []byte(address.Address)
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

type MultisigAccountInfo struct {
	Account          PublicAccount
	MinApproval      int32
	MinRemoval       int32
	Cosignatories    []*PublicAccount
	MultisigAccounts []*PublicAccount
}

type multisigAccountInfoDTO struct {
	Multisig struct {
		Account          string   `json:"account"`
		MinApproval      int32    `json:"minApproval"`
		MinRemoval       int32    `json:"minRemoval"`
		Cosignatories    []string `json:"cosignatories"`
		MultisigAccounts []string `json:"multisigAccounts"`
	} `json:"multisig"`
}

func (dto *multisigAccountInfoDTO) toStruct(networkType NetworkType) (*MultisigAccountInfo, error) {
	var wg sync.WaitGroup
	var cs []*PublicAccount
	var ms []*PublicAccount

	acc, err := NewPublicAccount(dto.Multisig.Account, networkType)
	if err != nil {
		return nil, err
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i, c := range dto.Multisig.Cosignatories {
			cs[i], err = NewPublicAccount(c, networkType)
		}
	}()

	go func() {
		defer wg.Done()
		for i, m := range dto.Multisig.MultisigAccounts {
			ms[i], err = NewPublicAccount(m, networkType)
		}
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return &MultisigAccountInfo{
		*acc,
		dto.Multisig.MinApproval,
		dto.Multisig.MinRemoval,
		cs,
		ms,
	}, nil
}

type MultisigAccountGraphInfo struct {
	MultisigAccounts map[int32][]*MultisigAccountInfo
}

type multisigAccountGraphInfoDTOEntry struct {
	Level     int32                    `json:"level"`
	Multisigs []multisigAccountInfoDTO `json:"multisigEntries"`
}

type multisigAccountGraphInfoDTOS []multisigAccountGraphInfoDTOEntry

func (dto multisigAccountGraphInfoDTOS) toStruct(networkType NetworkType) (*MultisigAccountGraphInfo, error) {
	var ms map[int32][]*MultisigAccountInfo
	var wg1 sync.WaitGroup
	var err error

	for _, m := range dto {
		wg1.Add(1)
		go func(m multisigAccountGraphInfoDTOEntry) {
			defer wg1.Done()
			var wg2 sync.WaitGroup
			var mdto []*MultisigAccountInfo

			for i, c := range m.Multisigs {
				wg2.Add(1)
				go func(i int, c multisigAccountInfoDTO) {
					defer wg2.Done()
					mdto[i], err = c.toStruct(networkType)
				}(i, c)
			}
			wg2.Wait()

			ms[m.Level] = mdto
		}(m)
	}
	wg1.Wait()
	if err != nil {
		return nil, err
	}

	return &MultisigAccountGraphInfo{ms}, nil
}

var addressError = errors.New("wrong address")

func NewPublicAccount(pKey string, networkType NetworkType) (*PublicAccount, error) {
	ad, err := NewAddressFromPublicKey(pKey, networkType)
	if err != nil {
		return nil, err
	}
	return &PublicAccount{ad, pKey}, nil
}

func NewAddress(address string, networkType NetworkType) *Address {
	address = strings.Replace(address, "-", "", -1)
	address = strings.ToUpper(address)
	return &Address{networkType, address}
}

func NewAddressFromRaw(address string) (*Address, error) {
	switch address[:1] {
	case "N":
		return NewAddress(address, MAIN_NET), nil
	case "T":
		return NewAddress(address, TEST_NET), nil
	case "M":
		return NewAddress(address, MIJIN), nil
	case "S":
		return NewAddress(address, MIJIN_TEST), nil
	}
	return nil, addressError
}

func NewAddressFromPublicKey(pKey string, networkType NetworkType) (*Address, error) {
	ad, err := crypto.GenerateEncodedAddress(pKey, uint8(networkType))
	if err != nil {
		return nil, err
	}
	return NewAddress(ad, networkType), nil
}

func NewAddressFromEncoded(encoded string) (*Address, error) {
	pH, err := hex.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	parsed := base32.StdEncoding.EncodeToString(pH)
	ad, err := NewAddressFromRaw(parsed)
	if err != nil {
		return nil, err
	}

	return ad, nil
}
