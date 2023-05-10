package model

import (
	"fmt"
)

type CartService interface {
	Account() float64
	AccoutWithChoose() float64
	AccountWithTicket(...Ticket)
}
type Cart struct {
	CartID      int64   `gorm:"primaryKey" json:"cart_id" form:"cart_id"`
	TotalCount  int64   `gorm:"comment:全选金额" json:"total_count" form:"total_count"`
	TotalAmount float64 `gorm:"comment:全部商品数量" json:"total_amount" form:"total_amount"`
	//存储
	Products []*OrderProduct `gorm:"foreignKey:CartRefer"`
}

// 需要创建OrderProudct表
type PassengerCart struct {
	PassengerID int64 `json:"passenger_id" form:"passenger_id"`
	Cart
}
type DriverCart struct {
	DriverID int64 `json:"driver_id" form:"driver_id"`
	Cart
}

// 计算购物车中商品总价(可以与AllIn合并)
func (c *Cart) Account() float64 {

	c.TotalCount = int64(len(c.Products))
	for k, v := range c.Products {
		fmt.Println(k)
		//c.Products[i].IsChosen=true
		c.TotalAmount += v.GetAmount()
	}
	return c.TotalAmount
}

// 计算购物车中选中的商品总价
func (c *Cart) AccountWithChoose() float64 {
	var sum float64 = 0
	for _, v := range c.Products {
		sum += v.GetChooseAmount()
	}
	return sum
}

func (c *Cart) AllIn() {
	n := len(c.Products)
	for i := 0; i < n; i++ {
		c.Products[i].IsChosen = true
	}
	c.Account()
}

// Account Chosen Products
//func (c *PassengerCart) Account() int64 {
//	var sum int64 = 0
//	for _, v := range c.Products {
//		if v.IsChoose {
//			sum += v.BuyPrize
//		}
//	}
//	return sum
//}
//func (c *DriverCart) Account() int64 {
//	var sum int64 = 0
//	for _, v := range c.Products {
//		if v.IsChoose {
//			sum += v.SupplyPrize
//		}
//	}
//	return sum
//}
