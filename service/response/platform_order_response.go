package response

//车主端查看订单信息响应
type OrderResponse struct {
	Response
	OrderInfos []FactoryProductsInfo
}
