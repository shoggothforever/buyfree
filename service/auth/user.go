package auth

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/utils"
	"context"
	"github.com/sirupsen/logrus"
)

//注册 后续可以增加修改密码功能
func SavePtUser(admin *model.Platform) (model.LoginInfo, error) {
	admin.Role = int(model.PLATFORMADMIN)
	admin.ID = utils.GetSnowFlake()
	logininfo, err := SavePtLoginInfo(admin)
	if err != nil {
		return model.LoginInfo{}, err
	}
	return logininfo, dal.Getdb().Model(&model.Platform{}).Create(&admin).Error

}
func SaveDrUser(admin *model.Driver) (model.LoginInfo, error) {
	admin.Role = int(model.DRIVER)
	admin.ID = utils.GetSnowFlake()
	logininfo, err := SaveDrLoginInfo(admin)
	if err != nil {
		return model.LoginInfo{}, err
	}
	return logininfo, dal.Getdb().Model(&model.Driver{}).Create(&admin).Error

}
func SaveFUser(admin *model.Factory) (model.LoginInfo, error) {
	admin.Role = int(model.FACTORYADMIN)
	admin.ID = utils.GetSnowFlake()
	logininfo, err := SaveFLoginInfo(admin)
	if err != nil {
		return model.LoginInfo{}, err
	}
	return logininfo, dal.Getdb().Model(&model.Factory{}).Create(&admin).Error

}
func SavePtLoginInfo(admin *model.Platform) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = admin.ID
	loginInfo.Salt = admin.PasswordSalt
	loginInfo.ROLE = model.PLATFORMADMIN
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

func SaveDrLoginInfo(admin *model.Driver) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = admin.ID
	loginInfo.Salt = admin.PasswordSalt
	loginInfo.ROLE = model.DRIVER
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
func SaveFLoginInfo(admin *model.Factory) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = admin.ID
	loginInfo.Salt = admin.PasswordSalt
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
