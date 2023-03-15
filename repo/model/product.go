package model

import (
	"github.com/google/uuid"
)

type Product struct {
	ID   int64 `gorm:"type:int;primaryKey;"`
	Name string
	//存货
	inventory int64
	//月销售量
	MonthlySales int64
	//销售价
	BuyPrize float64
	//批发价
	SupplyPrize float64
}
type OrderProduct struct {
	//外键
	CartRefer  uuid.UUID
	OrderRefer uuid.UUID
	IsChosen   bool
	Name       string
	Count      int64
	//添加到购物车时记得写入价格字段
	Prize float64
}

func (o *OrderProduct) GetAmount() float64 {
	price := o.Prize * float64(o.Count)
	return price
}
func (o *OrderProduct) GetChooseAmount() float64 {
	var price float64 = 0
	if o.IsChosen {
		price = o.Prize * float64(o.Count)
	}
	return price
}

type VDProduct struct {
	Product
	//售货机编号
	VDID uuid.UUID
}
type RepoProduct struct {
	Product
	//场站编号
	FactoryID uuid.UUID

	//Subscribers map[uuid.UUID]*Driver
}

//func newProduct(name string, in int64) *Product {
//	return &Product{
//		name,
//		in,
//		0,
//	}
//}

//func (p *Product) register(o design.Observer) {
//
//}
//func (p *Product) deregister(o design.Observer) {
//
//}
//func (p *Product) notifyall(o design.Observer) {
//
//}
