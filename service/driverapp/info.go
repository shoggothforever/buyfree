package driverapp

import (
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
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
// @Description 根据输入mode查看全部，代付款，待取货订单
// @Tags Driver/info
// @Accept json
// @Produce json
// @Param mode path int true "mode=0查看全部订单，mode=1查看待付款订单，mode=2查看待取货订单"
// @Success 200 {object} response.DriverOrdersResponse
// @Failure 400 {object} response.Response
// @Router /dr/infos/orderform/{mode} [get]
func (i *InfoController) GetOrders(c *gin.Context) {
	var DOrderforms []model.DriverOrderForm

	c.JSON(200, response.DriverOrderResponse{Response: response.Response{200, "获取订单信息成功"}, OrderInfos: DOrderforms})
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
	var address string
	var mobile string
	var distance float64
	var odinfos model.DriverOrderForm
	c.JSON(200, response.DriverOrderDetailResponse{response.Response{200, "获取订单信息成功"},
		address, mobile, distance, odinfos})
}
