package driverapp

import (
	"buyfree/dal"
	"buyfree/middleware"
	"buyfree/repo/model"
	"buyfree/service/response"
	"fmt"
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
// @Failure 400 {onject} response.Response
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
	fmt.Println(admin)
	var dev_ids []int64
	err := db.Model(&model.Device{}).Select("id").Where("owner_id = ?", admin.ID).Find(&dev_ids).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		logrus.Info("拂去用户设备信息失败", err)
		i.Error(c, 400, "获取车主设备信息失败")
	}
	var products []model.DeviceProduct
	err = db.Model(&model.DeviceProduct{}).Where("device_id in ?", dev_ids).Find(&products).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		logrus.Info("拂去用户设备商品信息失败", err)
		i.Error(c, 400, "获取车主设备商品信息失败")
	}
	c.JSON(200, response.InventoryResponse{response.Response{200, "库存信息:"}, products})
}
