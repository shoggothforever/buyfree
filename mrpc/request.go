package mrpc

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	Res        bool       `json:"res"`
}

func NewPayRequest(ptid int64, cash float64) *PayRequest {
	return &PayRequest{PlatFormID: ptid, Cash: cash, ReplyChan: make(ReplyQueue, 1)}
}

type OrderRequest struct {
	//TODO:添加相关项
	FactoryID   int64  `json:"factory_id,omitempty"`
	OrderID     int64  `json:"order_id,omitempty"`
	FactoryName string `json:"factory_name,omitempty"`

	ProductInfos *[]*model.OrderProduct `json:"product_infos,omitempty"`
	ReplyChan    ReplyQueue             `json:"reply_chan,omitempty"`
	Res          bool                   `json:"res,omitempty"`
	DoneChan     chan struct{}          `json:"res_chan"`
}

func NewOrderRequest(fid, oid int64, fname string, products *[]*model.OrderProduct) *OrderRequest {
	return &OrderRequest{FactoryID: fid, OrderID: oid, FactoryName: fname, ProductInfos: products, ReplyChan: make(ReplyQueue, 1), DoneChan: make(chan struct{}, 1)}
}

//------------------------------------------------------------------------------------------------------------------------
//封装每个req下的DO方法的操作
//对redis数据库进行操作,考虑退款操作
func (r *Reply) Handle() {

}

func (o *CountRequest) Handle() {
	o.ReplyChan <- true
	//fmt.Println("管道大小", len(o.ReplyChan))
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
	//fmt.Println("pay handle begin")
	rdb := dal.Getrdb()
	ctx := rdb.Context()
	var name string
	err := dal.Getdb().Model(&model.Platform{}).Select("name").Where("id= ?", p.PlatFormID).First(&name).Error
	if err != nil {
		fmt.Println(err)
		p.ReplyChan <- false
		p.Res = false
		return
	}
	scash := strconv.FormatFloat(p.Cash, 'f', 2, 64)
	_, err = utils.ModifySales(ctx, rdb, utils.Ranktype1, name, scash)
	if err != nil {
		fmt.Println(err)
		p.Res = false
		p.ReplyChan <- false

	} else {
		p.Res = true
		p.ReplyChan <- true

	}
}
func (o *OrderRequest) Handle() {
	//TODO 业务逻辑
	//处理一个场站的订单

	//查询场站商品库存信息，有一个商品库存不满足就直接判定为结算失败
	err := dal.Getdb().Transaction(func(tx *gorm.DB) error {
		for k, _ := range *o.ProductInfos {
			v := *(*o.ProductInfos)[k]
			fmt.Println(v)
			var inv int64
			terr := tx.Model(&model.FactoryProduct{}).Select("inventory").Where("factory_id = ? and name = ? and is_on_shelf =true and inventory>=?", v.FactoryID, v.Name, v.Count).UpdateColumn("inventory", gorm.Expr("inventory - ?", v.Count)).First(&inv).Error
			fmt.Println(fmt.Sprintf("%d场站%s商品库存数量%d", v.FactoryID, v.Name, inv))
			if terr != nil {
				logrus.Info(terr)
				return terr
			}
		}
		fmt.Println("订单编号：", o.OrderID)
		//TODO更新榜单信息
		for k, _ := range *o.ProductInfos {
			v := *(*o.ProductInfos)[k]
			fmt.Println(v)
			//var inv int64
			//terr := tx.Model(&model.FactoryProduct{}).Select("inventory").Where("factory_id = ? and name = ? and is_on_shelf =true ", v.FactoryID, v.Name).UpdateColumn("inventory", gorm.Expr("inventory - ?", v.Count)).First(&inv).Error
			fmt.Println(fmt.Sprintf("%d订单：%s商品营销额:%f", v.OrderRefer, v.Name, float64(v.Count)*v.Price))

			//if terr != nil {
			//	logrus.Info(terr)
			//	return terr
			//}
		}
		return nil
	})
	if err != nil {
		logrus.Info(err)
		fmt.Println(err)
		o.Res = false
		o.ReplyChan <- false
		o.DoneChan <- struct{}{}
		return
	}
	o.Res = true
	o.ReplyChan <- true
	o.DoneChan <- struct{}{}
}

//------------------------------------------------------------------------------------------------------------------------
// 实现每个Req的接口定义
func (r *Reply) Do(exitchan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	r.Handle()
	select {
	case val := <-r.Chan:
		fmt.Println("HandleReq res:", val)
		exitchan <- val
		return
	case <-ticker.C:
		fmt.Println("time out")
		exitchan <- false
		return
	}
}
func (o *CountRequest) Do(exitChan ReplyQueue) {
	ticker := time.NewTicker(TimeOut / 100)
	defer ticker.Stop()
	o.Handle()
	select {
	case val := <-o.ReplyChan:
		//fmt.Println("HandleCounter res:", val)
		o.ReplyChan <- val
		close(o.ReplyChan)
		exitChan <- val
		return
	case <-ticker.C:
		fmt.Println("time out")
		close(o.ReplyChan)
		exitChan <- false
		return
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
		//fmt.Println("HandlePay res:", val)
		//p.ReplyChan <- val
		close(p.ReplyChan)
		//fmt.Println("管道大小", len(p.ReplyChan))
		exitChan <- val
		return
	case <-ticker.C:
		//fmt.Println("time out")
		//p.ReplyChan <- false
		close(p.ReplyChan)
		fmt.Println("超时管道大小", len(p.ReplyChan))
		exitChan <- false
		return
	}
}
func (o *OrderRequest) Do(exitChan ReplyQueue) {
	ticker := time.NewTicker(TimeOut)
	defer ticker.Stop()
	o.Handle()
	select {
	case val := <-o.ReplyChan:
		fmt.Println("HandleOrderFormRequest res:", val)
		//o.ReplyChan <- val
		//fmt.Println(o.OrderID, "管道大小", len(o.ReplyChan))
		close(o.ReplyChan)
		exitChan <- val
		return
	case <-ticker.C:
		//fmt.Println("time out")
		close(o.ReplyChan)
		exitChan <- false
		return
	}
}

//----------------------------------------------------------------------------------------------------------------------
