package model

import (
	"github.com/google/uuid"
)

type CartService interface {
	Account() float64
	AccoutWithChoose() float64
	AccountWithTicket(...Ticket)
}
type Cart struct {
	CartID      uuid.UUID `gorm:"primaryKey"`
	TotalCount  int64
	TotalAmount float64
	//存储
	Products []*OrderProduct `gorm:"foreignKey:CartRefer"`
}

//需要创建OrderProudct表
type PassengerCart struct {
	//外键 Passenger.Uuid
	PassengerID uuid.UUID
	Cart
}
type DriverCart struct {
	//外键 Driver.Uuid
	DriverID    uuid.UUID
	FactoryName string
	//距离场站距离
	Distance int64
	Cart
}

func (c *Cart) Account() float64 {
	c.TotalCount = int64(len(c.Products))
	for _, v := range c.Products {
		c.TotalAmount += v.GetAmount()
	}
	return c.TotalAmount
}
func (c *Cart) AccountWithChoose() float64 {
	var sum float64 = 0
	for _, v := range c.Products {
		sum += v.GetChooseAmount()
	}
	return sum
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
