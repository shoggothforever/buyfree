package model

//广告管理
//想要存redis里的
type SalesData struct {
	DailySales    float64 `gorm:"comment:日销量" json:"daily_sales"`
	WeeklySales   float64 `gorm:"comment:周销量" json:"weekly_sales"`
	MonthlySales  float64 `gorm:"comment:月销量" json:"monthly_sales"`
	AnnuallySales float64 `gorm:"comment:年销售量" json:"annually_sales"`
	TotalSales    float64 `gorm:"comment:总销售量" json:"total_sales"`
}
type Vitality struct {
	DailyAdAna int64 `gorm:"comment:日活跃度" json:"daily_ad_ana"`
	WeeklyAd   int64 `gorm:"comment:周活" json:"weekly_ad"`
	MonthlyAd  int64 `gorm:"comment:月活" json:"monthly_ad"`
	AnnualAd   int64 `gorm:"comment:年活" json:"annual_ad"`
}
type AanlyInfo struct {
	SalesData
	Vitality
}
