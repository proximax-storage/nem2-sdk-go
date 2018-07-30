package sdk

import (
"errors"
)

// Models

// Transaction
type Transaction interface {

}

// AbstractTransaction
type AbstractTransaction struct {
	NetworkType               `json:"networkType"`
	TransactionInfo           `json:"transactionInfo"`
	Type      TransactionType `json:"type"`
	Version   uint64          `json:"version"`
	Fee       []uint64        `json:"fee"`
	Deadline  []uint64        `json:"deadline"`
	Signature string          `json:"signature"`
	Signer    string          `json:"signer"`

}

// Transaction Info
type TransactionInfo struct {
	Height              []uint64 `json:"height"`
	Index               uint32   `json:"index"`
	Id                  string   `json:"id"`
	Hash                string   `json:"hash"`
	MerkleComponentHash string   `json:"merkleComponentHash"`
	AggregateHash       string   `json:"aggregate_hash,omitempty"`
	AggregateId         string   `json:"aggregate_id,omitempty"`
}

// AggregateTransaction
type AggregateTransaction struct {
	AbstractTransaction
	InnerTransactions []Transaction                     `json:"transactions"`
	Cosignatures      []AggregateTransactionCosignature `json:"cosignatures"`
}

// MosaicDefinitionTransaction
type MosaicDefinitionTransaction struct {
	AbstractTransaction
	//NamespaceId
	MosaicProperties
	MosaicId   []uint64 `json:"mosaicId"`
	MosaicName string   `json:"name"`
}

// MosaicSupplyChangeTransaction
type MosaicSupplyChangeTransaction struct {
	AbstractTransaction
	MosaicSupplyType  `json:"direction"`
	MosaicId []uint64 `json:"mosaicId"`
	Delta    uint64   `json:"delta"`
}

// TransferTransaction
type TransferTransaction struct {
	AbstractTransaction
	Mosaics         `json:"mosaics"`
	Address string  `json:"recipient"`
	Message Message `json:"message"`
}

// ModifyMultisigAccountTransaction
type ModifyMultisigAccountTransaction struct {
	AbstractTransaction
	MinApprovalDelta int32                             `json:"minApprovalDelta"`
	MinRemovalDelta  int32                             `json:"minRemovalDelta"`
	Modifications    []MultisigCosignatoryModification `json:"modifications"`
}

// RegisterNamespaceTransaction
type RegisterNamespaceTransaction struct {
	AbstractTransaction
	//NamespaceId
	NamspaceName string   `json:"name"`
	Duration     []uint64 `json:"duration"`
	//ParentId NamespaceId
	//NamespaceType
}

// LockFundsTransaction
type LockFundsTransaction struct {
	AbstractTransaction
	Mosaic          `json:"Mosaic"`
	Duration uint64 `json:"duration"`
	Hash     string `json:"hash"`
}

// SecretLockTransaction
type SecretLockTransaction struct {
	AbstractTransaction
	Mosaic             `json:"mosaic"`
	Duration  []uint64 `json:"duration"`
	HashType           `json:"hashAlgorithm"`
	Secret    string   `json:"secret"`
	Recipient string   `json:"recipient"`
}

// SecretProofTransaction
type SecretProofTransaction struct {
	AbstractTransaction
	HashType      `json:"hashAlgorithm"`
	Secret string `json:"secret"`
	Proof  string `json:"proof"`
}

// SignedTransaction
type SignedTransaction struct {
	TransactionType `json:"transactionType"`
	Payload string  `json:"payload"`
	Hash    string  `json:"hash"`
}

// CosignatureSignedTransaction
type CosignatureSignedTransaction struct {
	ParentHash string `json:"parentHash"`
	Signature  string `json:"signature"`
	Signer     string `json:"signer"`
}

// AggregateTransactionCosignature
type AggregateTransactionCosignature struct {
	Signature string        `json:"signature"`
	Signer    PublicAccount `json:"signer"`
}

// MultisigCosignatoryModification
type MultisigCosignatoryModification struct {
	Type          MultisigCosignatoryModificationType `json:"type"`
	PublicAccount string                              `json:"cosignatoryPublicKey"`
}

// TransactionStatus
type TransactionStatus struct {
	Group    string   `json:"group"`
	Status   string   `json:"status"`
	Hash     string   `json:"hash"`
	Deadline []uint64 `json:"deadline"`
	Height   []uint64 `json:"height"`
}

// TransactionIds
type TransactionIds struct {
	Ids []string `json:"transactionIds"`
}

// TransactionHashes
type TransactionHashes struct {
	Hashes []string `json:"hashes"`
}

// Message
type Message struct {
	Type    int8   `json:"type"`
	Payload string `json:"payload"`
}

type transactionTypeStruct struct {
	transactionType TransactionType
	raw uint32
	hex uint32
}

var transactionTypes = []transactionTypeStruct{
	{AGGREGATE_COMPLETE, 16705, 0x4141},
	{AGGREGATE_BONDED, 16961, 0x4241},
	{MOSAIC_DEFINITION, 16717, 0x414d},
	{MOSAIC_SUPPLY_CHANGE, 16973, 0x424d},
	{MODIFY_MULTISIG_ACCOUNT, 16725, 0x4155},
	{REGISTER_NAMESPACE, 16718, 0x414e},
	{TRANSFER, 16724, 0x4154},
	{LOCK, 16716, 0x414C},
	{SECRET_LOCK, 16972, 0x424C},
	{SECRET_PROOF, 17228, 0x434C},
}

type TransactionType uint

// TransactionType enums
const (
	AGGREGATE_COMPLETE TransactionType = iota
	AGGREGATE_BONDED
	MOSAIC_DEFINITION
	MOSAIC_SUPPLY_CHANGE
	MODIFY_MULTISIG_ACCOUNT
	REGISTER_NAMESPACE
	TRANSFER
	LOCK
	SECRET_LOCK
	SECRET_PROOF
)

func TransactionTypeFromRaw(value uint32) (TransactionType, error){
	for _, t := range transactionTypes {
		if t.raw == value{
			return t.transactionType, nil
		}
	}
	return 0, transactionTypeError
}

func (t TransactionType) Hex() uint32{
	return transactionTypes[t].hex
}

func (t TransactionType) Raw() uint32{
	return transactionTypes[t].raw
}

// TransactionType error
var transactionTypeError = errors.New("wrong raw TransactionType int")

type MultisigCosignatoryModificationType uint8

const (
	ADD MultisigCosignatoryModificationType = 0
	REMOVE MultisigCosignatoryModificationType = 1
)

type HashType uint8

const SHA3_512 HashType = 0