package mrpc

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type DeviceProductInfo struct {
}

type HomeScanReq struct {
	DeviceID       int64
	ADUrls         []model.ADurl
	DeviceProducts []model.DeviceProductPartInfo
	Communicator
}

func NewHomeScanReq(id int64) *HomeScanReq {
	return &HomeScanReq{DeviceID: id, ADUrls: []model.ADurl{}, DeviceProducts: []model.DeviceProductPartInfo{}, Communicator: NewCommunicator()}
}
func (h *HomeScanReq) Handle() {
	fmt.Println("HomeScanHandler Begin...")
	ids := []int64{}
	db := dal.Getdb()
	err := db.Transaction(func(tx *gorm.DB) error {
		terr := db.Raw("select advertisement_id from ad_devices where device_id = ?", h.DeviceID).Find(&ids).Error
		if terr != nil {
			logrus.Info(terr)
			return terr
		}
		terr = db.Raw("select id,video_cover,play_url from advertisements where id in ?", ids).Find(&h.ADUrls).Error
		if terr != nil {
			logrus.Info(terr)
			return terr
		}
		return nil
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Info(err)
		h.Send(false)
		return
	}
	err = db.Raw("select name,pic,type,buy_price from device_products where device_id = ?", h.DeviceID).Find(&h.DeviceProducts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Info(err)
		h.Send(false)
	} else {
		h.Send(true)
	}

}

type PassengerPayReq struct {
	DeviceID  int64                    `json:"device_id,omitempty"`
	Name      string                   `json:"name,omitempty"`
	BuyPrice  float64                  `json:"buy_price,omitempty"`
	Orderform model.PassengerOrderForm `json:"order_form"`
	Communicator
}

func NewPassengerPayReq(id int64, name string, price float64) *PassengerPayReq {
	return &PassengerPayReq{DeviceID: id, Name: name, BuyPrice: price, Orderform: model.PassengerOrderForm{OrderForm: model.OrderForm{PlaceTime: time.Now()}}, Communicator: NewCommunicator()}
}
func (p *PassengerPayReq) Handle() {
	db := dal.Getdb()
	var dp model.DeviceProduct
	err := db.Model(&model.DeviceProduct{}).Where("device_id = ? and name = ? and inventory >0", p.DeviceID, p.Name).UpdateColumn("inventory", gorm.Expr("inventory - ?", 1)).First(&dp).Error

	if err != nil {
		logrus.Info("购买商品失败")
		p.Send(false)
		return
	} else {
		rdb := dal.Getrdb()
		ctx := rdb.Context()
		var oid int64
		var name string
		err = db.Model(&model.Device{}).Select("owner_id").Where("id = ?", p.DeviceID).UpdateColumn("inventory", gorm.Expr("inventory - ?", 1)).First(&oid).Error
		if err != nil {
			logrus.Info("购买商品失败")
			p.Send(false)
			return
		}
		var driver model.Driver
		err = db.Model(&model.Driver{}).Where("id = ?", oid).First(&driver).Error
		if err != nil {
			logrus.Info("获取车主信息失败")
			p.Send(false)
			return
		}
		utils.ModifySales(ctx, rdb, utils.Ranktype1, dp.Name)
		utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, name, p.Name, p.BuyPrice)
		p.Orderform.OrderID = utils.GetSnowFlake()
		p.Orderform.Cost = p.BuyPrice
		p.Orderform.State = 2
		p.Orderform.Location = driver.Address
		p.Orderform.PayTime = time.Now()

		p.Send(true)
	}
}
