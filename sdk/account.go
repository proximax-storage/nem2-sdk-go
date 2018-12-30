// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/proximax-storage/proximax-utils-go/net"
	"net/http"
)

type AccountService service

func (a *AccountService) GetAccountInfo(ctx context.Context, address *Address) (*AccountInfo, error) {
	if address == nil {
		return nil, ErrNilAddress
	}

	if len(address.Address) == 0 {
		return nil, ErrBlankAddress
	}

	url := net.NewUrl(fmt.Sprintf(accountRoute, address.Address))

	dto := &accountInfoDTO{}

	resp, err := a.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, dto)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return dto.toStruct()
}

// Gets AccountsInfo for different accounts.
func (a *AccountService) GetAccountsInfo(ctx context.Context, addresses []*Address) ([]*AccountInfo, error) {
	if len(addresses) == 0 {
		return nil, ErrEmptyAddressesIds
	}

	addrs := struct {
		Messages []string `json:"addresses"`
	}{
		Messages: make([]string, len(addresses)),
	}

	for i, address := range addresses {
		addrs.Messages[i] = address.Address
	}

	dtos := accountInfoDTOs(make([]*accountInfoDTO, 0))

	resp, err := a.client.DoNewRequest(ctx, http.MethodPost, accountsRoute, addrs, &dtos)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return dtos.toStruct()
}

// Gets a MultisigAccountInfo for an account.
func (a *AccountService) GetMultisigAccountInfo(ctx context.Context, address *Address) (*MultisigAccountInfo, error) {
	if address == nil {
		return nil, ErrNilAddress
	}

	url := net.NewUrl(fmt.Sprintf(multisigAccountRoute, address.Address))

	dto := &multisigAccountInfoDTO{}

	resp, err := a.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, dto)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return dto.toStruct(a.client.config.NetworkType)
}

// Gets a MultisigAccountGraphInfo for an account.
func (a *AccountService) GetMultisigAccountGraphInfo(ctx context.Context, address *Address) (*MultisigAccountGraphInfo, error) {
	if address == nil {
		return nil, ErrNilAddress
	}

	url := net.NewUrl(fmt.Sprintf(multisigAccountGraphInfoRoute, address.Address))

	dto := &multisigAccountGraphInfoDTOS{}

	resp, err := a.client.DoNewRequest(ctx, http.MethodGet, url.Encode(), nil, dto)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{404: ErrResourceNotFound, 409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return dto.toStruct(a.client.config.NetworkType)
}

// Gets an array of confirmed transactions for which an account is signer or receiver.
func (a *AccountService) Transactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, error) {
	return a.findTransactions(ctx, account, opt, accountTransactionsRoute)
}

// Gets an array of transactions for which an account is the recipient of a transaction.
// A transaction is said to be incoming with respect to an account if the account is the recipient of a transaction.
func (a *AccountService) IncomingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, error) {
	return a.findTransactions(ctx, account, opt, incomingTransactionsRoute)
}

// Gets an array of transactions for which an account is the sender a transaction.
// A transaction is said to be outgoing with respect to an account if the account is the sender of a transaction.
func (a *AccountService) OutgoingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, error) {
	return a.findTransactions(ctx, account, opt, outgoingTransactionsRoute)
}

// Gets the array of transactions for which an account is the sender or receiver and which have not yet been included in a block.
// Unconfirmed transactions are those transactions that have not yet been included in a block.
// Unconfirmed transactions are not guaranteed to be included in any block.
func (a *AccountService) UnconfirmedTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, error) {
	return a.findTransactions(ctx, account, opt, unconfirmedTransactionsRoute)
}

// Gets an array of transactions for which an account is the sender or has sign the transaction.
// A transaction is said to be aggregate bonded with respect to an account if there are missing signatures.
func (a *AccountService) AggregateBondedTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]*AggregateTransaction, error) {
	txs, err := a.findTransactions(ctx, account, opt, aggregateTransactionsRoute)
	if err != nil {
		return nil, err
	}

	atxs := make([]*AggregateTransaction, len(txs))
	for i, tx := range txs {
		atxs[i] = tx.(*AggregateTransaction)
	}

	return atxs, nil
}

func (a *AccountService) findTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption, path string) ([]Transaction, error) {
	if account == nil {
		return nil, ErrNilAccount
	}

	var b bytes.Buffer

	u, err := addOptions(fmt.Sprintf(transactionsByAccountRoute, account.PublicKey, path), opt)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.DoNewRequest(ctx, http.MethodGet, u, nil, &b)
	if err != nil {
		return nil, err
	}

	if err = handleResponseStatusCode(resp, map[int]error{409: ErrArgumentNotValid}); err != nil {
		return nil, err
	}

	return MapTransactions(&b)
}
