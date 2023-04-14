package response

import "buyfree/repo/model"

type ReplenishInfo struct {
	FactoryID   int64   `json:"factory_id"`
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
	DistanceInfo   FactoryDistanceReq   `json:"distance_info"`
	ProductDetails []*CartProductDetail `json:"product_details"`
}
type CartProductDetail struct {
	Name     string  `json:"name,omitempty"`
	Pic      string  `json:"pic"`
	Type     string  `json:"type,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Count    int64   `json:"count"`
	IsChosen bool    `json:"is_chosen"`
}
type CartResponse struct {
	Response
	//购物车所有价格
	TotalAmount float64 `json:"total_amount"`
	//购物车所有商品数量
	TotalCount int64        `json:"total_count"`
	Groups     []*CartGroup `json:"groups"`
}
