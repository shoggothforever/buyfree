package response

import "buyfree/repo/model"

type ReplenishInfo struct {
	FactorID    int64   `json:"factor_id"`
	FactoryName string  `json:"factory_name"`
	ProductName string  `json:"product_name"`
	Type        string  `json:"type"`
	Pic         string  `json:"pic"`
	Price       float64 `json:"price"`
	//传入商品件数
	Count int64 `json:"count"`
}
type ModifyResponse struct {
	Response
	FactorID int64

	ReplenishInfo
}
type InventoryResponse struct {
	Response
	Products []model.DeviceProduct
}
type CartGroup struct {
	DistanceInfo   FactoryDistanceInfo  `json:"distance_info"`
	ProductDetails []*CartProductDetail `json:"product_details"`
}
type CartProductDetail struct {
	Name  string  `json:"name,omitempty"`
	Pic   string  `json:"pic"`
	Type  string  `json:"type,omitempty"`
	Price float64 `json:"price,omitempty"`
	Count int64   `json:"count"`
}
type CartResponse struct {
	Response
	Groups []*CartGroup `json:"groups"`
}
