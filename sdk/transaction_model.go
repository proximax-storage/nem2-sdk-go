package sdk

import (
	"errors"
	"fmt"
	"strconv"
	"math/big"
	"bytes"
)

// Models

// Transaction
type Transaction interface {
	SignWith(account PublicAccount) (Signed, error)
	String() string
}

// AbstractTransaction
type AbstractTransaction struct {
	NetworkType
	*TransactionInfo
	Type             TransactionType
	Version          uint64
	Fee              *big.Int
	Deadline         *big.Int
	Signature        string
	Signer           string
}

func (tx *AbstractTransaction) String() string {
	return fmt.Sprintf(
		`
			"NetworkType": %s,
			"TransactionInfo": %s,
			"Type": %s,
			"Version": %d,
			"Fee": %d,
			"Deadline": %d,
			"Signature": %s,
			"Signer": %s
		`,
		tx.NetworkType,
		tx.TransactionInfo.String(),
		tx.Type,
		tx.Version,
		tx.Fee,
		tx.Deadline,
		tx.Signature,
		tx.Signer,
	)
}

type abstractTransactionDTO struct {
	NetworkType      `json:"networkType"`
	Type             uint32 `json:"type"`
	Version          uint64          `json:"version"`
	Fee              uint64DTO        `json:"fee"`
	Deadline         uint64DTO        `json:"deadline"`
	Signature        string          `json:"signature"`
	Signer           string          `json:"signer"`
}

func (dto *abstractTransactionDTO) toStruct(tInfo *TransactionInfo) (*AbstractTransaction, error) {
	t, err := TransactionTypeFromRaw(dto.Type)
	if err != nil {
		return nil, err
	}

	nt, err := ExtractNetworkType(dto.Version)
	if err != nil {
		return nil, err
	}

	tv, err := ExtractTransactionVersion(dto.Version)
	if err != nil {
		return nil, err
	}

	return &AbstractTransaction{
		nt,
		tInfo,
		t,
		tv,
		dto.Fee.toStruct(),
		dto.Deadline.toStruct(),
		dto.Signature,
		dto.Signer,
	}, nil
}

// Transaction Info
type TransactionInfo struct {
	Height              *big.Int
	Index               uint32
	Id                  string
	Hash                string
	MerkleComponentHash string
	AggregateHash       string
	AggregateId         string
}

func (ti *TransactionInfo) String() string {
	return fmt.Sprintf(
		`
			"Height": %d,
			"Index": %d,
			"Id": %s,
			"Hash": %s,
			"MerkleComponentHash:" %s,
			"AggregateHash": %s,
			"AggregateId": %s
		`,
		ti.Height,
		ti.Index,
		ti.Id,
		ti.Hash,
		ti.MerkleComponentHash,
		ti.AggregateHash,
		ti.AggregateId,
	)
}

type transactionInfoDTO struct {
	Height              uint64DTO `json:"height"`
	Index               uint32    `json:"index"`
	Id                  string    `json:"id"`
	Hash                string    `json:"hash"`
	MerkleComponentHash string    `json:"merkleComponentHash"`
	AggregateHash       string    `json:"aggregate_hash,omitempty"`
	AggregateId         string    `json:"aggregate_id,omitempty"`
}

func (dto *transactionInfoDTO) toStruct() *TransactionInfo {
	return &TransactionInfo{
		dto.Height.toStruct(),
		dto.Index,
		dto.Id,
		dto.Hash,
		dto.MerkleComponentHash,
		dto.AggregateHash,
		dto.AggregateId,
	}
}

// AggregateTransaction
type AggregateTransaction struct {
	AbstractTransaction
	InnerTransactions []Transaction
	Cosignatures      []*AggregateTransactionCosignature
}

func (tx *AggregateTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *AggregateTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"InnerTransactions": %s,
			"Cosignatures": %s
		`,
		tx.AbstractTransaction.String(),
		tx.InnerTransactions,
		tx.Cosignatures,
	)
}

type aggregateTransactionDTO struct {
	Tx struct {
		ADto              abstractTransactionDTO
		Cosignatures      []*AggregateTransactionCosignature `json:"cosignatures"`
		InnerTransactions []map[string]interface{}           `json:"transactions"`
	} `json:"transaction"`
	TDto transactionInfoDTO  `json:"meta"`
}

func (dto *aggregateTransactionDTO) toStruct() (*AggregateTransaction, error) {
	txsj, err := json.Marshal(dto.Tx.InnerTransactions)
	if err != nil {
		return nil, err
	}

	txs, err := MapTransactions(bytes.NewBuffer(txsj))
	if err != nil {
		return nil, err
	}

	atx, err := dto.Tx.ADto.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	return &AggregateTransaction{
		*atx,
		txs,
		dto.Tx.Cosignatures,
	}, nil
}

// MosaicDefinitionTransaction
type MosaicDefinitionTransaction struct {
	AbstractTransaction
	*MosaicProperties
	*NamespaceId
	MosaicId   *big.Int
	MosaicName string
}

func (tx *MosaicDefinitionTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *MosaicDefinitionTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"MosaicProperties": %s,
			"MosaicId": [ %d ],
			"MosaicName": %s
		`,
		tx.AbstractTransaction.String(),
		tx.MosaicProperties.String(),
		tx.MosaicId,
		tx.MosaicName,
	)
}

type mosaicDefinitionTransactionDTO struct {
	Tx struct {
		ADto       abstractTransactionDTO
		Properties string    `json:"properties"`
		ParentId   uint64DTO `json:"parentId"`
		MosaicId   uint64DTO  `json:"mosaicId"`
		MosaicName string    `json:"name"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *mosaicDefinitionTransactionDTO) toStruct() (*MosaicDefinitionTransaction, error) {
	return &MosaicDefinitionTransaction{}, nil
}

// MosaicSupplyChangeTransaction
type MosaicSupplyChangeTransaction struct {
	AbstractTransaction
	MosaicSupplyType
	MosaicId         *big.Int
	Delta            *big.Int
}

func (tx *MosaicSupplyChangeTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *MosaicSupplyChangeTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"MosaicSupplyType": %s,
			"MosaicId": [ %d ],
			"Delta": %d
		`,
		tx.AbstractTransaction.String(),
		tx.MosaicSupplyType.String(),
		tx.MosaicId,
		tx.Delta,
	)
}

type mosaicSupplyChangeTransactionDTO struct {
	Tx struct {
		ADto     abstractTransactionDTO
		MosaicSupplyType   `json:"direction"`
		MosaicId uint64DTO `json:"mosaicId"`
		Delta    uint64DTO `json:"delta"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *mosaicSupplyChangeTransactionDTO) toStruct() (*MosaicSupplyChangeTransaction, error) {
	return &MosaicSupplyChangeTransaction{}, nil
}

// TransferTransaction
type TransferTransaction struct {
	AbstractTransaction
	Message
	Mosaics []*Mosaic
	Address string
}

func (tx *TransferTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *TransferTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"Mosaics": %d,
			"Address": %s,
			"Message": %s,
		`,
		tx.AbstractTransaction.String(),
		tx.Mosaics,
		tx.Address,
		tx.Message.String(),
	)
}

type transferTransactionDTO struct {
	Tx struct {
		ADto     abstractTransactionDTO
		Message `json:"message"`
		Mosaics []*Mosaic `json:"mosaics"`
		Address string    `json:"recipient"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *transferTransactionDTO) toStruct() (*TransferTransaction, error) {
 return &TransferTransaction{}, nil
}

// ModifyMultisigAccountTransaction
type ModifyMultisigAccountTransaction struct {
	AbstractTransaction
	MinApprovalDelta int32
	MinRemovalDelta  int32
	Modifications    []*MultisigCosignatoryModification
}

func (tx *ModifyMultisigAccountTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *ModifyMultisigAccountTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"MinApprovalDelta": %d,
			"MinRemovalDelta": %d,
			"Modifications": %s 
		`,
		tx.AbstractTransaction.String(),
		tx.MinApprovalDelta,
		tx.MinRemovalDelta,
		tx.Modifications,
	)
}

type modifyMultisigAccountTransactionDTO struct {
	Tx struct {
		ADto     abstractTransactionDTO
		MinApprovalDelta int32                              `json:"minApprovalDelta"`
		MinRemovalDelta  int32                              `json:"minRemovalDelta"`
		Modifications    []*MultisigCosignatoryModification `json:"modifications"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *modifyMultisigAccountTransactionDTO) toStruct() (*ModifyMultisigAccountTransaction, error) {
	return &ModifyMultisigAccountTransaction{}, nil
}

// RegisterNamespaceTransaction
type RegisterNamespaceTransaction struct {
	AbstractTransaction
	NamespaceId
	NamspaceName string   `json:"name"`
	Duration     *big.Int `json:"duration"`
	ParentId     NamespaceId
	NamespaceType
}

func (tx *RegisterNamespaceTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *RegisterNamespaceTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"NamespaceName": %s,
			"Duration": %d
		`,
		tx.AbstractTransaction.String(),
		tx.NamspaceName,
		tx.Duration,
	)
}

type registerNamespaceTransactionDTO struct {
	Tx struct {
		ADto     abstractTransactionDTO
		NamespaceId
		NamespaceType
		NamspaceName string   `json:"name"`
		Duration     uint64DTO `json:"duration"`
		ParentId     NamespaceId

	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *registerNamespaceTransactionDTO) toStruct() (*RegisterNamespaceTransaction, error) {
	return &RegisterNamespaceTransaction{}, nil
}

// LockFundsTransaction
type LockFundsTransaction struct {
	AbstractTransaction
	Mosaic   `json:"mosaic"`
	Duration uint64 `json:"duration"`
	Hash     string `json:"hash"`
}

func (tx *LockFundsTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *LockFundsTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"Mosaic": %s,
			"Duration": %d,
			"Hash": %s
		`,
		tx.AbstractTransaction.String(),
		tx.Mosaic.String(),
		tx.Duration,
		tx.Hash,
	)
}

type lockFundsTransactionDTO struct {
	Tx struct {
		ADto     abstractTransactionDTO
		Mosaic   `json:"mosaic"`
		Duration uint64 `json:"duration"`
		Hash     string `json:"hash"`

	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *lockFundsTransactionDTO) toStruct() (*LockFundsTransaction, error) {
	return &LockFundsTransaction{}, nil
}

// SecretLockTransaction
type SecretLockTransaction struct {
	AbstractTransaction
	*Mosaic   `json:"mosaic"`
	HashType  `json:"hashAlgorithm"`
	Duration  *big.Int `json:"duration"`
	Secret    string   `json:"secret"`
	Recipient string   `json:"recipient"`
}

func (tx *SecretLockTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *SecretLockTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"Mosaic": %s,
			"Duration": %d,
			"HashType": %s,
			"Secret": %s,
			"Recipient": %s
		`,
		tx.AbstractTransaction.String(),
		tx.Mosaic.String(),
		tx.Duration,
		tx.HashType.String(),
		tx.Secret,
		tx.Recipient,
	)
}

type secretLockTransactionDTO struct {
	Tx struct {
		ADto     abstractTransactionDTO
		*Mosaic   `json:"mosaic"`
		HashType  `json:"hashAlgorithm"`
		Duration  uint64DTO `json:"duration"`
		Secret    string   `json:"secret"`
		Recipient string   `json:"recipient"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *secretLockTransactionDTO) toStruct() (*SecretLockTransaction, error) {
	return &SecretLockTransaction{}, nil
}

// SecretProofTransaction
type SecretProofTransaction struct {
	AbstractTransaction
	HashType `json:"hashAlgorithm"`
	Secret   string `json:"secret"`
	Proof    string `json:"proof"`
}

func (tx *SecretProofTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *SecretProofTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"HashType": %s,
			"Secret": %s,
			"Proof": %s
		`,
		tx.AbstractTransaction.String(),
		tx.HashType.String(),
		tx.Secret,
		tx.Proof,
	)
}

type secretProofTransactionDTO struct {
	AbstractTransaction
	HashType `json:"hashAlgorithm"`
	Secret   string `json:"secret"`
	Proof    string `json:"proof"`
}

func (dto *secretProofTransactionDTO) toStruct() (*SecretLockTransaction, error) {
	return &SecretLockTransaction{}, nil
}

type CosignatureTransaction struct {
	TransactionToCosign *AggregateTransaction
}

func (tx *CosignatureTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *CosignatureTransaction) String() string {
	return fmt.Sprintf(`"TransactionToCosign": %s`, tx.TransactionToCosign.String())
}

// Signed
type Signed interface {
}

// SignedTransaction
type SignedTransaction struct {
	TransactionType `json:"transactionType"`
	Payload         string `json:"payload"`
	Hash            string `json:"hash"`
}

// CosignatureSignedTransaction
type CosignatureSignedTransaction struct {
	ParentHash string `json:"parentHash"`
	Signature  string `json:"signature"`
	Signer     string `json:"signer"`
}

// AggregateTransactionCosignature
type AggregateTransactionCosignature struct {
	Signature string `json:"signature"`
	Signer    string `json:"signer"`
}

func (agt *AggregateTransactionCosignature) String() string {
	return fmt.Sprintf(
		`
			"Signature": %s,
			"Signer": %s
		`,
		agt.Signature,
		agt.Signer,
	)
}

// MultisigCosignatoryModification
type MultisigCosignatoryModification struct {
	Type          MultisigCosignatoryModificationType
	PublicAccount
}

func (m *MultisigCosignatoryModification) String() string {
	return fmt.Sprintf(
		`
			"Type": %s,
			"PublicAccount": %s
		`,
		m.Type.String(),
		m.PublicAccount,
	)
}

type multisigCosignatoryModificationDTO struct {
	Type          MultisigCosignatoryModificationType `json:"type"`
	PublicAccount string                              `json:"cosignatoryPublicKey"`
}

func (dto *multisigCosignatoryModificationDTO) toStruct() (*MultisigCosignatoryModification, error) {
	return &MultisigCosignatoryModification{}, nil
}

// TransactionStatus
type TransactionStatus struct {
	Group    string   `json:"group"`
	Status   string   `json:"status"`
	Hash     string   `json:"hash"`
	Deadline *big.Int `json:"deadline"`
	Height   *big.Int `json:"height"`
}

func (ts *TransactionStatus) String() string {
	return fmt.Sprintf(
		`
			"Group:" %s,
			"Status:" %s,
			"Hash": %s,
			"Deadline": %d,
			"Height": %d
		`,
		ts.Group,
		ts.Status,
		ts.Hash,
		ts.Deadline,
		ts.Height,
	)
}

type transactionStatusDTO struct {
	Group    string   `json:"group"`
	Status   string   `json:"status"`
	Hash     string   `json:"hash"`
	Deadline uint64DTO `json:"deadline"`
	Height   uint64DTO `json:"height"`
}

func (dto *transactionStatusDTO) toStruct() (*TransactionStatus, error) {
	return &TransactionStatus{}, nil
}

// TransactionIds
type TransactionIdsDTO struct {
	Ids []string `json:"transactionIds"`
}

// TransactionHashes
type TransactionHashesDTO struct {
	Hashes []string `json:"hashes"`
}

// Message
type Message struct {
	Type    int8   `json:"type"`
	Payload string `json:"payload"`
}

func (m *Message) String() string {
	return fmt.Sprintf(
		`
			"Type": %d,
			"Payload": %s
		`,
		m.Type,
		m.Payload,
	)
}

type transactionTypeStruct struct {
	transactionType TransactionType
	raw             uint32
	hex             uint32
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

func TransactionTypeFromRaw(value uint32) (TransactionType, error) {
	for _, t := range transactionTypes {
		if t.raw == value {
			return t.transactionType, nil
		}
	}
	return 0, transactionTypeError
}

func (t TransactionType) Hex() uint32 {
	return transactionTypes[t].hex
}

func (t TransactionType) Raw() uint32 {
	return transactionTypes[t].raw
}

func (t TransactionType) String() string {
	return fmt.Sprintf("%d", t.Raw())
}

// TransactionType error
var transactionTypeError = errors.New("wrong raw TransactionType int")

type MultisigCosignatoryModificationType uint8

func (t MultisigCosignatoryModificationType) String() string {
	return fmt.Sprintf("%d", t)
}

const (
	ADD    MultisigCosignatoryModificationType = 0
	REMOVE MultisigCosignatoryModificationType = 1
)

type HashType uint8

func (ht HashType) String() string {
	return fmt.Sprintf("%d", ht)
}

const SHA3_512 HashType = 0

func ExtractTransactionVersion(version uint64) (uint64, error) {
	res, err := strconv.ParseUint(strconv.FormatUint(version, 16)[2:4], 16, 32)
	if err != nil {
		return 0, err
	}
	return res, nil
}
