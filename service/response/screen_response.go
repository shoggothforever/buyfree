package response

import "buyfree/repo/model"

type ScreenInfo struct {
	DevNums        int64 `json:"dev_nums"`
	OnlineDevNums  int64 `json:"online_dev_nums"`
	OfflineDevNums int64 `json:"offline_dev_nums"`
	//营销额七日增长曲线
	SalesCurve [7]float64 `json:"sales"`
	model.SalesData
	ADList          [10]*model.Advertisement
	ProductRankList [10]*FactoryProductsInfo
}

type ScreenInfoResponse struct {
	Response
	ScreenInfo
}

type SaleStaticResponse struct {
	Response
	model.SalesData
	ProductsRank [10]model.ProductRank
}
