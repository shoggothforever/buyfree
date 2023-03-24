package utils

import (
	"buyfree/repo/model"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const MAXBUFFER int = 8192

type DealWithOrderForm func(o *model.OrderFormRequest)

var sem = make(chan bool, MAXBUFFER)
var TimeOut = time.Duration(500 * time.Millisecond)

/*
设计订单接口，使用装饰模式鉴别是购货订单还是退货订单
*/
//购物通道|补货通道
//var Orderchannel chan *model.OrderFormRequest = make(chan *model.OrderFormRequest)
//退款通道
//var Refundchannel chan *model.OrderFormRequest = make(chan *model.OrderFormRequest)
var Lock sync.Locker
var rect atomic.Value

//对redis数据库进行操作,考虑退款操作
func HandleOrderForm(p *model.OrderFormRequest) {
	//c := context.TODO()
	//rc := dal.Getrd()
	//y, m, d := GetDateKey()
	rect.Store(p)
	p.ReplySign <- true
	<-sem
}

func DetectSaleSign(work DealWithOrderForm, OrderChannel chan *model.OrderFormRequest, ofsquit chan bool) {
	var ticker = time.NewTicker(time.Duration(TimeOut))
	defer ticker.Stop()
	for {
		sem <- true
		select {
		case orderformrequest := <-OrderChannel:
			go work(orderformrequest) //购物
		case <-ofsquit:
			return
		case <-ticker.C:
			ofsquit <- true
		default:

		}
	}
}
func StartServer(work DealWithOrderForm) (service chan *model.OrderFormRequest, quit chan bool) {
	service = make(chan *model.OrderFormRequest, 512)
	quit = make(chan bool)
	go DetectSaleSign(work, service, quit)
	return service, quit
}
func ListenOrderFromRequest() {
	//TODO:迁移到main.go中去
	ordqueue, ofsquit := StartServer(HandleOrderForm)
	defer close(ordqueue)
	defer close(ofsquit)
	//ordqueue作为服务端暴露给客户端使用

	//服务结束
	<-ofsquit
	fmt.Println("DealWithOrderForms Service shutdown ...")
}

//func LISTEN(){
//	var ticker = time.NewTicker(time.Duration(time.Millisecond * 500))
//	defer ticker.Stop()
//	for {
//		select {
//		case product:=<-Orderchannel:{
//		//TODO:处理司机发来的订单信息
//		}
//		case product:=<-Refundchannel:{
//		//TODO:处理司机发来的信息
//		}
//		case <-ticker.C:{
//			fmt.Println("服务超时")
//		}
//		default:
//
//		}
//
//	}
//}
