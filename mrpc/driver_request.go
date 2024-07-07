package mrpc

import (
	"buyfree/dal"
	"buyfree/logger"
	"buyfree/repo/model"
	"buyfree/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

//购物通道|补货通道
//var Orderchannel chan *model.CountRequest = make(chan *model.CountRequest)
//退款通道
//var Refundchannel chan *model.CountRequest = make(chan *model.CountRequest)

type ScanRequest struct {
	DriverID int64  `json:"driver_id"`
	DeviceID *int64 `json:"device_id"`
	Communicator
}

func NewScanRequest(driver_id int64, pid *int64) *ScanRequest {
	return &ScanRequest{DriverID: driver_id, DeviceID: pid, Communicator: NewCommunicator()}
}

type DeviceAuthRequest struct {
	// 车主ID
	DriverID int64 `json:"driver_id,omitempty"`
	//车主设备ID
	DeviceID   int64  `json:"device_id"`
	DriverName string `json:"driver_name,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	Communicator
}

func NewDeviceAuthRequest(driver_id, device_id int64, driver_name, mobile string) *DeviceAuthRequest {
	return &DeviceAuthRequest{
		DriverID:     driver_id,
		DeviceID:     device_id,
		DriverName:   driver_name,
		Mobile:       mobile,
		Communicator: NewCommunicator(),
	}
}

type PayRequest struct {
	PlatFormID int64   `json:"platform_id,omitempty"`
	Cash       float64 `json:"cash,omitempty"`
	Communicator
}

func NewPayRequest(ptid int64, cash float64) *PayRequest {
	return &PayRequest{PlatFormID: ptid, Cash: cash, Communicator: NewCommunicator()}
}

type OrderRequest struct {
	FactoryID   int64  `json:"factory_id,omitempty"`
	OrderID     int64  `json:"order_id,omitempty"`
	FactoryName string `json:"factory_name,omitempty"`

	ProductInfos *[]*model.OrderProduct `json:"product_infos,omitempty"`
	Communicator
}

func NewOrderRequest(fid, oid int64, fname string, products *[]*model.OrderProduct) *OrderRequest {
	return &OrderRequest{FactoryID: fid, OrderID: oid, FactoryName: fname, ProductInfos: products, Communicator: NewCommunicator()}
}

type PassengerPayRequest struct {
	DriverID  int64
	DeviceID  int64
	ProductID int64
}

func NewPassengerPayRequest() *PassengerPayRequest {
	return &PassengerPayRequest{
		DeviceID:  0,
		ProductID: 0,
	}
}

//------------------------------------------------------------------------------------------------------------------------
//封装每个req下的DO方法的操作
//对redis数据库进行操作,考虑退款操作

func (s *ScanRequest) Handle() {
	var id int64
	dal.Getdb().Model(&model.Platform{}).Select("id").Take(&id)
	err := dal.Getdb().Model(&model.Driver{}).Where("id=?", s.DriverID).Update("platform_id", id).Error
	if err != nil {
		logrus.Info("用户认证失败")
		s.Send(false)
	} else {
		dal.Getdb().Model(&model.Device{}).Select("id").Where("is_activated = false").First(s.DeviceID)
		s.Send(true)
	}
}
func (d *DeviceAuthRequest) Handle() {
	err := dal.Getdb().Model(&model.Device{}).Where("id=?", d.DeviceID).UpdateColumns(map[string]interface{}{"is_activated": true, "is_online": true, "owner_id": d.DriverID}).Error
	if err != nil {
		logrus.Info("设备激活失败", err)
		d.Send(false)
	} else {
		d.Send(true)
	}
}
func (p *PayRequest) Handle() {
	if p.Cash == 0 {
		p.Send(true)
		return
	}
	//fmt.Println("pay handle begin")
	rdb := dal.Getrdb()
	ctx := rdb.Context()
	var name string
	err := dal.Getdb().Model(&model.Platform{}).Select("name").Where("id= ?", p.PlatFormID).First(&name).Error
	if err != nil {
		logger.Loger.Info(err)
		p.Send(false)
		return
	}
	scash := strconv.FormatFloat(p.Cash, 'f', 2, 64)
	_, err = utils.ModifySales(ctx, rdb, utils.Ranktype1, utils.PTNAME, scash)
	if err != nil {
		logger.Loger.Info(err)
		p.Send(false)
	} else {
		p.Send(true)

	}
}
func (o *OrderRequest) Handle() {
	//处理一个场站的订单

	//查询场站商品库存信息，有一个商品库存不满足就直接判定为结算失败
	err := dal.Getdb().Transaction(func(tx *gorm.DB) error {
		for k, _ := range *o.ProductInfos {
			v := *(*o.ProductInfos)[k]
			var fp model.FactoryProduct
			terr := tx.Model(&model.FactoryProduct{}).Where(
				"factory_id = ? and name = ? and is_on_shelf =true and inventory>=?", v.FactoryID, v.Name, v.Count).First(&fp).UpdateColumn(
				"inventory", gorm.Expr("inventory - ?", v.Count)).Error
			if terr != nil {
				var s string
				if terr == gorm.ErrRecordNotFound {
					s = fmt.Sprintf("%d场站%s商品库存不足,订单取消", v.FactoryID, v.Name)
				}
				logger.Loger.Info(s, terr)
				return terr
			}
			ms, _ := strconv.ParseInt(fp.MonthlySales, 10, 64)
			ts, _ := strconv.ParseInt(fp.TotalSales, 10, 64)
			ums := strconv.FormatInt(ms+v.Count, 10)
			uts := strconv.FormatInt(ts+v.Count, 10)
			merr := tx.Model(&model.FactoryProduct{}).Select("monthly_sales", "total_sales").Where(
				"factory_id = ? and name = ? and is_on_shelf =true", v.FactoryID, v.Name).UpdateColumns(map[string]interface{}{
				"monthly_sales": ums,
				"total_sales":   uts,
			}).Error
			if merr != nil {
				logger.Loger.Info("Update Factory_products error:\n", merr)
				return merr
			}
		}
		//TODO更新榜单信息
		rdb := dal.Getrdb()
		ctx := rdb.Context()
		for k, _ := range *o.ProductInfos {
			v := *(*o.ProductInfos)[k]
			//var inv int64
			//terr := tx.Model(&model.FactoryProduct{}).Select("inventory").Where("factory_id = ? and name = ? and is_on_shelf =true ", v.FactoryID, v.Name).UpdateColumn("inventory", gorm.Expr("inventory - ?", v.Count)).First(&inv).Error
			cost := float64(v.Count) * v.Price
			fmt.Println(fmt.Sprintf("%d订单：%s商品营销额:%f", v.OrderRefer, v.Name, float64(v.Count)*v.Price))

			utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, o.FactoryName, v.Name, cost)
			utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, utils.PTNAME, v.Name, cost)
		}
		return nil
	})
	if err != nil {
		logger.Loger.Info("OrderRequest Transaction error\n", err)
		o.Send(false)
		return
	} else {
		o.Send(true)
	}
}
