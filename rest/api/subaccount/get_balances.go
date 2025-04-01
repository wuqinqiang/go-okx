package subaccount

import (
	"github.com/wuqinqiang/go-okx/rest/api"
	"github.com/wuqinqiang/go-okx/rest/api/account"
)

func NewGetBalances(param *GetBalancesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/subaccount/balances",
		Method: api.MethodGet,
		Param:  param,
	}, &account.GetBalanceResponse{}
}

type GetBalancesParam struct {
	SubAcct string `url:"subAcct"`
}
