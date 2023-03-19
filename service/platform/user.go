package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/utils"
	"context"
	"github.com/sirupsen/logrus"
)

//注册 后续可以增加修改密码功能
func SavePtUser(ptadmin model.Platform) (model.LoginInfo, error) {
	ptadmin.Role = 3
	ptadmin.ID = utils.GetSnowFlake()
	logininfo, err := SavePtLoginInfo(ptadmin)
	if err != nil {
		return model.LoginInfo{}, err
	}
	return logininfo, dal.Getdb().Model(&model.Platform{}).Create(&ptadmin).Error

}

func SavePtLoginInfo(ptadmin model.Platform) (model.LoginInfo, error) {
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = ptadmin.ID
	loginInfo.Salt = ptadmin.PasswordSalt
	loginInfo.Password = utils.Messagedigest5(ptadmin.Password, ptadmin.PasswordSalt)
	loginInfo.Jwt, err = utils.GeneraterJwt(ptadmin.ID, ptadmin.Name, ptadmin.PasswordSalt)
	if err != nil {
		logrus.Info("JWT created fail")
		return model.LoginInfo{}, err
	}
	c := context.TODO()
	dal.Getrd().Set(c, loginInfo.Jwt, 1, utils.EXPIRE)
	return loginInfo, dal.Getdb().Model(&model.LoginInfo{}).Create(&loginInfo).Error
}
