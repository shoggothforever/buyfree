package passenger

import (
	"buyfree/dal"
	"buyfree/mrpc"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

type HomePageController struct {
	BasePaController
}

// @Summary 乘客端首页
// @Description 用户扫码打开小程序，
// @Tags Passenger/home
// @Accept json
// @Produce json
// @Param id path int true "扫码获取的设备id"
// @Success 200 {object} response.PassengerHomeResponse
// @Failure 400 {object} response.Response
// @Router /home/{id} [get]
func (h *HomePageController) GetStatic(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	homereq := mrpc.NewHomeScanReq(id)
	mrpc.DriverService.ReqChan <- homereq
	<-homereq.DoneChan
	if !homereq.Res {
		h.Error(c, 400, "获取信息失败")
	} else {
		c.JSON(200, response.PassengerHomeResponse{
			response.Response{200, "获取信息成功"},
			homereq.ADUrls,
			homereq.DeviceProducts,
		})
	}
}

// @Summary 购买商品
// @Description 点击商品购买按钮
// @Tags Passenger/home
// @Accept json
// @Produce json
// @Param PayInfo body response.PassengerPayInfo true "将首页的商品信息传入"
// @Success 200 {object} response.PayResponse
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /home/pay [post]
func (h *HomePageController) Pay(c *gin.Context) {
	var info response.PassengerPayInfo
	err := c.ShouldBind(&info)

	if err != nil {
		h.Error(c, 400, "传输数据有误")
		return
	}
	admin, ok := utils.GetPassengerInfo(c)
	if !ok {
		h.Error(c, 400, "获取用户信息失败")
		return
	}
	payreq := mrpc.NewPassengerPayReq(info.DeviceID, info.Name, info.BuyPrice)
	mrpc.DriverService.ReqChan <- payreq
	<-payreq.DoneChan
	if !payreq.Res {
		h.Error(c, 500, "处理支付信息失败")
		return
	} else {
		payreq.Orderform.PassengerID = admin.ID
		cerr := dal.Getdb().Model(&model.OrderForm{}).Create(&payreq.Orderform).Error
		if cerr != nil {
			h.Error(c, 400, "创建订单信息失败")
			return
		}
		var p model.OrderProduct
		p.Count = 1
		p.Name = info.Name
		p.Price = info.BuyPrice
		p.OrderRefer = payreq.Orderform.OrderID
		perr := dal.Getdb().Model(&model.OrderProduct{}).Create(&p)
		if perr != nil {
			h.Error(c, 400, "创建订单信息失败")
			return
		} else {
			c.JSON(200, response.PayResponse{response.Response{200, "支付成功"}})
		}
	}
}

// @Summary 获取订单信息
// @Description 根据输入mode查看全部，待付款，待取货订单
// @Tags Passenger/info
// @Accept json
// @Produce json
// @Param mode path int true "mode=0查看全部订单，mode=1查看待付款订单，mode=2查看待取货订单"
// @Success 200 {object} response.PassengerOrderFormResponse
// @Failure 400 {object} response.Response
// @Router /infos/orders/{mode} [get]
func (h *HomePageController) GetOrders(c *gin.Context) {
	mode := c.Param("mode")
	admin, ok := utils.GetPassengerInfo(c)
	if ok != true {
		h.Error(c, 400, "获取用户信息失败")
		return
	}
	var dofs []model.PassengerOrderForm
	if mode == "0" || mode == "1" || mode == "2" {
		err := dal.Getdb().Model(&model.PassengerOrderForm{}).Where("passenger_id = ? and state = ?", admin.ID, mode).Find(&dofs).Error
		if err != nil {
			h.Error(c, 400, "获取订单信息失败 1")
			return
		}
	} else {
		err := dal.Getdb().Model(&model.PassengerOrderForm{}).Where("passenger_id = ?", admin.ID).Find(&dofs).Error
		if err != nil {
			h.Error(c, 400, "获取订单信息失败 1")
			return
		}
	}
	n := len(dofs)
	logrus.Infof("获取到%d条订单信息", n)
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
		c.JSON(200, response.PassengerOrderFormResponse{
			response.Response{
				200,
				fmt.Sprintf("成功获取到%d条订单信息", len(dofs)),
			},
			dofs,
		})
	} else {
		h.Error(c, 400, "获取订单信息失败")
	}
	c.Set("Orders", dofs)
	c.Next()
}

// @Summary 获取单个订单信息
// @Description 传入订单ID
// @Tags Passenger/info
// @Accept json
// @Produce json
// @Param id path int true "订单编号"
// @Success 200 {object} response.PassengerOrderFormResponse
// @Failure 400 {object} response.Response
// @Router /infos/{id}/orders [get]
func (h *HomePageController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	var dofs []model.PassengerOrderForm
	err := dal.Getdb().Model(&model.PassengerOrderForm{}).Where("order_id = ?", id).Find(&dofs).Error
	if err != nil {
		logrus.Info(err)
		h.Error(c, 400, "获取订单信息失败")
		return
	}
	n := len(dofs)
	logrus.Info("获取到%d条订单信息", n)
	err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
		for i := 0; i < n; i++ {
			terr := tx.Model(&model.OrderProduct{}).Where("order_refer = ?", dofs[i].OrderID).Find(&dofs[i].ProductInfos).Error
			if terr != nil && terr != gorm.ErrRecordNotFound {
				return terr
			}
		}
		return nil
	})
	if err == nil {
		c.JSON(200, response.PassengerOrderFormResponse{
			response.Response{
				200,
				fmt.Sprintf("成功获取到%d条订单信息", len(dofs)),
			},
			dofs,
		})
	} else {
		logrus.Info(err)
		h.Error(c, 400, "获取订单信息失败")
	}
	c.Set("Orders", dofs)
	c.Next()
}
