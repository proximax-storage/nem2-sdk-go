package sdk

import (
	"fmt"
    "context"
	"net/http"
	"bytes"
)

type AccountService service

var mainAccountRoute = "account"

func (a *AccountService) GetAccountInfo(ctx context.Context, address *Address) (*AccountInfo, *http.Response, error) {
	dto := &accountInfoDTO{}

	resp, err := a.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s", mainTransactionRoute, address.Address),nil, dto)
	if err != nil {
		return nil, nil, err
	}

	return dto.toStruct(), resp, nil
}

func (a *AccountService) GetAccountsInfo(ctx context.Context, addresses *Addresses) ([]*AccountInfo, *http.Response, error) {
	var dtos []*accountInfoDTO
	var infos []*AccountInfo

	resp, err := a.client.DoNewRequest(ctx, "POST", mainTransactionRoute, addresses, dtos)
	if err != nil {
		return nil, nil, err
	}

	for i, dto := range dtos {
		infos[i] = dto.toStruct()
	}

	return infos, resp, nil
}

func (a *AccountService) GetMultisigAccountInfo(ctx context.Context, address *Address) (*MultisigAccountInfo, *http.Response, error) {
	dto := &multisigAccountInfoDTO{}
	resp, err := a.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s/multisig", mainTransactionRoute, address.Address), nil, dto)
	if err != nil {
		return nil, nil, err
	}

	info, err := dto.toStruct(a.client.config.NetworkType)
	if err != nil {
		return nil, nil, err
	}

	return info, resp, nil
}

func (a *AccountService) GetMultisigAccountGraphInfo(ctx context.Context, address *Address) (*MultisigAccountGraphInfo, *http.Response, error) {
	dto := &multisigAccountGraphInfoDTOS{}
	resp, err := a.client.DoNewRequest(ctx, "GET", fmt.Sprintf("%s/%s/multisig/graph", mainTransactionRoute, address.Address), nil, dto)
	if err != nil {
		return nil, nil, err
	}

	info, err := dto.toStruct(a.client.config.NetworkType)
	if err != nil {
		return nil, nil, err
	}

	return info, resp, nil
}

func (a *AccountService) Transactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, "/transactions")
}

func (a *AccountService) IncomingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, "/transactions/incoming")
}

func (a *AccountService) OutgoingTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, "/transactions/outgoing")
}

func (a *AccountService) UnconfirmedTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]Transaction, *http.Response, error) {
	return a.findTransactions(ctx, account, opt, "/transactions/unconfirmed")
}

func (a *AccountService) AggregateBondedTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption) ([]*AggregateTransaction, *http.Response, error) {
	var atxs []*AggregateTransaction
	txs, resp, err :=a.findTransactions(ctx, account, opt, "/transactions/aggregateBondedTransactions")

	for i, tx := range txs {
		atxs[i] = tx.(*AggregateTransaction)
	}

	return atxs, resp, err
}

func (a *AccountService) findTransactions(ctx context.Context, account *PublicAccount, opt *AccountTransactionsOption, path string) ([]Transaction, *http.Response, error) {
	var b bytes.Buffer

	u, err := addOptions(fmt.Sprintf("%s/%s/%s", mainTransactionRoute, account.PublicKey, path), opt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := a.client.DoNewRequest(ctx, "GET", u, nil, &b)
	if err != nil {
		return nil, nil, err
	}

	txs, err := MapTransactions(&b)
	if err != nil {
		return nil, nil, err
	}

	return txs, resp, nil
}