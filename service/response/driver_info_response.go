package response

import "buyfree/repo/model"

type DriverOrderFormResponse struct {
	Response
	OrderInfos []model.DriverOrderForm `json:"order_infos"`
}
type DriverOrdersResponse struct {
	Response
	Cash              float64                 `json:"cash"`
	FactoriesDistance []FactoryDistanceReq    `json:"factories_distance,omitempty"`
	OrderInfos        []model.DriverOrderForm `json:"order_infos,omitempty"`
}
type SubmitOrderForms struct {
	Cash              float64                 `json:"cash"`
	FactoriesDistance []FactoryDistanceReq    `json:"factories_distance,omitempty"`
	OrderInfos        []model.DriverOrderForm `json:"order_infos,omitempty"`
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
	Devices []model.Device `json:"devices"`
}
type LoadResponse struct {
	Response
	DevProducts []model.DeviceProduct `json:"dev_products"`
}
type BalanceResponse struct {
	Response
	Fund float64 `json:"fund"`
}
