// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import "errors"

// Catapult REST API errors

var (
	// TODO
	ErrCatapultRestAPIError = errors.New("")
	ErrResourceNotFound     = errors.New("resource is not found")
	ErrArgumentNotValid     = errors.New("argument is not valid")
	ErrInvalidRequest       = errors.New("request is not valid")
	ErrInternalError        = errors.New("response is nil")
)

// MosaicId API errors
var (
	ErrEmptyMosaicIds  = errors.New("list mosaics ids must not by empty")
	ErrNilMosaicId     = errors.New("mosaicId must not be nil")
	ErrNilMosaicAmount = errors.New("amount must be not nil")
)

// Namespace API errors
var (
	ErrNamespaceToManyPart = errors.New("too many parts")
	ErrNilIdNamespace      = errors.New("namespaceId must not be nil")
	ErrEmptyNamespaceIds   = errors.New("list namespace ids must not by empty")
)

var (
	ErrEmptyAddressesIds = errors.New("list of addresses should not be nil")
	ErrNilAddress        = errors.New("address is nil")
)
