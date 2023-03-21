package platform

import (
	"buyfree/dal"
	"buyfree/repo/gen"
	"buyfree/repo/model"
	"buyfree/service/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	BaseController
}

func (o *OrderController) GetFactoryOrders(c *gin.Context) {
	var dofs []*model.DriverOrderForm
	err := dal.Getdb().Model(&model.DriverOrderForm{}).Find(&dofs).Error
	if err != nil {
		o.Error(c, 400, "获取订单信息失败 1")
	}
	n := len(dofs)
	fmt.Printf("获取到%d条订单信息\n", n)
	ords := []response.OrderInfostruct{}
	for i := 0; i < n; i++ {

		//fmt.Println(dofs[i].OrderID)
		//dal.Getdb().Raw("select * from order_products as op where op.order_refer = (select order_id from driver_order_forms where order_id =?)",dofs[i].OrderID).Find(&products)
		products, err := gen.OrderProduct.GetAllOrderProductReferDOrder(dofs[i].OrderID)
		if err != nil {
			o.Error(c, 200, "获取订单信息失败 2")
			return
		}
		k := len(products)
		fmt.Printf("获取到%d条货品信息\n", k)
		//fmt.Println(products)
		factoryname := dofs[i].FactoryName
		for j := 0; j < k; j++ {
			var info response.OrderInfostruct
			//info[j].FactoryName = factoryname
			info.FactoryName = factoryname
			info.Name = products[j].Name
			info.Sku = products[j].Sku
			info.Pic = products[j].Pic
			info.Type = products[j].Type
			//TODO:展示在首页和上架就交给前端吧
			//info[j].State=
			//
			saleinfo, _ := gen.FactoryProduct.GetBySkuAndFName(info.Sku, info.FactoryName)
			info.Sales = saleinfo.TotalSales
			info.Inventory = saleinfo.Inventory
			//if products[j].OrderRefer == dofs[i].OrderID { //TODO:改为判断其他状态
			ords = append(ords, info)
			//}
		}
		//ords = append(ords, info)
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
		o.Error(c, 400, "获取订单信息失败 3")
	}
	c.Set("Orders", ords)
	c.Next()
}

func (o *OrderController) GetOnShelf(c *gin.Context) {
	ctxInfo, _ := c.Get("Orders")
	orders := ctxInfo.([]*response.OrderInfostruct)
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
	orders := ctxInfo.([]*response.OrderInfostruct)
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
func (o *OrderController) Getdownshelf(c *gin.Context) {
	ctxInfo, _ := c.Get("Orders")
	orders := ctxInfo.([]*response.OrderInfostruct)
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

func (o *OrderController) GetGoodinfo(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func (o *OrderController) TakeOn(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
