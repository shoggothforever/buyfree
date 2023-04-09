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
// @Success 201 {object} response.CartResponse
// @Failure 400 {object} response.Response
// @Router /dr/factory/cart [get]
func (ct *CartController) OpenCart(c *gin.Context) {
	//TODO:可以获取
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		ct.Error(c, 400, "获取车主信息失败")
		return
	}
	var cart model.DriverCart
	err := dal.Getdb().Model(&model.DriverCart{}).Where("driver_id = ?", admin.ID).First(&cart).Error
	if err != nil {
		logrus.Error("获取购物车信息失败", err)
		ct.Error(c, 200, "获取购物车信息失败")
		return
	}
	type nid struct {
		Name string `json:"name,omitempty"`
		ID   int64  `json:"id,omitempty"`
	}
	var nis []nid
	err = dal.Getdb().Raw("select name,id  from factories where id in (select distinct factory_id from order_products where cart_refer = ?)", cart.CartID).Find(&nis).Error
	if err != nil {
		logrus.Error("获取场站信息失败", err)
		ct.Error(c, 200, "获取场站信息失败")
		return
	}
	var gs []*response.CartGroup
	tx := dal.Getdb()
	var sum float64 = 0
	var cnt int64 = 0
	for _, v := range nis {
		var g response.CartGroup
		g.DistanceInfo.FactoryName = v.Name
		g.DistanceInfo.FactoryID = v.ID
		err = tx.Raw("select name,pic,type,count,price from order_products where cart_refer = ? and factory_id = ?", cart.CartID, v.ID).Find(&g.ProductDetails).Error
		if err == nil {
			gs = append(gs, &g)
			for _, gv := range g.ProductDetails {
				cnt++
				sum += float64(gv.Count) * gv.Price
			}
			//fmt.Println(gs)
		} else if err != gorm.ErrRecordNotFound {
			fmt.Println(err)
			ct.Error(c, 400, "获取购物车商品信息失败")
			return
		}
	}
	err = dal.Getdb().Model(&model.DriverCart{}).Where("cart_id = ?", cart.CartID).Updates(map[string]interface{}{
		"total_count":  cnt,
		"total_amount": sum,
	}).Error
	if err != nil {
		logrus.Error("获取购物车信息失败", err)
		ct.Error(c, 200, "获取购物车信息失败")
		return
	} else {
		c.JSON(200, response.CartResponse{
			response.Response{200, "获取购物车信息成功"},
			sum,
			cnt,
			gs})
	}
}
