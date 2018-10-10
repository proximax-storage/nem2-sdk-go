package sdk

import "errors"

var (
	errNamespaceToManyPart = errors.New("too many parts")
	errNilIdNamespace      = errors.New("id nust not null")
	errNullAddress         = errors.New("address is null")
	errNilMosaicId         = errors.New("mosaicId must be not null")
	errNilMosaicAmount     = errors.New("amount must be not null")
)
