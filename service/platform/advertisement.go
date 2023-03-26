package platform

import (
	"buyfree/dal"
	"buyfree/repo/gen"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ADController struct {
	BaseController
}

//TODO:swagger
// @Summary 获取所有广告信息
// @Description
// @Tags	Platform/Advertisement
// @Accept json
// @Accept mpfd
// @Produce json
// @Param page path int true "跳转到的页数，起始为第一页"
// @Success 200 {object} response.ADResponse
// @Failure 400 {object} response.Response
// @Router /pt/ads/list/{page} [get]
func (a *ADController) GetADList(c *gin.Context) {
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	var ads []model.Advertisement
	dal.Getdb().Model(model.Advertisement{}).Limit(20).Offset(int((page - 1) * 20)).Find(&ads)

	c.JSON(200, response.ADResponse{
		response.Response{
			200,
			"ok"},
		ads})
}

//TODO:swagger
// @Summary 添加广告信息
// @Description 按照Advertisement定义的内容传递json格式的数据
// @Tags	Platform/Advertisement
// @Accept json
// @Accept mpfd
// @Produce json
// @Success 201 {object} response.ADResponse
// @Failure 400 {object} response.Response
// @Router /pt/ads [post]
func (a *ADController) AddAD(c *gin.Context) {
	var ad model.Advertisement
	c.Bind(&ad)
	err := dal.Getdb().Model(model.Advertisement{}).Limit(20).Create(&ad).Error
	if err == nil {
		ad.ID = utils.IDWorker.NextId()
		c.JSON(200, response.ADResponse{
			response.Response{201,
				"添加广告信息成功"},
			[]model.Advertisement{ad},
		})
	} else {
		a.Error(c, 400, "添加广告信息失败")
	}
}

//TODO:swagger
// @Summary 获取单个广告信息
// @Description 传入广告ID
// @Tags	Platform/Advertisement
// @Accept json
// @Accept mpfd
// @Produce json
// @Param id path int true "广告ID"
// @Success 200 {object} response.ADResponse
// @Failure 400 {object} response.Response
// @Router /pt/ads/infos/{id} [get]
func (a *ADController) GetADContent(c *gin.Context) {
	//TODO:交给前端吧
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var ad model.Advertisement
	err := dal.Getdb().Model(&model.Advertisement{}).Where("id=?", id).First(&ad).Error
	if err != nil {
		a.Error(c, 400, "获取广告信息失败")
	} else {
		c.JSON(200, response.ADResponse{
			response.Response{
				200,
				"获取广告信息成功"},
			[]model.Advertisement{ad}})
	}
}

//TODO:swagger
// @Summary 获取单个广告效益
// @Description 传入广告ID
// @Tags	Platform/Advertisement
// @Accept json
// @Accept mpfd
// @Produce json
// @Param id path int true "广告ID"
// @Success 200 {object} response.ADEfficientResponse
// @Failure 400 {object} response.Response
// @Router /pt/ads/efficient/{id} [get]
func (a *ADController) GetADEfficient(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println(id)
	var ad model.Advertisement
	var err error
	var devices []model.Device
	devices, err = gen.Device.GetDeviceByAdvertiseID(id)
	if err != nil {
		a.Error(c, 400, "获取广告信息失败")
		return
	}
	fmt.Println(devices)
	n := len(devices)
	effinfo := make([]response.ADEfficientInfo, n, n)
	for i := 0; i < n; i++ {
		var driver *model.Driver
		dal.Getdb().Model(&model.Driver{}).Select("name", "car_id").Where("id = ?", devices[i].OwnerID).First(&driver)
		effinfo[i].DriverName = driver.Name
		effinfo[i].CarID = driver.CarID
		effinfo[i].DeviceID = devices[i].ID
		ad, err = gen.Advertisement.GetAdvertisementProfitAndPlayTimes(id, devices[i].ID)
		if err != nil {
			a.Error(c, 400, "获取广告信息失败")
			return
		}
		//fmt.Println(ad)
		effinfo[i].PlayedTimes = ad.PlayTimes
		effinfo[i].Profit = ad.Profit
		effinfo[i].Profit = ad.Profit
	}
	if err == nil {
		c.JSON(200, response.ADEfficientResponse{
			response.Response{
				200,
				"获取广告播放效果成功",
			}, effinfo,
		})
	}
}
