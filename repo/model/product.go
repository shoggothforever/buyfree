package model

type Product struct {
	ID        int64  `gorm:"primaryKey;" json:"id" form:"id"`
	FactoryID int64  `gorm:"comment:指向场站的编号" json:"factory_id" form:"factory_id"`
	Sku       string `gorm:"notnull;comment:库存控制最小可用单位" json:"sku" form:"sku"`
	//存货
	Inventory int64 `gorm:"notnull;comment:存货" json:"inventory" form:"inventory"`
	//产品名称
	Name string `gorm:"notnull;comment:产品名称" json:"name" form:"name"`
	//图片
	Pic string `gorm:"comment:图片" json:"pic" form:"pic"`
	//型号
	Type string `gorm:"notnull;comment:型号" json:"type" form:"type"`
	//销售价
	BuyPrice float64 `gorm:"notnull;comment:销售价" json:"buy_price" form:"buy_price"`
	//批发价
	SupplyPrice float64 `gorm:"notnull;comment:批发价" json:"supply_price" form:"supply_price"`
	SalesData
}
type DeviceProduct struct {
	//售货机编号
	DeviceID int64 `gorm:"comment:售货机编号" json:"device_id" form:"device_id"`
	//
	DriverID int64 `gorm:"车主id" json:"driver_id" form:"driver_id"`
	Product
}

func (d *DeviceProduct) Set(id, devid, drid, fid, inv int64, bprice, sprice float64, name, ty, sku, pic string) {
	d.ID = id
	d.DeviceID = devid
	d.DriverID = drid
	d.FactoryID = fid
	d.Inventory = inv
	d.BuyPrice = bprice
	d.SupplyPrice = sprice
	d.Name = name
	d.Type = ty
	d.Sku = sku
	d.Pic = pic
	d.SalesData = SalesData{"0", "0", "0", "0", "0"}
}

type DeviceProductPartInfo struct {
	Name string `json:"name" form:"name"`
	//图片
	Pic string `json:"pic" form:"pic"`
	//型号
	Type string `json:"type" form:"type"`
	//库存
	Inventory int64 `json:"inventory"`
	//销售价
	BuyPrice float64 `json:"buy_price" form:"buy_price"`
}

type FactoryProducts []FactoryProduct
type FactoryProduct struct {
	//场站名字
	FactoryName string `json:"factory_name" form:"factory_name"`
	Product
	//上架状态
	IsOnShelf bool `json:"is_on_shelf" form:"is_on_shelf"`
}

func (f *FactoryProduct) Set(id, fid int64, fname string, v *FactoryProduct) {
	f.FactoryName = fname
	f.Name = v.Name
	f.Pic = v.Pic
	f.Type = v.Type
	f.Sku = v.Sku
	f.FactoryID = fid
	f.Inventory = v.Inventory
	f.ID = id
	f.BuyPrice = v.BuyPrice
	f.SupplyPrice = v.SupplyPrice
	f.IsOnShelf = true
}

type OrderProducts []*OrderProduct

// 订单中的商品信息
type OrderProduct struct {
	//外键
	CartRefer  int64   `gorm:"comment:所属购物车;" json:"cart_refer" form:"cart_refer"`
	FactoryID  int64   `gorm:"comment:所属场站" json:"factory_id" form:"factory_id"`
	OrderRefer int64   `gorm:"comment:所属订单" json:"order_refer" form:"order_refer"`
	IsChosen   bool    `gorm:"comment:场站是否上线该产品 1-上线 0-下线" json:"is_chosen" form:"is_chosen"`
	Name       string  `gorm:"notnull;comment:商品名称" json:"name" form:"name"`
	Sku        string  `gorm:"notnull;comment:库存控制最小可用单位" json:"sku" form:"sku"`
	Pic        string  `gorm:"comment:图片" json:"pic" form:"pic"`
	Type       string  `gorm:"notnull;comment:商品型号" json:"type" form:"type"`
	Count      int64   `gorm:"notnull;comment:需求量" json:"count" form:"count"`
	Price      float64 `gorm:"notnull;comment:价格,根据所属购物车种类赋予不同类型的价格，用户购物车内该值为零售价,车主购物车内该值为批发价" json:"price" form:"price"`
}

func (o *OrderProduct) Set(ore, cr, fr, cnt int64, pr float64, bi bool, name, sku, pic, Type string) {
	o.OrderRefer = ore
	o.CartRefer = cr
	o.FactoryID = fr
	o.Count = cnt
	o.Price = pr
	o.IsChosen = bi
	o.Name = name
	o.Sku = sku
	o.Pic = pic
	o.Type = Type
}

type CartProduct struct {
	//外键
	CartRefer  int64   `gorm:"comment:所属购物车;" json:"cart_refer" form:"cart_refer"`
	FactoryID  int64   `gorm:"comment:所属场站" json:"factory_id" form:"factory_id"`
	OrderRefer int64   `gorm:"comment:所属订单" json:"order_refer" form:"order_refer"`
	IsChosen   bool    `gorm:"comment:场站是否上线该产品 1-上线 0-下线" json:"is_chosen" form:"is_chosen"`
	Name       string  `gorm:"notnull;comment:商品名称" json:"name" form:"name"`
	Sku        string  `gorm:"notnull;comment:库存控制最小可用单位" json:"sku" form:"sku"`
	Pic        string  `gorm:"comment:图片" json:"pic" form:"pic"`
	Type       string  `gorm:"notnull;comment:商品型号" json:"type" form:"type"`
	Count      int64   `gorm:"notnull;comment:需求量" json:"count" form:"count"`
	Price      float64 `gorm:"notnull;comment:价格,根据所属购物车种类赋予不同类型的价格，用户购物车内该值为零售价,车主购物车内该值为批发价" json:"price" form:"price"`
}

func (o *CartProduct) Set(ore, cr, fr, cnt int64, pr float64, bi bool, name, sku, pic, Type string) {
	o.OrderRefer = ore
	o.CartRefer = cr
	o.FactoryID = fr
	o.Count = cnt
	o.Price = pr
	o.IsChosen = bi
	o.Name = name
	o.Sku = sku
	o.Pic = pic
	o.Type = Type
}
func (c CartProduct) Trans(o OrderProduct) {
	c.OrderRefer = o.OrderRefer
	c.CartRefer = o.CartRefer
	c.FactoryID = o.FactoryID
	c.Count = o.Count
	c.Price = o.Price
	c.IsChosen = o.IsChosen
	c.Name = o.Name
	c.Sku = o.Sku
	c.Pic = o.Pic
	c.Type = o.Type
}
func (o *OrderProduct) GetAmount() float64 {
	price := o.Price * float64(o.Count)
	return price
}
func (o *OrderProduct) GetChooseAmount() float64 {
	var price float64 = 0
	if o.IsChosen {
		price = o.Price * float64(o.Count)
	}
	return price
}

type ProductRank struct {
	Score  float64     `json:"score" form:"score"`
	Member interface{} `json:"member" form:"member"`
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
