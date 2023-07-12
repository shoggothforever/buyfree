package mrpc

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
var DriverService *WorkerPool

func init() {
	PlatFormService = NewWorkerPool(WORKERNUMS)
	DriverService = NewWorkerPool(WORKERNUMS)
}

type ReplyQueue chan bool

var TimeOut = time.Duration(5000 * time.Millisecond)
var GlobalCnt int64 = 0

type Handler func()
type Req interface {
	//Refund()
	//exitChan 向Worker传递工作处理结束的信息,handle传递工作的处理方法
	Do(exitChan ReplyQueue, handle Handler)
	//Handle() 由decorator实现
	Handle()
	Done()
	Result() bool
}

// 装饰着模式中的 Component
type Communicator struct {
	//客户端验证结果
	Res bool
	//通知worker工作处理结果
	ReplyChan ReplyQueue
	//客户端通信管道，告知任务完成
	DoneChan chan struct{}
}

func NewCommunicator() Communicator {
	return Communicator{
		Res:       *new(bool),
		ReplyChan: make(ReplyQueue, 1),
		DoneChan:  make(chan struct{}, 1),
	}
}

// 实现每个Req的接口定义
func (c *Communicator) Do(exitchan ReplyQueue, handle Handler) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	//实现运行时多态
	handle()
	select {
	case val := <-c.ReplyChan:
		fmt.Println("HandleReq res:", val)
		close(c.ReplyChan)
		exitchan <- val
		return
	case <-ticker.C:
		fmt.Println("time out")
		close(c.ReplyChan)
		exitchan <- false
		return
	}
}
func (c *Communicator) Send(sig bool) {
	c.ReplyChan <- sig
	c.Res = sig
	c.DoneChan <- struct{}{}
}
func (c *Communicator) Done() {
	<-c.DoneChan
}
func (c *Communicator) Result() bool {
	return c.Res
}

// 定义worker，用于处理请求
type ReqQueue chan Req
type Worker struct {
	ReqChan   ReqQueue
	ReplyChan ReplyQueue
}

// 定义worker的Reqchan缓冲为1
func NewWorker() *Worker {
	return &Worker{ReqChan: make(ReqQueue, 1), ReplyChan: make(ReplyQueue, 1)}
}

type WorkerQueue chan *Worker
type WorkerPool struct {
	PoolSize   int
	ReqChan    ReqQueue
	WorkerChan WorkerQueue
	ReplyChan  ReplyQueue
}

// 定义WorkerPool 的Reqchan 的缓冲为Worker的数量
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
					//fmt.Println(ok)
					//atomic.AddInt64(&GlobalCnt, 1)
					//fmt.Println("第", GloabalCnt, "号任务开始执行")
					req.Do(w.ReplyChan, req.Handle)
				}
			}
		}
	}()
}
func (p *WorkerPool) PutReq(r Req) {
	p.ReqChan <- r
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
					worker.ReqChan <- req
					go func(worker *Worker, p *WorkerPool) {
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
					}(worker, p)
				}
			}
		}
	}()
}

func PutDriverReq(r Req) {
	PlatFormService.PutReq(r)
	r.Done()
}
func PutPassengerReq(r Req) {
	DriverService.PutReq(r)
	r.Done()
}
