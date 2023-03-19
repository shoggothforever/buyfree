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
	CartID      int64   `gorm:"primaryKey"`
	TotalCount  int64   `gorm:"comment:全选金额"`
	TotalAmount float64 `gorm:"comment:全部商品数量"`
	//存储
	Products []*OrderProduct `gorm:"foreignKey:CartRefer"`
}

//需要创建OrderProudct表
type PassengerCart struct {
	//外键 Passenger.Uuid
	PassengerID int64
	Cart
}
type DriverCart struct {
	//外键 Driver.Uuid
	DriverID string

	FactoryName string `gorm:"comment:购物场站名称"`
	//距离场站距离
	Distance int64 `gorm:"comment:距离场站距离"`
	Cart
}

//计算购物车中商品总价(可以与AllIn合并)
func (c *Cart) Account() float64 {

	c.TotalCount = int64(len(c.Products))
	for k, v := range c.Products {
		fmt.Println(k)
		//c.Products[i].IsChosen=true
		c.TotalAmount += v.GetAmount()
	}
	return c.TotalAmount
}

//计算购物车中选中的商品总价
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
