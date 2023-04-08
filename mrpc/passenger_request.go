package mrpc

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DeviceProductInfo struct {
}

type HomeScanReq struct {
	Device_id      int64
	ADUrls         []model.ADurl
	DeviceProducts []model.DeviceProductPartInfo
	Communicator
}

func NewHomeScanReq(id int64) *HomeScanReq {
	return &HomeScanReq{Device_id: id, ADUrls: []model.ADurl{}, DeviceProducts: []model.DeviceProductPartInfo{}, Communicator: NewCommunicator()}
}
func (h *HomeScanReq) Handle() {
	fmt.Println("HomeScanHandler Begin...")
	ids := []int64{}
	db := dal.Getdb()
	err := db.Transaction(func(tx *gorm.DB) error {
		terr := db.Raw("select advertisement_id from ad_devices where device_id = ?", h.Device_id).Find(&ids).Error
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
	err = db.Raw("select name,pic,type,buy_price from device_products where device_id = ?", h.Device_id).Find(&h.DeviceProducts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.Info(err)
		h.Send(false)
	} else {
		h.Send(true)
	}

}
