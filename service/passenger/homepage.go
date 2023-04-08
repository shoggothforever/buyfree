package passenger

import (
	"buyfree/dal"
	"buyfree/mrpc"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type HomePageController struct {
	BasePaController
}

// @Summary 乘客端首页
// @Description 用户扫码打开小程序，
// @Tags Driver
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
// @Tags Driver
// @Accept json
// @Produce json
// @Param PayInfo body response.PassengerPayInfo true "将首页的商品信息传入"
// @Success 200 {object} response.PayResponse
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /goods/{id} [get]
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
		h.Error(c, 500, "处理支付信息失败有误")
		return
	} else {
		payreq.Orderform.PassengerID = admin.ID
		cerr := dal.Getdb().Model(&model.OrderForm{}).Create(&payreq.Orderform).Error
		if cerr != nil {
			h.Error(c, 400, "创建订单信息失败")
			return
		}
		c.JSON(200, response.PayResponse{response.Response{200, "支付成功"}})
	}
}
