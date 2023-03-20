package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	//var dofs []*model.DriverOrderForm
	//
	//err := dal.Getdb().Model(&model.DriverOrderForm{}).Find(&dofs).Error
	//n := len(dofs)
	////ords := []*model.OrderProduct
	//for i := 0; i < n; i++ {
	//	products, err := gen.OrderProduct.GetAllOrderProductReferDOrder(dofs[i].OrderID)
	//	if err != nil {
	//		c.JSON(200, response.Response{
	//			400,
	//			"获取订单商品信息失败",
	//		})
	//		return
	//	}
	//	k := len(products)
	//	pros := make([]*model.OrderProduct, k)
	//	for j := 0; j < k; j++ {
	//		pros[j].Name = products[j].Name
	//	}
	//	//ords=append(ords,pros)
	//}
	//if err == nil {
	//	//c.JSON(200, response.OrderResponse{
	//	//	response.Response{
	//	//		200,
	//	//		"成功获取所有订单信息",
	//	//	},
	//	//	//ords,
	//	//})
	//} else {
	//	c.JSON(200, response.Response{
	//		400,
	//		"获取订单信息失败",
	//	})
	//}
}

func GetOnShelf(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func Getsoldout(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func Getdownshelf(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetGoodinfo(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func TakeOn(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
