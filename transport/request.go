package transport

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

/*
设计订单接口，使用装饰模式鉴别是购货订单还是退货订单
设计设备激活order
设计车主激活order
*/
//购物通道|补货通道
//var Orderchannel chan *model.CountRequest = make(chan *model.CountRequest)
//退款通道
//var Refundchannel chan *model.CountRequest = make(chan *model.CountRequest)

func practice() {
	PlatFormService = NewWorkerPool(WORKERNUMS)
	for i := 0; i < WORKERNUMS; i++ {
		worker := NewWorker()
		worker.Run()
	}
	go func() {
		for {
			select {
			case req := <-PlatFormService.ReqChan:
				{
					worker := <-PlatFormService.WorkerChan
					worker.ReqChan <- req
					if val := <-worker.ReplyChan; val == true {
						PlatFormService.WorkerChan <- worker
					} else {
						close(worker.ReqChan)
						close(worker.ReplyChan)
						worker = nil
						newworker := NewWorker()
						PlatFormService.WorkerChan <- newworker
						newworker.Run()

					}
				}

			}
		}
	}()
}

type Reply struct {
	Chan ReplyQueue `json:"reply_chan,omitempty"`
}
type CountRequest struct {
	CountInfo *SingleOrderForm
	//回复信号
	ReplyChan ReplyQueue
}

func NewCountRequest(s *SingleOrderForm) *CountRequest {
	req := &CountRequest{CountInfo: s, ReplyChan: make(ReplyQueue, 1)}
	return req
}

type SingleOrderForm struct {
	//订单车主ID
	UserID int64 `json:"user_id" form:"user_id"`
	//订单编码
	OrderID string `json:"order_id" form:"order_id"`
	//花费
	Cost int64 `json:"cost" form:"cost"`
	//订单状态 订单状态 2-已完成 1-待取货 0-未支付
	State int64 `json:"state" form:"state"`
	//支付时存储位置(购物时获取车主位置）
	Location string `json:"location" form:"location"`
	//true:订货，false:退货
	IsReplenishment bool `json:"is_replenishment" form:"is_replenishment"`
}

type ScanRequest struct {
	DriverID  int64  `json:"driver_id"`
	DeviceID  *int64 `json:"device_id"`
	ReplyChan ReplyQueue
}

func NewScanRequest(driver_id int64, pid *int64) *ScanRequest {
	return &ScanRequest{DriverID: driver_id, DeviceID: pid, ReplyChan: make(ReplyQueue, 1)}
}

type DeviceAuthRequest struct {
	// 车主ID
	DriverID int64 `json:"driver_id,omitempty"`
	//
	DeviceID     int64                    `json:"device_id"`
	DriverName   string                   `json:"driver_name,omitempty"`
	Mobile       string                   `json:"mobile,omitempty"`
	AuthResponse *response.DriverAuthInfo `json:"auth_response"`
	ReplyChan    ReplyQueue               `json:"reply_chan"`
}

func NewDeviceAuthRequest(driver_id, device_id int64, driver_name, mobile string) *DeviceAuthRequest {
	return &DeviceAuthRequest{
		DriverID:     driver_id,
		DeviceID:     device_id,
		DriverName:   driver_name,
		Mobile:       mobile,
		AuthResponse: &response.DriverAuthInfo{},
		ReplyChan:    make(ReplyQueue, 1),
	}
}

//------------------------------------------------------------------------------------------------------------------------
//封装每个req下的DO方法的操作
//对redis数据库进行操作,考虑退款操作
func (r *Reply) Handle() {

}

func (o *CountRequest) Handle() {
	//TODO 根据数据库操作结果返回对应的result
	//c := context.Background()
	//rc := dal.Getrd()
	//y, m, d := GetDateKey()
	//rect.Store(o)

	//fmt.Println("cost:", o.OrderInfo.Cost)
	//rdb := dal.Getrdb()
	//c := rdb.Context()
	//res, err := rdb.Set(c, "goroutine", o.OrderInfo.Cost, 3e11).Result()
	//fmt.Println(res, err, o.OrderInfo.Cost)
	{
		o.ReplyChan <- true
	}
}
func (s *ScanRequest) Handle() {
	err := dal.Getdb().Model(&model.Driver{}).Where("id=?", s.DriverID).Error
	fmt.Println(s.DriverID)
	if err != nil {
		logrus.Info("用户认证失败")
		s.ReplyChan <- false
	} else {
		*s.DeviceID = utils.IDWorker.NextId()
		s.ReplyChan <- true
	}
}
func (d *DeviceAuthRequest) Handle() {
	err := dal.Getdb().Model(&model.Device{}).Where("id=?", d.DeviceID).Update("is_activated", true).Error
	fmt.Println(d.DriverID)
	if err != nil {
		logrus.Info("设备激活失败", err)
		d.ReplyChan <- false
	} else {
		//d.AuthResponse.Name = d.DriverName
		//d.AuthResponse.DriverID = d.DriverID
		//d.AuthResponse.DeviceID = d.DeviceID
		//d.AuthResponse.Mobile = d.Mobile
		d.ReplyChan <- true
	}
}

//------------------------------------------------------------------------------------------------------------------------
// 实现每个Req的接口定义
func (r *Reply) Do(exitchan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	r.Handle()
	select {
	case val := <-r.Chan:
		//fmt.Println("HandleOrderForm res:", val)
		exitchan <- val
		return
	case <-ticker.C:
		fmt.Println("time out")
		exitchan <- false
		return
	}
}

func (o *CountRequest) Do(exitchan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	o.Handle()
	select {
	case val := <-o.ReplyChan:
		//fmt.Println("HandleOrderForm res:", val)
		exitchan <- val
		return
		//case <-ticker.C:
		//	fmt.Println("time out")
		//	exitchan <- false
		//	return
	}
}
func (s *ScanRequest) Do(exitchan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	s.Handle()
	select {
	case val := <-s.ReplyChan:
		fmt.Println("HandleScan res:", val)
		s.ReplyChan <- val
		close(s.ReplyChan)
		//ticker.Stop()
		exitchan <- val
		return
	case <-ticker.C:
		//fmt.Println("time out")
		s.ReplyChan <- false
		close(s.ReplyChan)
		exitchan <- false
		return
	}
}
func (d *DeviceAuthRequest) Do(exitchan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	d.Handle()
	select {
	case val := <-d.ReplyChan:
		fmt.Println("HandleDeviceAuth res:", val)
		d.ReplyChan <- val
		close(d.ReplyChan)
		exitchan <- val
		return
	case <-ticker.C:
		//fmt.Println("time out")
		d.ReplyChan <- false
		close(d.ReplyChan)
		exitchan <- false
		return
	}
}

//----------------------------------------------------------------------------------------------------------------------
