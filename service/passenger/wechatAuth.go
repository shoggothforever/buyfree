package passenger

import (
	"buyfree/config"
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"strconv"
)

var loginurl = "appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

type WeiXinAuthController struct {
	BasePaController
}

// @Summary 用户登录
// @Description 小程序提供code，操作流程按照https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/login.html指示，前端调用wx.login()获取临时登录凭证code，以formData的格式传到后端服务器
// @Tags Passenger/Authentication
// @Accept json
// @Produce json
// @Param code formData string true "登录凭证"
// @Success 200 {object} response.WeiXinLoginResponse
// @Failure 500 {object} response.Response
// @Failure 40029 {object} response.Response
// @Failure 45011 {object} response.Response
// @Failure 40226 {object} response.Response
// @Router /login [post]
func (w *WeiXinAuthController) Login(c *gin.Context) {
	code := c.PostForm("code")
	fmt.Println(code)
	client := resty.New()
	//var res interface{}
	var authres response.WeiXinLoginInfo

	ret, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryString(fmt.Sprintf(loginurl, config.APPID, config.APPSECRET, code)).
		//SetResult(&res).
		Get("https://api.weixin.qq.com/sns/jscode2session")
	if err != nil {
		w.Error(c, 500, "服务器处理错误")
	}
	err = utils.Json.Unmarshal(ret.Body(), &authres)
	if err != nil {
		w.Error(c, 500, "服务器处理错误")
	}
	if authres.ErrCode == 0 {
		openid, _ := strconv.ParseInt(authres.UnionID, 10, 64)
		var token = utils.DoubleMessagedigest5(authres.SessionKey, authres.OpenID)
		info := model.NewLoginInfo(openid, 0, authres.OpenID, "", "", token)
		err := dal.Getdb().Model(&model.LoginInfo{}).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"jwt"}),
		}).Create(&info).Error
		if err != nil {
			w.Error(c, 500, "服务器处理错误")
			return
		} else {

			rdb := dal.Getrdb()
			if _, rerr := rdb.SetEX(rdb.Context(), token, 1, utils.WechatExpire).Result(); rerr != nil {
				logrus.Info("存储用户登录状态失败")
				w.Error(c, 500, "服务器处理错误")
			}
			c.JSON(200, response.WeiXinLoginResponse{
				response.Response{200, "请求验证信息成功"},
				authres.OpenID,
				authres.UnionID,
				authres.ErrCode,
				authres.ErrMsg,
				token,
			})
		}
	} else if authres.ErrCode == 40029 {
		w.Error(c, 40029, "js_code无效")
	} else if authres.ErrCode == 45011 {
		w.Error(c, 45011, "高风险等级用户，小程序登录拦截")
	} else if authres.ErrCode == 40226 {
		w.Error(c, 40226, "高风险等级用户，小程序登录拦截")
	} else {
		w.Error(c, 500, "系统繁忙，请开发者稍候再试")
	}
}
