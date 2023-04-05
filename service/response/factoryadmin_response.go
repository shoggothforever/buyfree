package response

import "buyfree/repo/model"

type FactoryRegisterResponse struct {
	Response
	FactoryID   int64  `json:"user_id"`
	FactoryName string `json:"factory_name"`
	//地址
	Address string `json:"address"`
	//经度
	Longitude string `json:"longitude" binding:"required"`
	//纬度
	Latitude string `json:"latitude" binding:"required"`
	//描述（选填）
	Description string `json:"description,omitempty"`
}

type FactoryProductsModifyResponse struct {
	Response
	Products []model.FactoryProduct `json:"products"`
}
