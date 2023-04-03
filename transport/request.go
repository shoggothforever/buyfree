package transport

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
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
	Iterator int64 `json:"iterator"`
	//回复信号
	ReplyChan ReplyQueue
}

func NewCountRequest(it int64) *CountRequest {
	req := &CountRequest{Iterator: it, ReplyChan: make(ReplyQueue, 1)}
	return req
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

type PayRequest struct {
	PlatFormID int64      `json:"platform_id,omitempty"`
	Cash       float64    `json:"cash,omitempty"`
	ReplyChan  ReplyQueue `json:"reply_chan,omitempty"`
}

func NewPayRequest(ptid int64, cash float64) *PayRequest {
	return &PayRequest{PlatFormID: ptid, Cash: cash, ReplyChan: make(ReplyQueue, 1)}
}

//------------------------------------------------------------------------------------------------------------------------
//封装每个req下的DO方法的操作
//对redis数据库进行操作,考虑退款操作
func (r *Reply) Handle() {

}

func (o *CountRequest) Handle() {
	o.ReplyChan <- true
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
func (p *PayRequest) Handle() {
	rdb := dal.Getrdb()
	ctx := rdb.Context()
	var name string
	err := dal.Getdb().Model(&model.Platform{}).Select("name").Where("id= ?", p.PlatFormID).First(&name).Error
	if err != nil {
		p.ReplyChan <- false
		return
	}
	scash := strconv.FormatFloat(p.Cash, 'f', 2, 64)
	_, err = utils.ModifySales(ctx, rdb, utils.Ranktype1, name, scash)
	if err != nil {
		p.ReplyChan <- false
	} else {
		p.ReplyChan <- true
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

func (o *CountRequest) Do(exitChan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	o.Handle()
	select {
	case val := <-o.ReplyChan:
		//fmt.Println("HandleOrderForm res:", val)
		exitChan <- val
		return
		//case <-ticker.C:
		//	fmt.Println("time out")
		//	exitchan <- false
		//	return
	}
}
func (s *ScanRequest) Do(exitChan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	s.Handle()
	select {
	case val := <-s.ReplyChan:
		fmt.Println("HandleScan res:", val)
		s.ReplyChan <- val
		close(s.ReplyChan)
		//ticker.Stop()
		exitChan <- val
		return
	case <-ticker.C:
		//fmt.Println("time out")
		s.ReplyChan <- false
		close(s.ReplyChan)
		exitChan <- false
		return
	}
}
func (d *DeviceAuthRequest) Do(exitChan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	d.Handle()
	select {
	case val := <-d.ReplyChan:
		fmt.Println("HandleDeviceAuth res:", val)
		d.ReplyChan <- val
		close(d.ReplyChan)
		exitChan <- val
		return
	case <-ticker.C:
		//fmt.Println("time out")
		d.ReplyChan <- false
		close(d.ReplyChan)
		exitChan <- false
		return
	}
}
func (p *PayRequest) Do(exitChan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	p.Handle()
	select {
	case val := <-p.ReplyChan:
		fmt.Println("HandleDeviceAuth res:", val)
		p.ReplyChan <- val
		close(p.ReplyChan)
		exitChan <- val
		return
	case <-ticker.C:
		//fmt.Println("time out")
		p.ReplyChan <- false
		close(p.ReplyChan)
		exitChan <- false
		return
	}
}

//----------------------------------------------------------------------------------------------------------------------
