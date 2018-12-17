package sdk

import (
	"encoding/base32"
	"encoding/hex"
	"github.com/proximax-storage/nem2-sdk-go/crypto"
	"sync"
)

var addressNet = map[uint8]NetworkType{
	'N': MainNet,
	'T': TestNet,
	'M': Mijin,
	'S': MijinTest,
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
	ms := make([]*Mosaic, len(dto.Account.Mosaics))
	for i, m := range dto.Account.Mosaics {
		msc, err := m.toStruct()
		if err != nil {
			return nil, err
		}

		ms[i] = msc
	}

	add, err := NewAddressFromEncoded(dto.Account.Address)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{
		add,
		dto.Account.AddressHeight.toBigInt(),
		dto.Account.PublicKey,
		dto.Account.PublicKeyHeight.toBigInt(),
		dto.Account.Importance.toBigInt(),
		dto.Account.ImportanceHeight.toBigInt(),
		ms,
	}, nil
}

type accountInfoDTOs []*accountInfoDTO

func (a *accountInfoDTOs) toStruct() ([]*AccountInfo, error) {
	accsDTO := *a
	accountInfos := make([]*AccountInfo, 0, len(accsDTO))

	for _, dto := range accsDTO {
		accountInfo, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		accountInfos = append(accountInfos, accountInfo)
	}

	return accountInfos, nil
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
	cs := make([]*PublicAccount, len(dto.Multisig.Cosignatories))
	ms := make([]*PublicAccount, len(dto.Multisig.MultisigAccounts))

	acc, err := NewAccountFromPublicKey(dto.Multisig.Account, networkType)
	if err != nil {
		return nil, err
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i, c := range dto.Multisig.Cosignatories {
			cs[i], err = NewAccountFromPublicKey(c, networkType)
		}
	}()

	go func() {
		defer wg.Done()
		for i, m := range dto.Multisig.MultisigAccounts {
			ms[i], err = NewAccountFromPublicKey(m, networkType)
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

type addresses []*Address

func (ref *addresses) MarshalJSON() (buf []byte, err error) {
	buf = []byte(`{"addresses":[`)
	for i, address := range *ref {
		b := []byte(`"` + address.Address + `"`)
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, b...)
	}

	buf = append(buf, ']', '}')
	return
}

func (ref *addresses) UnmarshalJSON(buf []byte) error {
	return nil
}

// generateEncodedAddress convert publicKey to address
func generateEncodedAddress(pKey string, version NetworkType) (string, error) {
	// step 1: sha3 hash of the public key
	pKeyD, err := hex.DecodeString(pKey)
	if err != nil {
		return "", err
	}
	sha3PublicKeyHash, err := crypto.HashesSha3_256(pKeyD)
	if err != nil {
		return "", err
	}
	// step 2: ripemd160 hash of (1)
	ripemd160StepOneHash, err := crypto.HashesRipemd160(sha3PublicKeyHash)
	if err != nil {
		return "", err
	}

	// step 3: add version byte in front of (2)
	versionPrefixedRipemd160Hash := append([]byte{uint8(version)}, ripemd160StepOneHash...)

	// step 4: get the checksum of (3)
	stepThreeChecksum, err := GenerateChecksum(versionPrefixedRipemd160Hash)
	if err != nil {
		return "", err
	}

	// step 5: concatenate (3) and (4)
	concatStepThreeAndStepSix := append(versionPrefixedRipemd160Hash, stepThreeChecksum...)

	// step 6: base32 encode (5)
	return base32.StdEncoding.EncodeToString(concatStepThreeAndStepSix), nil
}
