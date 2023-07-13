package driverapp

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HomePageController struct {
	BaseDrController
}

// @Summary 车主端首页
// @Description 展示销售数据(日收入，日环比，周环比，本月收入，今日广告收入以及播放次数，两件热销商品)
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} response.HomePageResponse
// @Failure 400 {object} response.Response
// @Router /dr/home [get]
func (h *HomePageController) GetStatic(c *gin.Context) {
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		h.Error(c, 400, "获取车主信息失败")
		return
	}
	//fmt.Println(admin)
	rdb := dal.Getrdb()
	h.rwm.RLock()
	defer h.rwm.RUnlock()
	array, err := utils.GetHomeStatic(c, rdb, admin.Name)
	fmt.Println(array, err)
	if err != nil {
		h.Error(c, 400, "获取车主端首页信息失败")
		return
	}
	//fmt.Println(array)
	var static response.HomeStatic
	static.ADDailySales = array[5]
	static.MonthlySales = array[4]
	static.DailySales = array[0]
	if array[1] == 0 {
		static.DailyRatio = 0
	} else {
		static.DailyRatio = (array[0] - array[1]) / array[1]
	}

	if array[3] == 0 {
		static.WeeklyRatio = 0
	} else {
		static.WeeklyRatio = (array[2] - array[3]) / array[3]
	}
	var ids []int64
	//fmt.Println(admin.ID)
	db := dal.Getdb()
	err = db.Raw("select id from devices where owner_id = ? ", admin.ID).Find(&ids).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		fmt.Println(err)
		h.Error(c, 400, "没有绑定设备信息")
	}
	//fmt.Println(ids)
	if len(ids) != 0 {
		err = db.Raw("select sum(play_times),sum(profit) from ad_devices where device_id in ?", ids).Row().Scan(&static.ADPlayTimes, &static.ADDailySales)
		if err != gorm.ErrRecordNotFound && err != nil {
			logrus.Info(err)
			//h.Error(c, 400, "无法获取车主端广告信息")
		}
	}
	err = db.Where("device_id in ?", ids).Order("monthly_sales DESC").Limit(2).Find(&static.ProductRankList).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		logrus.Info(err)
		h.Error(c, 400, "无法获取车主端商品排行信息")
	} else if err == gorm.ErrRecordNotFound {
		static.ProductRankList = make([]model.DeviceProduct, 1)
		static.ProductRankList[0] = model.DeviceProduct{-1, -1, model.Product{
			ID:          0,
			FactoryID:   0,
			Sku:         "",
			Inventory:   0,
			Name:        "",
			Pic:         "",
			Type:        "",
			BuyPrice:    0,
			SupplyPrice: 0,
			SalesData:   model.SalesData{},
		}}
	}
	//fmt.Println(static)
	c.JSON(200, response.HomePageResponse{response.Response{200, "首页信息:"}, static})
}

// @Summary 实时更新车主地理位置信息
// @Description 传入车主地理位置信息
// @Tags Driver
// @Accept json
// @Produce json
// @Param geo body model.Geo true "实时更新车主的位置信息,传入经度，纬度,address可传"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /dr/ping [post]
func (h *HomePageController) Ping(c *gin.Context) {
	var info model.Geo
	err := c.ShouldBind(&info)
	if err != nil {
		h.Error(c, 400, "传输数据格式错误")
		return
	}
	admin, ok := utils.GetDriveInfo(c)
	if ok != true {
		h.Error(c, 400, "获取车主信息失败")
		return
	}
	rdb := dal.Getrdb()
	h.rwm.Lock()
	defer h.rwm.Unlock()
	_, err = rdb.Do(c, "geoadd", utils.DRIVERLOCATION, info.Longitude, info.Latitude, admin.Name).Result()
	if err != nil {
		h.Error(c, 400, "更新车主位置信息失败")
	} else {
		c.JSON(200, response.Response{200, "更新车主位置信息成功"})
	}
}
