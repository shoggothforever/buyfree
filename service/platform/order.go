package platform

import (
	"buyfree/dal"
	"buyfree/repo/gen"
	"buyfree/repo/model"
	"buyfree/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	BasePtController
}

// @Summary 获取车主订单信息(车主在该场站下的订单)
// @Description	传入字段mode，获取对应订单信息
// @Tags	Orderform
// @Accept json
// @Accept mpfd
// @Produce json
// @Param mode path int true "按照不同模式获取订单信息，mode={0:未支付,1:未完成,2:完成,传入其他任意数值代表获取全部订单信息}"
// @Success 200 {object} response.OrderResponse
// @Failure 400 {object} response.Response
// @Router /pt/orders/factory/{mode} [get]
func (o *OrderController) GetDriverOrders(c *gin.Context) {
	//mode =2-已完成 1-待取货 0-未支付 else 全部
	mode := c.Param("mode")
	var dofs []*model.DriverOrderForm
	if mode == "0" || mode == "1" || mode == "2" {
		err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("state = ?", mode).Find(&dofs).Error
		if err != nil {
			o.Error(c, 400, "获取订单信息失败 1")
			return
		}
	} else {
		err := dal.Getdb().Model(&model.DriverOrderForm{}).Find(&dofs).Error
		if err != nil {
			o.Error(c, 400, "获取订单信息失败 1")
			return
		}
	}
	n := len(dofs)
	fmt.Printf("获取到%d条订单信息\n", n)
	ords := []response.FactoryProductsInfo{}
	for i := 0; i < n; i++ {
		products, err := gen.OrderProduct.GetAllOrderProductReferDOrder(dofs[i].OrderID)
		if err != nil {
			o.Error(c, 400, "获取订单信息失败 2")
			return
		}
		k := len(products)
		fmt.Printf("获取到%d条货品信息\n", k)
		factoryname := dofs[i].FactoryName
		infos := make([]response.FactoryProductsInfo, k)
		for j := 0; j < k; j++ {
			var info response.FactoryProductsInfo
			infos[j].FactoryName = factoryname
			infos[j].Name = products[j].Name
			infos[j].Sku = products[j].Sku
			infos[j].Pic = products[j].Pic
			infos[j].Type = products[j].Type
			//TODO:展示在首页和上架就交给前端吧,获取订单中的商品在场站的上下架状态，根据factoryID 和 商品SKU在场站的商品表中查询对应的状态信息
			infos[j].IsOnShelf = products[j].IsChosen
			//saleinfo, _ := gen.FactoryProduct.GetBySkuAndFName(info.Sku, info.FactoryName)
			var saleinfo model.FactoryProduct
			err := dal.Getdb().Model(&model.FactoryProduct{}).Select("total_sales").Where("sku = ? and factory_name = ?", info.Sku, info.FactoryName).First(&saleinfo.TotalSales).Error
			if err != gorm.ErrRecordNotFound && err != nil {
				o.Error(c, 400, "获取订单信息失败 2")
				return
			}
			infos[j].TotalSales = saleinfo.TotalSales
			infos[j].Inventory = saleinfo.Inventory
		}
		ords = append(ords, infos...)
	}
	if len(ords) != 0 {
		c.JSON(200, response.OrderResponse{
			response.Response{
				200,
				"成功获取所有订单信息",
			},
			ords,
		})
	} else {
		c.JSON(200, response.OrderResponse{
			response.Response{
				200,
				"暂无相关订单信息",
			},
			ords,
		})
	}
	c.Set("Orders", ords)
	c.Next()
}
