package response

import "buyfree/repo/model"

type HomeStatic struct {
	DailySales      float64 `json:"daily_sales,omitempty"`
	DailyRatio      float64 `json:"daily_ratio,omitempty"`
	WeeklyRatio     float64 `json:"weekly_ratio,omitempty"`
	MonthlySales    float64 `json:"monthly_sales,omitempty"`
	ADDailySales    float64 `json:"ad_daily_sales,omitempty"`
	ADPlayTimes     int64   `json:"ad_play_times,omitempty"`
	ProductRankList []*model.DeviceProduct
}
type HomePageResponse struct {
	Response
	HomeStatic
}
