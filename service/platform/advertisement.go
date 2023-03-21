package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ADController struct {
	BaseController
}

func (a *ADController) GetADList(c *gin.Context) {

	var ads []model.Advertisement
	dal.Getdb().Model(model.Advertisement{}).Find(&ads)

	c.JSON(200, response.ADResponse{
		response.Response{
			200,
			"ok"},
		ads})
}
func (a *ADController) AddAD(c *gin.Context) {
	var ad []model.Advertisement
	c.ShouldBind(&ad)
	if len(ad) == 0 {
		a.Error(c, 400, "添加广告信息失败")
		return
	}
	err := dal.Getdb().Model(model.Advertisement{}).Create(&ad[0]).Error
	if err == nil {
		c.JSON(200, response.ADResponse{
			response.Response{200,
				"ok"},
			ad,
		})
	} else {
		a.Error(c, 400, "添加广告信息失败")
	}
}
func (a *ADController) GetADContent(c *gin.Context) {
	//TODO:交给前端吧
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var ad model.Advertisement
	err := dal.Getdb().Model(&model.Advertisement{}).Where("id=?", id).First(&ad).Error
	if err != nil {
		a.Error(c, 400, "获取广告信息失败")
	}
	c.JSON(200, response.ADInfoResponse{
		response.Response{
			200,
			"ok"},
		ad})
}

func (a *ADController) GetADEfficient(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var ad model.Advertisement
	//var drivers []*model.Driver
	//var driver *model.Driver
	var devices []*model.Device
	var effinfo []*response.ADEfficientInfo
	dal.Getdb().Model(&model.Advertisement{}).Where("id = ?", id).First(&ad)

	//TODO:好好写原生SQL语句
	dal.Getdb().Model(&model.Device{}).Raw("select * from devices as d where d. id in "+
		"(select device_id where advertisement_id = ? )", id).Find(&devices)
	c.JSON(200, response.ADEfficientResponse{
		response.Response{
			200,
			"获取广告播放效果成功",
		}, effinfo,
	})
}
