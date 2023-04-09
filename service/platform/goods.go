package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
// @Param factory_name path string false "根据场站名字获取场站所有商品信息，默认获取所有商品信息"
// @Param mode path int true "按照不同模式获取库存商品信息，mode={0:未上架,1:上架,传入其他数据获取所有商品信息}"
// @Success 200 {object} response.FactoryProductsResponse
// @Failure 400 {object} response.Response
// @Router /pt/products/{mode}/factory/{factory_name}/ [get]
func (o *GoodsController) PGetAllProducts(c *gin.Context) {
	fname := c.Param("factory_name")
	mode := c.Param("mode")
	var sf bool
	if mode == "1" {
		sf = true
	} else if mode == "0" {
		sf = false
	}
	var infos []response.FactoryProductsInfo
	if len(fname) != 0 && (mode == "0" || mode == "1") {
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
	} else if mode == "0" || mode == "1" {
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

// @Summary  平台获取场站获取商品信息
// @Description 传如对应场站名以及商品名
// @Tags Products
// @Accept json
// @Accept mpfd
// @Produce json
// @Param factory_name path string true "场站名"
// @Param product_name path string true "商品名"
// @Success 200 {object} response.FactoryGoodsResponse
// @Failure 400 {object} response.Response
// @Router /pt/products/infos/{factory_name}/{product_name} [get]
func (o *GoodsController) PGetGoodsInfo(c *gin.Context) {
	factoryName := c.Param("factory_name")
	productName := c.Param("product_name")
	var product model.FactoryProduct
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("factory_name = ? and name = ?", factoryName, productName).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		logrus.Info(err)
		o.Error(c, 404, "不存在该商品，请输入正确的信息")
		return
	} else if err != nil {
		logrus.Info(err)
		o.Error(c, 404, "查询失败")
		return
	}
	c.JSON(200, response.FactoryGoodsResponse{
		response.Response{200, "成功获取对应信息"},
		product,
	})
}

// @Summary  调整场站商品上下架状态
// @Description 传入场站名与商品名
// @Tags Products
// @Accept json
// @Accept mpfd
// @Produce json
// @Param Info body response.UnionNameInfo true "传入factory_name + product_name 联合主键"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /pt/products/turn [patch]
func (o *GoodsController) TurnOver(c *gin.Context) {
	var info response.UnionNameInfo
	err := c.ShouldBind(&info)
	if err != nil {
		o.Error(c, 400, "传入信息格式有误")
		return
	}
	var pr model.FactoryProduct
	err = dal.Getdb().Model(&model.FactoryProduct{}).Where("factory_name = ? and name = ?", info.FactoryName, info.ProductName).UpdateColumn("is_on_shelf", gorm.Expr("not is_on_shelf")).First(&pr).Error
	var msg string
	if err == nil {
		if pr.IsOnShelf == true {
			msg = "商品上架成功"
		} else {
			msg = "商品下架成功"
		}
		c.JSON(200, response.Response{
			200,
			msg,
		})
	} else {
		logrus.Info(err)
		o.Error(c, 400, "调整商品上/下架信息失败")
	}
}
