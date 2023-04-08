package response

import "buyfree/repo/model"

type PassengerHomeResponse struct {
	Response
	ADUrls             []model.ADurl                 `json:"ad_urls"`
	DeviceProductInfos []model.DeviceProductPartInfo `json:"device_product_infos"`
}
