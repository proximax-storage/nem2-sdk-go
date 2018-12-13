package sdk

import (
	"encoding/binary"
	"github.com/proximax-storage/nem2-sdk-go/utils"
	"golang.org/x/crypto/sha3"
	"math/big"
)

func bigIntToNamespaceId(bigInt *big.Int) *NamespaceId {
	if bigInt == nil {
		return nil
	}

	nsId := NamespaceId(*bigInt)

	return &nsId
}

func namespaceIdToBigInt(nsId *NamespaceId) *big.Int {
	if nsId == nil {
		return nil
	}

	return (*big.Int)(nsId)
}

type namespaceIdDTO uint64DTO

func (dto *namespaceIdDTO) toStruct() (*NamespaceId, error) {
	return NewNamespaceId(uint64DTO(*dto).toBigInt())
}

// namespaceNameDTO temporary struct for reading responce & fill NamespaceName
type namespaceNameDTO struct {
	NamespaceId uint64DTO
	Name        string
	ParentId    uint64DTO
}

func (ref *namespaceNameDTO) toStruct() (*NamespaceName, error) {
	nsId, err := NewNamespaceId(ref.NamespaceId.toBigInt())
	if err != nil {
		return nil, err
	}

	parentId, err := NewNamespaceId(ref.ParentId.toBigInt())
	if err != nil {
		return nil, err
	}

	return &NamespaceName{
		nsId,
		ref.Name,
		parentId,
	}, nil
}

func namespaceNameDTOsToNamespaceNames(namespaceNameDTOs []*namespaceNameDTO) ([]*NamespaceName, error) {
	nsNames := make([]*NamespaceName, 0, len(namespaceNameDTOs))

	if len(namespaceNameDTOs) == 0 {
		return nsNames, nil
	}

	for _, dto := range namespaceNameDTOs {
		nsName, err := dto.toStruct()
		if err != nil {
			return nil, err
		}

		nsNames = append(nsNames, nsName)
	}

	return nsNames, nil
}

// namespaceDTO temporary struct for reading responce & fill NamespaceInfo
type namespaceDTO struct {
	NamespaceId   uint64DTO
	FullName      string
	Type          int
	Depth         int
	Level0        *uint64DTO
	Level1        *uint64DTO
	Level2        *uint64DTO
	ParentId      uint64DTO
	MosaicIds     []uint64DTO
	SubNamespaces []uint64DTO
	Owner         string
	OwnerAddress  string
	StartHeight   uint64DTO
	EndHeight     uint64DTO
}

// namespaceInfoDTO temporary struct for reading response & fill NamespaceInfo
type namespaceInfoDTO struct {
	Meta      namespaceMosaicMetaDTO
	Namespace namespaceDTO
}

//toStruct create & return new NamespaceInfo from namespaceInfoDTO
func (ref *namespaceInfoDTO) toStruct() (*NamespaceInfo, error) {
	nsId, err := NewNamespaceId(ref.Namespace.NamespaceId.toBigInt())
	if err != nil {
		return nil, err
	}

	pubAcc, err := NewAccountFromPublicKey(ref.Namespace.Owner, NetworkType(ref.Namespace.Type))
	if err != nil {
		return nil, err
	}

	parentId, err := NewNamespaceId(ref.Namespace.ParentId.toBigInt())
	if err != nil {
		return nil, err
	}

	levels, err := ref.extractLevels()
	if err != nil {
		return nil, err
	}

	subNses := make([]*NamespaceId, 0, len(ref.Namespace.SubNamespaces))

	for _, nsIdDTO := range ref.Namespace.SubNamespaces {
		nsId, err := NewNamespaceId(nsIdDTO.toBigInt())
		if err != nil {
			return nil, err
		}

		subNses = append(subNses, nsId)
	}

	mscIds := make([]*MosaicId, 0, len(ref.Namespace.MosaicIds))

	for _, mscIdDTO := range ref.Namespace.MosaicIds {
		mscId, err := NewMosaicId(mscIdDTO.toBigInt())
		if err != nil {
			return nil, err
		}

		mscIds = append(mscIds, mscId)
	}

	return &NamespaceInfo{
		NamespaceId:     nsId,
		FullName:        ref.Namespace.FullName,
		Active:          ref.Meta.Active,
		Index:           ref.Meta.Index,
		MetaId:          ref.Meta.Id,
		TypeSpace:       NamespaceType(ref.Namespace.Type),
		Depth:           ref.Namespace.Depth,
		Levels:          levels,
		ParentId:        parentId,
		Owner:           pubAcc,
		MosaicIds:       mscIds,
		SubNamespaceIds: subNses,
		StartHeight:     ref.Namespace.StartHeight.toBigInt(),
		EndHeight:       ref.Namespace.EndHeight.toBigInt(),
	}, nil
}

func (ref *namespaceInfoDTO) extractLevels() ([]*NamespaceId, error) {
	levels := make([]*NamespaceId, 0)

	if ref.Namespace.Level0 != nil {
		nsName, err := NewNamespaceId(ref.Namespace.Level0.toBigInt())
		if err != nil {
			return nil, err
		}

		levels = append(levels, nsName)
	}

	if ref.Namespace.Level1 != nil {
		nsName, err := NewNamespaceId(ref.Namespace.Level1.toBigInt())
		if err != nil {
			return nil, err
		}

		levels = append(levels, nsName)
	}

	if ref.Namespace.Level2 != nil {
		nsName, err := NewNamespaceId(ref.Namespace.Level2.toBigInt())
		if err != nil {
			return nil, err
		}

		levels = append(levels, nsName)
	}

	return levels, nil
}

func namespaceDTOsToNamespaceInfos(res []*namespaceInfoDTO) ([]*NamespaceInfo, error) {
	nsInfos := make([]*NamespaceInfo, 0, len(res))

	for _, nsInfoDTO := range res {
		nsInfo, err := nsInfoDTO.toStruct()
		if err != nil {
			return nil, err
		}

		nsInfos = append(nsInfos, nsInfo)
	}

	return nsInfos, nil
}

func generateId(name string, parentId *big.Int) (*big.Int, error) {
	b := make([]byte, 8)

	if parentId.Int64() != 0 {
		b = parentId.Bytes()
	}

	utils.ReverseByteArray(b)

	result := sha3.New256()

	if _, err := result.Write(b); err != nil {
		return nil, err
	}

	if _, err := result.Write([]byte(name)); err != nil {
		return nil, err
	}

	t := result.Sum(nil)

	return uint64DTO{binary.LittleEndian.Uint32(t[0:4]), binary.LittleEndian.Uint32(t[4:8])}.toBigInt(), nil
}
