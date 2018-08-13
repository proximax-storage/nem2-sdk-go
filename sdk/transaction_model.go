package sdk

import (
	"bytes"
	jsonLib "encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"
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
	Type      TransactionType
	Version   uint64
	Fee       *big.Int
	Deadline  time.Time
	Signature string
	Signer    *PublicAccount
}

func (tx *AbstractTransaction) String() string {
	return fmt.Sprintf(
		`
			"NetworkType": %s,
			"TransactionInfo": %s,
			"Type": %s,
			"Version": %d,
			"Fee": %d,
			"Deadline": %s,
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

type AbstractTransactionDTO struct {
	NetworkType `json:"networkType"`
	Type        uint32     `json:"type"`
	Version     uint64     `json:"version"`
	Fee         *uint64DTO `json:"fee"`
	Deadline    *uint64DTO `json:"deadline"`
	Signature   string     `json:"signature"`
	Signer      string     `json:"signer"`
}

func (dto *AbstractTransactionDTO) toStruct(tInfo *TransactionInfo) (*AbstractTransaction, error) {
	t, err := TransactionTypeFromRaw(dto.Type)
	if err != nil {
		return nil, err
	}

	nt := ExtractNetworkType(dto.Version)

	tv, err := ExtractTransactionVersion(dto.Version)
	if err != nil {
		return nil, err
	}

	pa, err := NewPublicAccount(dto.Signer, nt)
	if err != nil {
		return nil, err
	}

	return &AbstractTransaction{
		nt,
		tInfo,
		t,
		tv,
		dto.Fee.toStruct(),
		time.Unix(dto.Deadline.toStruct().Int64(), int64(time.Millisecond)),
		dto.Signature,
		pa,
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
	Height              *uint64DTO `json:"height"`
	Index               uint32     `json:"index"`
	Id                  string     `json:"id"`
	Hash                string     `json:"hash"`
	MerkleComponentHash string     `json:"merkleComponentHash"`
	AggregateHash       string     `json:"aggregateHash,omitempty"`
	AggregateId         string     `json:"aggregateId,omitempty"`
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
		AbstractTransactionDTO
		Cosignatures      []*aggregateTransactionCosignatureDTO `json:"cosignatures"`
		InnerTransactions []map[string]interface{}              `json:"transactions"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
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

	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	as := make([]*AggregateTransactionCosignature, len(dto.Tx.Cosignatures))
	for i, a := range dto.Tx.Cosignatures {
		as[i], err = a.toStruct(atx.NetworkType)
	}
	if err != nil {
		return nil, err
	}

	return &AggregateTransaction{
		*atx,
		txs,
		as,
	}, nil
}

// MosaicDefinitionTransaction
type MosaicDefinitionTransaction struct {
	AbstractTransaction
	*MosaicProperties
	*NamespaceId
	*MosaicId
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
			"MosaicId": [ %s ],
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
		AbstractTransactionDTO
		Properties mosaicPropertiesDTO `json:"properties"`
		ParentId   *uint64DTO          `json:"parentId"`
		MosaicId   *uint64DTO          `json:"mosaicId"`
		MosaicName string              `json:"name"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *mosaicDefinitionTransactionDTO) toStruct() (*MosaicDefinitionTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	m, err := NewMosaicId(dto.Tx.MosaicId.toStruct(), "")
	if err != nil {
		return nil, err
	}

	return &MosaicDefinitionTransaction{
		*atx,
		dto.Tx.Properties.toStruct(),
		NewNamespaceId(dto.Tx.ParentId.toStruct(), ""),
		m,
		dto.Tx.MosaicName,
	}, nil
}

// MosaicSupplyChangeTransaction
type MosaicSupplyChangeTransaction struct {
	AbstractTransaction
	MosaicSupplyType
	*MosaicId
	Delta *big.Int
}

func (tx *MosaicSupplyChangeTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *MosaicSupplyChangeTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"MosaicSupplyType": %s,
			"MosaicId": [ %v ],
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
		AbstractTransactionDTO
		MosaicSupplyType `json:"direction"`
		MosaicId         *uint64DTO `json:"mosaicId"`
		Delta            *uint64DTO `json:"delta"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *mosaicSupplyChangeTransactionDTO) toStruct() (*MosaicSupplyChangeTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	m, err := NewMosaicId(dto.Tx.MosaicId.toStruct(), "")
	if err != nil {
		return nil, err
	}

	return &MosaicSupplyChangeTransaction{
		*atx,
		dto.Tx.MosaicSupplyType,
		m,
		dto.Tx.Delta.toStruct(),
	}, nil
}

// TransferTransaction
type TransferTransaction struct {
	AbstractTransaction
	Message
	Mosaics Mosaics
	*Address
}

func (tx *TransferTransaction) SignWith(account PublicAccount) (Signed, error) {
	return SignedTransaction{}, nil
}

func (tx *TransferTransaction) String() string {
	return fmt.Sprintf(
		`
			"AbstractTransaction": %s,
			"Mosaics": %s,
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
		AbstractTransactionDTO
		Message `json:"message"`
		Mosaics []*mosaicDTO `json:"mosaics"`
		Address string       `json:"recipient"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *transferTransactionDTO) toStruct() (*TransferTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	txs := make(Mosaics, len(dto.Tx.Mosaics))
	for i, tx := range dto.Tx.Mosaics {
		txs[i], err = tx.toStruct()
	}
	if err != nil {
		return nil, err
	}

	a, err := NewAddressFromEncoded(dto.Tx.Address)
	if err != nil {
		return nil, err
	}

	return &TransferTransaction{
		*atx,
		dto.Tx.Message,
		txs,
		a,
	}, nil
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
		AbstractTransactionDTO
		MinApprovalDelta int32                                 `json:"minApprovalDelta"`
		MinRemovalDelta  int32                                 `json:"minRemovalDelta"`
		Modifications    []*multisigCosignatoryModificationDTO `json:"modifications"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *modifyMultisigAccountTransactionDTO) toStruct() (*ModifyMultisigAccountTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	ms := make([]*MultisigCosignatoryModification, len(dto.Tx.Modifications))
	for i, m := range dto.Tx.Modifications {
		ms[i], err = m.toStruct(atx.NetworkType)
	}
	if err != nil {
		return nil, err
	}

	return &ModifyMultisigAccountTransaction{
		*atx,
		dto.Tx.MinApprovalDelta,
		dto.Tx.MinRemovalDelta,
		ms,
	}, nil
}

// RegisterNamespaceTransaction
type RegisterNamespaceTransaction struct {
	AbstractTransaction
	*NamespaceId
	NamespaceType
	NamspaceName string
	Duration     *big.Int
	ParentId     *NamespaceId
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
		AbstractTransactionDTO
		Id            namespaceIdDTO `json:"namespaceId"`
		NamespaceType `json:"namespaceType"`
		NamspaceName  string    `json:"name"`
		Duration      uint64DTO `json:"duration"`
		ParentId      namespaceIdDTO
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *registerNamespaceTransactionDTO) toStruct() (*RegisterNamespaceTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	var d *big.Int
	n := &NamespaceId{}
	if dto.Tx.NamespaceType == RootNamespace {
		d = dto.Tx.Duration.toStruct()
	} else {
		d = big.NewInt(0)
		n = dto.Tx.ParentId.toStruct()
	}

	return &RegisterNamespaceTransaction{
		*atx,
		dto.Tx.Id.toStruct(),
		dto.Tx.NamespaceType,
		dto.Tx.NamspaceName,
		d,
		n,
	}, nil
}

// LockFundsTransaction
type LockFundsTransaction struct {
	AbstractTransaction
	*Mosaic
	Duration *big.Int
	*SignedTransaction
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
		AbstractTransactionDTO
		Mosaic   mosaicDTO `json:"mosaic"`
		Duration uint64DTO `json:"duration"`
		Hash     string    `json:"hash"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *lockFundsTransactionDTO) toStruct() (*LockFundsTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	m, err := dto.Tx.Mosaic.toStruct()
	if err != nil {
		return nil, err
	}

	return &LockFundsTransaction{
		*atx,
		m,
		dto.Tx.Duration.toStruct(),
		&SignedTransaction{LOCK, "", dto.Tx.Hash},
	}, nil
}

// SecretLockTransaction
type SecretLockTransaction struct {
	AbstractTransaction
	*Mosaic
	HashType
	Duration  *big.Int
	Secret    string
	Recipient *Address
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
		AbstractTransactionDTO
		Mosaic    *mosaicDTO `json:"mosaic"`
		MosaicId  *uint64DTO `json:"mosaicId"`
		HashType  `json:"hashAlgorithm"`
		Duration  uint64DTO `json:"duration"`
		Secret    string    `json:"secret"`
		Recipient string    `json:"recipient"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *secretLockTransactionDTO) toStruct() (*SecretLockTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	m, err := dto.Tx.Mosaic.toStruct()
	if err != nil {
		return nil, err
	}

	a, err := NewAddressFromEncoded(dto.Tx.Recipient)
	if err != nil {
		return nil, err
	}

	return &SecretLockTransaction{
		*atx,
		m,
		dto.Tx.HashType,
		dto.Tx.Duration.toStruct(),
		dto.Tx.Secret,
		a,
	}, nil
}

// SecretProofTransaction
type SecretProofTransaction struct {
	AbstractTransaction
	HashType
	Secret string
	Proof  string
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
	Tx struct {
		AbstractTransactionDTO
		HashType `json:"hashAlgorithm"`
		Secret   string `json:"secret"`
		Proof    string `json:"proof"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *secretProofTransactionDTO) toStruct() (*SecretProofTransaction, error) {
	atx, err := dto.Tx.AbstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	return &SecretProofTransaction{
		*atx,
		dto.Tx.HashType,
		dto.Tx.Secret,
		dto.Tx.Proof,
	}, nil
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
	Signature string
	Signer    *PublicAccount
}

type aggregateTransactionCosignatureDTO struct {
	Signature string `json:"signature"`
	Signer    string
}

func (dto *aggregateTransactionCosignatureDTO) toStruct(networkType NetworkType) (*AggregateTransactionCosignature, error) {
	acc, err := NewPublicAccount(dto.Signer, networkType)
	if err != nil {
		return nil, err
	}
	return &AggregateTransactionCosignature{
		dto.Signature,
		acc,
	}, nil
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
	Type MultisigCosignatoryModificationType
	*PublicAccount
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

func (dto *multisigCosignatoryModificationDTO) toStruct(networkType NetworkType) (*MultisigCosignatoryModification, error) {
	acc, err := NewPublicAccount(dto.PublicAccount, networkType)
	if err != nil {
		return nil, err
	}

	return &MultisigCosignatoryModification{
		dto.Type,
		acc,
	}, nil
}

// TransactionStatus
type TransactionStatus struct {
	Group    string    `json:"group"`
	Status   string    `json:"status"`
	Hash     string    `json:"hash"`
	Deadline time.Time `json:"deadline"`
	Height   *big.Int  `json:"height"`
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
		ts.Deadline.Unix(),
		ts.Height,
	)
}

type transactionStatusDTO struct {
	Group    string    `json:"group"`
	Status   string    `json:"status"`
	Hash     string    `json:"hash"`
	Deadline uint64DTO `json:"deadline"`
	Height   uint64DTO `json:"height"`
}

func (dto *transactionStatusDTO) toStruct() (*TransactionStatus, error) {
	return &TransactionStatus{
		dto.Group,
		dto.Status,
		dto.Hash,
		time.Unix(dto.Deadline.toStruct().Int64(), int64(time.Millisecond)),
		dto.Height.toStruct(),
	}, nil
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
	ADD MultisigCosignatoryModificationType = iota
	REMOVE
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

func TransactionTypeFromRaw(value uint32) (TransactionType, error) {
	for _, t := range transactionTypes {
		if t.raw == value {
			return t.transactionType, nil
		}
	}
	return 0, transactionTypeError
}

func MapTransactions(b *bytes.Buffer) ([]Transaction, error) {
	var wg sync.WaitGroup
	var err error

	m := []jsonLib.RawMessage{}

	json.Unmarshal(b.Bytes(), &m)

	tx := make([]Transaction, len(m))
	for i, t := range m {
		wg.Add(1)
		go func(i int, t jsonLib.RawMessage) {
			defer wg.Done()
			json.Marshal(t)
			tx[i], err = MapTransaction(bytes.NewBuffer([]byte(t)))
		}(i, t)
	}
	wg.Wait()

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func MapTransaction(b *bytes.Buffer) (Transaction, error) {
	rawT := struct {
		Transaction struct {
			Type uint32
		}
	}{}

	err := json.Unmarshal(b.Bytes(), &rawT)
	if err != nil {
		return nil, err
	}

	t, err := TransactionTypeFromRaw(rawT.Transaction.Type)
	if err != nil {
		return nil, err
	}

	switch t {
	case AGGREGATE_BONDED:
		mapAggregateTransaction(b)
	case AGGREGATE_COMPLETE:
		mapAggregateTransaction(b)
	case MOSAIC_DEFINITION:
		dto := mosaicDefinitionTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case MOSAIC_SUPPLY_CHANGE:
		dto := mosaicSupplyChangeTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case MODIFY_MULTISIG_ACCOUNT:
		dto := modifyMultisigAccountTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case REGISTER_NAMESPACE:
		dto := registerNamespaceTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case TRANSFER:
		dto := transferTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case LOCK:
		dto := lockFundsTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case SECRET_LOCK:
		dto := secretLockTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	case SECRET_PROOF:
		dto := secretProofTransactionDTO{}

		err := json.Unmarshal(b.Bytes(), &dto)
		if err != nil {
			return nil, err
		}

		tx, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		return tx, nil
	}

	return nil, nil
}

func mapAggregateTransaction(b *bytes.Buffer) (*AggregateTransaction, error) {
	dto := aggregateTransactionDTO{}

	err := json.Unmarshal(b.Bytes(), &dto)
	if err != nil {
		return nil, err
	}

	tx, err := dto.toStruct()
	if err != nil {
		return nil, err
	}
	return tx, nil
}
