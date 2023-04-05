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
	loginInfo.ROLE = model.FACTORYADMIN
	loginInfo.Password = utils.Messagedigest5(admin.Password, admin.PasswordSalt)
	loginInfo.Jwt, err = utils.GeneraterJwt(admin.ID, admin.Name, admin.PasswordSalt)
	loginInfo.UserName = admin.Name
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
func (f *FactoryadminController) Register(c *gin.Context) {
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

// @Summary 场站添加商品信息
// @Description 添加的商品是场站中没有的，使用前请先查看场站商品列表
// @Tags Platform/factory
// @Accept json
// @Produce json
// @Param factory_name path string true "场站名字"
// @Param factoryInfo body model.FactoryProducts true "sku 可以和name值相同 必填项:name,pic,type,sku,inventory,buy_price,supply_price"
// @Success 200 {object} response.FactoryProductsModifyResponse
// @Failure	400 {object} response.Response
// @Router /pt/factory-admin/{factory_name}/prdoucts [post]
func (f *FactoryadminController) Add(c *gin.Context) {
	var products []model.FactoryProduct
	err := c.ShouldBind(&products)
	if err != nil {
		f.Error(c, 400, "传入数据格式错误")
		return
	}
	fmt.Println(products)
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
		fmt.Println(fpros[k])
	}
	err = dal.Getdb().Model(&model.FactoryProduct{}).Save(&fpros).Error
	if err != nil {
		logrus.Info(err)
		f.Error(c, 400, "添加商品信息失败")
		return
	} else {
		c.JSON(200, response.FactoryProductsModifyResponse{
			response.Response{200, "添加商品信息成功"}, fpros,
		})
	}

}

// @Summary 场站更新商品库存信息
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
func (f *FactoryadminController) AddInv(c *gin.Context) {
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
