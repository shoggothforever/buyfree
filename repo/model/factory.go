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
	//供应的商品
	Products map[string]RepoProduct
}

func (f *Factory) deliver(pro_id uuid.UUID, d *Driver) {

}

//func (r *RepoProduct) register(o design.Observer) {
//	d,err:=o.(*Driver)
//}
//func (r *RepoProduct) deregister(o *design.Observer) {
//	r.Subscribers.
//}
//
//func (r *RepoProduct)notifyAll(){
//
//}
