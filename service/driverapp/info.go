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
	admin, ok := utils.GetDriveInfo(c)
	if !ok {
		i.Error(c, 400, "获取用户信息失败")
		return
	}
	var res response.DriverDeviceResponse
	res.Response = response.Response{200, "获取信息成功"}
	err := dal.Getdb().Raw("select * from devices where owner_id =?", admin.ID).Find(&res.Devices).Error
	if err != nil {
		i.Error(c, 400, "获取设备信息失败")
		return
	}
	for k, v := range res.Devices {
		err = dal.Getdb().Raw("select * from device_products where device_id = ?", v.ID).Find(&res.Devices[k].Products).Error
		if err != nil {
			i.Error(c, 400, "获取设备信息失败")
			return
		}
		err = dal.Getdb().Raw("select * from advertisements as a inner join (select advertisement_id from ad_devices where device_id=?) as b  on a.id = b.advertisement_id", v.ID).Find(&res.Devices[k].Advertisements).Error
		if err != nil {
			i.Error(c, 400, "获取设备信息失败")
			return
		}
	}
	c.JSON(200, res)
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
	var mobile string
	var distance string
	rdb := dal.Getrdb()

	err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("order_id = ?", id).First(&odinfos).Error
	if err != nil {
		logger.Loger.Info(err)
		i.Error(c, 400, "获取订单信息失败")
		return
	}
	dal.Getdb().Raw("select mobile,address from factories where id = ?", odinfos[0].FactoryID).Row().Scan(&mobile, &faddress)
	//if err != nil {
	//	logger.Loger.Info(err)
	//	i.Error(c, 400, "获取场站信息失败")
	//	return
	//}
	err = dal.Getdb().Model(&model.OrderProduct{}).Where("order_refer = ?", id).Find(&odinfos[0].ProductInfos).Error
	if err != nil {
		logger.Loger.Info(err)
		i.Error(c, 400, "获取订单信息失败")
	} else {
		idistance, _ := utils.LocDist(c, rdb, utils.LOCATION, admin.Name, odinfos[0].FactoryName, "m")
		if idistance != nil {
			distance = idistance.(string)
		}
		c.JSON(200, response.DriverOrderDetailResponse{
			Response:       response.Response{200, "成功获取订单信息"},
			FactoryAddress: faddress,
			Distance:       distance,
			ReserveMobile:  mobile,
			OrderInfos:     odinfos[0],
		})
	}
}

//// @Summary 获取余额信息
//// @Description 余额组成:未结算广告收入+未体现设备收入
//// @Tags Driver/balance
//// @Accept json
//// @Produce json
//// @Success 200 {object} response.BalanceResponse
//// @Failure 400 {object} response.Response
//// @Router /dr/infos/balance [get]
//func (i *InfoController) GetBalance(c *gin.Context) {
//	admin, ok := utils.GetDriveInfo(c)
//	if ok != true {
//		i.Error(c, 400, "获取车主信息失败")
//		return
//	}
//	var ids []int64
//	err := dal.Getdb().Raw("select id from devices where owner_id = ?", admin.ID).Find(&ids).Error
//	if err != nil {
//		i.Error(c, 400, "获取设备信息失败")
//		return
//	}
//	var sum1, sum2, sum float64
//	err = dal.Getdb().Raw("select sum(profit) from devices where owner_id = ?", admin.ID).First(&sum1).Error
//	if err != nil {
//		i.Error(c, 400, "获取设备收益信息失败")
//		return
//	}
//	err = dal.Getdb().Raw("select sum(profit) from ad_devices where device_id in ?", ids).First(&sum2).Error
//	if err != nil {
//		i.Error(c, 400, "获取广告收益信息失败")
//		return
//	}
//	sum = sum1 + sum2
//	c.JSON(200, response.BalanceResponse{response.Response{200, "获取账户余额成功"}, sum})
//}

// @Summary 提取余额信息（只支持全部提现)
// @Description 余额组成:未结算广告收入+未体现设备收入
// @Tags Driver/balance
// @Accept json
// @Produce json
// @Success 200 {object} response.BalanceResponse
// @Failure 400 {object} response.Response
// @Router /dr/infos/withdraw [get]
func (i *InfoController) Withdraw(c *gin.Context) {
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	var ids []int64
	err := dal.Getdb().Raw("select id from devices where owner_id = ?", admin.ID).Find(&ids).Error
	if err != nil {
		i.Error(c, 400, "获取设备信息失败")
		return
	}
	var sum1, sum2, sum, cash float64
	err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			terr := dal.Getdb().Model(&model.Device{}).Select("profit").Where("id = ? ", id).First(&cash).Update("profit", 0).Error
			if terr != nil {
				logger.Loger.Info(terr)
				i.Error(c, 400, "获取设备收益信息失败")
				return terr
			}
			sum1 += cash
			var cashs []float64
			terr = dal.Getdb().Model(&model.Ad_Device{}).Select("profit").Where("device_id = ? ", id).Find(&cashs).Update("profit", 0).Error
			if terr != nil {
				logger.Loger.Info(terr)
				i.Error(c, 400, "获取广告收益信息失败")
				return terr
			}
			for _, v := range cashs {
				sum2 += v
			}

		}
		return nil
	})

	if err != nil {
		i.Error(c, 400, "获取账户余额失败")
	} else {
		sum = sum1 + sum2
		c.JSON(200, response.BalanceResponse{response.Response{200, "提现成功"}, sum})
	}
}
