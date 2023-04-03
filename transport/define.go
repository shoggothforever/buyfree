package transport

import (
	"fmt"
	"time"
)

const (
	MAXBUFFER     int           = 2048
	WORKERTIMEOUT time.Duration = 3e9
	WORKERNUMS    int           = 1000
)

var PlatFormService *WorkerPool

func init() {
	PlatFormService = NewWorkerPool(WORKERNUMS)

}

type ReplyQueue chan bool
type DealWithOrderForm func(o *CountRequest)

var sem = make(chan bool, MAXBUFFER)
var TimeOut = time.Duration(5000 * time.Millisecond)
var GlobalCnt int64 = 0

type Req interface {
	//Refund()
	Do(exitchan ReplyQueue)
	Handle()
}
type ReqQueue chan Req
type Worker struct {
	ReqChan   ReqQueue
	ReplyChan ReplyQueue
}

func NewWorker() *Worker {
	return &Worker{ReqChan: make(ReqQueue), ReplyChan: make(ReplyQueue)}
}

type WorkerQueue chan *Worker
type WorkerPool struct {
	PoolSize   int
	ReqChan    ReqQueue
	WorkerChan WorkerQueue
	ReplyChan  ReplyQueue
}

func NewWorkerPool(size int) *WorkerPool {
	return &WorkerPool{
		PoolSize:   size,
		ReqChan:    make(ReqQueue, size),
		WorkerChan: make(WorkerQueue, size),
		ReplyChan:  make(ReplyQueue, 1),
	}
}
func (w *Worker) Run() {
	go func() {
		var req Req
		var ok bool
		for {
			select {
			case req, ok = <-w.ReqChan:
				{
					if !ok {
						w.ReplyChan <- false
						return
					}
					//atomic.AddInt64(&GlobalCnt, 1)
					//fmt.Println("第", GloabalCnt, "号任务开始执行")
					req.Do(w.ReplyChan)
				}
			}

		}

	}()
}

func (p *WorkerPool) Run() {
	fmt.Println("WorkerPool 初始化")
	for i := 0; i < p.PoolSize; i++ {
		worker := NewWorker()
		p.WorkerChan <- worker
		worker.Run()
	}
	//var cnt int64
	go func() {
		for {
			select {
			case req := <-p.ReqChan:
				{
					//fmt.Println("消费者接收到的任务编号", req.(*CountRequest).OrderInfo.Cost)
					worker := <-p.WorkerChan
					//atomic.AddInt64(&cnt, 1)
					//fmt.Println("获取到", cnt, worker)
					worker.ReqChan <- req
					//<-worker.ReplyChan
					res := <-worker.ReplyChan
					if res == true {
						//p.ReplyChan <- true
						p.WorkerChan <- worker
					} else {
						//工人系统异常关闭工人的发送管道,工人的线程也会随之关闭,将执行失败的任务再次送回管道（可以设定重试次数）
						close(worker.ReqChan)
						<-worker.ReplyChan
						//p.ReplyChan <- false
						//p.ReqChan <- req
						worker = nil
						newworker := NewWorker()
						p.WorkerChan <- newworker
						newworker.Run()
					}
				}
			}
		}
	}()
}

//func DetectSaleSign(work DealWithOrderForm, OrderChannel chan *CountRequest, ofsquit chan bool) {
//	var ticker = time.NewTicker(time.Duration(TimeOut))
//	defer ticker.Stop()
//	for {
//		sem <- true
//		select {
//		case orderformrequest := <-OrderChannel:
//			go work(orderformrequest) //购物
//		case <-ofsquit:
//			return
//		case <-ticker.C:
//			ofsquit <- true
//		default:
//
//		}
//	}
//}
//func StartServer(work DealWithOrderForm) (service chan *CountRequest, quit chan bool) {
//	service = make(chan *CountRequest, 512)
//	quit = make(chan bool)
//	go DetectSaleSign(work, service, quit)
//	return service, quit
//}
//func ListenOrderFromRequest() {
//	//TODO:迁移到main.go中去
//	ordqueue, ofsquit := StartServer(HandleOrderForm)
//	defer close(ordqueue)
//	defer close(ofsquit)
//	//ordqueue作为服务端暴露给客户端使用
//
//	//服务结束
//	<-ofsquit
//	fmt.Println("DealWithOrderForms Service shutdown ...")
//}
