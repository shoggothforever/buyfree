package model

type Product struct {
	ID int64 `gorm:"primaryKey;" json:"id"`
	//存货
	Inventory int64 `gorm:"comment:存货" json:"inventory"`
	//月销售量
	MonthlySales int64 `gorm:"comment:月销" json:"monthly_sales"`

	FactoryRefer string `gorm:"comment:指向场站的编号" json:"factory_refer"`
	//库存单位
	Sku string `gorm:"comment:库存控制最小可用单位" json:"sku"`
	//产品名称
	Name string `gorm:"comment:产品名称" json:"name"`
	//图片
	Pic string `gorm:"comment:图片" json:"pic"`
	//型号
	Type string `gorm:"comment:型号" json:"type"`
	//销售价
	BuyPrize float64 `gorm:"comment:销售价" json:"buy_prize"`
	//批发价
	SupplyPrize float64 `gorm:"comment:批发价" json:"supply_prize"`
}
type DeviceProduct struct {
	Product
	//售货机编号
	DeviceID int64 `gorm:"comment:售货机编号" json:"device_id"`
}
type RepoProduct struct {
	Product
	//场站名字
	FactoryName string `json:"factory_name"`
}
type OrderProduct struct {
	//外键
	CartRefer    int64   `gorm:"comment:所属购物车" json:"cart_refer"`
	FactoryRefer int64   `gorm:"comment:所属场站" json:"factory_refer"`
	OrderRefer   string  `gorm:"comment:所属订单" json:"order_refer"`
	IsChosen     bool    `gorm:"comment:场站是否上线该产品 1-上线 0-下线" json:"is_chosen"`
	Name         string  `gorm:"comment:商品名称" json:"name"`
	Pic          string  `gorm:"comment:图片" json:"pic"`
	Type         string  `gorm:"comment:商品型号" json:"type"`
	Count        int64   `gorm:"comment:需求量" json:"count"`
	Prize        float64 `gorm:"comment:价格,根据所属购物车种类赋予不同类型的价格，用户购物车内该值为零售价,车主购物车内该值为批发价" json:"prize"`
}
type OrderProductImplement interface {
	Refund()
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
