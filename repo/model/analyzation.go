package model

//广告管理
//想要存redis里的
type SalesData struct {
	DailySales    float64 `gorm:"comment:日销量"`
	WeeklySales   float64 `gorm:"comment:周销量"`
	MonthlySales  float64 `gorm:"comment:月销量"`
	AnnuallySales float64 `gorm:"comment:年销售量"`
	TotalSales    float64 `gorm:"comment:总销售量"`
}
type Vitality struct {
	DailyAdAna int64 `gorm:"comment:日活跃度"`
	WeeklyAd   int64 `gorm:"comment:周活"`
	MonthlyAd  int64 `gorm:"comment:月活"`
	AnnualAd   int64 `gorm:"comment:年活"`
}
type AanlyInfo struct {
	SalesData
	Vitality
}
