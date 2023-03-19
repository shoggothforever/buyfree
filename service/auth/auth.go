package auth

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/platform"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var ptadmin model.Platform
	c.Bind(&ptadmin)
	fmt.Println(ptadmin)
	fmt.Println(ptadmin.ID)
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
}
