package sdk

import (
	"encoding/base32"
	"encoding/hex"
	"github.com/proximax-storage/nem2-crypto-go"
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
	var (
		ms  = make([]*Mosaic, len(dto.Account.Mosaics))
		err error
	)

	for idx, m := range dto.Account.Mosaics {
		ms[idx], err = m.toStruct()
		if err != nil {
			return nil, err
		}
	}

	add, err := NewAddressFromEncoded(dto.Account.Address)
	if err != nil {
		return nil, err
	}

	return &AccountInfo{
		Address:          add,
		AddressHeight:    dto.Account.AddressHeight.toBigInt(),
		PublicKey:        dto.Account.PublicKey,
		PublicKeyHeight:  dto.Account.PublicKeyHeight.toBigInt(),
		Importance:       dto.Account.Importance.toBigInt(),
		ImportanceHeight: dto.Account.ImportanceHeight.toBigInt(),
		Mosaics:          ms,
	}, nil
}

type accountInfoDTOs []*accountInfoDTO

func (a accountInfoDTOs) toStruct() ([]*AccountInfo, error) {
	var (
		accountInfos = make([]*AccountInfo, len(a))
		err          error
	)

	for idx, dto := range a {
		accountInfos[idx], err = dto.toStruct()
		if err != nil {
			return nil, err
		}
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
	cs := make([]*PublicAccount, len(dto.Multisig.Cosignatories))
	ms := make([]*PublicAccount, len(dto.Multisig.MultisigAccounts))

	acc, err := NewAccountFromPublicKey(dto.Multisig.Account, networkType)
	if err != nil {
		return nil, err
	}

	for i, c := range dto.Multisig.Cosignatories {
		cs[i], err = NewAccountFromPublicKey(c, networkType)
		if err != nil {
			return nil, err
		}
	}

	for i, m := range dto.Multisig.MultisigAccounts {
		ms[i], err = NewAccountFromPublicKey(m, networkType)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return &MultisigAccountInfo{
		Account:          *acc,
		MinApproval:      dto.Multisig.MinApproval,
		MinRemoval:       dto.Multisig.MinRemoval,
		Cosignatories:    cs,
		MultisigAccounts: ms,
	}, nil
}

type multisigAccountGraphInfoDTOEntry struct {
	Level     int32                    `json:"level"`
	Multisigs []multisigAccountInfoDTO `json:"multisigEntries"`
}

type multisigAccountGraphInfoDTOS []multisigAccountGraphInfoDTOEntry

func (dto multisigAccountGraphInfoDTOS) toStruct(networkType NetworkType) (*MultisigAccountGraphInfo, error) {
	var (
		ms  = make(map[int32][]*MultisigAccountInfo)
		err error
	)

	for _, m := range dto {
		mAccInfos := make([]*MultisigAccountInfo, len(m.Multisigs))

		for idx, c := range m.Multisigs {
			mAccInfos[idx], err = c.toStruct(networkType)
			if err != nil {
				return nil, err
			}
		}

		ms[m.Level] = mAccInfos
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
