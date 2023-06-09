package model

type DriverInfo struct {
	User
	CarID      string `gorm:"comment:车牌号" json:"car_id" form:"car_id"`
	Mobile     string `gorm:"comment:手机号" json:"mobile" form:"mobile"`
	IDCard     string `gorm:"comment:身份证" json:"id_card" form:"id_card"`
	PlatformID int64  `json:"platform_id" form:"platform_id"`
	IsAuth     bool   `gorm:"comment:1为已认证，0为未认证" json:"is_auth" form:"is_auth"`
}

// Observer Driver
type Driver struct {
	DriverInfo
	Geo
	Devices []*Device `gorm:"foreignKey:OwnerID;comment:持有售货机"`
	//购物车信息
	Cart *DriverCart `gorm:"foreignKey:DriverID;comment:补货购物车" `
	//购物订单
	DriverOrderForms *[]DriverOrderForm `gorm:"foreignKey:DriverID;补货订单" `
}
type LocationInfo struct {
	Name string `json:"name"`
	Geo  `json:"geo"`
}

// lng经度 lat纬度
func (d *Driver) GetLocation(lng, lat string) (Location string) {
	//TODO 调用Api
	//Location=Api(lng,lat)
	return Location
}
