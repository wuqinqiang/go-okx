package wallet

import (
	"errors"
	"github.com/iaping/go-okx/rest/api"
)

func NewTransactionDetailByTxHash(param *TransactionDetailByTxHashParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/wallet/post-transaction/transaction-detail-by-txhash",
		Method: api.MethodGet,
		Param:  param,
	}, &TransactionDetailByTxHashResp{}
}

type TransactionDetailByTxHashParam struct {
	ChainIndex string `url:"chainIndex,omitempty"`
	TxHash     string `url:"txHash,omitempty"`
}

type TransactionDetailByTxHashResp struct {
	api.Response
	Data []Trade `json:"data"`
}

type Trade struct {
	ChainIndex   string `json:"chainIndex"`
	Height       string `json:"height"`
	TxTime       string `json:"txTime"`
	Txhash       string `json:"txhash"`
	GasLimit     string `json:"gasLimit"`
	GasUsed      string `json:"gasUsed"`
	GasPrice     string `json:"gasPrice"`
	TxFee        string `json:"txFee"`
	Nonce        string `json:"nonce"`
	Symbol       string `json:"symbol"`
	Amount       string `json:"amount"`
	TxStatus     string `json:"txStatus"`
	MethodId     string `json:"methodId"`
	L1OriginHash string `json:"l1OriginHash"`
	FromDetails  []struct {
		Address      string `json:"address"`
		VinIndex     string `json:"vinIndex"`
		PreVoutIndex string `json:"preVoutIndex"`
		TxHash       string `json:"txHash"`
		IsContract   bool   `json:"isContract"`
		Amount       string `json:"amount"`
	} `json:"fromDetails"`
	ToDetails []struct {
		Address    string `json:"address"`
		VoutIndex  string `json:"voutIndex"`
		IsContract bool   `json:"isContract"`
		Amount     string `json:"amount"`
	} `json:"toDetails"`
	InternalTransactionDetails []struct {
		From           string `json:"from"`
		To             string `json:"to"`
		IsFromContract bool   `json:"isFromContract"`
		IsToContract   bool   `json:"isToContract"`
		Amount         string `json:"amount"`
		State          string `json:"state"`
	} `json:"internalTransactionDetails"`
	TokenTransferDetails []Token `json:"tokenTransferDetails"`
}

type Token struct {
	Amount               string `json:"amount"`
	From                 string `json:"from"`
	IsFromContract       bool   `json:"isFromContract"`
	IsToContract         bool   `json:"isToContract"`
	Symbol               string `json:"symbol"`
	To                   string `json:"to"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

func (t *TransactionDetailByTxHashResp) GetTokenDetail() (Token, error) {
	if len(t.Data) == 0 {
		return Token{}, errors.New("no supported chains found")
	}
	tran := t.Data[0]
	if len(tran.TokenTransferDetails) == 0 {
		return Token{}, errors.New("token transfer details not found")
	}

	return tran.TokenTransferDetails[0], nil
}
