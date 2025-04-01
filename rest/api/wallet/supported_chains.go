package wallet

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewSupportedChains() (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/wallet/chain/supported-chains",
		Method: api.MethodGet,
		Param:  nil,
	}, &SupportedChainsResp{}
}

type SupportedChainsResp struct {
	api.Response
	Data []Chain `json:"data"`
}

func (c *SupportedChainsResp) GetByChainIndex(chainIndex string) (Chain, bool) {
	for i := range c.Data {
		if c.Data[i].ChainIndex == chainIndex {
			return c.Data[i], true
		}
	}
	return Chain{}, false
}

type Chain struct {
	Name       string `json:"name"`
	LogoUrl    string `json:"logoUrl"`
	ShortName  string `json:"shortName"`
	ChainIndex string `json:"chainIndex"`
}
