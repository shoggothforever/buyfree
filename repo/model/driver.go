package model

import "github.com/google/uuid"

type DriverInfo struct {
	CarID       string
	PhoneNumber string
	User
}

//Observer Driver
type Driver struct {
	Location string
	DriverInfo
	mVdMachine []*VdMachine `gorm:"foreignKey:OwnerID"`
	//购物车信息
	Cart DriverCart `gorm:"foreignkey:DriverKey"`
	//购物订单
	DriverOrderForms DriverOrderForm `gorm:"foreignkey:DriverKey"`
}
type Replenish struct {
	FactorID  uuid.UUID
	ProductID uuid.UUID
	nums      int64
}

//TODO
func (d *Driver) replenishment(r []*Replenish) error {
	return nil
}

//func (d *Driver) GetID() uuid.UUID {
//	return d.Uuid
//}

type SubscibeResponse struct {
}
