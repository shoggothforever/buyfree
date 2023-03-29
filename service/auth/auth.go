package auth

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
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
	var admin model.Platform
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&admin)
	var password string = admin.Password
	dal.Getdb().Raw("select id,password_salt from platforms where name = ? and role = ?", admin.Name, model.PLATFORMADMIN).First(&admin)
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
	var admin model.Driver
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&admin)
	var password string = admin.Password
	//查找数据库获得用户的密码盐
	dal.Getdb().Raw("select id,password_salt from drivers where name = ? and role = ?", admin.Name, model.DRIVER).First(&admin)
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
