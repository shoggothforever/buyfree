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
	Cost int `gorm:"comment:花费"`
	//订单状态
	State ORDERSTATE `gorm:"type:smallint;comment:订单状态 2-已完成 1-待取货 0-未支付"`
	//支付时存储位置(购物时获取车主位置）
	Location string `gorm:"comment:支付时存储位置(购物时获取车主位置）"`
	//支付时存储车主车牌号
	DriverCarID string `gorm:"comment:支付时存储车主车牌号"`
	//下单时间
	Placetime time.Time `gorm:"comment:下单时间"`
	//支付时间
	Paytime time.Time `gorm:"comment:支付时间"`
	//商品信息
	ProductInfo []*OrderProduct `gorm:"foreignKey:OrderRefer"`
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
	CarID string `gorm:"foreignKey:PassengerID"`
	//备注
	Comment string `gorm:"comment:备注"`
	//自取时间
	GetTime time.Time `gorm:"comment:自取时间"`
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
