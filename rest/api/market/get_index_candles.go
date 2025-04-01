package market

import "github.com/wuqinqiang/go-okx/rest/api"

func NewGetIndexCandles(param *GetCandlesParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/market/index-candles",
		Method: api.MethodGet,
		Param:  param,
	}, &GetCandlesResponse{}
}
