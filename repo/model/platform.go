package model

type Platform struct {
	Admin
	//登记的司机
	AuthorizedDrivers []*Driver        `gorm:"foreignkey:PlatformID" json:"authorized_drivers"`
	Devices           []*Device        `gorm:"foreignkey:PlatformID" json:"devices"`
	Advertisements    []*Advertisement `gorm:"foreignkey:PlatformID" json:"advertisements"`
}
