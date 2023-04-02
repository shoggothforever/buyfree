package response

import "buyfree/repo/model"

//车主端查看订单信息响应
type OrderResponse struct {
	Response
	OrderForm []FactoryProductsInfo
}
type SubmitResponse struct {
	Response
	OrderForm model.DriverOrderForm
}
