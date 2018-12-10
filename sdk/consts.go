// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import "regexp"

var XemMosaicId, _ = NewMosaicIdFromName("nem:xem")

// const routers path for methods AccountService
const (
	mainAccountRoute              = "account"
	transactionsRoute             = "transactions"
	incomingTransactionsRoute     = "transactions/incoming"
	outgoingTransactionsRoute     = "transactions/outgoing"
	unconfirmedTransactionsRoute  = "transactions/unconfirmed"
	aggregateTransactionsRoute    = "transactions/aggregateBondedTransactions"
	multisigAccountInfoRoute      = "multisig"
	multisigAccountGraphInfoRoute = "multisig/graph"
)

// const routers path for methods NamespaceService
const (
	pathNamespacesFromAccounts = "/account/namespaces"
	pathNamespace              = "/namespace/"
	pathNamespacenames         = "/namespace/names"
	pathNamespacesFromAccount  = "/account/%s/namespaces"
)

// const routers path for methods MosaicService
const (
	pathMosaic              = "/mosaic/"
	pathMosaicNames         = "/mosaic/names"
	pathMosaicFromNamespace = "/namespace/%s/mosaics/"
)

// const routers path for methods BlockchainService
const (
	pathBlockHeight         = "/chain/height"
	pathBlockByHeight       = "/block/%s"
	pathBlockScore          = "/chain/score"
	pathBlockGetTransaction = "/block/%s/transactions"
	pathBlockInfo           = "/blocks/%s/limit/%s"
	pathBlockStorage        = "/diagnostic/storage"
)

// const routers path for methods MosaicService
const (
	pathNetwork = "/network"
)

// const routers path for methods TransactionService
const (
	mainTransactionRoute               = "transaction"
	announceAggreagateRoute            = "partial"
	announceAggreagateCosignatureRoute = "cosignature"
	transactionStatusRoute             = "status"
	transactionStatusesRoute           = "statuses"
)

type NamespaceType uint8

const (
	Root NamespaceType = iota
	Sub
)

// regValidNamespace check namespace on valid symbols
var regValidNamespace = regexp.MustCompile(`^[a-z0-9][a-z0-9-_]*$`)
