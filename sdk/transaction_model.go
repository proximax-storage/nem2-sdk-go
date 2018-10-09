package sdk

import (
	"bytes"
	"encoding/base32"
	"encoding/hex"
	jsonLib "encoding/json"
	"errors"
	"fmt"
	"github.com/google/flatbuffers/go"
	"github.com/proximax-storage/nem2-sdk-go/crypto"
	"github.com/proximax-storage/nem2-sdk-go/transactions"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Models
// Transaction
type Transaction interface {
	GetAbstractTransaction() *abstractTransaction
	String() string
	generateBytes() ([]byte, error)
}

// abstractTransaction
type abstractTransaction struct {
	NetworkType
	*TransactionInfo
	*Deadline
	Type      TransactionType
	Version   uint64
	Fee       *big.Int
	Signature string
	Signer    *PublicAccount
}

func (tx *abstractTransaction) IsUnconfirmed() bool {
	return tx.TransactionInfo != nil && tx.TransactionInfo.Height.Int64() == 0 && tx.TransactionInfo.Hash == tx.TransactionInfo.MerkleComponentHash
}

func (tx *abstractTransaction) IsConfirmed() bool {
	return tx.TransactionInfo != nil && tx.TransactionInfo.Height.Int64() > 0
}

func (tx *abstractTransaction) HasMissingSignatures() bool {
	return tx.TransactionInfo != nil && tx.TransactionInfo.Height.Int64() == 0 && tx.TransactionInfo.Hash != tx.TransactionInfo.MerkleComponentHash
}

func (tx *abstractTransaction) IsUnannounced() bool {
	return tx.TransactionInfo == nil
}

func (tx *abstractTransaction) String() string {
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

func (tx *abstractTransaction) generateVectors(builder *flatbuffers.Builder) (v uint64, signatureV, signerV, dV, fV flatbuffers.UOffsetT, err error) {
	v, err = strconv.ParseUint(strconv.FormatUint(uint64(tx.NetworkType), 16)+"0"+strconv.FormatUint(tx.Version, 16), 16, 32)
	signatureV = transactions.TransactionBufferCreateByteVector(builder, make([]byte, 64))
	signerV = transactions.TransactionBufferCreateByteVector(builder, make([]byte, 32))
	dV = transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(big.NewInt(tx.Deadline.GetInstant())))
	fV = transactions.TransactionBufferCreateUint32Vector(builder, []uint32{0, 0})
	return
}

func (tx *abstractTransaction) buildVectors(builder *flatbuffers.Builder, v uint64, signatureV, signerV, dV, fV flatbuffers.UOffsetT) {
	transactions.TransactionBufferAddSignature(builder, signatureV)
	transactions.TransactionBufferAddSigner(builder, signerV)
	transactions.TransactionBufferAddVersion(builder, v)
	transactions.TransactionBufferAddType(builder, tx.Type.Hex())
	transactions.TransactionBufferAddFee(builder, fV)
	transactions.TransactionBufferAddDeadline(builder, dV)
}

type abstractTransactionDTO struct {
	NetworkType `json:"networkType"`
	Type        uint32     `json:"type"`
	Version     uint64     `json:"version"`
	Fee         *uint64DTO `json:"fee"`
	Deadline    *uint64DTO `json:"deadline"`
	Signature   string     `json:"signature"`
	Signer      string     `json:"signer"`
}

func (dto *abstractTransactionDTO) toStruct(tInfo *TransactionInfo) (*abstractTransaction, error) {
	t, err := TransactionTypeFromRaw(dto.Type)
	if err != nil {
		return nil, err
	}

	nt := ExtractNetworkType(dto.Version)

	tv, err := ExtractVersion(dto.Version)
	if err != nil {
		return nil, err
	}

	pa, err := NewPublicAccount(dto.Signer, nt)
	if err != nil {
		return nil, err
	}

	var d *Deadline
	if dto.Deadline != nil {
		d = &Deadline{time.Unix(0, dto.Deadline.toBigInt().Int64()*int64(time.Millisecond))}
	}

	var f *big.Int
	if dto.Fee != nil {
		f = dto.Fee.toBigInt()
	}

	return &abstractTransaction{
		nt,
		tInfo,
		d,
		t,
		tv,
		f,
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
	height := big.NewInt(0)
	if dto.Height != nil {
		height = dto.Height.toBigInt()
	}
	return &TransactionInfo{
		height,
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
	abstractTransaction
	InnerTransactions []Transaction
	Cosignatures      []*AggregateTransactionCosignature
}

// Create an aggregate complete transaction
func NewCompleteAggregateTransaction(deadline *Deadline, innerTxs []Transaction, networkType NetworkType) (*AggregateTransaction, error) {
	if innerTxs == nil {
		return nil, errors.New("innerTransactions must not be nil")
	}
	return &AggregateTransaction{
		abstractTransaction: abstractTransaction{
			Type:        AggregateCompleted,
			Version:     2,
			Deadline:    deadline,
			NetworkType: networkType,
		},
		InnerTransactions: innerTxs,
	}, nil
}

func NewBondedAggregateTransaction(deadline *Deadline, innerTxs []Transaction, networkType NetworkType) (*AggregateTransaction, error) {
	if innerTxs == nil {
		return nil, errors.New("innerTransactions must not be nil")
	}
	return &AggregateTransaction{
		abstractTransaction: abstractTransaction{
			Type:        AggregateBonded,
			Version:     2,
			Deadline:    deadline,
			NetworkType: networkType,
		},
		InnerTransactions: innerTxs,
	}, nil
}

func (tx *AggregateTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *AggregateTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"InnerTransactions": %s,
			"Cosignatures": %s
		`,
		tx.abstractTransaction.String(),
		tx.InnerTransactions,
		tx.Cosignatures,
	)
}

func (tx *AggregateTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)

	var txsb []byte
	for _, itx := range tx.InnerTransactions {
		txb, err := toAggregateTransactionBytes(itx)
		if err != nil {
			return nil, err
		}
		txsb = append(txsb, txb...)
	}
	tV := transactions.TransactionBufferCreateByteVector(builder, txsb)

	v, signatureV, signerV, dV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.AggregateTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 120+4+len(txsb))
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, dV, fV)
	transactions.AggregateTransactionBufferAddTransactionsSize(builder, len(txsb))
	transactions.AggregateTransactionBufferAddTransactions(builder, tV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return aggregateTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type aggregateTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		Cosignatures      []*aggregateTransactionCosignatureDTO `json:"cosignatures"`
		InnerTransactions []map[string]interface{}              `json:"transactions"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *aggregateTransactionDTO) toStruct() (*AggregateTransaction, error) {
	txsr, err := json.Marshal(dto.Tx.InnerTransactions)
	if err != nil {
		return nil, err
	}

	txs, err := MapTransactions(bytes.NewBuffer(txsr))
	if err != nil {
		return nil, err
	}

	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
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

	for _, tx := range txs {
		iatx := tx.GetAbstractTransaction()
		iatx.Deadline = atx.Deadline
		iatx.Signature = atx.Signature
		iatx.Fee = atx.Fee
		if iatx.TransactionInfo == nil {
			iatx.TransactionInfo = atx.TransactionInfo
		}
	}

	return &AggregateTransaction{
		*atx,
		txs,
		as,
	}, nil
}

// MosaicDefinitionTransaction
type MosaicDefinitionTransaction struct {
	abstractTransaction
	*MosaicProperties
	*NamespaceId
	*MosaicId
	MosaicName string
}

func NewMosaicDefinitionTransaction(deadline *Deadline, mosaicName string, namespaceName string, mosaicProps *MosaicProperties, networkType NetworkType) (*MosaicDefinitionTransaction, error) {
	if mosaicName == "" {
		return nil, errors.New("mosaicName must not be empty")
	}
	if namespaceName == "" {
		return nil, errors.New("namespaceName must not be empty")
	}
	if mosaicProps == nil {
		return nil, errors.New("mosaicProps must not be nil")
	}

	id, err := generateMosaicId(namespaceName, mosaicName)
	if err != nil {
		return nil, err
	}

	return &MosaicDefinitionTransaction{
		abstractTransaction: abstractTransaction{
			Version:     2,
			Deadline:    deadline,
			Type:        MosaicDefinition,
			NetworkType: networkType,
		},
		MosaicName:  mosaicName,
		NamespaceId: NewNamespaceId(nil, namespaceName),
		MosaicId: &MosaicId{
			Id:       id,
			FullName: "",
		},
		MosaicProperties: mosaicProps,
	}, nil
}

func (tx *MosaicDefinitionTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *MosaicDefinitionTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"MosaicProperties": %s,
			"MosaicId": [ %s ],
			"MosaicName": %s
		`,
		tx.abstractTransaction.String(),
		tx.MosaicProperties.String(),
		tx.MosaicId,
		tx.MosaicName,
	)
}

func (tx *MosaicDefinitionTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)
	f := 0
	if tx.MosaicProperties.SupplyMutable {
		f += 1
	}
	if tx.MosaicProperties.Transferable {
		f += 2
	}
	if tx.MosaicProperties.LevyMutable {
		f += 4
	}

	mV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.MosaicId.Id))
	dV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.MosaicProperties.Duration))
	nV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.NamespaceId.Id))

	n := builder.CreateString(tx.MosaicName)

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.MosaicDefinitionTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 149+len(tx.MosaicName))
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.MosaicDefinitionTransactionBufferAddMosaicId(builder, mV)
	transactions.MosaicDefinitionTransactionBufferAddParentId(builder, nV)
	transactions.MosaicDefinitionTransactionBufferAddMosaicNameLength(builder, len(tx.MosaicName))
	transactions.MosaicDefinitionTransactionBufferAddNumOptionalProperties(builder, 1)
	transactions.MosaicDefinitionTransactionBufferAddFlags(builder, f)
	transactions.MosaicDefinitionTransactionBufferAddDivisibility(builder, tx.MosaicProperties.Divisibility)
	transactions.MosaicDefinitionTransactionBufferAddMosaicName(builder, n)
	transactions.MosaicDefinitionTransactionBufferAddIndicateDuration(builder, 2)
	transactions.MosaicDefinitionTransactionBufferAddDuration(builder, dV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)
	return mosaicDefinitionTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type mosaicDefinitionTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		Properties mosaicPropertiesDTO `json:"properties"`
		ParentId   *uint64DTO          `json:"parentId"`
		MosaicId   *uint64DTO          `json:"mosaicId"`
		MosaicName string              `json:"name"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *mosaicDefinitionTransactionDTO) toStruct() (*MosaicDefinitionTransaction, error) {
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	m, err := NewMosaicId(dto.Tx.MosaicId.toBigInt(), "")
	if err != nil {
		return nil, err
	}

	return &MosaicDefinitionTransaction{
		*atx,
		dto.Tx.Properties.toStruct(),
		NewNamespaceId(dto.Tx.ParentId.toBigInt(), ""),
		m,
		dto.Tx.MosaicName,
	}, nil
}

// MosaicSupplyChangeTransaction
type MosaicSupplyChangeTransaction struct {
	abstractTransaction
	MosaicSupplyType
	*MosaicId
	Delta *big.Int
}

func NewMosaicSupplyChangeTransaction(deadline *Deadline, mosaicId *MosaicId, supplyType MosaicSupplyType, delta *big.Int, networkType NetworkType) (*MosaicSupplyChangeTransaction, error) {
	if mosaicId == nil {
		return nil, errors.New("mosaicId must not be nil")
	}
	if !(supplyType == Increase || supplyType == Decrease) {
		return nil, errors.New("supplyType must not be nil")
	}
	if delta == nil {
		return nil, errors.New("delta must not be nil")
	}

	return &MosaicSupplyChangeTransaction{
		abstractTransaction: abstractTransaction{
			Version:     2,
			Deadline:    deadline,
			Type:        MosaicSupplyChange,
			NetworkType: networkType,
		},
		MosaicId:         mosaicId,
		MosaicSupplyType: supplyType,
		Delta:            delta,
	}, nil
}

func (tx *MosaicSupplyChangeTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *MosaicSupplyChangeTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"MosaicSupplyType": %s,
			"MosaicId": [ %v ],
			"Delta": %d
		`,
		tx.abstractTransaction.String(),
		tx.MosaicSupplyType.String(),
		tx.MosaicId,
		tx.Delta,
	)
}

func (tx *MosaicSupplyChangeTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)

	mV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.MosaicId.Id))
	dV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Delta))

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.MosaicSupplyChangeTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 137)
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.MosaicSupplyChangeTransactionBufferAddMosaicId(builder, mV)
	transactions.MosaicSupplyChangeTransactionBufferAddDirection(builder, uint8(tx.MosaicSupplyType))
	transactions.MosaicSupplyChangeTransactionBufferAddDelta(builder, dV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return mosaicSupplyChangeTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type mosaicSupplyChangeTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		MosaicSupplyType `json:"direction"`
		MosaicId         *uint64DTO `json:"mosaicId"`
		Delta            *uint64DTO `json:"delta"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *mosaicSupplyChangeTransactionDTO) toStruct() (*MosaicSupplyChangeTransaction, error) {
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	m, err := NewMosaicId(dto.Tx.MosaicId.toBigInt(), "")
	if err != nil {
		return nil, err
	}

	return &MosaicSupplyChangeTransaction{
		*atx,
		dto.Tx.MosaicSupplyType,
		m,
		dto.Tx.Delta.toBigInt(),
	}, nil
}

// TransferTransaction
type TransferTransaction struct {
	abstractTransaction
	*Message
	Mosaics   Mosaics
	Recipient *Address
}

// Create a transfer transaction
func NewTransferTransaction(deadline *Deadline, recipient *Address, mosaics Mosaics, message *Message, networkType NetworkType) (*TransferTransaction, error) {
	if recipient == nil {
		return nil, errors.New("recipient must not be nil")
	}
	if mosaics == nil {
		return nil, errors.New("mosaics must not be nil")
	}
	if message == nil {
		return nil, errors.New("message must not be nil, but could be with empty payload")
	}

	return &TransferTransaction{
		abstractTransaction: abstractTransaction{
			Version:     3,
			Deadline:    deadline,
			Type:        Transfer,
			NetworkType: networkType,
		},
		Recipient: recipient,
		Mosaics:   mosaics,
		Message:   message,
	}, nil
}

func (tx *TransferTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

// Convert an aggregate transaction to an inner transaction including transaction signer
func (tx *TransferTransaction) ToAggregate(publicAccount *PublicAccount) *abstractTransaction {
	tx.Signer = publicAccount
	return &tx.abstractTransaction
}

func (tx *TransferTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"Mosaics": %s,
			"Address": %s,
			"Message": %s,
		`,
		tx.abstractTransaction.String(),
		tx.Mosaics,
		tx.Recipient,
		tx.Message.String(),
	)
}

func (tx *TransferTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)

	ml := len(tx.Mosaics)
	mb := make([]flatbuffers.UOffsetT, ml)
	for i, mos := range tx.Mosaics {
		id := transactions.MosaicBufferCreateIdVector(builder, fromBigInt(mos.MosaicId.Id))
		am := transactions.MosaicBufferCreateAmountVector(builder, fromBigInt(mos.Amount))
		transactions.MosaicBufferStart(builder)
		transactions.MosaicBufferAddId(builder, id)
		transactions.MosaicBufferAddAmount(builder, am)
		mb[i] = transactions.MosaicBufferEnd(builder)
	}

	p := []byte(tx.Payload)
	pl := len(p)
	mp := transactions.TransactionBufferCreateByteVector(builder, p)
	transactions.MessageBufferStart(builder)
	transactions.MessageBufferAddType(builder, tx.Message.Type)
	transactions.MessageBufferAddPayload(builder, mp)
	m := transactions.TransactionBufferEnd(builder)

	r, err := base32.StdEncoding.DecodeString(tx.Recipient.Address)
	if err != nil {
		return nil, err
	}

	rV := transactions.TransactionBufferCreateByteVector(builder, r)
	mV := transactions.TransactionBufferCreateUOffsetVector(builder, mb)

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.TransferTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 149+(16*ml)+pl)
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.TransferTransactionBufferAddRecipient(builder, rV)
	transactions.TransferTransactionBufferAddNumMosaics(builder, ml)
	transactions.TransferTransactionBufferAddMessageSize(builder, pl+1)
	transactions.TransferTransactionBufferAddMessage(builder, m)
	transactions.TransferTransactionBufferAddMosaics(builder, mV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return transferTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type transferTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		messageDTO `json:"message"`
		Mosaics    []*mosaicDTO `json:"mosaics"`
		Address    string       `json:"recipient"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *transferTransactionDTO) toStruct() (*TransferTransaction, error) {
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
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
		dto.Tx.messageDTO.toStruct(),
		txs,
		a,
	}, nil
}

// ModifyMultisigAccountTransaction
type ModifyMultisigAccountTransaction struct {
	abstractTransaction
	MinApprovalDelta int32
	MinRemovalDelta  int32
	Modifications    []*MultisigCosignatoryModification
}

func NewModifyMultisigAccountTransaction(deadline *Deadline, minApprovalDelta int32, minRemovalDelta int32, modifications []*MultisigCosignatoryModification, networkType NetworkType) (*ModifyMultisigAccountTransaction, error) {
	if modifications == nil {
		return nil, errors.New("modifications must not be nil")
	}

	return &ModifyMultisigAccountTransaction{
		abstractTransaction: abstractTransaction{
			Version:     3,
			Deadline:    deadline,
			Type:        ModifyMultisig,
			NetworkType: networkType,
		},
		MinRemovalDelta:  minRemovalDelta,
		MinApprovalDelta: minApprovalDelta,
		Modifications:    modifications,
	}, nil
}

func (tx *ModifyMultisigAccountTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *ModifyMultisigAccountTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"MinApprovalDelta": %d,
			"MinRemovalDelta": %d,
			"Modifications": %s 
		`,
		tx.abstractTransaction.String(),
		tx.MinApprovalDelta,
		tx.MinRemovalDelta,
		tx.Modifications,
	)
}

func (tx *ModifyMultisigAccountTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)
	msb := make([]flatbuffers.UOffsetT, len(tx.Modifications))
	for i, m := range tx.Modifications {
		b, err := utils.HexDecodeStringOdd(m.PublicAccount.PublicKey)
		if err != nil {
			return nil, err
		}
		pV := transactions.TransactionBufferCreateByteVector(builder, b)
		transactions.CosignatoryModificationBufferStart(builder)
		transactions.CosignatoryModificationBufferAddType(builder, uint8(m.Type))
		transactions.CosignatoryModificationBufferAddCosignatoryPublicKey(builder, pV)
		msb[i] = transactions.TransactionBufferEnd(builder)
	}

	mV := transactions.TransactionBufferCreateUOffsetVector(builder, msb)

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.ModifyMultisigAccountTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 123+(33*len(tx.Modifications)))
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.ModifyMultisigAccountTransactionBufferAddMinApprovalDelta(builder, tx.MinApprovalDelta)
	transactions.ModifyMultisigAccountTransactionBufferAddMinRemovalDelta(builder, tx.MinRemovalDelta)
	transactions.ModifyMultisigAccountTransactionBufferAddNumModifications(builder, len(tx.Modifications))
	transactions.ModifyMultisigAccountTransactionBufferAddModifications(builder, mV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return modifyMultisigAccountTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type modifyMultisigAccountTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		MinApprovalDelta int32                                 `json:"minApprovalDelta"`
		MinRemovalDelta  int32                                 `json:"minRemovalDelta"`
		Modifications    []*multisigCosignatoryModificationDTO `json:"modifications"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *modifyMultisigAccountTransactionDTO) toStruct() (*ModifyMultisigAccountTransaction, error) {
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
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
	abstractTransaction
	*NamespaceId
	NamespaceType
	NamspaceName string
	Duration     *big.Int
	ParentId     *NamespaceId
}

func NewRegisterRootNamespaceTransaction(deadline *Deadline, namespaceName string, duration *big.Int, networkType NetworkType) (*RegisterNamespaceTransaction, error) {
	if namespaceName == "" {
		return nil, errors.New("namespaceName must not be nil")
	}
	if namespaceName == "" {
		return nil, errors.New("namespaceName must not be nil")
	}
	if duration == nil {
		return nil, errors.New("duration must not be nil")
	}

	id, err := generateNamespaceId(namespaceName)
	if err != nil {
		return nil, err
	}

	return &RegisterNamespaceTransaction{
		abstractTransaction: abstractTransaction{
			Version:     2,
			Deadline:    deadline,
			Type:        RegisterNamespace,
			NetworkType: networkType,
		},
		NamspaceName:  namespaceName,
		NamespaceId:   &NamespaceId{Id: id},
		NamespaceType: Root,
		Duration:      duration,
	}, nil
}

func NewRegisterSubNamespaceTransaction(deadline *Deadline, namespaceName string, parentId *NamespaceId, networkType NetworkType) (*RegisterNamespaceTransaction, error) {
	if namespaceName == "" {
		return nil, errors.New("namespaceName must not be nil")
	}
	if namespaceName == "" {
		return nil, errors.New("namespaceName must not be nil")
	}
	if parentId == nil {
		return nil, errors.New("parentId must not be nil")
	}

	id, err := generateId(namespaceName, parentId.Id)
	if err != nil {
		return nil, err
	}

	return &RegisterNamespaceTransaction{
		abstractTransaction: abstractTransaction{
			Version:     2,
			Deadline:    deadline,
			Type:        RegisterNamespace,
			NetworkType: networkType,
		},
		NamspaceName:  namespaceName,
		NamespaceId:   &NamespaceId{Id: id, FullName: namespaceName},
		NamespaceType: Sub,
		ParentId:      parentId,
	}, nil
}

func (tx *RegisterNamespaceTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *RegisterNamespaceTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"NamespaceName": %s,
			"Duration": %d
		`,
		tx.abstractTransaction.String(),
		tx.NamspaceName,
		tx.Duration,
	)
}

func (tx *RegisterNamespaceTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)

	nV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Id))
	var dV flatbuffers.UOffsetT
	if tx.NamespaceType == Root {
		dV = transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Duration))
	} else {
		dV = transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.ParentId.Id))
	}
	n := builder.CreateString(tx.NamspaceName)

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.RegisterNamespaceTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 138+len(tx.NamspaceName))
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.RegisterNamespaceTransactionBufferAddNamespaceType(builder, uint8(tx.NamespaceType))
	transactions.RegisterNamespaceTransactionBufferAddDurationParentId(builder, dV)
	transactions.RegisterNamespaceTransactionBufferAddNamespaceId(builder, nV)
	transactions.RegisterNamespaceTransactionBufferAddNamespaceNameSize(builder, len(tx.NamspaceName))
	transactions.RegisterNamespaceTransactionBufferAddNamespaceName(builder, n)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return registerNamespaceTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type registerNamespaceTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		Id            namespaceIdDTO `json:"namespaceId"`
		NamespaceType `json:"namespaceType"`
		NamspaceName  string    `json:"name"`
		Duration      uint64DTO `json:"duration"`
		ParentId      namespaceIdDTO
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *registerNamespaceTransactionDTO) toStruct() (*RegisterNamespaceTransaction, error) {
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
	if err != nil {
		return nil, err
	}

	var d *big.Int
	n := &NamespaceId{}
	if dto.Tx.NamespaceType == Root {
		d = dto.Tx.Duration.toBigInt()
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
	abstractTransaction
	*Mosaic
	Duration *big.Int
	*SignedTransaction
}

func NewLockFundsTransaction(deadline *Deadline, mosaic *Mosaic, duration *big.Int, signedTx *SignedTransaction, networkType NetworkType) (*LockFundsTransaction, error) {
	if mosaic == nil {
		return nil, errors.New("mosaic must not be nil")
	}
	if duration == nil {
		return nil, errors.New("duration must not be nil")
	}
	if signedTx == nil {
		return nil, errors.New("signedTx must not be nil")
	}
	if signedTx.TransactionType != AggregateBonded {
		return nil, errors.New("signedTx must be of type AggregateBonded")
	}

	return &LockFundsTransaction{
		abstractTransaction: abstractTransaction{
			Version:     3,
			Deadline:    deadline,
			Type:        Lock,
			NetworkType: networkType,
		},
		Mosaic:            mosaic,
		Duration:          duration,
		SignedTransaction: signedTx,
	}, nil
}

func (tx *LockFundsTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *LockFundsTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"Mosaic": %s,
			"Duration": %d,
			"Hash": %s
		`,
		tx.abstractTransaction.String(),
		tx.Mosaic.String(),
		tx.Duration,
		tx.Hash,
	)
}

func (tx *LockFundsTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)

	mv := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Mosaic.MosaicId.Id))
	maV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Mosaic.Amount))
	dV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Duration))

	h, err := hex.DecodeString(tx.SignedTransaction.Hash)
	if err != nil {
		return nil, err
	}
	hV := transactions.TransactionBufferCreateByteVector(builder, h)

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.LockFundsTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 176)
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.LockFundsTransactionBufferAddMosaicId(builder, mv)
	transactions.LockFundsTransactionBufferAddMosaicAmount(builder, maV)
	transactions.LockFundsTransactionBufferAddDuration(builder, dV)
	transactions.LockFundsTransactionBufferAddHash(builder, hV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return lockFundsTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type lockFundsTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		Mosaic   mosaicDTO `json:"mosaic"`
		Duration uint64DTO `json:"duration"`
		Hash     string    `json:"hash"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *lockFundsTransactionDTO) toStruct() (*LockFundsTransaction, error) {
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
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
		dto.Tx.Duration.toBigInt(),
		&SignedTransaction{Lock, "", dto.Tx.Hash},
	}, nil
}

// SecretLockTransaction
type SecretLockTransaction struct {
	abstractTransaction
	*Mosaic
	HashType
	Duration  *big.Int
	Secret    string
	Recipient *Address
}

func NewSecretLockTransaction(deadline *Deadline, mosaic *Mosaic, duration *big.Int, hashType HashType, secret string, recipient *Address, networkType NetworkType) (*SecretLockTransaction, error) {
	if mosaic == nil {
		return nil, errors.New("mosaic must not be nil")
	}
	if duration == nil {
		return nil, errors.New("duration must not be nil")
	}
	if secret == "" {
		return nil, errors.New("secret must not be empty")
	}
	if recipient == nil {
		return nil, errors.New("recipient must not be nil")
	}

	return &SecretLockTransaction{
		abstractTransaction: abstractTransaction{
			Version:     3,
			Deadline:    deadline,
			Type:        SecretLock,
			NetworkType: networkType,
		},
		Mosaic:    mosaic,
		Duration:  duration,
		HashType:  hashType,
		Secret:    secret, //TODO Add secret validation
		Recipient: recipient,
	}, nil
}

func (tx *SecretLockTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *SecretLockTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"Mosaic": %s,
			"Duration": %d,
			"HashType": %s,
			"Secret": %s,
			"Recipient": %s
		`,
		tx.abstractTransaction.String(),
		tx.Mosaic.String(),
		tx.Duration,
		tx.HashType.String(),
		tx.Secret,
		tx.Recipient,
	)
}

func (tx *SecretLockTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)

	mV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Mosaic.MosaicId.Id))
	maV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Mosaic.Amount))
	dV := transactions.TransactionBufferCreateUint32Vector(builder, fromBigInt(tx.Duration))

	s, err := hex.DecodeString(tx.Secret)
	if err != nil {
		return nil, err
	}
	sV := transactions.TransactionBufferCreateByteVector(builder, s)

	addr, err := base32.StdEncoding.DecodeString(tx.Recipient.Address)
	if err != nil {
		return nil, err
	}
	rV := transactions.TransactionBufferCreateByteVector(builder, addr)

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.SecretLockTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 234)
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.SecretLockTransactionBufferAddMosaicId(builder, mV)
	transactions.SecretLockTransactionBufferAddMosaicAmount(builder, maV)
	transactions.SecretLockTransactionBufferAddDuration(builder, dV)
	transactions.SecretLockTransactionBufferAddHashAlgorithm(builder, byte(tx.HashType))
	transactions.SecretLockTransactionBufferAddSecret(builder, sV)
	transactions.SecretLockTransactionBufferAddRecipient(builder, rV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return secretLockTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type secretLockTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
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
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
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
		dto.Tx.Duration.toBigInt(),
		dto.Tx.Secret,
		a,
	}, nil
}

// SecretProofTransaction
type SecretProofTransaction struct {
	abstractTransaction
	HashType
	Secret string
	Proof  string
}

func NewSecretProofTransaction(deadline *Deadline, hashType HashType, secret string, proof string, networkType NetworkType) (*SecretProofTransaction, error) {
	if proof == "" {
		return nil, errors.New("proof must not be empty")
	}
	if secret == "" {
		return nil, errors.New("secret must not be empty")
	}

	return &SecretProofTransaction{
		abstractTransaction: abstractTransaction{
			Version:     3,
			Deadline:    deadline,
			Type:        SecretProof,
			NetworkType: networkType,
		},
		HashType: hashType,
		Secret:   secret, //TODO Add secret validation
		Proof:    proof,
	}, nil
}

func (tx *SecretProofTransaction) GetAbstractTransaction() *abstractTransaction {
	return &tx.abstractTransaction
}

func (tx *SecretProofTransaction) String() string {
	return fmt.Sprintf(
		`
			"abstractTransaction": %s,
			"HashType": %s,
			"Secret": %s,
			"Proof": %s
		`,
		tx.abstractTransaction.String(),
		tx.HashType.String(),
		tx.Secret,
		tx.Proof,
	)
}

func (tx *SecretProofTransaction) generateBytes() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)

	s, err := hex.DecodeString(tx.Secret)
	if err != nil {
		return nil, err
	}
	sV := transactions.TransactionBufferCreateByteVector(builder, s)

	p, err := hex.DecodeString(tx.Proof)
	if err != nil {
		return nil, err
	}
	pV := transactions.TransactionBufferCreateByteVector(builder, p)

	v, signatureV, signerV, deadlineV, fV, err := tx.abstractTransaction.generateVectors(builder)
	if err != nil {
		return nil, err
	}

	transactions.SecretProofTransactionBufferStart(builder)
	transactions.TransactionBufferAddSize(builder, 187+len(p))
	tx.abstractTransaction.buildVectors(builder, v, signatureV, signerV, deadlineV, fV)
	transactions.SecretProofTransactionBufferAddHashAlgorithm(builder, byte(tx.HashType))
	transactions.SecretProofTransactionBufferAddSecret(builder, sV)
	transactions.SecretProofTransactionBufferAddProofSize(builder, len(p))
	transactions.SecretProofTransactionBufferAddProof(builder, pV)
	t := transactions.TransactionBufferEnd(builder)
	builder.Finish(t)

	return secretProofTransactionSchema().serialize(builder.FinishedBytes()), nil
}

type secretProofTransactionDTO struct {
	Tx struct {
		abstractTransactionDTO
		HashType `json:"hashAlgorithm"`
		Secret   string `json:"secret"`
		Proof    string `json:"proof"`
	} `json:"transaction"`
	TDto transactionInfoDTO `json:"meta"`
}

func (dto *secretProofTransactionDTO) toStruct() (*SecretProofTransaction, error) {
	atx, err := dto.Tx.abstractTransactionDTO.toStruct(dto.TDto.toStruct())
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

func NewCosignatureTransaction(txToCosign *AggregateTransaction) (*CosignatureTransaction, error) {
	if txToCosign == nil {
		return nil, errors.New("txToCosign must not be nil")
	}
	return &CosignatureTransaction{txToCosign}, nil
}

func (tx *CosignatureTransaction) String() string {
	return fmt.Sprintf(`"TransactionToCosign": %s`, tx.TransactionToCosign.String())
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
	*Deadline
	Group  string
	Status string
	Hash   string
	Height *big.Int
}

func (ts *TransactionStatus) String() string {
	return fmt.Sprintf(
		`
			"Group:" %s,
			"Status:" %s,
			"Hash": %s,
			"Deadline": %s,
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
	Group    string    `json:"group"`
	Status   string    `json:"status"`
	Hash     string    `json:"hash"`
	Deadline uint64DTO `json:"deadline"`
	Height   uint64DTO `json:"height"`
}

func (dto *transactionStatusDTO) toStruct() (*TransactionStatus, error) {
	return &TransactionStatus{
		&Deadline{time.Unix(dto.Deadline.toBigInt().Int64(), int64(time.Millisecond))},
		dto.Group,
		dto.Status,
		dto.Hash,
		dto.Height.toBigInt(),
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

var TimestampNemesisBlock = time.Unix(1459468800, 0)

// Deadline
type Deadline struct {
	time.Time
}

func (d *Deadline) GetInstant() int64 {
	return (d.Time.UnixNano() / 1e6) - (TimestampNemesisBlock.UnixNano() / 1e6)
}

// Create deadline model
func NewDeadline(d time.Duration) *Deadline {
	return &Deadline{time.Now().Add(d)}
}

// Message
type Message struct {
	Type    uint8
	Payload string
}

// The transaction message of 1024 characters.
func NewPlainMessage(payload string) *Message {
	return &Message{0, payload}
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

type messageDTO struct {
	Payload string `json:"payload"`
}

func (m *messageDTO) toStruct() *Message {
	b, err := hex.DecodeString(m.Payload)
	if err != nil {
		return &Message{0, m.Payload}
	}

	return &Message{0, string(b)}
}

type transactionTypeStruct struct {
	transactionType TransactionType
	raw             uint32
	hex             uint16
}

var transactionTypes = []transactionTypeStruct{
	{AggregateCompleted, 16705, 0x4141},
	{AggregateBonded, 16961, 0x4241},
	{MosaicDefinition, 16717, 0x414d},
	{MosaicSupplyChange, 16973, 0x424d},
	{ModifyMultisig, 16725, 0x4155},
	{RegisterNamespace, 16718, 0x414e},
	{Transfer, 16724, 0x4154},
	{Lock, 16716, 0x414C},
	{SecretLock, 16972, 0x424C},
	{SecretProof, 17228, 0x434C},
}

type TransactionType uint16

// TransactionType enums
const (
	AggregateCompleted TransactionType = iota
	AggregateBonded
	MosaicDefinition
	MosaicSupplyChange
	ModifyMultisig
	RegisterNamespace
	Transfer
	Lock
	SecretLock
	SecretProof
)

func (t TransactionType) Hex() uint16 {
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
	Add MultisigCosignatoryModificationType = iota
	Remove
)

type HashType uint8

func (ht HashType) String() string {
	return fmt.Sprintf("%d", ht)
}

const SHA3_512 HashType = 0

func ExtractVersion(version uint64) (uint64, error) {
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

	var m []jsonLib.RawMessage

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
	case AggregateBonded:
		return mapAggregateTransaction(b)
	case AggregateCompleted:
		return mapAggregateTransaction(b)
	case MosaicDefinition:
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
	case MosaicSupplyChange:
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
	case ModifyMultisig:
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
	case RegisterNamespace:
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
	case Transfer:
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
	case Lock:
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
	case SecretLock:
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
	case SecretProof:
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

func createTransactionHash(p string) (string, error) {
	b, err := hex.DecodeString(p)
	if err != nil {
		return "", err
	}
	sb := make([]byte, len(b)-36)
	copy(sb[:32], b[4:32+4])
	copy(sb[32:], b[68:])

	r, err := crypto.HashesSha3_256(sb)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString(r)), nil
}

func toAggregateTransactionBytes(tx Transaction) ([]byte, error) {
	if tx.GetAbstractTransaction().Signer == nil {
		return nil, fmt.Errorf("some of the transaction does not have a signer")
	}
	sb, err := hex.DecodeString(tx.GetAbstractTransaction().Signer.PublicKey)
	if err != nil {
		return nil, err
	}
	b, err := tx.generateBytes()
	if err != nil {
		return nil, err
	}

	rB := make([]byte, len(b)-64-16)
	copy(rB[4:32+4], sb[:32])
	copy(rB[32+4:32+4+4], b[100:104])
	copy(rB[32+4+4:32+4+4+len(b)-120], b[100+2+2+16:100+2+2+16+len(b)-120])

	s := big.NewInt(int64(len(b) - 64 - 16)).Bytes()
	utils.ReverseByteArray(s)

	copy(rB[:len(s)], s)

	return rB, nil
}

func signTransactionWith(tx Transaction, a *Account) (*SignedTransaction, error) {
	s := crypto.NewSignerFromKeyPair(a.KeyPair, nil)
	b, err := tx.generateBytes()
	if err != nil {
		return nil, err
	}
	sb := make([]byte, len(b)-100)
	copy(sb, b[100:])
	signature, err := s.Sign(sb)
	if err != nil {
		return nil, err
	}

	p := make([]byte, len(b))
	copy(p[:4], b[:4])
	copy(p[4:64+4], signature.Bytes())
	copy(p[64+4:64+4+32], a.KeyPair.PublicKey.Raw)
	copy(p[100:], b[100:])

	ph := hex.EncodeToString(p)
	h, err := createTransactionHash(ph)
	if err != nil {
		return nil, err
	}
	return &SignedTransaction{tx.GetAbstractTransaction().Type, strings.ToUpper(ph), h}, nil
}

func signTransactionWithCosignatures(tx *AggregateTransaction, a *Account, cosignatories []*Account) (*SignedTransaction, error) {
	stx, err := signTransactionWith(tx, a)
	if err != nil {
		return nil, err
	}

	p := stx.Payload

	b, err := hex.DecodeString(stx.Hash)
	if err != nil {
		return nil, err
	}

	for _, cos := range cosignatories {
		s := crypto.NewSignerFromKeyPair(cos.KeyPair, nil)
		sb, err := s.Sign(b)
		if err != nil {
			return nil, err
		}
		p += cos.KeyPair.PublicKey.String() + hex.EncodeToString(sb.Bytes())
	}

	pb, err := hex.DecodeString(p)
	if err != nil {
		return nil, err
	}

	s := big.NewInt(int64(len(pb))).Bytes()
	utils.ReverseByteArray(s)

	copy(pb[:len(s)], s)

	return &SignedTransaction{tx.Type, hex.EncodeToString(pb), stx.Hash}, nil
}

func signCosignatureTransaction(a *Account, tx *CosignatureTransaction) (*CosignatureSignedTransaction, error) {
	s := crypto.NewSignerFromKeyPair(a.KeyPair, nil)
	b, err := hex.DecodeString(tx.TransactionToCosign.TransactionInfo.Hash)
	if err != nil {
		return nil, err
	}

	sb, err := s.Sign(b)
	if err != nil {
		return nil, err
	}

	return &CosignatureSignedTransaction{tx.TransactionToCosign.TransactionInfo.Hash, hex.EncodeToString(sb.Bytes()), a.PublicAccount.PublicKey}, nil
}
