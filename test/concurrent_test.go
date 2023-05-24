package test

import (
	"buyfree/mrpc"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestOrderWorkerPool(t *testing.T) {
	//num := 2W //开启 2万个线程
	//debug.SetMaxThreads(num + 1000) //设置最大线程数
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	p := mrpc.NewWorkerPool(mrpc.WORKERNUMS)
	p.Run()
	var i int64 = 0
	go func() {
		//要对并发量做限制
		for {
			//fmt.Scanln(&i)
			i++
			orderreq := &mrpc.CountRequest{Iterator: i, Communicator: mrpc.NewCommunicator()}
			//fmt.Println("任务:", i)
			p.ReqChan <- orderreq //数据传进去会被自动执行Do()方法，具体对数据的处理自己在Do()方法中定义![](../../../车主端.png)
			//fmt.Println(<-orderreq.ReplyChan)
			orderreq = nil
		}
	}()
	//循环打印输出当前进程的Goroutine 个数
	for {
		fmt.Println("当前迭代的全局变量global值：", mrpc.GlobalCnt)
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(3 * time.Second)
	}
}
