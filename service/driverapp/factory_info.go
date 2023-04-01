package driverapp

import (
	"buyfree/dal"
	"buyfree/middleware"
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
	iadmin, ok := c.Get(middleware.DRADMIN)
	if ok != true {
		i.Error(c, 400, "获取用户信息失败")
		return
	}
	admin, ok := iadmin.(model.Driver)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	rdb := dal.Getrdb()
	db := dal.Getdb()
	ctx := rdb.Context()
	//更新车主位置信息
	rdb.Do(ctx, "geoadd", utils.LOCATION, locinfo.Longitude, locinfo.Latitude, admin.CarID)
	ires, err := utils.LocRadiusWithDist(ctx, rdb, utils.LOCATION, locinfo.Longitude, locinfo.Latitude, "10", "km")
	if err != nil {
		i.Error(c, 400, "附近场站信息获取失败,请传入正确的地理信息")
		return
	}
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
// @Description 传入场站名字和距离获取场站具体信息
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param distance_info body response.FactoryDistanceInfo true "对应场站的名字和距离"
// @Success 200 {object} response.FactoryDetailResponse
// @Failure 400 {onject} response.Response
// @Router /dr/factory/infos [post]
func (i *FactoryController) Detail(c *gin.Context) {
	iadmin, ok := c.Get(middleware.DRADMIN)
	if ok != true {
		i.Error(c, 400, "获取用户信息失败")
		return
	}
	admin, ok := iadmin.(model.Driver)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	var disinfo response.FactoryDistanceInfo
	err := c.ShouldBind(&disinfo)
	fmt.Println(disinfo)
	if err != nil || disinfo.FactoryName == "" {
		i.Error(c, 400, "场站名和距离获取失败")
		return
	}
	var fa response.FactoryDetail
	{
		err = dal.Getdb().Model(&model.Factory{}).Select("address", "description").Where("name=?", disinfo.FactoryName).First(&fa).Error
		if err != nil {
			logrus.Info("获取场站信息失败", err)
			i.Error(c, 400, "无法获取场站信息")
			return
		}
	}
	var details []*response.FactoryProductDetail
	{
		err = dal.Getdb().Raw("select fp.name,inventory,"+
			"dv.m_inventory,pic,type,monthly_sales,supply_price  "+
			"from factory_products as fp,(select dp.name,sum(dp.inventory)"+
			" as m_inventory from device_products dp where device_id in"+
			"(select id from devices where owner_id =?) and dp.name in "+
			"(select name from factory_products where factory_name=?)"+
			"group by (dp.name)) as dv where is_on_shelf=true "+
			"and factory_name='cat' and fp.name=dv.name", admin.ID, fa.Name).Find(&details).Error
		if err != nil {
			logrus.Info("获取场站商品信息失败", err)
			i.Error(c, 400, "无法获取商品场站信息")
			return
		}
	}

	c.JSON(200, response.FactoryDetailResponse{
		response.Response{200, "成功获取该场站商品信息"},
		disinfo,
		fa,
		details,
	})

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
