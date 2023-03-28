package response

import "buyfree/repo/model"

type HomeStatic struct {
	DailySales      float64 `json:"daily_sales"`
	ADList          []model.Advertisement
	ProductRankList []model.ProductRank
}
type HomePageResponse struct {
	Response
	HomeStatic
}
