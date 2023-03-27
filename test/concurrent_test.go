package test

import (
	"buyfree/repo/model"
	"buyfree/utils"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestOrderWorkerPool(t *testing.T) {
	num := 200 //开启 2万个线程
	//debug.SetMaxThreads(num + 1000) //设置最大线程数
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	p := utils.NewOrderWorkerPool(num)
	p.Run()
	var i int64 = 0
	go func() {
		for {
			fmt.Scanln(&i)
			//i++
			order := &utils.OrderFormRequest{OrderInfo: &model.SingleOrderForm{Cost: i}, Replychan: make(utils.ReplyQueue)}
			p.OrderChan <- order //数据传进去会被自动执行Do()方法，具体对数据的处理自己在Do()方法中定义
			order = nil
		}
	}()
	//循环打印输出当前进程的Goroutine 个数
	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(5 * time.Second)
	}
}
