package response

import (
	"buyfree/repo/model"
)

type FactoryProductsInfo struct {
	FactoryName string `gorm:"factory_name" json:"factory_name"`
	Sku         string `gorm:"sku" json:"sku"`
	//商品名称
	Name       string `gorm:"name" json:"name"`
	Type       string `gorm:"type" json:"type"`
	IsOnShelf  bool   `gorm:"is_on_shelf" json:"is_on_shelf"`
	Pic        string `gorm:"pic" json:"pic"`
	TotalSales string `gorm:"total_sales" json:"total_sales"`
	Inventory  int64  `gorm:"inventory" json:"inventory"`
}
type FactoryProductsResponse struct {
	Response
	Products []FactoryProductsInfo
}
type UnionNameInfo struct {
	FactoryName string `json:"factory_name,omitempty"`
	ProductName string `json:"product_name,omitempty"`
}
type FactoryGoodsResponse struct {
	Response
	model.FactoryProduct
}
