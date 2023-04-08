package response

import "buyfree/repo/model"

type PassengerHomeResponse struct {
	Response
	ADUrls             []model.ADurl                 `json:"ad_urls"`
	DeviceProductInfos []model.DeviceProductPartInfo `json:"device_product_infos"`
}

type PassengerPayInfo struct {
	DeviceID int64   `json:"device_id,omitempty"`
	Name     string  `json:"name,omitempty"`
	BuyPrice float64 `json:"buy_price,omitempty"`
}

type PassengerOrderFormResponse struct {
	Response
	PassengerOrderForms []model.PassengerOrderForm
}
