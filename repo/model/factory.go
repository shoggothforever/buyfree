package model

import "github.com/google/uuid"

//type Subject interface {
//	register(o *design.Observer)
//	deregister(o *design.Observer)
//	notifyAll()
//}
//

type Factory struct {
	User
	//场站地理位置
	Address string `gorm:"comment:场站位置信息" json:"address" form:"address"`
	//供应的商品
	Products   []*FactoryProduct  `gorm:"foreignkey:FactoryID" `
	OrderForms []*DriverOrderForm `gorm:"foreignkey:FactoryID"`
}

func (f *Factory) deliver(pro_id uuid.UUID, d *Driver) {

}

//func (r *FactoryProduct) register(o design.Observer) {
//	d,err:=o.(*Driver)
//}
//func (r *FactoryProduct) deregister(o *design.Observer) {
//	r.Subscribers.
//}
//
//func (r *FactoryProduct)notifyAll(){
//
//}
