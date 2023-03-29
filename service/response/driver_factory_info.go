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
	FactoryName  string                   `json:"factory_name" json:"factoryName,omitempty"`
	ProductViews []FactoryProductOverview `json:"productViews,omitempty"`
}

type FactoryInfoResponse struct {
	Response
	FactoryInfos []FactoryInfo
}
