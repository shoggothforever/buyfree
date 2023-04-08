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

// abstract
type OrderForm struct {
	//订单编码
	OrderID int64 `gorm:"primarykey" json:"order_id"`
	//花费
	Cost float64 `gorm:"comment:花费" json:"cost"`
	//订单状态
	State ORDERSTATE `gorm:"type:smallint;comment:订单状态 2-已完成 1-待取货 0-未支付" json:"state"`
	//支付时存储位置(购物时获取车主位置）
	Location string `gorm:"comment:支付时存储位置(购物时获取车主位置）" json:"location"`
	//下单时间（创建时更新即可）
	PlaceTime time.Time `gorm:"comment:下单时间" json:"place_time"`
	//支付时间（更改操作先于发送订单请求，支付时更新即可）
	PayTime time.Time `gorm:"comment:支付时间" json:"pay_time"`
	//商品信息
	ProductInfos []*OrderProduct `gorm:"foreignKey:OrderRefer"`
}

// 需要关联创表
type PassengerOrderForm struct {
	//Passenger外键
	PassengerID int64 `json:"passenger_id"`
	//支付时存储车主车牌号
	DriverCarID string `gorm:"comment:支付时存储车主车牌号" json:"car_id"`
	OrderForm
}

func (p *PassengerOrderForm) Set() {
	p.PassengerID = p.PassengerID
}

// 需要关联创表,其实是场站存放的司机发起的订单
type DriverOrderForm struct {
	FactoryID   int64  `gorm:"comment:指向factory.id" json:"factory_id"`
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

func (d *DriverOrderForm) Set(fid, did int64, cost float64, fname, carid, loc string) {
	d.FactoryID = fid
	d.DriverID = did
	d.Cost = cost
	d.FactoryName = fname
	d.CarID = carid
	// d.comment =comment string
	// d.GetTime = gtime time.time
	d.PlaceTime = time.Now()
	d.Location = loc
	d.State = 0

}

type ReplenInfo struct {
	DriverID        string
	AllCount        int64
	WaitCount       int64
	FinishCount     int64
	WaitOrderForm   []*DriverOrderForm
	FinishOrderForm []*DriverOrderForm
}

//工作站
