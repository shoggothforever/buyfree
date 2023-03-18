package model

type Platform struct {
	Admin
	//登记的司机
	AuthorizedDrivers []*Driver        `gorm:"foreignkey:PlatformID"`
	Devices           []*DEVICE        `gorm:"foreignkey:PlatformID"`
	Advertisements    []*Advertisement `gorm:"foreignkey:PlatformID"`
}
