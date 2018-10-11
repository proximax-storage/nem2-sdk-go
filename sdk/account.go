// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type AccountService service

func (a *AccountService) GetAccountInfo(ctx context.Context, address *Address) (*AccountInfo, *http.Response, error) {
	dto := &accountInfoDTO{}

	resp, err := a.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s", mainAccountRoute, address.Address), nil, dto)
	if err != nil {
		return nil, resp, err
	}

	acc, err := dto.toStruct()
	if err != nil {
		return nil, resp, err
	}

	return acc, resp, nil
}

// Gets AccountsInfo for different accounts.
func (a *AccountService) GetAccountsInfo(ctx context.Context, addresses []*Address) ([]*AccountInfo, *http.Response, error) {
	ads := make([]string, len(addresses))
	for i, ad := range addresses {
		ads[i] = ad.Address
	}
	addrs := struct {
		Messages []string `json:"addresses"`
	}{ads}

	dtos := make([]*accountInfoDTO, len(addresses))
	resp, err := a.client.DoNewRequest(ctx, "POST", mainAccountRoute, addrs, &dtos)
	if err != nil {
		return nil, resp, err
	}

	infos := make([]*AccountInfo, len(dtos))
	for i, dto := range dtos {
		infos[i], err = dto.toStruct()
	}
	if err != nil {
		return nil, resp, err
	}

	return infos, resp, nil
}

// Gets a MultisigAccountInfo for an account.
func (a *AccountService) GetMultisigAccountInfo(ctx context.Context, address *Address) (*MultisigAccountInfo, *http.Response, error) {
	dto := &multisigAccountInfoDTO{}
	resp, err := a.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", mainAccountRoute, address.Address, multisigAccountInfoRoute), nil, dto)
	if err != nil {
		return nil, resp, err
	}

	info, err := dto.toStruct(a.client.config.NetworkType)
	if err != nil {
		return nil, resp, err
	}

	return info, resp, nil
}

// Gets a MultisigAccountGraphInfo for an account.
func (a *AccountService) GetMultisigAccountGraphInfo(ctx context.Context, address *Address) (*MultisigAccountGraphInfo, *http.Response, error) {
	dto := &multisigAccountGraphInfoDTOS{}
	resp, err := a.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", mainAccountRoute, address.Address, multisigAccountGraphInfoRoute), nil, dto)
	if err != nil {
		return nil, resp, err
	}

	info, err := dto.toStruct(a.client.config.NetworkType)
	if err != nil {
		return nil, resp, err
	}

	return info, resp, nil
}

// Gets an array of confirmed transactions for which an account is signer or receiver.
func (a *AccountService) Transactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, transactionsRoute)
}

// Gets an array of transactions for which an account is the recipient of a transaction.
// A transaction is said to be incoming with respect to an account if the account is the recipient of a transaction.
func (a *AccountService) IncomingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, incomingTransactionsRoute)
}

// Gets an array of transactions for which an account is the sender a transaction.
// A transaction is said to be outgoing with respect to an account if the account is the sender of a transaction.
func (a *AccountService) OutgoingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, outgoingTransactionsRoute)
}

// Gets the array of transactions for which an account is the sender or receiver and which have not yet been included in a block.
// Unconfirmed transactions are those transactions that have not yet been included in a block.
// Unconfirmed transactions are not guaranteed to be included in any block.
func (a *AccountService) UnconfirmedTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, unconfirmedTransactionsRoute)
}

// Gets an array of transactions for which an account is the sender or has sign the transaction.
// A transaction is said to be aggregate bonded with respect to an account if there are missing signatures.
func (a *AccountService) AggregateBondedTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]*AggregateTransaction, *http.Response, error) {
	txs, resp, err := a.findTransactions(ctx, account, opt, aggregateTransactionsRoute)

	atxs := make([]*AggregateTransaction, len(txs))
	for i, tx := range txs {
		atxs[i] = tx.(*AggregateTransaction)
	}

	return atxs, resp, err
}

func (a *AccountService) findTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption, path string) ([]Transaction, *http.Response, error) {
	var b bytes.Buffer

	u, err := addOptions(fmt.Sprintf("%s/%s/%s", mainAccountRoute, account.PublicKey, path), opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := a.client.DoNewRequest(ctx, "GET", u, nil, &b)
	if err != nil {
		return nil, resp, err
	}

	txs, err := MapTransactions(&b)
	if err != nil {
		return nil, resp, err
	}

	return txs, resp, nil
}
