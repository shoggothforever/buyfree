package driverapp

import (
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type FactoryController struct {
	BaseDrController
}

// @Summary 场站信息
// @Description 按照场站距离展示数据，距离进的排名靠前
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Success 200 {object} response.FactoryInfoResponse
// @Failure 400 {onject} response.Response
// @Router /dr/factory [get]
func (i *FactoryController) Get(c *gin.Context) {
	var detail []response.FactoryInfo

	c.JSON(200, response.FactoryInfoResponse{
		response.Response{200, "成功获取附近场站信息"},
		detail,
	})
	c.Next()
}

// @Summary 场站详情
// @Description 根据场站ID获取场站具体信息
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param id path int true "场站ID"
// @Success 200 {object} response.FactoryInfoResponse
// @Failure 400 {onject} response.Response
// @Router /dr/factory/infos/{id} [get]
func (i *FactoryController) Detail(c *gin.Context) {
	//id := c.Param("id")
	var orderform model.DriverOrderForm
	var detail []response.FactoryInfo

	c.JSON(200, response.FactoryInfoResponse{
		response.Response{200, "成功获取该场站商品信息"},
		detail,
	})
	c.Set("Orderform", orderform)
	c.Next()
}

// @Summary 订单界面
// @Description 点击结算，展示订单信息
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param OrderForm body model.DriverOrderForm true "车主订单信息"
// @Success 200 {object} response.OrderResponse
// @Failure 400 {onject} response.Response
// @Router /dr/order [post]
func (i *FactoryController) Order(c *gin.Context) {
	//id := c.Param("id")
	//var orderform model.DriverOrderForm
	//var detail []response.FactoryInfo
	//
	//c.JSON(200, response.FactoryInfoResponse{
	//	response.Response{200, "成功获取该场站商品信息"},
	//	detail,
	//})
	//c.Set("Orderform", orderform)
	//c.Set("id", id)
	var OrderInfos []response.FactoryProductsInfo
	c.JSON(200, response.OrderResponse{
		response.Response{200, "支付成功"},
		OrderInfos,
	})
	c.Next()
}

// @Summary 补货订单结算
// @Description 结算
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param OrderForm body model.DriverOrderForm true "传入订单信息"
// @Success 200 {object} response.PayResponse
// @Failure 400 {onject} response.Response
// @Router /dr/order/pay [put]
func (i *FactoryController) Pay(c *gin.Context) {
	//id, _ := c.Get("id")
	//TODO:业务逻辑
	//fmt.Println(id)
	c.JSON(200, response.PayResponse{
		response.Response{200, "支付成功"},
	})

}
