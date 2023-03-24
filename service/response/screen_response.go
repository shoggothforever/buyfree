package response

import "buyfree/repo/model"

type ScreenInfo struct {
	DevNums        int64 `json:"dev_nums"`
	OnlineDevNums  int64 `json:"online_dev_nums"`
	OfflineDevNums int64 `json:"offline_dev_nums"`
	//营销额七日增长曲线
	SalesCurve [7]int64 `json:"sales"`
	model.SalesData
	ADList          []model.Advertisement
	ProductRankList []model.ProductRank
}

type ScreenInfoResponse struct {
	Response
	ScreenInfo
}

type SaleStaticResponse struct {
	Response
	model.SalesData
	ProductsRank []model.ProductRank
}
