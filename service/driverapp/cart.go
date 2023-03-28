package driverapp

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type CartController struct {
	BaseDrController
}

// @Summary 购物车界面
// @Description 展示所有添加的商品，点击结算，跳转到订单提交界面
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param id path int true "车主id"
// @Success 201 {object} response.CartResponse
// @Failure 400 {onject} response.Response
// @Router /dr/order/cart/{id} [get]
func (ct *CartController) GetCart(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var cart model.DriverCart
	err := dal.Getdb().Model(&model.DriverCart{}).Where("driver_id = ?", id).First(&cart).Error
	if err != nil {
		logrus.Error("获取购物车信息失败", err)
		ct.Error(c, 200, "获取购物车信息失败")
		return
	}
	err = dal.Getdb().Model(&model.OrderProduct{}).Where("cart_refer = ?", cart.CartID).Order("factory_id").Find(&cart.Products).Error
	if err != nil {
		logrus.Error("获取购物车商品信息失败", err)
		ct.Error(c, 200, "获取购物车商品信息失败")
		return
	}
	c.JSON(200, response.CartResponse{response.Response{200, "获取购物车信息成功"}, cart})
}
