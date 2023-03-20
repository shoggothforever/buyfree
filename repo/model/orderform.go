package model

import (
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
	OrderID string `gorm:"primarykey" json:"order_id"`
	//花费
	Cost int64 `gorm:"comment:花费" json:"cost"`
	//订单状态
	State ORDERSTATE `gorm:"type:smallint;comment:订单状态 2-已完成 1-待取货 0-未支付" json:"state"`
	//支付时存储位置(购物时获取车主位置）
	Location string `gorm:"comment:支付时存储位置(购物时获取车主位置）" json:"location"`

	//下单时间
	Placetime time.Time `gorm:"comment:下单时间" json:"place_time"`
	//支付时间
	Paytime time.Time `gorm:"comment:支付时间" json:"pay_time"`
	//商品信息
	ProductInfo []*OrderProduct `gorm:"foreignKey:OrderRefer"`
}

//需要关联创表
type PassengerOrderForm struct {
	//Passenger外键
	PassengerID int64 `json:"passenger_id"`
	//支付时存储车主车牌号
	DriverCarID string `gorm:"comment:支付时存储车主车牌号" json:"car_id"`
	OrderForm
}

//需要关联创表
type DriverOrderForm struct {
	FactoryRefer string `gorm:"comment:指向factory.id" json:"factory_id"`

	FactoryName string `gorm:"comment:订单发货场站名" json:"factory_name"`
	//Driver外键
	DriverID int64 `json:"driver_id"`
	//车牌号
	CarID string `json:"car_id"`
	//备注
	Comment string `gorm:"comment:备注" json:"comment"`
	//自取时间
	GetTime time.Time `gorm:"comment:自取时间" json:"get_time"`
	OrderForm
}

type ReplenInfo struct {
	DriverID        string
	AllCount        int64
	WaitCount       int64
	FinishCount     int64
	WaitOrderForm   []*DriverOrderForm
	FinishOrderForm []*DriverOrderForm
}
