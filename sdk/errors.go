// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import "errors"

var (
	errNamespaceToManyPart = errors.New("too many parts")
	errNilIdNamespace      = errors.New("id nust not null")
	errEmptyNamespaceIds   = errors.New("list namespace ids must not by empty")
	errEmptyMosaicIds      = errors.New("list namespace ids must not by empty")
	errNullAddress         = errors.New("address is null")
	errNilMosaicId         = errors.New("mosaicId must be not null")
	errNilMosaicAmount     = errors.New("amount must be not null")
)
