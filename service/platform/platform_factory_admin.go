package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
// @Description 添加一个或多个的商品
// @Tags Platform/factory
// @Accept json
// @Produce json
// @Param factory_name path string true "场站名字"
// @Param factoryInfo body model.FactoryProducts true "sku 可以和name值相同 必填项:name,pic,type,sku,inventory,buy_price,supply_price"
// @Success 200 {object} response.FactoryProductsModifyResponse
// @Failure	400 {object} response.Response
// @Router /pt/factory-admin/{factory_name}/prdoucts [post]
func (f *FactoryadminController) PAdd(c *gin.Context) {
	var products []model.FactoryProduct
	err := c.ShouldBind(&products)
	if err != nil {
		f.Error(c, 400, "传入数据格式错误")
		return
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
	n := len(products)
	fpros := make([]model.FactoryProduct, n)
	for k, v := range products {
		fpros[k].Set(utils.GetSnowFlake(), fid, fname, &v)
		//fmt.Println(fpros[k])
	}
	err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
		var id int64
		for _, v := range fpros {
			terr := tx.Model(&model.FactoryProduct{}).Select("id").Where("factory_name = ? and name = ?", fname, v.Name).First(&id).Error
			if terr != nil && terr != gorm.ErrRecordNotFound {
				logrus.Info(terr)
				f.Error(c, 400, "添加商品信息失败")
				return terr
			} else if terr == gorm.ErrRecordNotFound {
				cerr := tx.Model(&model.FactoryProducts{}).Create(&v).Error
				if cerr != nil {
					logrus.Info(cerr)
					f.Error(c, 400, "添加商品信息失败")
					return cerr
				}
			} else {
				uerr := tx.Model(&model.FactoryProducts{}).Where("id = ?", id).UpdateColumn("inventory", gorm.Expr("inventory + ?", v.Inventory)).Error
				if uerr != nil {
					logrus.Info(uerr)
					f.Error(c, 400, "添加商品信息失败")
					return uerr
				}
			}
		}
		return nil
	})
	if err == nil {
		c.JSON(200, response.FactoryProductsModifyResponse{
			response.Response{200, "添加商品信息成功"}, fpros,
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
// @Description 添加一个或多个的商品
// @Tags Factory
// @Accept json
// @Produce json
// @Param factoryInfo body model.FactoryProducts true "sku 可以和name值相同 必填项:name,pic,type,sku,inventory,buy_price,supply_price"
// @Success 200 {object} response.FactoryProductsModifyResponse
// @Failure	400 {object} response.Response
// @Router /fa/inventory [post]
func (f *FactoryadminController) Add(c *gin.Context) {
	var products []model.FactoryProduct
	err := c.ShouldBind(&products)
	if err != nil {
		f.Error(c, 400, "传入数据格式错误")
		return
	}
	admin, ok := utils.GetFactoryInfo(c)
	if !ok {
		f.Error(c, 400, "获取场站信息失败")
	}
	fname := admin.Name
	fid := admin.ID
	n := len(products)
	fpros := make([]model.FactoryProduct, n)
	for k, v := range products {
		fpros[k].Set(utils.GetSnowFlake(), fid, fname, &v)
	}
	err = dal.Getdb().Transaction(func(tx *gorm.DB) error {
		var id int64
		for _, v := range fpros {
			terr := tx.Model(&model.FactoryProduct{}).Select("id").Where("factory_name = ? and name = ?", fname, v.Name).First(&id).Error
			if terr != nil && terr != gorm.ErrRecordNotFound {
				logrus.Info(terr)
				f.Error(c, 400, "查找商品信息失败")
				return terr
			} else if terr == gorm.ErrRecordNotFound {
				cerr := tx.Model(&model.FactoryProducts{}).Create(&v).Error
				if cerr != nil {
					logrus.Info(cerr)
					f.Error(c, 400, "添加商品信息失败")
					return cerr
				}
			} else {
				uerr := tx.Model(&model.FactoryProducts{}).Where("id = ?", id).UpdateColumn("inventory", gorm.Expr("inventory + ?", v.Inventory)).Error
				if uerr != nil {
					logrus.Info(uerr)
					f.Error(c, 400, "更新商品信息失败")
					return uerr
				}
			}
		}
		return nil
	})
	if err == nil {
		c.JSON(200, response.FactoryProductsModifyResponse{
			response.Response{200, "添加商品信息成功"}, fpros,
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
// @Router /fa/inventory/{product_name}/{inv} [patch]
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
// @Description	传入场站名，获取该场站所有商品信息
// @Tags	Factory
// @Accept json
// @Accept mpfd
// @Produce json
// @Param mode path int true "按照不同模式获取库存商品信息，mode={0:未上架,1:上架,传入其他数据获取所有商品信息}"
// @Success 200 {object} response.FactoryProductsResponse
// @Failure 400 {object} response.Response
// @Router /fa/infos/all/:mode [get]
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
// @Router /fa/infos/detail/:product_name [get]
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
