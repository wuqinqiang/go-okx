package convert

import "github.com/wuqinqiang/go-okx/rest/api"

func NewGetCurrencies() (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/asset/convert/currencies",
		Method: api.MethodGet,
	}, &GetCurrenciesResponse{}
}

type GetCurrenciesResponse struct {
	api.Response
	Data []Currency `json:"data"`
}

type Currency struct {
	Ccy string `json:"ccy"`
	Min string `json:"min"`
	Max string `json:"max"`
}
