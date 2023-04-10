package auth

import (
	"buyfree/dal"
	"buyfree/logger"
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
	var detectors []model.Platform
	c.ShouldBind(&admin)
	dal.Getdb().Model(&model.Platform{}).Where("name = ? and role = ?", admin.Name, model.PLATFORMADMIN).First(&detectors)
	if len(detectors) == 0 {
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
			logger.Loger.Info(err)
			c.JSON(200, response.LoginResponse{
				response.Response{
					500,
					"注册失败"},
				-1,
				""})
		}
	} else {
		c.JSON(200, response.LoginResponse{response.Response{200, "该用户名已存在，请更换用户名"}, -1, ""})
	}
	c.Next()
}

// PlatformAccount godoc
// @Summary 平台用户登录
// @Description 	Input user_name and password
// @Tags			User
// @accept			json
// @Produce			json
// @Param loginInfo body model.LoginInfo true "输入昵称，密码"
// @Success			200 {object} response.LoginResponse
// @Failure			500 {object} response.LoginResponse
// @Router			/pt/login [post]
func PlatformLogin(c *gin.Context) {
	var l []model.LoginInfo
	var info model.LoginInfo
	var linfo model.LoginInfo
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&info)
	var password = info.Password
	err := dal.Getdb().Raw("select user_id,password,salt from login_infos where user_name = ? and role = ?", info.UserName, model.PLATFORMADMIN).First(&linfo).Error
	if err != nil {
		logger.Loger.Info(err)
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"登录失败,请检查输入账户名密码是否正确"},
			-1,
			"",
		})
		return
	}
	jwt, _ := utils.GeneraterJwt(linfo.UserID, info.UserName, linfo.Salt)
	psw := utils.Messagedigest5(password, linfo.Salt)
	if psw == linfo.Password {
		dal.Getdb().Model(&model.LoginInfo{}).Where("user_id = ?", linfo.UserID).UpdateColumn("jwt", jwt).First(&l)
	}
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
				"登录失败,请检查输入账户名密码是否正确"},
			-1,
			"",
		})
	}
	c.Next()
}

// PlatformAccount godoc
// @Summary 获取平台信息
// @Description 	传入jwt/token 获取用户信息
// @Tags			User
// @accept			application/x-www-form-urlencoded
// @Produce			json
// @Param	jwt formData string true "鉴权信息"
// @Success			200 {object} response.PtInfoResponse
// @Failure			400 {object} response.Response
// @Router			/pt/userinfo [post]
func PlatformUserInfo(c *gin.Context) {
	jwt := c.PostForm("jwt")
	//jwt := c.GetHeader("Authorization")
	//if len(jwt) > 7 {
	//	jwt = jwt[7:]
	//}
	db := dal.Getdb()
	var admin model.Platform
	var info model.LoginInfo
	admin.Name = info.UserName
	admin.Password = info.Password
	err := db.Model(&model.LoginInfo{}).Where("jwt = ?", jwt).First(&info).Error
	if err != nil {
		c.JSON(200, response.Response{400, "鉴权信息失效，无法获取平台数据"})
		return
	}
	err = db.Model(&model.Platform{}).Where("id = ?", info.UserID).First(&admin).Error
	if err != nil {
		c.JSON(200, response.Response{400, "查找平台信息失败"})
		return
	} else {
		c.JSON(200, response.PtInfoResponse{response.Response{200, "获取平台信息成功"}, admin})
	}
}

// @Summary 车主用户注册
// @Description 	Input info as model.User
// @Tags			User
// @accept			json
// @Produce			json
// @Param	RegisterInfo body model.Driver true "一定要填入已有的平台ID,用户名，密码，password_salt为可选项"
// @Success			200 {object} response.LoginResponse
// @failure			500 {object} response.LoginResponse
// @Router			/dr/register [post]
func DriverRegister(c *gin.Context) {
	//一定要定义成值类型，在bind里要传地址
	var admin model.Driver
	var detectors []model.Driver
	c.ShouldBind(&admin)
	dal.Getdb().Model(&model.Driver{}).Where("name = ? and role = ?", admin.Name, model.DRIVER).First(&detectors)
	if len(detectors) == 0 {
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
	} else {
		c.JSON(200, response.LoginResponse{response.Response{200, "该用户名已存在，请更换用户名"}, -1, ""})
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
	var linfo model.LoginInfo
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&info)
	var password = info.Password
	err := dal.Getdb().Raw("select user_id,password,salt from login_infos where user_name = ? and role = ?", info.UserName, model.DRIVER).First(&linfo).Error
	if err != nil {
		logger.Loger.Info(err)
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"登录失败,请检查输入账户名密码是否正确"},
			-1,
			"",
		})
		return
	}
	jwt, _ := utils.GeneraterJwt(linfo.UserID, info.UserName, linfo.Salt)
	psw := utils.Messagedigest5(password, linfo.Salt)
	if psw == linfo.Password {
		dal.Getdb().Model(&model.LoginInfo{}).Where("user_id = ?", linfo.UserID).UpdateColumn("jwt", jwt).First(&l)
	}
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

// PlatformAccount godoc
// @Summary 获取车主用户信息
// @Description 	传入jwt/token 获取用户信息
// @Tags			User
// @accept			application/x-www-form-urlencoded
// @Produce			json
// @Param jwt formData string true "鉴权信息"
// @Success			200 {object} response.DrInfoResponse
// @Failure			400 {object} response.Response
// @Router			/dr/userinfo [post]
func DriverUserInfo(c *gin.Context) {
	//jwt := c.GetHeader("Authorization")
	//if len(jwt) > 7 {
	//	jwt = jwt[7:]
	//}
	jwt := c.PostForm("jwt")
	logger.Loger.Info(jwt)
	db := dal.Getdb()
	var admin model.Driver
	var info model.LoginInfo
	err := db.Model(&model.LoginInfo{}).Where("jwt = ?", jwt).First(&info).Error
	if err != nil {
		c.JSON(200, response.Response{400, "鉴权信息失效，无法获取车主数据"})
		return
	}
	err = db.Model(&model.Driver{}).Where("id = ?", info.UserID).First(&admin).Error
	if err != nil {
		c.JSON(200, response.Response{400, "查找车主信息失败"})
		return
	}
	c.JSON(200, response.DrInfoResponse{response.Response{200, "获取车主信息成功"}, admin})

}

// @Summary 场站用户注册
// @Description 	Input info as model.User
// @Tags			User
// @accept			json
// @Produce			json
// @Param	RegisterInfo body model.Factory true "用户名(name)，密码(password)，GEO(longitude,latitude)为必填项,password_salt为可选项"
// @Success			200 {object} response.LoginResponse
// @failure			500 {object} response.LoginResponse
// @Router			/fa/register [post]
func FactoryRegister(c *gin.Context) {
	//一定要定义成值类型，在bind里要传地址
	var admin model.Factory
	var detectors []model.Factory
	c.ShouldBind(&admin)
	dal.Getdb().Model(&model.Factory{}).Where("name = ? and role = ?", admin.Name, model.FACTORYADMIN).First(&detectors)
	if len(detectors) == 0 {
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
	} else {
		c.JSON(200, response.LoginResponse{response.Response{200, "该用户名已存在，请更换用户名"}, -1, ""})
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
	var info model.LoginInfo
	var linfo model.LoginInfo
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&info)
	var password = info.Password
	err := dal.Getdb().Raw("select user_id,password,salt from login_infos where user_name = ? and role = ?", info.UserName, model.FACTORYADMIN).First(&linfo).Error
	if err != nil {
		logger.Loger.Info(err)
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"登录失败,请检查输入账户名密码是否正确"},
			-1,
			"",
		})
		return
	}
	jwt, _ := utils.GeneraterJwt(linfo.UserID, info.UserName, linfo.Salt)
	psw := utils.Messagedigest5(password, linfo.Salt)
	if psw == linfo.Password {
		dal.Getdb().Model(&model.LoginInfo{}).Where("user_id = ?", linfo.UserID).UpdateColumn("jwt", jwt).First(&l)
	}
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

// @Summary 获取场站用户信息
// @Description 	传入jwt/token 获取用户信息
// @Tags			User
// @accept			application/x-www-form-urlencoded
// @Produce			json
// @Param jwt formData string true "鉴权信息"
// @Success			200 {object} response.FaInfoResponse
// @Failure			400 {object} response.Response
// @Router			/fa/userinfo [post]
func FactoryUserInfo(c *gin.Context) {
	jwt := c.PostForm("jwt")
	//jwt := c.GetHeader("Authorization")
	//if len(jwt) > 7 {
	//	jwt = jwt[7:]
	//}
	logger.Loger.Info(jwt)
	db := dal.Getdb()
	var admin model.Factory
	var info model.LoginInfo
	err := db.Model(&model.LoginInfo{}).Where("jwt = ?", jwt).First(&info).Error
	if err != nil {
		logger.Loger.Info(err)
		c.JSON(200, response.Response{400, "鉴权信息失效，无法获取用户数据"})
		return
	}
	err = db.Model(&model.Factory{}).Where("id = ?", info.UserID).First(&admin).Error
	if err != nil {
		logger.Loger.Info(err)
		c.JSON(200, response.Response{400, "查找场站信息失败"})
		return
	}
	err = db.Model(&model.FactoryProduct{}).Where("factory_id = ?", admin.ID).Find(&admin.Products).Error
	if err != nil {
		logger.Loger.Info(err)
		c.JSON(200, response.Response{400, "获取场站商品信息失败"})
		return
	}
	c.JSON(200, response.FaInfoResponse{response.Response{200, "获取场站信息成功"}, admin})

}
