package model

type DriverInfo struct {
	User
	CarID      string `gorm:"comment:车牌号"`
	Mobile     string `gorm:"comment:手机号"`
	IDCard     string `gorm:"comment:身份证"`
	PlatformID int64
	IsAuth     bool `gorm:"comment:1为已认证，0为未认证"`
}

//Observer Driver
type Driver struct {
	DriverInfo
	Location string `gorm:"comment:地理位置"`

	Devices []*Device `gorm:"foreignKey:OwnerID;comment:持有售货机"`
	//购物车信息
	Cart *DriverCart `gorm:"foreignKey:DriverID;comment:补货购物车"`
	//购物订单
	DriverOrderForms *DriverOrderForm `gorm:"foreignKey:DriverID;补货订单"`
}
type Replenish struct {
	FactorID int64
	//TODO 携带补货需求信息
	ProductID string
	nums      int64
}

// lng经度 lat纬度
func (d *Driver) GetLocation(lng, lat string) (Location string) {
	//TODO 调用Api
	//Location=Api(lng,lat)
	return Location
}

//TODO	发送补货请求
func (d *Driver) replenishment(r []*Replenish) error {
	return nil
}

//func (d *Driver) GetID() uuid.UUID {
//	return d.Uuid
//}

type SubscibeResponse struct {
}
