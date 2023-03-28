package response

import "buyfree/repo/model"

type FactoryInfo struct {
	Distance float64 `json:"distance"`
	model.Factory
	FactoryProducts []model.FactoryProduct
}

type FactoryInfoResponse struct {
	Response
	FactoryInfos []FactoryInfo
}
