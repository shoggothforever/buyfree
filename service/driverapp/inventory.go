package driverapp

import (
	"buyfree/dal"
	"buyfree/middleware"
	"buyfree/repo/model"
	"buyfree/service/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InventoryController struct {
	BaseDrController
}

// @Summary 车主设备库存
// @Description 展示车主拥有的设备库存数据
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} response.InventoryResponse
// @Failure 400 {object} response.Response
// @Router /dr/inventory [get]
func (i *InventoryController) GetInventory(c *gin.Context) {
	db := dal.Getdb()
	iadmin, ok := c.Get(middleware.DRADMIN)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	admin, ok := iadmin.(model.Driver)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	var dev_ids []int64
	i.rwm.RLock()
	defer i.rwm.RUnlock()
	err := db.Model(&model.Device{}).Select("id").Where("owner_id = ?", admin.ID).Find(&dev_ids).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		logrus.Info("获取用户设备信息失败", err)
		i.Error(c, 400, "获取车主设备信息失败")
		return
	}
	var products []model.DeviceProduct
	err = db.Model(&model.DeviceProduct{}).Where("device_id in ?", dev_ids).Find(&products).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		logrus.Info("获取用户设备商品信息失败", err)
		i.Error(c, 400, "获取车主设备商品信息失败")
	} else {
		c.JSON(200, response.InventoryResponse{Response: response.Response{Code: 200, Msg: "库存信息:"}, Products: products})
	}
}

// @Summary 车主单个设备库存
// @Description 乘客端和用户端交互的接口,(用户端扫码，获取该设备的商品信息）
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} response.InventoryResponse
// @Failure 400 {object} response.Response
// @Router /dr/inventory/{device_id} [get]
func (i *InventoryController) GetDeviceByScan(c *gin.Context) {
	dev_id := c.Param("device_id")
	db := dal.Getdb()
	var products []model.DeviceProduct
	i.rwm.RLock()
	defer i.rwm.RUnlock()
	err := db.Model(&model.DeviceProduct{}).Where("device_id = ?", dev_id).Find(&products).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		logrus.Info("获取用户设备商品信息失败", err)
		i.Error(c, 400, "获取车主设备商品信息失败")
	} else {
		c.JSON(200, response.InventoryResponse{Response: response.Response{Code: 200, Msg: "库存信息:"}, Products: products})
	}
}
