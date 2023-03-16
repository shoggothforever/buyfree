package model

import "github.com/google/uuid"

type DriverInfo struct {
	CarID  string
	Mobile string
	User
}

//Observer Driver
type Driver struct {
	Location string
	DriverInfo
	Devices []*DEVICE `gorm:"foreignKey:OwnerID"`
	//购物车信息
	Cart *DriverCart `gorm:"foreignKey:DriverID"`
	//购物订单
	DriverOrderForms *DriverOrderForm `gorm:"foreignKey:DriverID"`
}
type Replenish struct {
	FactorID uuid.UUID
	//TODO 携带补货需求信息
	ProductID uuid.UUID
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
