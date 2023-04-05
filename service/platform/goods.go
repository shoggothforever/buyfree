package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GoodsController struct {
	BasePtController
}

// @Summary 获取所有商品
// @Description	传入场站名，获取该场站所有商品信息
// @Tags	Products
// @Accept json
// @Accept mpfd
// @Produce json
// @Param factory_name path string true "根据场站名字获取场站所有商品信息，默认获取所有商品信息"
// @Param mode path int true "按照不同模式获取订单信息，mode={0:未上架,1:上架,传入其他数据获取所有商品信息}"
// @Success 200 {object} response.FactoryProductsResponse
// @Failure 400 {object} response.Response
// @Router /pt/products/{mode}/factory/{factory_name}/ [get]
func (o *GoodsController) GetAllProducts(c *gin.Context) {
	fname := c.Param("factory_name")
	mode := c.Param("mode")
	var sf bool
	if mode == "1" {
		sf = true
	} else {
		sf = false
	}
	var infos []response.FactoryProductsInfo
	if len(fname) != 0 && len(mode) != 0 {
		err := dal.Getdb().Raw("select * from factory_products where factory_name = ? and is_on_shelf = ?", fname, sf).Find(&infos).Error
		if err != nil {
			o.Error(c, 400, "获取商品信息失败")
			return
		}
	} else if len(fname) != 0 {
		err := dal.Getdb().Raw("select * from factory_products where factory_name = ?", fname).Find(&infos).Error
		if err != nil {
			o.Error(c, 400, "获取商品信息失败")
			return
		}
	} else if len(mode) != 0 {
		err := dal.Getdb().Raw("select * from factory_products where is_on_shelf = ?", sf).Find(&infos).Error
		if err != nil {
			o.Error(c, 400, "获取商品信息失败")
			return
		}
	} else {
		err := dal.Getdb().Model(&model.FactoryProduct{}).Find(&infos).Error
		if err != nil {
			o.Error(c, 200, "获取商品信息失败")
		}
	}
	//n := len(infos)
	//infos := make([]response.FactoryOrderInfo, n)
	//for j := 0; j < n; j++ {
	//	factoryname := products[j].FactoryName
	//	infos[j].FactoryName = factoryname
	//	infos[j].Name = products[j].Name
	//	infos[j].Sku = products[j].Sku
	//	infos[j].Pic = products[j].Pic
	//	infos[j].Type = products[j].Type
	//	//TODO:展示在首页和上架就交给前端吧,获取订单中的商品在场站的上下架状态，根据factoryID 和 商品SKU在场站的商品表中查询对应的状态信息
	//	infos[j].IsOnShelf = products[j].IsOnShelf
	//	infos[j].TotalSales = products[j].TotalSales
	//	infos[j].Inventory = products[j].Inventory
	//}
	if len(infos) != 0 {
		c.JSON(200, response.FactoryProductsResponse{
			response.Response{
				200,
				"成功获取所有商品信息",
			},
			infos,
		})
	} else {
		c.JSON(200, response.FactoryProductsResponse{
			response.Response{
				200,
				"暂无相关商品信息",
			},
			infos,
		})
	}
	c.Set("Orders", infos)
	c.Next()
}

// @Summary  获取商品信息
// @Description 输入商品SKU,获取场站中对应商品的详细信息
// @Tags Products
// @Accept json
// @Accept mpfd
// @Produce json
// @Param sku path string true "sku 指向唯一的场站中的商品信息"
// @Success 200 {object} response.FactoryGoodsResponse
// @Failure 400 {object} response.Response
// @Router /pt/products/infos/{sku} [get]
func (o *GoodsController) GetGoodsInfo(c *gin.Context) {
	sku := c.Param("sku")
	var product model.FactoryProduct
	//todo:要么确定sku的获取策略，要么就全用id代替
	//err := dal.Getdb().Model(&model.FactoryProduct{}).Where("id = ?", sku).First(&product).Error
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("sku = ?", sku).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		o.Error(c, 404, "不存在该商品，请输入正确的信息")
		return
	} else if err != nil {
		o.Error(c, 404, "查询失败")
		return
	}
	c.JSON(200, response.FactoryGoodsResponse{
		response.Response{200, "成功获取对应信息"},
		product,
	})
}

// @Summary  上架场站商品
// @Description 输入商品id,获取场站中对应商品的详细信息
// @Tags Products
// @Accept json
// @Accept mpfd
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} response.FactoryGoodsResponse
// @Failure 400 {object} response.Response
// @Router /pt/products/on/{id} [patch]
func (o *GoodsController) OnShelfGoods(c *gin.Context) {

	id := c.Param("id")
	var pr model.FactoryProduct
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("id = ?", id).Update("is_on_shelf", true).First(&pr).Error
	if err == nil {
		c.JSON(200, response.FactoryGoodsResponse{
			response.Response{
				200,
				"商品上架成功",
			}, pr,
		})
	} else {
		o.Error(c, 400, "商品上架失败")
	}
}

// @Summary  下架场站商品
// @Description 输入商品id,获取场站中对应商品的详细信息
// @Tags Products
// @Accept json
// @Accept mpfd
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} response.FactoryGoodsResponse
// @Failure 400 {object} response.Response
// @Router /pt/products/down/{id} [patch]
func (o *GoodsController) DownShelfGoods(c *gin.Context) {

	id := c.Param("id")
	var pr model.FactoryProduct
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("id = ?", id).Update("is_on_shelf", false).First(&pr).Error
	if err == nil {
		c.JSON(200, response.FactoryGoodsResponse{
			response.Response{
				200,
				"商品下架成功",
			}, pr,
		})
	} else {
		o.Error(c, 400, "商品下架失败")
	}
}
