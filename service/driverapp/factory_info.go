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

type FactoryController struct {
	BaseDrController
}

// @Summary 场站信息

// @Description 按照场站距离展示数据，距离近的排名靠前
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param locInfo body model.Geo true "传入进行该操作时的司机地理位置信息，Address为可选项,其余为必填项"
// @Success 200 {object} response.FactoryInfoResponse
// @Failure 400 {onject} response.Response
// @Router /dr/factory [post]
func (i *FactoryController) FactoryOverview(c *gin.Context) {
	var locinfo model.Geo
	err := c.ShouldBind(&locinfo)
	if err != nil || locinfo.Longitude == "" || locinfo.Latitude == "" {
		i.Error(c, 400, "地理信息获取失败,请传入正确的地理信息")
		return
	}
	fmt.Println(locinfo)
	//可能不需要获取车主信息
	//iadmin, ok := c.Get(middleware.DRADMIN)
	//if ok != true {
	//	i.Error(c, 400, "获取用户信息失败")
	//	return
	//}
	//admin, ok := iadmin.(model.Driver)
	//if ok != true {
	//	i.Error(c, 400, "获取车主信息失败")
	//	return
	//}

	rdb := dal.Getrdb()
	db := dal.Getdb()
	ctx := rdb.Context()
	ires, err := utils.LocRadiusWithDist(ctx, rdb, utils.LOCATION, locinfo.Longitude, locinfo.Latitude, "10", "km")
	if err != nil {
		i.Error(c, 400, "附近场站信息获取失败,请传入正确的地理信息")
		return
	}
	//t.Log(res, len(res.([]interface{})))
	res := ires.([]interface{})
	n := len(res)
	views := make([]response.FactoryInfo, n)
	for k, iv := range res {
		v := iv.([]interface{})
		views[k].FactoryName = v[0].(string)
		views[k].Distance = v[1].(string)
		err := db.Raw("select pic,name from factory_products where factory_name = ?", views[k].FactoryName).Limit(5).Find(&views[k].ProductViews).Error
		if err != gorm.ErrRecordNotFound && err != nil {
			logrus.Info("获取"+views[k].FactoryName+"的商品信息失败", err)
			continue
		}
		if len(views[k].ProductViews) == 0 {
			defaultfp := response.FactoryProductOverview{"void", "void"}
			views[k].ProductViews = []response.FactoryProductOverview{defaultfp}
		}
		fmt.Println(views[k].ProductViews)
	}
	c.JSON(200, response.FactoryInfoResponse{
		response.Response{200, "成功获取附近场站信息"},
		views,
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
// @Success 201 {object} response.OrderResponse
// @Failure 400 {onject} response.Response
// @Router /dr/order/submit [post]
func (i *FactoryController) Order(c *gin.Context) {
	//var orderform model.DriverOrderForm
	//id := strconv.FormatInt(utils.IDWorker.NextId(), 10)
	var OrderForm model.DriverOrderForm
	c.ShouldBind(&OrderForm)
	OrderForm.State = 0
	//var Prdoucts []model.OrderProduct

	c.JSON(200, response.DriverOrdersResponse{
		response.Response{200, "订单提交成功"},
		[]model.DriverOrderForm{OrderForm},
	})
	c.Next()
}

// @Summary 补货订单结算
// @Description 结算
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param OrderForm body model.DriverOrderForm true "传入订单信息"
// @Success 201 {object} response.PayResponse
// @Failure 400 {onject} response.Response
// @Router /dr/order/pay [put]
func (i *FactoryController) Pay(c *gin.Context) {
	//id, _ := c.GetInventory("id")
	//TODO:业务逻辑
	//fmt.Println(id)
	c.JSON(200, response.PayResponse{
		response.Response{201, "支付成功"},
	})

}
