package platform

import (
	"buyfree/dal"
	"buyfree/middleware"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

type ADController struct {
	BasePtController
}

// TODO:swagger
// @Summary 获取该平台的所有广告信息
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
	//iadmin, ok := c.Get(middleware.PTADMIN)
	//if ok != true {
	//	a.Error(c, 400, "获取用户信息失败")
	//	return
	//}
	//admin := iadmin.(model.Platform)
	var ads []model.Advertisement
	//dal.Getdb().Model(model.Advertisement{}).Limit(20).Where("platform_id = ? ", admin.ID).Offset(int((page - 1) * 20)).Find(&ads)
	//err := dal.Getdb().Raw("select * from advertisements inner join "+
	//	"(select id from advertisements where platform_id = ? order by profit desc limit 20 offset ?  )as lim using (id)", admin.ID, (int((page - 1) * 20))).Find(&ads).Error
	err := dal.Getdb().Raw("select * from advertisements inner join "+
		"(select id from advertisements order by profit desc limit 20 offset ?  )as lim using (id)", (int((page - 1) * 20))).Find(&ads).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		logrus.Info(err)
		a.Error(c, 400, "获取广告信息失败")
		return
	}
	c.JSON(200, response.ADResponse{
		response.Response{
			200,
			"获取广告信息成功"},
		ads})
}

// TODO:swagger
// @Summary 添加广告信息
// @Description 按照Advertisement定义的内容传递json格式的数据,无需传入平台ID
// @Tags	Platform/Advertisement
// @Accept json
// @Accept mpfd
// @Produce json
// @Param ADInfo body model.Advertisement true "传入广告描述，投放资金，广告主，预期播放次数，广告视频地址"
// @Success 201 {object} response.ADResponse
// @Failure 400 {object} response.Response
// @Router /pt/ads [post]
func (a *ADController) AddAD(c *gin.Context) {
	var ad model.Advertisement
	c.Bind(&ad)
	iadmin, ok := c.Get(middleware.PTADMIN)
	if ok != true {
		a.Error(c, 400, "获取用户信息失败")
		return
	}
	admin := iadmin.(model.Platform)
	ad.PlatformID = admin.ID
	err := dal.Getdb().Model(model.Advertisement{}).Create(&ad).Error
	ad.Profit = 0
	ad.PlayTimes = 0
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

// TODO:swagger
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

// TODO:swagger
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
	//fmt.Println(id)
	var ad model.Advertisement
	//var err error
	var devices []model.Device
	//devices, err := gen.Device.GetDeviceByAdvertiseID(id)
	//获取所有投放该广告的设备
	//err := dal.Getdb().Raw("select * from devices as d where d.id in (select device_id from ad_devices where advertisement_id = ?)", id).Find(&devices).Error
	// 使用in可能会产生性能问题，
	err := dal.Getdb().Raw("select * from devices as d where exists (select d.id from ad_devices where advertisement_id = ? and device_id = d.id)", id).Find(&devices).Error
	fmt.Println(devices, err)
	if err != nil || devices == nil || len(devices) == 0 {
		a.Error(c, 400, "获取广告信息失败")
		return
	}
	n := len(devices)
	var effinfos []response.ADEfficientInfo
	for i := 0; i < n; i++ {
		var driver *model.Driver
		var effinfo response.ADEfficientInfo
		err := dal.Getdb().Model(&model.Driver{}).Where("id = ?", devices[i].OwnerID).First(&driver).Error
		if err != gorm.ErrRecordNotFound && err != nil {
			continue
		}
		effinfo.DriverName = driver.Name
		effinfo.CarID = driver.CarID
		effinfo.DeviceID = devices[i].ID
		//ad, err = gen.Advertisement.GetAdvertisementProfitAndPlayTimes(id, devices[i].ID)
		err = dal.Getdb().Raw("select play_times,profit from ad_devices where advertisement_id=? and  device_id=?", id, devices[i].ID).Scan(&ad).Error
		if err != gorm.ErrRecordNotFound && err != nil {
			continue
		}
		effinfo.PlayedTimes = ad.PlayTimes
		effinfo.Profit = ad.Profit
		effinfos = append(effinfos, effinfo)
	}
	if len(effinfos) != 0 {
		c.JSON(200, response.ADEfficientResponse{
			response.Response{
				200,
				"获取广告播放效果成功",
			}, effinfos,
		})
	}
}
