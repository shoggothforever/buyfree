package response

import "buyfree/repo/model"

type ScreenInfo struct {
	DevNums        int64 `json:"dev_nums"`
	OnlineDevNums  int64 `json:"online_dev_nums"`
	OfflineDevNums int64 `json:"offline_dev_nums"`
	//营销额七日增长曲线
	SalesCurve [7]float64 `json:"sales"`
	model.SalesData
	ADList          []*model.Advertisement
	ProductRankList [5]*FactoryOrderInfo
}

type ScreenInfoResponse struct {
	Response
	ScreenInfo
}

type SaleStaticResponse struct {
	model.SalesData
	ProductsRank [5]model.FactoryProduct
	DevicesRank  [5]model.Device
}
