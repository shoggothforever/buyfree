package response

import (
	"buyfree/repo/model"
	"time"
)

// 平台端查看订单信息响应

type DriverInfo struct {
	Name    string `json:"name" form:"name"`
	CarID   string `json:"car_id" form:"car_id"`
	Mobile  string `json:"mobile" form:"mobile"`
	Comment string `json:"comment"`
}
type OrderProductInfo struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Sku   string  `json:"sku"`
	Pic   string  `json:"pic"`
	Count int64   `json:"count"`
	Price float64 `json:"price"`
}
type FactoryOrderInfo struct {
	OrderID          int64 `json:"order_id"`
	DriverInfo       DriverInfo
	OrderProductInfo []OrderProductInfo
	//司机取货时间,只有已完成的订单会有该属性
	GetTime time.Time `json:"get_time"`
	//订单状态 0：待付款，1：待取货，2：已完成（已关闭）
	State int64 `json:"state"`
	//订单付款
	Cost float64 `json:"cost"`
}
type OrderResponse struct {
	Response
	OrderForms []FactoryOrderInfo
}
type SubmitResponse struct {
	Response
	OrderForm model.DriverOrderForm
}
type OrderFormResponse struct {
	Response
}
