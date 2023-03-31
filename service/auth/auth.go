package auth

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RegisterInfo struct {
	Name, Password, PasswordSalt string
}

// PlatformAccount godoc
// @Summary 平台用户注册
// @Description 	Input info as model.User
// @Tags			User
// @accept			json
// @Produce			json
// @Param	RegisterInfo body model.Platform true "只需要用户名，密码，password_salt为可选项"
// @Success			200 {object} response.LoginResponse
// @failure			500 {object} response.LoginResponse
// @Router			/pt/register [post]
func PlatformRegister(c *gin.Context) {
	//一定要定义成值类型，在bind里要传地址
	var admin model.Platform

	c.ShouldBind(&admin)
	logininfo, err := SavePtUser(&admin)
	if err == nil {
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"注册成功",
			},
			logininfo.UserID,
			logininfo.Jwt})
		c.Set("Authorization", "Bearer:"+logininfo.Jwt)
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"注册失败"},
			-1,
			""})
	}
	c.Next()
}

// PlatformAccount godoc
// @Summary 平台用户登录
// @Description 	Input user's nickname and password
// @Tags			User
// @accept			json
// @Produce			json
// @Param loginInfo body model.LoginInfo true "输入昵称，密码"
// @Success			200 {object} response.LoginResponse
// @Failure			400 {object} response.LoginResponse
// @Router			/pt/login [post]
func PlatformLogin(c *gin.Context) {
	var l []model.LoginInfo
	var info model.LoginInfo
	var admin model.Platform
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&info)
	fmt.Println(info)
	var password string = info.Password
	dal.Getdb().Raw("select id,password_salt from platforms where name = ? and role = ?", info.UserName, model.PLATFORMADMIN).First(&admin)
	psw := utils.Messagedigest5(password, admin.PasswordSalt)
	dal.Getdb().Model(&model.LoginInfo{}).Where("user_id = ? and password = ?", admin.ID, psw).First(&l)
	if len(l) != 0 {
		c.Set("name", admin.Name)
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"登录成功"},
			l[0].UserID,
			l[0].Jwt,
		})
		c.Set("Authorization", "Bearer:"+l[0].Jwt)
		dal.Getrdb().Set(c, l[0].Jwt, 1, utils.EXPIRE)
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"登录失败"},
			-1,
			"",
		})
	}
	c.Next()
}

// PlatformAccount godoc
// @Summary 获取用户信息
// @Description 	传入jwt/token 获取用户信息
// @Tags			User
// @accept			application/x-www-form-urlencoded
// @Produce			json
// @Param jwt formData string true "鉴权信息"
// @Success			200 {object} model.Platform
// @Failure			400 {object} response.Response
// @Router			/pt/userinfo [post]
func PlatformUserInfo(c *gin.Context) {
	jwt := c.PostForm("jwt")
	db := dal.Getdb()
	var admin model.Platform
	var info model.LoginInfo
	err := db.Model(&model.LoginInfo{}).Where("jwt = ?", jwt).First(&info).Error
	if err != nil {
		c.JSON(200, response.Response{400, "鉴权信息失效，无法获取用户数据"})
		return
	}
	err = db.Model(&model.Platform{}).Where("id = ?", info.UserID).First(&admin).Error
	if err != nil {
		c.JSON(200, response.Response{400, "查找用户信息失败"})
		return
	}
	c.JSON(200, admin)

}

// @Summary 车主用户注册
// @Description 	Input info as model.User
// @Tags			User
// @accept			json
// @Produce			json
// @Param	RegisterInfo body model.Driver true "一定要填入已有的平台ID,用户名，密码，password_salt为可选项"
// @Success			200 {object} response.LoginResponse
// @failure			400 {object} response.LoginResponse
// @Router			/dr/register [post]
func DriverRegister(c *gin.Context) {
	//一定要定义成值类型，在bind里要传地址
	var admin model.Driver
	c.ShouldBind(&admin)
	logininfo, err := SaveDrUser(&admin)
	if err == nil {
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"注册成功",
			},
			logininfo.UserID,
			logininfo.Jwt})
		c.Set("Authorization", "Bearer:"+logininfo.Jwt)
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"注册失败"},
			-1,
			""})
	}
	c.Next()
}

// @Summary 车主用户登录
// @Description 	Input user's nickname and password
// @Tags			User
// @accept			json
// @Produce			json
// @Param loginInfo body model.LoginInfo true "输入昵称，密码"
// @Success			200 {object} response.LoginResponse
// @Failure			500 {object} response.LoginResponse
// @Router			/dr/login [post]
func DriverLogin(c *gin.Context) {
	var l []model.LoginInfo
	var info model.LoginInfo
	var admin model.Driver
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&info)
	var password string = info.Password
	//查找数据库获得用户的密码盐
	dal.Getdb().Raw("select id,password_salt from drivers where name = ? and role = ?", info.UserName, model.DRIVER).First(&admin)
	psw := utils.Messagedigest5(password, admin.PasswordSalt)
	dal.Getdb().Model(&model.LoginInfo{}).Where("user_id = ? and password = ?", admin.ID, psw).First(&l)
	if len(l) != 0 {
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"登录成功"},
			l[0].UserID,
			l[0].Jwt,
		})
		c.Set("Authorization", "Bearer:"+l[0].Jwt)
		dal.Getrdb().Set(c, l[0].Jwt, 1, utils.EXPIRE)
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"登录失败"},
			-1,
			"",
		})
	}
	c.Next()
}

// @Summary 场站用户注册
// @Description 	Input info as model.User
// @Tags			User
// @accept			json
// @Produce			json
// @Param	RegisterInfo body model.Factory true "填入用户名，密码，password_salt为可选项"
// @Success			200 {object} response.LoginResponse
// @failure			400 {object} response.LoginResponse
// @Router			/fa/register [post]
func FactoryRegister(c *gin.Context) {
	//一定要定义成值类型，在bind里要传地址
	var admin model.Factory
	c.ShouldBind(&admin)
	rdb := dal.Getrdb()
	ctx := rdb.Context()
	//向redis中写入场站的地理位置信息
	utils.LocAdd(ctx, rdb, utils.LOCATION, admin.Longitude, admin.Latitude, admin.Name)
	logininfo, err := SaveFUser(&admin)
	if err == nil {
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"注册成功",
			},
			logininfo.UserID,
			logininfo.Jwt})
		c.Set("Authorization", "Bearer:"+logininfo.Jwt)
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"注册失败"},
			-1,
			""})
	}
	c.Next()
}

// @Summary 场站用户登录
// @Description 	Input user's nickname and password
// @Tags			User
// @accept			json
// @Produce			json
// @Param loginInfo body model.LoginInfo true "输入昵称，密码"
// @Success			200 {object} response.LoginResponse
// @Failure			500 {object} response.LoginResponse
// @Router			/fa/login [post]
func FactoryLogin(c *gin.Context) {
	var l []model.LoginInfo
	var admin model.Driver
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&admin)
	var password string = admin.Password
	//查找数据库获得用户的密码盐
	dal.Getdb().Raw("select id,password_salt from factorys where name = ? and role = ?", admin.Name, model.FACTORYADMIN).First(&admin)
	psw := utils.Messagedigest5(password, admin.PasswordSalt)
	dal.Getdb().Model(&model.LoginInfo{}).Where("user_id = ? and password = ?", admin.ID, psw).First(&l)
	if len(l) != 0 {
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"登录成功"},
			l[0].UserID,
			l[0].Jwt,
		})
		c.Set("Authorization", "Bearer:"+l[0].Jwt)
		dal.Getrdb().Set(c, l[0].Jwt, 1, utils.EXPIRE)
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"登录失败"},
			-1,
			"",
		})
	}
	c.Next()
}
