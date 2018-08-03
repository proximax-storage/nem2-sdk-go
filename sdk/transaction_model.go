package sdk

import (
	"bytes"
	jsonLib "encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

// Models

// Transaction
type Transaction interface {
	SignWith(account PublicAccount) (Signed, error)
	String() string
}

// AbstractTransaction
type AbstractTransaction struct {
	NetworkType     `json:"networkType"`
	TransactionInfo `json:"transactionInfo"`
	Type            TransactionType `json:"type"`
	Version         uint64          `json:"version"`
	Fee             []uint64        `json:"fee"`
	Deadline        []uint64        `json:"deadline"`
	Signature       string          `json:"signature"`
	Signer          string          `json:"signer"`
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

// AggregateTransaction
type AggregateTransaction struct {
	AbstractTransaction
	InnerTransactions []Transaction                     `json:"transactions"`
	Cosignatures      []AggregateTransactionCosignature `json:"cosignatures"`
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

// MosaicDefinitionTransaction
type MosaicDefinitionTransaction struct {
	AbstractTransaction
	//NamespaceId
	MosaicProperties
	MosaicId   []uint64 `json:"mosaicId"`
	MosaicName string   `json:"name"`
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

// MosaicSupplyChangeTransaction
type MosaicSupplyChangeTransaction struct {
	AbstractTransaction
	MosaicSupplyType `json:"direction"`
	MosaicId         []uint64 `json:"mosaicId"`
	Delta            uint64   `json:"delta"`
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

// tpl struct for encoding server responce
type MosaicDTO struct {
	MosaicId []uint64 `json:"id"`
	Amount   []uint64 `json:"amount"`
}
type MosaicsDTO []MosaicDTO

func (ref MosaicsDTO) String() string {
	s := "["
	for i, m := range ref {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf(
			`
			"MosaicId": %d,
			"Amount": %d 
		`,
			m.MosaicId,
			m.Amount,
		)
	}
	return s + "]"
}

// TransferTransaction
type TransferTransaction struct {
	AbstractTransaction
	Message `json:"message"`
	Mosaics MosaicsDTO `json:"mosaics"`
	Address string     `json:"recipient"`
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

// ModifyMultisigAccountTransaction
type ModifyMultisigAccountTransaction struct {
	AbstractTransaction
	MinApprovalDelta int32                             `json:"minApprovalDelta"`
	MinRemovalDelta  int32                             `json:"minRemovalDelta"`
	Modifications    []MultisigCosignatoryModification `json:"modifications"`
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

// RegisterNamespaceTransaction
type RegisterNamespaceTransaction struct {
	AbstractTransaction
	//NamespaceId
	NamspaceName string   `json:"name"`
	Duration     []uint64 `json:"duration"`
	//ParentId NamespaceId
	//NamespaceType
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

// SecretLockTransaction
type SecretLockTransaction struct {
	AbstractTransaction
	Mosaic    `json:"mosaic"`
	HashType  `json:"hashAlgorithm"`
	Duration  []uint64 `json:"duration"`
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

type CosignatureTransaction struct {
	TransactionToCosign AggregateTransaction
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
	Type          MultisigCosignatoryModificationType `json:"type"`
	PublicAccount string                              `json:"cosignatoryPublicKey"`
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

// TransactionStatus
type TransactionStatus struct {
	Group    string   `json:"group"`
	Status   string   `json:"status"`
	Hash     string   `json:"hash"`
	Deadline []uint64 `json:"deadline"`
	Height   []uint64 `json:"height"`
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
		mapAggregateTransaction(b, AGGREGATE_BONDED)
	case AGGREGATE_COMPLETE:
		mapAggregateTransaction(b, AGGREGATE_COMPLETE)
	case MOSAIC_DEFINITION:
		rawTx := struct {
			Tx struct {
				MosaicDefinitionTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.MosaicDefinitionTransaction

		aTx, err := mapAbstractTransaction(b, MOSAIC_DEFINITION)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case MOSAIC_SUPPLY_CHANGE:
		rawTx := struct {
			Tx struct {
				MosaicSupplyChangeTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.MosaicSupplyChangeTransaction

		aTx, err := mapAbstractTransaction(b, MOSAIC_SUPPLY_CHANGE)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case MODIFY_MULTISIG_ACCOUNT:
		rawTx := struct {
			Tx struct {
				ModifyMultisigAccountTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.ModifyMultisigAccountTransaction

		aTx, err := mapAbstractTransaction(b, MODIFY_MULTISIG_ACCOUNT)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case REGISTER_NAMESPACE:
		rawTx := struct {
			Tx struct {
				RegisterNamespaceTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.RegisterNamespaceTransaction

		aTx, err := mapAbstractTransaction(b, REGISTER_NAMESPACE)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case TRANSFER:
		rawTx := struct {
			Tx struct {
				TransferTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.TransferTransaction

		aTx, err := mapAbstractTransaction(b, TRANSFER)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case LOCK:
		rawTx := struct {
			Tx struct {
				LockFundsTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.LockFundsTransaction

		aTx, err := mapAbstractTransaction(b, LOCK)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case SECRET_LOCK:
		rawTx := struct {
			Tx struct {
				SecretLockTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.SecretLockTransaction

		aTx, err := mapAbstractTransaction(b, SECRET_LOCK)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	case SECRET_PROOF:
		rawTx := struct {
			Tx struct {
				SecretProofTransaction
			} `json:"transaction"`
		}{}

		err := json.Unmarshal(b.Bytes(), &rawTx)
		if err != nil {
			return nil, err
		}

		tx := rawTx.Tx.SecretProofTransaction

		aTx, err := mapAbstractTransaction(b, SECRET_PROOF)
		if err != nil {
			return nil, err
		}

		tx.AbstractTransaction = *aTx

		return &tx, nil
	}

	return nil, nil
}

func mapAbstractTransaction(b *bytes.Buffer, t TransactionType) (*AbstractTransaction, error) {
	rawTx := struct {
		Tx struct {
			AbstractTransaction
		} `json:"transaction"`
		TransactionInfo `json:"meta"`
	}{}

	err := json.Unmarshal(b.Bytes(), &rawTx)
	if err != nil {
		return nil, err
	}

	aTx := rawTx.Tx.AbstractTransaction
	aTx.Type = t
	aTx.TransactionInfo = rawTx.TransactionInfo

	nt, err := ExtractNetworkType(aTx.Version)
	if err != nil {
		return nil, err
	}

	tv, err := ExtractTransactionVersion(aTx.Version)

	aTx.Version = tv
	aTx.NetworkType = nt
	return &aTx, nil
}

func mapAggregateTransaction(b *bytes.Buffer, t TransactionType) (*AggregateTransaction, error) {
	rawTx := struct {
		Tx struct {
			AggregateTransaction
		} `json:"transaction"`
	}{}

	err := json.Unmarshal(b.Bytes(), &rawTx)
	if err != nil {
		return nil, err
	}

	tx := rawTx.Tx.AggregateTransaction

	aTx, err := mapAbstractTransaction(b, t)
	if err != nil {
		return nil, err
	}

	tx.AbstractTransaction = *aTx

	return &tx, nil
}
