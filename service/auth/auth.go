package auth

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/platform"
	"buyfree/service/response"
	"buyfree/utils"
	"github.com/gin-gonic/gin"
)

// PlatformAccount godoc
// @Summary 平台用户注册
// @Description 	Input info as model.Platform
// @Tags			Platform
// @accept			json
// @Produce			json
// @Success			200 {object} response.LoginResponse
// @failure			500 {object} response.LoginResponse
// @Router			/pt/register [post]
func Register(c *gin.Context) {
	//一定要定义成值类型，在bind里要传地址
	var ptadmin model.Platform
	c.ShouldBind(&ptadmin)
	//fmt.Println(ptadmin)
	//fmt.Println(ptadmin.ID)
	logininfo, err := platform.SavePtUser(ptadmin)
	if err == nil {
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"注册成功",
			},
			logininfo.UserID,
			logininfo.Jwt})
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"注册失败"},
			-1,
			""})
	}
}

// PlatformAccount godoc
// @Summary 平台用户登录
// @Description 	Input user's nickname and password
// @Tags			Platform
// @accept			json
// @Produce			json
// @Success			200 {object} response.LoginResponse
// @Failure			500 {object} response.LoginResponse
// @Router			/pt/login [post]
func Login(c *gin.Context) {
	var l []model.LoginInfo
	var pt model.Platform
	//输入昵称，密码 需要用户id和盐
	c.ShouldBind(&pt)
	var password string = pt.Password
	dal.Getdb().Raw("select id,password_salt from platforms where name = ?", pt.Name).First(&pt)
	psw := utils.Messagedigest5(password, pt.PasswordSalt)
	dal.Getdb().Model(&model.LoginInfo{}).Where("user_id = ? and password = ?", pt.ID, psw).First(&l)
	if len(l) != 0 {
		c.Set("name", pt.Name)
		c.JSON(200, response.LoginResponse{
			response.Response{
				200,
				"登录成功"},
			l[0].UserID,
			l[0].Jwt,
		})
	} else {
		c.JSON(200, response.LoginResponse{
			response.Response{
				500,
				"登录失败"},
			-1,
			"",
		})
	}
	c.Set("name", pt.Name)
	c.Next()
}
