package trade

import "github.com/wuqinqiang/go-okx/rest/api"

func NewPostAmendBatchOrders(param []*PostAmendOrderParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/amend-batch-orders",
		Method: api.MethodPost,
		Param:  param,
	}, &PostAmendOrderResponse{}
}
