package sdk

import (
	"errors"
)

// Models
// Transaction Status
type TransactionStatus struct {
	Group    string   `json:"group"`
	Status   string   `json:"status"`
	Hash     string   `json:"hash"`
	Deadline []uint64 `json:"deadline"`
	Height   []uint64 `json:"height"`
}

// Transaction
type Transaction struct {
	Type            TransactionType `json:"transactionType"`
	NetworkType     NetworkType     `json:"networkType"`
	Version         *uint64         `json:"version"`
	Fee             *uint64         `json:"fee"`
	Deadline        []uint64        `json:"deadline"`
	Signature       string          `json:"signature"`
	Signer          PublicAccount   `json:"signer"`
	TransactionInfo TransactionInfo `json:"transactionInfo"`
}

// Transaction Info
type TransactionInfo struct {
	Height              *uint64 `json:"height"`
	Index               *uint32 `json:"index"`
	Id                  string  `json:"id"`
	Hash                string  `json:"hash"`
	MerkleComponentHash string  `json:"merkleComponentHash"`
	AggregateHash       string  `json:"aggregate_hash"`
	AggregateId         string  `json:"aggregate_id"`
}

type SignedTransaction struct {
	Payload string `json:"payload"`
	Hash string `json:"hash"`
	TransactionType TransactionType `json:"transactionType"`
}

type TransactionIds struct {
	Ids []string `json:"transactionIds, hashes"`
}

type TransactionType int

// TransactionType enums
const (
	AGGREGATE_COMPLETE TransactionType = 0x4141
	AGGREGATE_BONDED TransactionType = 0x4241
	MOSAIC_DEFINITION TransactionType = 0x414d
	MOSAIC_SUPPLY_CHANGE TransactionType = 0x424d
	MODIFY_MULTISIG_ACCOUNT TransactionType = 0x4155
	REGISTER_NAMESPACE TransactionType = 0x414e
	TRANSFER TransactionType = 0x4154
	LOCK TransactionType = 0x414C
	SECRET_LOCK TransactionType = 0x424C
	SECRET_PROOF TransactionType = 0x434C
)

// TransactionType error
var transactionTypeError = errors.New("wrong raw TransactionType int")

// Get TransactionType from raw value
func TransactionTypeFromRaw(value int) (TransactionType, error){
	switch value {
	case 16724:
		return TRANSFER, nil
	case 16718:
		return REGISTER_NAMESPACE, nil
	case 16717:
		return MOSAIC_DEFINITION, nil
	case 16973:
		return MOSAIC_SUPPLY_CHANGE, nil
	case 16725:
		return MODIFY_MULTISIG_ACCOUNT, nil
	case 16716:
		return LOCK, nil
	case 16972:
		return SECRET_LOCK, nil
	case 17228:
		return SECRET_PROOF, nil
	case 16705:
		return AGGREGATE_COMPLETE, nil
	case 16961:
		return AGGREGATE_BONDED, nil
	default:
		return 0, transactionTypeError
	}
}