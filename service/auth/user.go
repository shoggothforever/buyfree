package auth

import (
	"buyfree/dal"
	"buyfree/logger"
	"buyfree/repo/model"
	"buyfree/utils"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"
)

// 注册 后续可以增加修改密码功能
func SavePtUser(admin *model.Platform) (model.LoginInfo, error) {
	admin.Role = int(model.PLATFORMADMIN)
	admin.ID = utils.GetSnowFlake()
	err := dal.Getdb().Model(&model.Platform{}).Omit("password", "password_salt").Create(&admin).Error
	if err != nil {
		logger.Loger.Info("创建用户信息失败", err)
		return model.LoginInfo{}, err
	}
	logininfo, err := SavePtLoginInfo(admin)
	if err != nil {
		return model.LoginInfo{}, err
	}
	return logininfo, err

}
func SaveDrUser(admin *model.Driver) (model.LoginInfo, error) {
	admin.Role = int(model.DRIVER)
	admin.ID = utils.GetSnowFlake()
	err := dal.Getdb().Model(&model.Driver{}).Omit("password", "password_salt").Create(&admin).Error
	if err != nil {
		logger.Loger.Info("创建用户信息失败", err)
		return model.LoginInfo{}, err
	}
	logininfo, err := SaveDrLoginInfo(admin)
	if err != nil {
		logger.Loger.Info("创建用户登录表失败", err)
		return model.LoginInfo{}, err
	}
	cart := model.DriverCart{DriverID: admin.ID, Cart: model.Cart{CartID: utils.GetSnowFlake(), TotalCount: 0, TotalAmount: 0}}

	cerr := dal.Getdb().Model(&model.DriverCart{}).Create(&cart).Error
	if cerr != nil {
		logrus.Info("创建购物车信息失败", err)
	}
	return logininfo, err

}
func SaveFUser(admin *model.Factory) (model.LoginInfo, error) {
	admin.Role = int(model.FACTORYADMIN)
	admin.ID = utils.GetSnowFlake()
	err := dal.Getdb().Model(&model.Factory{}).Omit("password", "password_salt").Create(&admin).Error
	if err != nil {
		logger.Loger.Info("创建用户信息失败", err)
		return model.LoginInfo{}, err
	}
	logininfo, err := SaveFLoginInfo(admin)
	if err != nil {
		logger.Loger.Info("创建用户登录表失败", err)
		return model.LoginInfo{}, err
	}
	return logininfo, err

}
func SavePtLoginInfo(admin *model.Platform) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = admin.ID
	loginInfo.Salt = admin.PasswordSalt
	loginInfo.ROLE = model.PLATFORMADMIN
	loginInfo.Password = utils.Messagedigest5(admin.Password, admin.PasswordSalt)
	loginInfo.Jwt, err = utils.GeneraterJwt(admin.ID, admin.Name, admin.PasswordSalt)
	loginInfo.UserName = admin.Name
	if err != nil {
		logrus.Info("JWT created fail")
		return model.LoginInfo{}, err
	}
	c := context.TODO()
	dal.Getrdb().Set(c, loginInfo.Jwt, model.PLATFORMADMIN, utils.EXPIRE)
	return loginInfo, dal.Getdb().Model(&model.LoginInfo{}).Create(&loginInfo).Error
}

func SaveDrLoginInfo(admin *model.Driver) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = admin.ID
	loginInfo.Salt = admin.PasswordSalt
	loginInfo.ROLE = model.DRIVER
	loginInfo.Password = utils.Messagedigest5(admin.Password, admin.PasswordSalt)
	loginInfo.Jwt, err = utils.GeneraterJwt(admin.ID, admin.Name, admin.PasswordSalt)
	loginInfo.UserName = admin.Name
	if err != nil {
		logrus.Info("JWT created fail")
		return model.LoginInfo{}, err
	}
	c := context.TODO()
	dal.Getrdb().Set(c, loginInfo.Jwt, model.DRIVER, utils.EXPIRE)
	return loginInfo, dal.Getdb().Model(&model.LoginInfo{}).Create(&loginInfo).Error
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
	dal.Getrdb().Set(c, loginInfo.Jwt, model.FACTORYADMIN, utils.EXPIRE)
	return loginInfo, dal.Getdb().Model(&model.LoginInfo{}).Create(&loginInfo).Error
}

func SaveLoginInfo[T *model.Passenger | *model.Driver | *model.Factory | *model.Platform](admin T) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	adminvalue := reflect.ValueOf(admin).Elem()
	loginInfo.UserID = adminvalue.FieldByName("ID").Interface().(int64)
	loginInfo.Salt = adminvalue.FieldByName("PasswordSale").Interface().(string)
	loginInfo.UserName = adminvalue.FieldByName("Name").Interface().(string)
	password := adminvalue.FieldByName("Password").Interface().(string)
	loginInfo.Password = utils.Messagedigest5(password, loginInfo.Salt)
	loginInfo.Jwt, err = utils.GeneraterJwt(loginInfo.UserID, loginInfo.UserName, loginInfo.Salt)
	if err != nil {
		logrus.Info("JWT created fail")
		return model.LoginInfo{}, err
	}
	switch adminvalue.Type().String() {
	case "model.Platform":
		loginInfo.ROLE = model.PLATFORMADMIN
	case "model.Driver":
		loginInfo.ROLE = model.DRIVER
	case "model.Passenger":
		loginInfo.ROLE = model.PASSENGER
	case "model.Factory":
		loginInfo.ROLE = model.FACTORYADMIN
	default:
		return model.LoginInfo{}, errors.New("传入数据类型错误")
	}
	c := context.TODO()
	dal.Getrdb().Set(c, loginInfo.Jwt, loginInfo.ROLE, utils.EXPIRE)
	return loginInfo, dal.Getdb().Model(&model.LoginInfo{}).Create(&loginInfo).Error
}
