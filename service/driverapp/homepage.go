package driverapp

import (
	"buyfree/dal"
	"buyfree/middleware"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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
	iadmin, ok := c.Get(middleware.DRADMIN)
	if ok != true {
		h.Error(c, 400, "获取用户信息失败")
		return
	}
	rdb := dal.Getrdb()
	admin := iadmin.(model.Driver)
	fmt.Println(admin)
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
			fmt.Println(err)
			h.Error(c, 400, "无法获取车主端广告信息")
		}
	}
	err = db.Raw("select * from device_products where device_id in ? order by monthly_sales DESC limit 2", ids).Find(&static.ProductRankList).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		fmt.Println(err)
		h.Error(c, 400, "无法获取车主端商品排行信息")
	}
	fmt.Println(static)
	c.JSON(200, response.HomePageResponse{response.Response{200, "首页信息:"}, static})
}
