package model

//广告管理
//想要存redis里的
type SalesData struct {
	DailySales    string `gorm:"comment:日销量" json:"daily_sales" form:"daily_sales"`
	WeeklySales   string `gorm:"comment:周销量" json:"weekly_sales" form:"weekly_sales"`
	MonthlySales  string `gorm:"comment:月销量" json:"monthly_sales" form:"monthly_sales"`
	AnnuallySales string `gorm:"comment:年销售量" json:"annually_sales" form:"annually_sales"`
	TotalSales    string `gorm:"comment:总销售量" json:"total_sales" form:"total_sales"`
}
type Vitality struct {
	DailyAdAna int64 `gorm:"comment:日活跃度" json:"daily_ad_ana" form:"daily_ad_ana"`
	WeeklyAd   int64 `gorm:"comment:周活" json:"weekly_ad" form:"weekly_ad"`
	MonthlyAd  int64 `gorm:"comment:月活" json:"monthly_ad" form:"monthly_ad"`
	AnnualAd   int64 `gorm:"comment:年活" json:"annual_ad" form:"annual_ad"`
}
type AanlyInfo struct {
	SalesData
	Vitality
}
