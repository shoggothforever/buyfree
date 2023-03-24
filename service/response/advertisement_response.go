package response

import (
	"buyfree/repo/model"
)

//获取广告信息以及添加广告的响应
type ADResponse struct {
	Response
	ADInfos []model.Advertisement
}

//广告效益信息
type ADEfficientInfo struct {
	DeviceID int64 `json:"device_id"`
	//已经播放次数
	PlayedTimes int64 `json:"played_times"`
	//收益
	Profit float64 `json:"profit"`
	//车主车牌号
	CarID string `json:"car_id"`
	//车主姓名
	DriverName string `json:"driver_name"`
}

//广告效益相应
type ADEfficientResponse struct {
	Response
	ADEfficientInfos []ADEfficientInfo
}
