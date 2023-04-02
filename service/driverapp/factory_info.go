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
	"gorm.io/gorm/clause"
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
// @Failure 400 {object} response.Response
// @Router /dr/factory [post]
func (i *FactoryController) FactoryOverview(c *gin.Context) {
	var locinfo model.Geo
	err := c.ShouldBind(&locinfo)
	if err != nil || locinfo.Longitude == "" || locinfo.Latitude == "" {
		i.Error(c, 400, "地理信息获取失败,请传入正确的地理信息")
		return
	}
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	rdb := dal.Getrdb()
	db := dal.Getdb()
	ctx := rdb.Context()
	//更新车主位置信息
	rdb.Do(ctx, "geoadd", utils.DRIVERLOCATION, locinfo.Longitude, locinfo.Latitude, admin.CarID)
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
// @Param distance_info body response.FactoryDistanceReq true "对应场站的名字和距离"
// @Success 200 {object} response.FactoryDetailResponse
// @Failure 400 {object} response.Response
// @Router /dr/factory/infos [post]
func (i *FactoryController) Detail(c *gin.Context) {
	var disinfo response.FactoryDistanceReq
	err := c.ShouldBind(&disinfo)
	fmt.Println(disinfo)
	if err != nil || disinfo.FactoryName == "" {
		i.Error(c, 400, "场站名和距离获取失败")
		return
	}
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	var fa response.FactoryDetail
	{
		err = dal.Getdb().Model(&model.Factory{}).Select("id", "address", "description").Where("name=?", disinfo.FactoryName).First(&fa).Error
		if err != nil {
			logrus.Info("获取场站信息失败", err)
			i.Error(c, 400, "无法获取场站信息")
			return
		}
	}
	fmt.Println(fa)
	var details []*response.FactoryProductDetail
	{
		err = dal.Getdb().Raw("select fp.name,inventory,"+
			"dv.m_inventory,pic,type,monthly_sales,supply_price  "+
			"from factory_products as fp,(select dp.name,sum(dp.inventory)"+
			" as m_inventory from device_products dp where device_id in"+
			"(select id from devices where owner_id =?) and dp.name in "+
			"(select name from factory_products where factory_name=?)"+
			"group by (dp.name)) as dv where is_on_shelf=true "+
			"and factory_name=? and fp.name=dv.name", admin.ID, disinfo.FactoryName, disinfo.FactoryName).Find(&details).Error
		//fmt.Println(details)
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

// @Summary 添加/移除出购物车
// @Description 点击事件:传入商品信息,修改商品在购物车中的数量，选项(1：增加1件，-1：减少1件，0：清空）
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param productInfo body response.ReplenishInfo true "传入场站名以及ID，商品名称、型号、价格、选购数量、图片信息"
// @Success 200 {object} response.Response
// @Failure 200 {object} response.Response
// @Router /dr/order/replenish [post]
func (i *FactoryController) Modify(c *gin.Context) {
	var info response.ReplenishInfo
	err := c.ShouldBind(&info)
	fmt.Println("获取到的记录是", info)
	if err != nil {
		logrus.Info("传入信息错误", err)
		i.Error(c, 400, "获取传入信息失败")
		return
	}
	var cartrefer int64
	admin, ok := utils.GetDriveInfo(c)
	if !ok {
		i.Error(c, 400, "获取用户信息失败")
		return
	}
	//获取购物车编号
	err = dal.Getdb().Model(&model.DriverCart{}).Select("cart_id").Where("driver_id = ?", admin.ID).First(&cartrefer).Error
	fmt.Println("购物车编号", cartrefer)
	if err != nil {
		logrus.Info("获取购物车信息失败", err)
		i.Error(c, 400, "获取购物车信息失败")
		return
	}
	var op model.OrderProduct

	err = dal.Getdb().Model(&model.OrderProduct{}).Where("name= ? and factory_id =? and cart_refer = ?", info.ProductName, info.FactoryID, cartrefer).First(&op).Error
	fmt.Println(op, err)
	if err == gorm.ErrRecordNotFound {
		fmt.Println("获取到的记录是", info)
		if info.Count < 0 {
			i.Error(c, 403, "未定义操作")
			return
		}
		op.Name = info.ProductName
		op.Pic = info.Pic
		op.Count = info.Count
		op.Price = info.Price
		op.Type = info.Type
		op.FactoryID = info.FactoryID
		op.IsChosen = true
		op.CartRefer = cartrefer
		err = dal.Getdb().Create(&op).Error
		if err != nil {
			fmt.Println(err)
			i.Error(c, 403, "添加商品信息失败")
			return
		}
	} else if err != nil {
		i.Error(c, 400, "操作数据库失败")
		return
	} else {
		var terr error
		dal.Getdb().Transaction(func(tx *gorm.DB) error {
			if op.Count == 0 {
				terr = dal.Getdb().Model(&model.OrderProduct{}).Where("name= ? and factory_id =? and cart_refer = ?", info.ProductName, info.FactoryID, cartrefer).Update("count", 0).Error
				fmt.Println(terr)
				return terr
			} else if op.Count+info.Count > 0 {
				terr = dal.Getdb().Model(&model.OrderProduct{}).Where("name= ? and factory_id =? and cart_refer = ?", info.ProductName, info.FactoryID, cartrefer).UpdateColumn("count", gorm.Expr("count + ?", info.Count)).Error
				fmt.Println(terr)
				return terr
			} else if op.Count+info.Count <= 0 {
				terr = dal.Getdb().Model(&model.OrderProduct{}).Where("name= ? and factory_id =? and cart_refer = ?", info.ProductName, info.FactoryID, cartrefer).Delete(&op).Error
			}
			return nil
		})
		if terr != nil {
			fmt.Println(terr)
			i.Error(c, 400, "更新购物车信息失败失败")
			return
		}
	}
	if err != nil {
		c.JSON(200, response.Response{400, "操作数据库失败"})
	} else {
		c.JSON(200, response.Response{200, "成功添加/减少商品"})
	}
}

// @Summary 选中购物车中的商品
// @Description 点击事件:选中商品，倒置商品选中状态
// @Tags Driver/Replenish
// @Accept json
// @Produce json
// @Param productInfo body response.ReplenishInfo true "传入场站名,商品名称"
// @Success 200 {object} response.Response
// @Failure 200 {object} response.Response
// @Router /dr/order/choose [put]
func (i *FactoryController) Choose(c *gin.Context) {
	var info response.ReplenishInfo
	err := c.ShouldBind(&info)
	fmt.Println("获取到的记录是", info)
	if err != nil {
		logrus.Info("传入信息错误", err)
		i.Error(c, 400, "获取传入信息失败")
		return
	}
	var cartrefer, fid int64
	admin, ok := utils.GetDriveInfo(c)
	if !ok {
		i.Error(c, 400, "获取用户信息失败")
		return
	}
	//获取购物车编号
	err = dal.Getdb().Model(&model.DriverCart{}).Select("cart_id").Where("driver_id = ?", admin.ID).First(&cartrefer).Error
	fmt.Println("购物车编号", cartrefer)
	if err != nil {
		logrus.Info("获取购物车信息失败", err)
		i.Error(c, 400, "获取购物车信息失败原因：获取购物车id失败")
		return
	}
	if err = dal.Getdb().Model(&model.Factory{}).Select("id").Where("name = ?", info.FactoryName).First(&fid).Error; err != nil {
		if err != nil {
			fmt.Println(err)
			i.Error(c, 400, "操作购物车商品信息失败原因：获取场站id失败")
			return
		}
	}
	var op model.OrderProduct
	err = dal.Getdb().Model(&model.OrderProduct{}).Select("is_chosen").Where("name= ? and factory_id =? and cart_refer = ?", info.ProductName, fid, cartrefer).First(&op).Update("is_chosen", !op.IsChosen).Error
	fmt.Println(op, err)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			i.Error(c, 403, "传入数据有误")
			return
		}
		i.Error(c, 400, "操作数据库失败")
		return
	}
	if !op.IsChosen {
		c.JSON(200, response.Response{200, "成功选中商品"})
	} else {
		c.JSON(200, response.Response{200, "成功取消选中商品"})
	}
}

// @Summary 仅生成单个场站的订单订单信息
// @Description 使用选中的商品生成订单，从购物车界面跳转到提交订单界面（暂时为未支付状态，设置了30分钟的过期时间，需要等待服务端验签，用户支付完毕）
// @Tags Driver/Pay
// @Accept json
// @Produce json
// @Param DistanceInfos body response.FactoryDistanceReq true "包含附近场站信息，已经获取了,直接打包传入"
// @Success 201 {object} response.DriverOrdersResponse
// @Failure 400 {object} response.Response
// @Router /dr/order/submit [post]
func (i *FactoryController) Submit(c *gin.Context) {
	var finfo response.FactoryDistanceReq
	err := c.ShouldBind(&finfo)
	if err != nil {
		i.Error(c, 400, "获取位置信息失败")
		return
	}
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	var om model.DriverOrderForm
	om.OrderID = utils.IDWorker.NextId()
	var cartrefer, fid int64
	//获取购物车编号
	err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&model.DriverCart{}).Select("cart_id").Where("driver_id = ?", admin.ID).First(&cartrefer).Error
		if err != nil {
			fmt.Println(err)
			i.Error(c, 400, "获取购物车信息失败")
			return err
		}

		err = tx.Model(&model.Factory{}).Select("id").Where("name = ?", finfo.FactoryName).First(&fid).Error
		if err != nil {
			fmt.Println(err)
			i.Error(c, 400, "获取场站信息失败")
			return err
		}

		err = tx.Model(&model.OrderProduct{}).Where("cart_refer= ? and factory_id = ? and is_chosen = true", cartrefer, fid).Update("order_refer", om.OrderID).Find(&om.ProductInfos).Error
		if err != nil {
			fmt.Println(err)
			i.Error(c, 400, "获取商品信息失败")
			return err
		} else if len(om.ProductInfos) == 0 {
			i.Error(c, 400, "未选中商品信息")
			return gorm.ErrRecordNotFound
		}
		var ferr error
		for _, pv := range om.ProductInfos {
			ferr = tx.Model(&model.OrderProduct{}).Where("cart_refer = ? and factory_id = ? and name = ?", cartrefer, fid, pv.Name).Update("is_chosen", false).Error
			if ferr != nil {
				fmt.Println(ferr)
				return ferr
			}
		}
		return nil
	})

	if err == nil {
		var cost float64 = 0
		for _, v := range om.ProductInfos {
			cost += float64(v.Count) * v.Price
		}
		om.Set(fid, admin.ID, cost, finfo.FactoryName, admin.CarID, admin.Address)
		if cerr := dal.Getdb().Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "order_id"}},
			DoNothing: true},
		).Omit("ProductInfos").Create(&om).Error; cerr != nil {
			fmt.Println(cerr)
			dal.Getdb().Model(&model.DriverOrderForm{}).Delete(&om)
			i.Error(c, 400, "创建订单表失败")
			return
		}
		c.JSON(200, response.DriverOrdersResponse{
			Response:        response.Response{200, "订单提交成功"},
			FactoryDistance: []response.FactoryDistanceReq{finfo},
			OrderInfos:      []model.DriverOrderForm{om},
		})
	} else if err != gorm.ErrRecordNotFound {
		i.Error(c, 400, "订单提交失败")
	}
	c.Next()
}

// @Summary 生成多个场站的订单信息
// @Description 使用选中的商品生成订单，从购物车界面跳转到提交订单界面（暂时为未支付状态，设置了30分钟的过期时间，需要等待服务端验签，用户支付完毕）
// @Tags Driver/Pay
// @Accept json
// @Produce json
// @Param DistanceInfos body response.FactoryDistanceInfos false "附近场站信息，已经获取了，打包后直接传入"
// @Success 201 {object} response.DriverOrdersResponse
// @Failure 400 {object} response.Response
// @Router /dr/order/submit2 [post]
func (i *FactoryController) SubmitMany(c *gin.Context) {
	var finfos []response.FactoryDistanceReq
	err := c.ShouldBind(&finfos)
	fmt.Println(finfos)
	if err != nil {
		i.Error(c, 400, "获取位置信息失败")
		return
	}
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	var oms []model.DriverOrderForm
	var cartrefer int64
	//获取购物车编号
	err = dal.Getdb().Model(&model.DriverCart{}).Select("cart_id").Where("driver_id = ?", admin.ID).First(&cartrefer).Error
	if err != nil {
		fmt.Println(err)
		i.Error(c, 400, "获取购物车信息失败")
		return
	}
	for _, v := range finfos {
		err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
			var om model.DriverOrderForm
			om.OrderID = utils.IDWorker.NextId()
			var fid int64
			err = tx.Model(&model.Factory{}).Select("id").Where("name = ?", v.FactoryName).First(&fid).Error
			if err != nil {
				fmt.Println(err)
				i.Error(c, 400, "获取场站信息失败")
				return err
			}
			terr := tx.Model(&model.OrderProduct{}).Where("cart_refer= ? and factory_id = ? and is_chosen = true", cartrefer, fid).Update("order_refer",
				om.OrderID).Find(&om.ProductInfos).Error
			if terr != nil && terr != gorm.ErrRecordNotFound {
				fmt.Println(terr)
				return terr
			} else if len(om.ProductInfos) == 0 {
				fmt.Println("没有选择该场站的商品", cartrefer, fid)
				return gorm.ErrRecordNotFound
			}
			var ferr error
			for _, pv := range om.ProductInfos {
				ferr = tx.Model(&model.OrderProduct{}).Where("cart_refer = ? and factory_id = ? and name = ?", cartrefer, fid, pv.Name).Update("is_chosen", false).Error
				if ferr != nil {
					fmt.Println(ferr)
					return ferr
				}
			}
			var cost float64 = 0
			for _, pv := range om.ProductInfos {
				cost += float64(pv.Count) * pv.Price
			}
			om.Set(fid, admin.ID, cost, v.FactoryName, admin.CarID, admin.Address)
			oms = append(oms, om)
			return nil
		})
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Println(err)
			i.Error(c, 400, "订单提交失败:获取商品信息失败")
			return
		}
	}
	if err == nil {
		if cerr := dal.Getdb().Clauses(clause.OnConflict{Columns: []clause.Column{
			{Name: "order_id"}}, DoNothing: true}).Omit("ProductInfos").Create(&oms).Error; cerr != nil {
			fmt.Println(cerr)
			i.Error(c, 400, "创建订单表失败")
			return
		}
		c.JSON(200, response.DriverOrdersResponse{
			Response:        response.Response{200, "订单提交成功"},
			FactoryDistance: finfos,
			OrderInfos:      oms,
		})
	} else if err == gorm.ErrRecordNotFound {
		fmt.Println(oms)
		i.Error(c, 400, "无选中商品，无需创建表单")
	} else {
		i.Error(c, 400, "订单提交失败:获取商品信息失败")
	}
	c.Next()
}

// @Summary 补货订单结算
// @Description 结算
// @Tags Driver/Pay
// @Accept json
// @Produce json
// @Param OrderForm body model.DriverOrderForm true "传入订单信息"
// @Success 201 {object} response.PayResponse
// @Failure 400 {object} response.Response
// @Router /dr/order/pay [put]
func (i *FactoryController) Pay(c *gin.Context) {
	//TODO:业务逻辑
	var OrderForm model.DriverOrderForm
	err := c.ShouldBind(&OrderForm)
	if err != nil {
		i.Error(c, 400, "获取订单信息失败")
		return
	}
	OrderForm.State = 0
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		i.Error(c, 400, "获取车主信息失败")
		return
	}
	fmt.Println(admin)
	c.JSON(200, response.PayResponse{
		response.Response{201, "支付成功"},
	})

}
