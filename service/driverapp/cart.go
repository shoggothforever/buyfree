package driverapp

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CartController struct {
	BaseDrController
}

// @Summary 购物车界面
// @Description "传入查看附近场站信息获取到的所有场站距离信息，要打开购物车必须先进入场站界面“
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param DistanceInfos body response.FactoryDistanceInfos false "附近场站信息，已经获取了，打包后直接传入"
// @Success 201 {object} response.CartResponse
// @Failure 400 {object} response.Response
// @Router /dr/factory/cart [post]
func (ct *CartController) OpenCart(c *gin.Context) {
	//TODO:可以获取
	var infos []response.FactoryDistanceInfo
	err := c.ShouldBind(&infos)
	if err != nil {
		ct.Error(c, 400, "获取附近场站信息失败")
		return
	}
	//fmt.Println(infos)
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		ct.Error(c, 400, "获取车主信息失败")
		return
	}

	var cart model.DriverCart
	err = dal.Getdb().Model(&model.DriverCart{}).Where("driver_id = ?", admin.ID).First(&cart).Error
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
	//fmt.Println(cart, err)
	var gs []*response.CartGroup

	var fid int64
	tx := dal.Getdb()
	var sum float64 = 0
	var cnt int64 = 0
	for _, v := range infos {
		var g response.CartGroup
		g.DistanceInfo.Distance = v.Distance
		g.DistanceInfo.FactoryName = v.FactoryName
		if err := tx.Model(&model.Factory{}).Select("id").Where("name = ?", v.FactoryName).First(&fid).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			} else {
				fmt.Println(err)
				ct.Error(c, 400, "获取购物车商品信息失败")
				return
			}
		} else {
			fmt.Println("场站编号是", fid)
			g.DistanceInfo.FactoryID = fid
			err = tx.Raw("select name,pic,type,count,price from order_products where cart_refer = ? and factory_id = ?", cart.CartID, fid).Find(&g.ProductDetails).Error
			if err == nil {
				gs = append(gs, &g)
				for _, gv := range g.ProductDetails {
					cnt++
					sum += float64(gv.Count) * gv.Price
				}
				fmt.Println(gs)
			} else if err != gorm.ErrRecordNotFound {
				fmt.Println(err)
				ct.Error(c, 400, "获取购物车商品信息失败")
				return
			}
		}
	}
	dal.Getdb().Model(&model.DriverCart{}).Where("cart_id = ?", cart.CartID).Updates(map[string]interface{}{
		"total_count":  cnt,
		"total_amount": sum,
	})
	c.JSON(200, response.CartResponse{
		response.Response{200, "获取购物车信息成功"},
		gs})
}
