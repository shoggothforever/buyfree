package platform

import (
	"buyfree/dal"
	"buyfree/logger"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

type FactoryadminController struct {
	BasePtController
}

func SaveFUser(admin *model.Factory) (model.LoginInfo, error) {
	admin.Role = int(model.FACTORYADMIN)
	admin.ID = utils.GetSnowFlake()
	logininfo, err := SaveFLoginInfo(admin)
	if err != nil {
		return model.LoginInfo{}, err
	}
	//TODO:对密码加密再存储，现在为了方便就先不管了
	return logininfo, dal.Getdb().Model(&model.Factory{}).Create(&admin).Error

}
func SaveFLoginInfo(admin *model.Factory) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = admin.ID
	loginInfo.Salt = admin.PasswordSalt
	loginInfo.UserName = admin.Name
	loginInfo.ROLE = model.FACTORYADMIN
	loginInfo.Password = utils.Messagedigest5(admin.Password, admin.PasswordSalt)
	loginInfo.Jwt, err = utils.GeneraterJwt(admin.ID, admin.Name, admin.PasswordSalt)
	if err != nil {
		logrus.Info("JWT created fail")
		return model.LoginInfo{}, err
	}
	c := context.TODO()
	dal.Getrdb().Set(c, loginInfo.Jwt, 1, utils.EXPIRE)
	return loginInfo, dal.Getdb().Model(&model.LoginInfo{}).Create(&loginInfo).Error
}

// @Summary 平台登记场站信息
// @Description 填入场站信息
// @Tags Platform/factory
// @Accept json
// @Produce json
// @Param factoryInfo body model.Factory true "必填项:name,address,longitude,latitude"
// @Success 200 {object} response.FactoryRegisterResponse
// @Failure	400 {object} response.Response
// @Router /pt/factory-admin/register [post]
func (f *FactoryadminController) PRegister(c *gin.Context) {
	var admin model.Factory
	err := c.ShouldBind(&admin)
	if err != nil {
		f.Error(c, 400, "传入数据格式错误")
		return
	}
	rdb := dal.Getrdb()
	ctx := rdb.Context()
	//向redis中写入场站的地理位置信息
	utils.LocAdd(ctx, rdb, utils.LOCATION, admin.Longitude, admin.Latitude, admin.Name)
	var logininfo model.LoginInfo
	err = dal.Getdb().Model(&model.LoginInfo{}).Where("role =2 and user_name = ? ", admin.Name).First(&logininfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		f.Error(c, 400, "场站信息登记失败")
		return
	}
	if logininfo.UserName != "" {
		fmt.Println(err)
		f.Error(c, 400, "场站信息已经登记了")
		return
	}
	logininfo, err = SaveFUser(&admin)
	if err == nil {
		c.JSON(200, response.FactoryRegisterResponse{
			response.Response{200, "平台登记成功"},
			logininfo.UserID,
			logininfo.UserName,
			admin.Address,
			admin.Longitude,
			admin.Latitude,
			admin.Description,
		})
	} else {
		fmt.Println(err)
		f.Error(c, 400, "场站信息登记失败")
	}
}

// @Summary 平台为场站添加商品信息
// @Description 添加一个商品
// @Tags Platform/factory
// @Accept json
// @Produce json
// @Param factory_name path string true "场站名字"
// @Param factoryInfo body model.FactoryProduct true "sku 可以和name值相同 必填项:name,pic,type,sku,inventory,buy_price,supply_price"
// @Success 200 {object} response.FactoryProductsModifyResponse
// @Failure	400 {object} response.Response
// @Router /pt/factory-admin/{factory_name}/prdoucts [post]
func (f *FactoryadminController) PAdd(c *gin.Context) {
	var product model.FactoryProduct
	err := c.ShouldBind(&product)
	if err != nil {
		f.Error(c, 400, "传入数据格式错误")
		return
	}
	if product.Name == "" {
		product.Name = c.PostForm("name")
		product.Pic = c.PostForm("pic")
		product.Type = c.PostForm("type")
		product.Sku = c.PostForm("sku")
		product.BuyPrice, _ = strconv.ParseFloat(c.PostForm("buy_price"), 64)
		product.SupplyPrice, _ = strconv.ParseFloat(c.PostForm("supply_price"), 64)
		product.Inventory, _ = strconv.ParseInt(c.PostForm("inventory"), 10, 64)
	}
	//fmt.Println(products)
	fname := c.Param("factory_name")
	var fid int64
	err = dal.Getdb().Model(&model.Factory{}).Select("id").Where("name = ? ", fname).First(&fid).Error
	if err != nil {
		logrus.Info(err)
		f.Error(c, 400, "获取场站信息失败")
		return
	}
	//n := len(product)
	//fpros := make([]model.FactoryProduct, n)
	//for k, v := range product {
	//	fpros[k].Set(utils.GetSnowFlake(), fid, fname, &v)
	//	//fmt.Println(fpros[k])
	//}
	//err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
	//	var id int64
	//	for _, v := range fpros {
	//		terr := tx.Model(&model.FactoryProduct{}).Select("id").Where("factory_name = ? and name = ?", fname, v.Name).First(&id).Error
	//		if terr != nil && terr != gorm.ErrRecordNotFound {
	//			logrus.Info(terr)
	//			f.Error(c, 400, "添加商品信息失败")
	//			return terr
	//		} else if terr == gorm.ErrRecordNotFound {
	//			cerr := tx.Model(&model.FactoryProducts{}).Create(&v).Error
	//			if cerr != nil {
	//				logrus.Info(cerr)
	//				f.Error(c, 400, "添加商品信息失败")
	//				return cerr
	//			}
	//		} else {
	//			uerr := tx.Model(&model.FactoryProducts{}).Where("id = ?", id).UpdateColumn("inventory", gorm.Expr("inventory + ?", v.Inventory)).Error
	//			if uerr != nil {
	//				logrus.Info(uerr)
	//				f.Error(c, 400, "添加商品信息失败")
	//				return uerr
	//			}
	//		}
	//	}
	//	return nil
	//})

	//fmt.Println(fpros[k])
	err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
		terr := tx.Model(&model.FactoryProduct{}).Select("id").Where("factory_name = ? and name = ?", fname, product.Name).First(&product.ID).Error
		if terr != nil && terr != gorm.ErrRecordNotFound {
			logrus.Info(terr)
			f.Error(c, 400, "添加商品信息失败")
			return terr
		} else if terr == gorm.ErrRecordNotFound {
			product.Set(utils.GetSnowFlake(), fid, fname, &product)
			cerr := tx.Model(&model.FactoryProducts{}).Create(&product).Error
			if cerr != nil {
				logrus.Info(cerr)
				f.Error(c, 400, "添加商品信息失败")
				return cerr
			}
		} else {
			product.Set(product.ID, fid, fname, &product)
			uerr := tx.Model(&model.FactoryProducts{}).Select("inventory").Where("id = ?", product.ID).UpdateColumn("inventory", gorm.Expr("inventory + ?", product.Inventory)).First(&product.Inventory).Error
			if uerr != nil {
				logrus.Info(uerr)
				f.Error(c, 400, "添加商品信息失败")
				return uerr
			}
		}
		return nil
	})
	if err == nil {
		c.JSON(200, response.FactoryProductsModifyResponse{
			response.Response{200, "添加商品信息成功"}, []model.FactoryProduct{product},
		})
	} else {
		logrus.Info(err)
		f.Error(c, 400, "添加商品信息失败")
	}

}

// @Summary 平台为场站更新商品库存信息
// @Description 传入增加的库存量
// @Tags Platform/factory
// @Accept json
// @Accept mpfd
// @Produce json
// @Param factory_name path string true "场站名字"
// @Param product_name path string true "商品名字"
// @Param inv path int64 true "增加的库存值"
// @Success 200 {object} response.FactoryProductsModifyResponse
// @Failure	400 {object} response.Response
// @Router /pt/factory-admin/{factory_name}/prdoucts/{product_name}/{inv} [patch]
func (f *FactoryadminController) PAddInv(c *gin.Context) {
	inv := c.Param("inv")
	fname := c.Param("factory_name")
	pname := c.Param("product_name")
	var pro model.FactoryProduct
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("factory_name = ? and name = ? ", fname, pname).UpdateColumn("inventory", gorm.Expr("inventory + ?", inv)).First(&pro).Error
	if err != nil {
		logrus.Info(err)
		f.Error(c, 400, "修改库存信息失败")
		return
	} else {
		c.JSON(200, response.FactoryProductsModifyResponse{
			response.Response{200, "添加商品信息成功"}, []model.FactoryProduct{pro},
		})
	}
}

// @Summary 场站添加商品信息
// @Description 添加一个商品
// @Tags Factory
// @Accept json
// @Produce json
// @Param factoryInfo body model.FactoryProduct true "sku 可以和name值相同 必填项:name,pic(直接传入文件名即可),type,sku,inventory,buy_price,supply_price"
// @Success 200 {object} response.FactoryProductsModifyResponse
// @Failure	400 {object} response.Response
// @Router /fa/inventory [post]
func (f *FactoryadminController) Add(c *gin.Context) {
	var product model.FactoryProduct
	err := c.ShouldBind(&product)
	if product.Name == "" {
		product.Name = c.PostForm("name")
		product.Pic = c.PostForm("pic")
		product.Type = c.PostForm("type")
		product.Sku = c.PostForm("sku")
		product.BuyPrice, _ = strconv.ParseFloat(c.PostForm("buy_price"), 64)
		product.SupplyPrice, _ = strconv.ParseFloat(c.PostForm("supply_price"), 64)
		product.Inventory, _ = strconv.ParseInt(c.PostForm("inventory"), 10, 64)
	}
	if err != nil {
		f.Error(c, 400, "传入数据格式错误")
		return
	}
	admin, ok := utils.GetFactoryInfo(c)
	if !ok {
		f.Error(c, 400, "获取场站信息失败")
		return
	}
	fname := admin.Name
	fid := admin.ID
	//n := len(products)
	//fpros := make([]model.FactoryProduct, n)
	//for k, v := range products {
	//	fpros[k].Set(utils.GetSnowFlake(), fid, fname, &v)
	//}
	//err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
	//	var id int64
	//	for k, v := range fpros {
	//		terr := tx.Model(&model.FactoryProduct{}).Select("id").Where("factory_name = ? and name = ?", fname, v.Name).First(&id).Error
	//		if terr != nil && terr != gorm.ErrRecordNotFound {
	//			logrus.Info(terr)
	//			f.Error(c, 400, "查找商品信息失败")
	//			return terr
	//		} else if terr == gorm.ErrRecordNotFound {
	//			cerr := tx.Model(&model.FactoryProducts{}).Create(&v).Error
	//			if cerr != nil {
	//				logrus.Info(cerr)
	//				f.Error(c, 400, "添加商品信息失败")
	//				return cerr
	//			}
	//		} else {
	//			fpros[k].Product.ID = id
	//			uerr := tx.Model(&model.FactoryProducts{}).Select("inventory").Where("id = ?", id).UpdateColumn("inventory", gorm.Expr("inventory + ?", v.Inventory)).First(&fpros[k].Product.Inventory).Error
	//			if uerr != nil {
	//				logrus.Info(uerr)
	//				f.Error(c, 400, "更新商品信息失败")
	//				return uerr
	//			}
	//		}
	//	}
	//	return nil
	//})
	product.Set(utils.GetSnowFlake(), fid, fname, &product)
	err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
		var id int64
		terr := tx.Model(&model.FactoryProduct{}).Select("id").Where("factory_name = ? and name = ?", fname, product.Name).First(&id).Error
		if terr != nil && terr != gorm.ErrRecordNotFound {
			logrus.Info(terr)
			f.Error(c, 400, "查找商品信息失败")
			return terr
		} else if terr == gorm.ErrRecordNotFound {
			cerr := tx.Model(&model.FactoryProducts{}).Create(&product).Error
			if cerr != nil {
				logrus.Info(cerr)
				f.Error(c, 400, "添加商品信息失败")
				return cerr
			}
		} else {
			product.ID = id
			uerr := tx.Model(&model.FactoryProducts{}).Select("inventory").Where("id = ?", id).UpdateColumn("inventory", gorm.Expr("inventory + ?", product.Inventory)).First(&product.Inventory).Error
			if uerr != nil {
				logrus.Info(uerr)
				f.Error(c, 400, "更新商品信息失败")
				return uerr
			}
		}
		return nil
	})
	if err == nil {
		c.JSON(200, response.FactoryProductsModifyResponse{
			response.Response{200, "添加商品信息成功"}, []model.FactoryProduct{product},
		})
	} else {
		logrus.Info(err)
		f.Error(c, 400, "添加商品信息失败")
	}

}

// @Summary 场站更新商品库存信息
// @Description 传入增加的库存量
// @Tags Factory
// @Accept json
// @Accept mpfd
// @Produce json
// @Param product_name path string true "商品名字"
// @Param inv path int64 true "增加的库存值"
// @Success 200 {object} response.FactoryProductsModifyResponse
// @Failure	400 {object} response.Response
// @Router /fa/inventory/add/{product_name}/{inv} [patch]
func (f *FactoryadminController) AddInv(c *gin.Context) {
	admin, ok := utils.GetFactoryInfo(c)
	if !ok {
		f.Error(c, 400, "获取场站信息失败")
	}
	inv := c.Param("inv")
	pname := c.Param("product_name")
	var pro model.FactoryProduct
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("factory_name = ? and name = ? ", admin.Name, pname).UpdateColumn("inventory", gorm.Expr("inventory + ?", inv)).First(&pro).Error
	if err != nil {
		logrus.Info(err)
		f.Error(c, 400, "修改库存信息失败")
		return
	} else {
		c.JSON(200, response.FactoryProductsModifyResponse{
			response.Response{200, "添加商品信息成功"}, []model.FactoryProduct{pro},
		})
	}
}

// @Summary 获取所有商品
// @Description	获取本场站所有商品信息
// @Tags	Factory
// @Accept json
// @Accept mpfd
// @Produce json
// @Param mode path int true "按照不同模式获取库存商品信息，mode={0:未上架,1:上架,传入其他数据获取所有商品信息}"
// @Success 200 {object} response.FactoryProductsResponse
// @Failure 400 {object} response.Response
// @Router /fa/infos/all/{mode} [get]
func (f *FactoryadminController) GetAllProducts(c *gin.Context) {
	admin, ok := utils.GetFactoryInfo(c)
	if !ok {
		f.Error(c, 400, "获取场站信息失败")
		return
	}
	fname := admin.Name
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
			f.Error(c, 400, "获取商品信息失败")
			return
		}
	} else if len(fname) != 0 {
		err := dal.Getdb().Raw("select * from factory_products where factory_name = ?", fname).Find(&infos).Error
		if err != nil {
			f.Error(c, 400, "获取商品信息失败")
			return
		}
	} else if mode == "0" || mode == "1" {
		err := dal.Getdb().Raw("select * from factory_products where is_on_shelf = ?", sf).Find(&infos).Error
		if err != nil {
			f.Error(c, 400, "获取商品信息失败")
			return
		}
	} else {
		err := dal.Getdb().Model(&model.FactoryProduct{}).Find(&infos).Error
		if err != nil {
			f.Error(c, 200, "获取商品信息失败")
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
	c.Next()
}

// @Summary  获取商品信息
// @Description
// @Tags Factory
// @Accept json
// @Accept mpfd
// @Produce json
// @Param product_name path string true "商品名"
// @Success 200 {object} response.FactoryGoodsResponse
// @Failure 400 {object} response.Response
// @Router /fa/infos/detail/{product_name} [get]
func (f *FactoryadminController) GetGoodsInfo(c *gin.Context) {
	admin, ok := utils.GetFactoryInfo(c)
	if !ok {
		f.Error(c, 400, "获取场站信息失败")
		return
	}
	factoryName := admin.Name
	productName := c.Param("product_name")
	var product model.FactoryProduct
	err := dal.Getdb().Model(&model.FactoryProduct{}).Where("factory_name = ? and name = ?", factoryName, productName).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		logrus.Info(err)
		f.Error(c, 404, "不存在该商品，请输入正确的信息")
		return
	} else if err != nil {
		logrus.Info(err)
		f.Error(c, 404, "查询失败")
		return
	}
	c.JSON(200, response.FactoryGoodsResponse{
		response.Response{200, "成功获取对应信息"},
		product,
	})
}

//// @Summary 获取车主订单信息(车主在该场站下的订单)
//// @Description	传入字段mode，获取对应订单信息
//// @Tags	Orderform
//// @Accept json
//// @Accept mpfd
//// @Produce json
//// @Param factory_name path string true "场站名"
//// @Param mode path int true "按照不同模式获取订单信息，mode={0:未支付,1:未完成,2:完成,传入其他任意数值代表获取全部订单信息}"
//// @Param page path int true "默认第一页，一页20个数据"
//// @Success 200 {object} response.OrderResponse
//// @Failure 400 {object} response.Response
//// @Router /pt/fa-admin/{factory_name}/orders/{mode}/{page} [get]
//func (o *FactoryadminController) GetDriverOrders(c *gin.Context) {
//	page := c.Param("page")
//	factory_name := c.Param("factory_name")
//	//mode =2-已完成 1-待取货 0-未支付 else 全部
//	mode := c.Param("mode")
//	var dofs []*model.DriverOrderForm
//	if mode == "0" || mode == "1" || mode == "2" {
//		err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("state = ?", mode).Find(&dofs).Error
//		if err != nil {
//			o.Error(c, 400, "获取订单信息失败 1")
//			return
//		}
//	} else {
//		err := dal.Getdb().Model(&model.DriverOrderForm{}).Find(&dofs).Error
//		if err != nil {
//			o.Error(c, 400, "获取订单信息失败 1")
//			return
//		}
//	}
//	n := len(dofs)
//	fmt.Printf("获取到%d条订单信息\n", n)
//	ords := []response.FactoryProductsInfo{}
//	for i := 0; i < n; i++ {
//		var products []model.OrderProduct
//		//products, err := gen.OrderProduct.GetAllOrderProductReferDOrder(dofs[i].OrderID)
//		err := dal.Getdb().Raw("select * from order_products where order_refer =? limit 3", dofs[i].OrderID).Find(&products).Error
//		if err != nil {
//			o.Error(c, 400, "获取订单信息失败 2")
//			return
//		}
//		k := len(products)
//		fmt.Printf("获取到%d条货品信息\n", k)
//		factoryname := dofs[i].FactoryName
//		infos := make([]response.FactoryProductsInfo, k)
//		for j := 0; j < k; j++ {
//			var info response.FactoryProductsInfo
//			infos[j].FactoryName = factoryname
//			infos[j].Name = products[j].Name
//			infos[j].Sku = products[j].Sku
//			infos[j].Pic = products[j].Pic
//			infos[j].Type = products[j].Type
//			//TODO:展示在首页和上架就交给前端吧,获取订单中的商品在场站的上下架状态，根据factoryID 和 商品SKU在场站的商品表中查询对应的状态信息
//			infos[j].IsOnShelf = products[j].IsChosen
//			//saleinfo, _ := gen.FactoryProduct.GetBySkuAndFName(info.Sku, info.FactoryName)
//			var saleinfo model.FactoryProduct
//			err := dal.Getdb().Model(&model.FactoryProduct{}).Select("total_sales").Where("sku = ? and factory_name = ?", info.Sku, info.FactoryName).First(&saleinfo.TotalSales).Error
//			if err != gorm.ErrRecordNotFound && err != nil {
//				o.Error(c, 400, "获取订单信息失败 2")
//				return
//			}
//			infos[j].TotalSales = saleinfo.TotalSales
//			infos[j].Inventory = saleinfo.Inventory
//		}
//		ords = append(ords, infos...)
//	}
//	if len(ords) != 0 {
//		c.JSON(200, response.OrderResponse{
//			response.Response{
//				200,
//				"成功获取所有订单信息",
//			},
//			ords,
//		})
//	} else {
//		c.JSON(200, response.OrderResponse{
//			response.Response{
//				200,
//				"暂无相关订单信息",
//			},
//			ords,
//		})
//	}
//	c.Set("Orders", ords)
//	c.Next()
//}

// @Summary 获取车主订单信息(车主在该场站下的订单)
// @Description	传入字段mode，获取对应订单信息
// @Tags	Factory
// @Accept json
// @Accept mpfd
// @Produce json
// @Param mode path int true "按照不同模式获取订单信息，mode={0:未支付,1:未完成,2:完成,传入其他任意数值代表获取全部订单信息}"
// @Param page path int true "默认第一页，一页20个数据"
// @Success 200 {object} response.OrderResponse
// @Failure 400 {object} response.Response
// @Router /fa/infos/orders/{mode}/{page} [get]
func (o *FactoryadminController) GetDriverOrders(c *gin.Context) {
	spage := c.Param("page")
	page, err := strconv.Atoi(spage)
	if err != nil || page < 1 {
		page = 1
	}
	//mode =2-已完成 1-待取货 0-未支付 else 全部
	mode := c.Param("mode")
	admin, ok := utils.GetFactoryInfo(c)
	if !ok {
		o.Error(c, 400, "获取场站信息失败")
		return
	}
	fid := admin.ID
	var dofs []*model.DriverOrderForm
	if mode == "0" || mode == "1" || mode == "2" {
		err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("factory_id = ? and state = ?", fid, mode).Offset((page - 1) * 20).Limit(20).Find(&dofs).Error
		if err != nil {
			o.Error(c, 400, "获取订单信息失败 1")
			return
		}
	} else {
		err := dal.Getdb().Model(&model.DriverOrderForm{}).Where("factory_id = ?", fid).Offset((page - 1) * 20).Limit(20).Find(&dofs).Error
		if err != nil {
			o.Error(c, 400, "获取订单信息失败 1")
			return
		}
	}
	n := len(dofs)
	fmt.Printf("获取到%d条订单信息\n", n)
	ords := make([]response.FactoryOrderInfo, n)
	for i := 0; i < n; i++ {
		//products, err := gen.OrderProduct.GetAllOrderProductReferDOrder(dofs[i].OrderID)
		var dr model.Driver
		err = dal.Getdb().Raw("select name,mobile,car_id from drivers where id = ?", dofs[i].DriverID).Find(&dr).Error
		if err != nil {
			logger.Loger.Info("获取车主信息失败", err)
			o.Error(c, 400, "获取车主信息失败")
			return
		}
		err := dal.Getdb().Raw("select name,type,sku,pic,count,price from order_products where order_refer = ?", dofs[i].OrderID).Find(&ords[i].OrderProductInfo).Error
		if err != nil {
			logger.Loger.Info("获取订单商品失败", err)
			o.Error(c, 400, "获取订单商品信息失败 2")
			return
		}
		ords[i].OrderID = dofs[i].OrderID

		ords[i].DriverInfo.Name = dr.Name
		ords[i].DriverInfo.Mobile = dr.Mobile
		ords[i].DriverInfo.CarID = dr.CarID
		ords[i].PayInfo.PlaceTime = dofs[i].PlaceTime
		ords[i].PayInfo.PayTime = dofs[i].PlaceTime
		ords[i].PayInfo.GetTime = dofs[i].GetTime
		ords[i].PayInfo.Cost = dofs[i].Cost
		ords[i].PayInfo.State = int64(dofs[i].State)
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
				400,
				"暂无相关订单信息",
			},
			ords,
		})
	}
}
