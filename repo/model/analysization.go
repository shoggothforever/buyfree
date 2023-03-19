package model

//广告管理
type AanlyInfo struct {
	DailySales    float64 `gorm:"comment:日销量"`
	DailyAdAna    int64   `gorm:"comment:日活跃度"`
	WeeklySales   float64 `gorm:"comment:周销量"`
	WeeklyAd      int64   `gorm:"comment:周活"`
	MonthlySales  float64 `gorm:"comment:月销量"`
	MonthlyAd     int64   `gorm:"comment:月活"`
	AnnuallySales float64 `gorm:"comment:年销售量"`
	AnnualAd      int64   `gorm:"comment:年活"`
}
