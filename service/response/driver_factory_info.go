package response

type FactoryProductOverview struct {
	Pic  string `json:"pic,omitempty"`
	Name string `json:"name,omitempty"`
}

type FactoryInfo struct {
	//场站距离(调用API）
	Distance string `json:"distance,omitempty"`
	////场站经纬度
	//Lontitude float64 `json:"lontitude"`
	//Latitude  float64 `json:"latitude"`
	//场站名
	FactoryName  string                   `json:"factory_name" json:"factory_name,omitempty"`
	ProductViews []FactoryProductOverview `json:"product_views,omitempty"`
}
type FactoryInfoResponse struct {
	Response
	FactoryInfos []FactoryInfo `json:"factory_infos"`
}
type FactoryDistanceInfos []*FactoryDistanceInfo
type FactoryDistanceInfo struct {
	FactoryName string `json:"factory_name"`
	FactoryID   int64  `json:"factory_id"`
	Distance    string `json:"distance"`
}
type FactoryDetail struct {
	ID          int64  `json:"id"`
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
}
type FactoryProductDetail struct {
	Name         string  `json:"name,omitempty"`
	Pic          string  `json:"pic"`
	Type         string  `json:"type,omitempty"`
	Inventory    int64   `json:"inventory,omitempty"`
	MInventory   int64   `json:"m_inventory,omitempty"`
	MonthlySales int64   `json:"monthly_sales,omitempty"`
	SupplyPrice  float64 `json:"supply_price,omitempty"`
}
type FactoryDetailResponse struct {
	Response
	DistanceInfo   FactoryDistanceInfo     `json:"distance_info"`
	FactoryDetail  FactoryDetail           `json:"factory_detail"`
	ProductDetails []*FactoryProductDetail `json:"product_details,omitempty"`
}
