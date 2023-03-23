package response

import (
	"buyfree/repo/model"
)

type ADResponse struct {
	Response
	ADInfos []model.Advertisement
}

type ADEfficientInfo struct {
	DeviceID    int64   `json:"device_id"`
	PlayedTimes int64   `json:"played_times"`
	Profit      float64 `json:"profit"`
	CarID       string  `json:"car_id"`
	DriverName  string  `json:"driver_name"`
}
type ADEfficientResponse struct {
	Response
	ADEfficientInfos []ADEfficientInfo
}
