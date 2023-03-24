package response

import "buyfree/repo/model"

type FactoryProductsInfo struct {
	FactoryName string  `gorm:"factory_name" json:"factory_name"`
	Sku         string  `gorm:"sku" json:"sku"`
	Name        string  `gorm:"name" json:"name"`
	Type        string  `gorm:"type" json:"type"`
	IsOnShelf   bool    `gorm:"is_on_shelf" json:"is_on_shelf"`
	Pic         string  `gorm:"pic" json:"pic"`
	TotalSales  float64 `gorm:"total_sales" json:"total_sales"`
	Inventory   int64   `gorm:"inventory" json:"inventory"`
}
type FactoryProductsResponse struct {
	Response
	Products []FactoryProductsInfo
}
type FactoryGoodsResponse struct {
	Response
	model.FactoryProduct
}
