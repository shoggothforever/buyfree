package model

type Product struct {
	ID int64 `gorm:"primaryKey;"`
	//存货
	inventory int64 `gorm:"comment:存货"`
	//月销售量
	MonthlySales int64 `gorm:"comment:月销"`
	FactoryRefer string
	//库存单位
	Sku string `gorm:"comment:库存控制最小可用单位"`
	//产品名称
	Name string `gorm:"comment:产品名称"`
	//型号
	Type string `gorm:"comment:型号"`
	//销售价
	BuyPrize float64 `gorm:"comment:销售价"`
	//批发价
	SupplyPrize float64 `gorm:"comment:批发价"`
}
type OrderProduct struct {
	//外键
	CartRefer    int64   `gorm:"comment:所属购物车"`
	FactoryRefer int64   `gorm:"comment:所属场站"`
	OrderRefer   string  `gorm:"comment:所属订单"`
	IsChosen     bool    `gorm:"comment:场站是否上线该产品 1-上线 0-下线"`
	Name         string  `gorm:"comment:商品名称"`
	Type         string  `gorm:"comment:商品型号"`
	Count        int64   `gorm:"comment:需求量"`
	Prize        float64 `gorm:"comment:价格,根据所属购物车种类赋予不同类型的价格，用户购物车内该值为零售价,车主购物车内该值为批发价"`
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

type DeviceProduct struct {
	Product
	//售货机编号
	DeviceID int64 `gorm:"comment:售货机编号"`
}
type RepoProduct struct {
	Product
	//场站编号
	FactoryID int64

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
