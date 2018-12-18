// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import "errors"

type RespErr struct {
	msg string
}

func newRespError(msg string) error {
	return &RespErr{msg: msg}
}

func (r *RespErr) Error() string {
	return r.msg
}

// Catapult REST API errors

var (
	ErrResourceNotFound              = newRespError("resource is not found")
	ErrArgumentNotValid              = newRespError("argument is not valid")
	ErrInvalidRequest                = newRespError("request is not valid")
	ErrInternalError                 = newRespError("response is nil")
	ErrNotAcceptedResponseStatusCode = newRespError("not accepted response status code")
)

// Mosaic errors
var (
	ErrEmptyMosaicIds      = errors.New("list mosaics ids must not by empty")
	ErrNilMosaicId         = errors.New("mosaicId must not be nil")
	ErrNilMosaicAmount     = errors.New("amount must be not nil")
	ErrInvalidMosaicName   = errors.New("mosaic name is invalid")
	ErrNilMosaicProperties = errors.New("mosaic properties must not be nil")
)

// Namespace errors
var (
	ErrNamespaceTooManyPart = errors.New("too many parts")
	ErrNilNamespaceId       = errors.New("namespaceId is nil or zero")
	ErrEmptyNamespaceIds    = errors.New("list namespace ids must not by empty")
	ErrInvalidNamespaceName = errors.New("namespace name is invalid")
)

// Blockchain errors
var (
	ErrNilOrZeroHeight = errors.New("block height should not be nil or zero")
	ErrNilOrZeroLimit  = errors.New("limit should not be nil or zero")
)

var (
	ErrEmptyAddressesIds = errors.New("list of addresses should not be empty")
	ErrNilAddress        = errors.New("address is nil")
	ErrBlankAddress      = errors.New("address is blank")
	ErrNilAccount        = errors.New("account should not be nil")
	ErrInvalidAddress    = errors.New("wrong address")
)
