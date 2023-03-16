package model

import (
	"github.com/google/uuid"
	"time"
)

type ORDERSTATE int

const (
	CANCLE ORDERSTATE = iota
	WAIT
	DONE
)

//abstract
type OrderForm struct {
	//订单编码
	OrderID uuid.UUID `gorm:"primarykey"`
	//花费
	Cost int
	//订单状态
	State ORDERSTATE `gorm:"type:smallint"`
	//支付时存储位置
	Location string
	//支付时存储车主车牌号
	DriverCarID string
	//商品信息
	ProductInfo []*OrderProduct `gorm:"foreignKey:OrderRefer"`
	//下单时间
	Placetime time.Time
	//支付时间
	Paytime time.Time
}

//需要关联创表
type PassengerOrderForm struct {
	//Passenger外键
	PassengerID uuid.UUID
	OrderForm
}

//需要关联创表
type DriverOrderForm struct {
	//Driver外键
	DriverID uuid.UUID
	//车牌号
	CarID string
	//备注
	Comment string
	//自取时间
	GetTime time.Time
	OrderForm
}

type ReplenInfo struct {
	DriverID        uuid.UUID
	AllCount        int
	WaitCount       int
	FinishCount     int
	WaitOrderForm   []*DriverOrderForm
	FinishOrderForm []*DriverOrderForm
}
