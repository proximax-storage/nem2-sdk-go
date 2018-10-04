package sdk

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type AccountService service

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

func (a *AccountService) Transactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, transactionsRoute)
}

func (a *AccountService) IncomingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, incomingTransactionsRoute)
}

func (a *AccountService) OutgoingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, outgoingTransactionsRoute)
}

func (a *AccountService) UnconfirmedTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, unconfirmedTransactionsRoute)
}

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
