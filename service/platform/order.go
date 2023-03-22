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
	BaseController
}

// @Summary 获取场站订单信息(车主在该场站下的订单)
// @Description	传入字段mode，获取对应订单信息
// @Tags	Orderform
// @Accept json
// @Accept mpfd
// @Produce json
// @Param mode path int true "按照不同模式获取订单信息，mode={0:未支付,1:未完成,2:完成,传入其他任意数值代表获取全部订单信息}"
// @Success 200 {object} response.OrderResponse
// @Failure 400 {object} response.Response
// @Router /pt/orders/factory/{mode} [get]
func (o *OrderController) GetFactoryOrders(c *gin.Context) {
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
	ords := []response.FactoryOrderInfo{}
	for i := 0; i < n; i++ {
		products, err := gen.OrderProduct.GetAllOrderProductReferDOrder(dofs[i].OrderID)
		if err != nil {
			o.Error(c, 400, "获取订单信息失败 2")
			return
		}
		k := len(products)
		fmt.Printf("获取到%d条货品信息\n", k)
		factoryname := dofs[i].FactoryName
		infos := make([]response.FactoryOrderInfo, k)
		for j := 0; j < k; j++ {
			var info response.FactoryOrderInfo
			infos[j].FactoryName = factoryname
			infos[j].Name = products[j].Name
			infos[j].Sku = products[j].Sku
			infos[j].Pic = products[j].Pic
			infos[j].Type = products[j].Type
			//TODO:展示在首页和上架就交给前端吧,获取订单中的商品在场站的上下架状态，根据factoryID 和 商品SKU在场站的商品表中查询对应的状态信息
			//infos[j].State = products[j].IsChosen
			saleinfo, _ := gen.FactoryProduct.GetBySkuAndFName(info.Sku, info.FactoryName)
			infos[j].Sales = saleinfo.TotalSales
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

func (o *OrderController) GetOnShelf(c *gin.Context) {
	ctxInfo, _ := c.Get("Orders")
	orders := ctxInfo.([]*response.FactoryOrderInfo)
	var ords []*response.OrderResponse
	//TODO:交给前端写吧
	n := len(ords)
	if n != 0 {
		for i := 0; i < n; i++ {
			fmt.Println(orders)
		}
	}
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func (o *OrderController) Getsoldout(c *gin.Context) {
	ctxInfo, _ := c.Get("Orders")
	orders := ctxInfo.([]*response.FactoryOrderInfo)
	var ords []*response.OrderResponse
	//TODO:交给前端写吧
	n := len(ords)
	if n != 0 {
		for i := 0; i < n; i++ {
			fmt.Println(orders)
		}
	}
	c.JSON(200, response.Response{
		200,
		"ok"})
}

// @Summary  获取商品信息
// @Description 输入商品SKU,获取场站中对应商品的详细信息
// @Tags Orderform
// @Accept json
// @Accept mpfd
// @Produce json
// @Param sku path string true "sku 指向唯一的场站中的商品信息"
// @Success 200 {object} response.FactoryGoodsResponse
// @Failure 400 {object} response.Response
// @Router /pt/orders/infos/{sku} [get]
func (o *OrderController) GetGoodinfo(c *gin.Context) {
	//TODO:交给前端吧
	sku := c.Param("sku")
	var product model.FactoryProduct
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("sku = ?", sku).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		o.Error(c, 404, "不存在该商品，请输入正确的信息")
	} else if err != nil {
		o.Error(c, 404, "查询失败")
		return
	}
	//c.JSON(200, gin.H{
	//	"Code":    200,
	//	"Msg":     "成功获取对应信息",
	//	"Product": product,
	//})
	c.JSON(200, response.FactoryGoodsResponse{
		response.Response{200, "成功获取对应信息"},
		product,
	})
}

func (o *OrderController) ModifyGoods(c *gin.Context) {
	//TODO:交给前端吧
	//mode := c.Param("mode")
	c.JSON(200, response.Response{
		200,
		"ok"})
}
