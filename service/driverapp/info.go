package driverapp

import (
	"buyfree/dal"
	"buyfree/logger"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InfoController struct {
	BaseDrController
}

// @Summary 我的信息界面获取车主设备信息
// @Description 获取激活设备信息
// @Tags Driver/info
// @Accept json
// @Produce json
// @Success 200 {object} response.DriverDeviceResponse
// @Failure 400 {object} response.Response
// @Router /dr/infos/devices [get]
func (i *InfoController) Getdevice(c *gin.Context) {

}

// @Summary 获取订单信息
// @Description 根据输入mode查看全部，待付款，待取货订单
// @Tags Driver/info
// @Accept json
// @Produce json
// @Param mode path int true "mode=0查看未支付订单，mode=1查看待取货订单，mode=2查看已完成订单，mode=else 返回全部订单"
// @Success 200 {object} response.DriverOrderFormResponse
// @Failure 400 {object} response.Response
// @Router /dr/infos/orderform/{mode} [get]
func (it *InfoController) GetOrders(c *gin.Context) {
	mode := c.Param("mode")
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		it.Error(c, 400, "获取车主信息失败")
		return
	}
	var dofs []model.DriverOrderForm
	if mode == "0" || mode == "1" || mode == "2" {
		err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("driver_id = ? and state = ?", admin.ID, mode).Find(&dofs).Error
		if err != nil {
			it.Error(c, 400, "获取订单信息失败")
			return
		}
	} else {
		err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("driver_id = ?", admin.ID).Find(&dofs).Error
		if err != nil {
			it.Error(c, 400, "获取订单信息失败")
			return
		}
	}
	n := len(dofs)
	logger.Loger.Info("获取到%d条订单信息\n", n)
	err := dal.Getdb().Transaction(func(tx *gorm.DB) error {
		for i := 0; i < n; i++ {
			terr := tx.Model(&model.OrderProduct{}).Where("order_refer = ?", dofs[i].OrderID).Find(&dofs[i].ProductInfos).Error
			if terr != nil && terr != gorm.ErrRecordNotFound {
				return terr
			}
		}
		return nil
	})

	if err == nil {
		c.JSON(200, response.DriverOrderFormResponse{
			response.Response{
				200,
				fmt.Sprintf("成功获取到%d条订单信息", len(dofs)),
			},
			dofs,
		})
	} else {
		it.Error(c, 400, "获取订单信息失败")
	}

	c.Set("Orders", dofs)
	c.Next()
}

// @Summary 获取订单信息
// @Description 根据传入id查看具体的订单信息
// @Tags Driver/info
// @Accept json
// @Produce json
// @Param id path string true "订单的编号"
// @Success 200 {object} response.DriverOrderDetailResponse
// @Failure 400 {object} response.Response
// @Router /dr/infos/orderdetail/{id} [get]
func (i *InfoController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	var odinfos []model.DriverOrderForm
	var faddress string
	var distance float64
	rdb := dal.Getrdb()

	err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("order_id = ?", id).First(&odinfos).Error
	if err != nil {
		logger.Loger.Info(err)
		i.Error(c, 400, "获取订单信息失败")
		return
	}
	err = dal.Getdb().Model(&model.OrderProduct{}).Where("order_refer = ?", id).Find(&odinfos[0].ProductInfos).Error
	if err != nil {
		logger.Loger.Info(err)
		i.Error(c, 400, "获取订单信息失败")
	} else {
		idistance, _ := utils.LocDist(c, rdb, utils.LOCATION, admin.Name, odinfos[0].FactoryName, "m")
		if idistance != nil {
			distance = idistance.(float64)
		}
		c.JSON(200, response.DriverOrderDetailResponse{
			Response:       response.Response{200, "成功获取订单信息"},
			FactoryAddress: faddress,
			Distance:       distance,
			ReserveMobile:  admin.Mobile,
			OrderInfos:     odinfos[0],
		})
	}
}
