package utils

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MAXBUFFER     int           = 2048
	WORKERTIMEOUT time.Duration = 3e9
)

type ReplyQueue chan bool
type DealWithOrderForm func(o *OrderFormRequest)

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

type OrderReq interface {
	//Refund()
	Pay(exitchan ReplyQueue)
}

type OrderFormRequest struct {
	OrderInfo *model.SingleOrderForm
	//回复信号
	Replychan ReplyQueue
}

//对redis数据库进行操作,考虑退款操作
func HandleOrderForm(o *OrderFormRequest) {
	//TODO
	//c := context.TODO() 数据库操作
	//rc := dal.Getrd()
	//y, m, d := GetDateKey()
	//rect.Store(o)

	fmt.Println(o.OrderInfo.Cost)
	rdb := dal.Getrdb()
	c := rdb.Context()
	res, err := rdb.Set(c, "goroutine", o.OrderInfo.Cost, 3e11).Result()
	fmt.Println(res, err, o.OrderInfo.Cost)
	<-o.Replychan
}

func (o *OrderFormRequest) Pay(exitchan ReplyQueue) {
	ticker := time.NewTicker(WORKERTIMEOUT)
	HandleOrderForm(o)
	select {
	case <-o.Replychan:
		exitchan <- true
		return
	case <-ticker.C:
		exitchan <- false
		return
	}
}
func (o *OrderFormRequest) Refund() {

}
func NewOrderFormRequest(s *model.SingleOrderForm) *OrderFormRequest {
	req := &OrderFormRequest{OrderInfo: s, Replychan: make(chan bool)}
	return req
}

type OrderReqQueue chan OrderReq
type OrderWorker struct {
	OrderChan OrderReqQueue
	ReplyChan ReplyQueue
}

func NewOrderWorker() *OrderWorker {
	return &OrderWorker{OrderChan: make(OrderReqQueue), ReplyChan: make(ReplyQueue)}
}

func (w *OrderWorker) Run(wp *OrderWorkerPool) {
	go func() {
		//ticker := time.NewTicker(WORKERTIMEOUT)
		wp.WorkerChan <- w
		var odreq OrderReq
		select {
		case odreq = <-w.OrderChan:
			{
				odreq.Pay(w.ReplyChan)
			}
		case val, ok := <-w.ReplyChan:
			{
				if !ok || val == false {
					w.ReplyChan = nil
					wp.OrderChan <- odreq
				}
				return
			}
			//case <-ticker.C:
			//	{
			//		logrus.Info("timeout while dealing with orderForm")
			//		return
			//	}
			//
		}

	}()
}

type OrderWorkerQueue chan *OrderWorker
type OrderWorkerPool struct {
	PoolSize   int
	OrderChan  OrderReqQueue
	WorkerChan OrderWorkerQueue
}

func NewOrderWorkerPool(size int) *OrderWorkerPool {
	return &OrderWorkerPool{
		PoolSize:   size,
		OrderChan:  make(OrderReqQueue, size),
		WorkerChan: make(OrderWorkerQueue, size),
	}
}
func (p *OrderWorkerPool) Run() {
	fmt.Println("WorkerPool 初始化")
	for i := 0; i < p.PoolSize; i++ {
		worker := NewOrderWorker()
		worker.Run(p)
	}
	go func() {
		for {
			select {
			case ordreq := <-p.OrderChan:
				{
					worker := <-p.WorkerChan
					worker.OrderChan <- ordreq
					//<-worker.ReplyChan
					worker.Run(p)
				}
			}
		}
	}()
}
func DetectSaleSign(work DealWithOrderForm, OrderChannel chan *OrderFormRequest, ofsquit chan bool) {
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
func StartServer(work DealWithOrderForm) (service chan *OrderFormRequest, quit chan bool) {
	service = make(chan *OrderFormRequest, 512)
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
