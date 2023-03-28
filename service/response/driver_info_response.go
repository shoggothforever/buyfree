package response

import "buyfree/repo/model"

type DriverOrderResponse struct {
	Response
	OrderInfos []model.DriverOrderForm
}
type DriverOrdersResponse struct {
	Response
	OrderInfos []model.DriverOrderForm
}
type DriverOrderDetailResponse struct {
	Response
	FactoryAddress string  `json:"factory_address" form:"factory_address"`
	ReserveMobile  string  `json:"reserve_mobile" form:"reserve_mobile"`
	Distance       float64 `json:"distance" form:"distance"`
	OrderInfos     model.DriverOrderForm
}
type DriverDeviceResponse struct {
	Response
	Devices []model.Device
}
