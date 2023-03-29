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
// @Description 展示销售数据
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
	err = dal.Getdb().Raw("select id from devices where owner_id = ? ", admin.ID).First(&static.ADPlayTimes).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		fmt.Println(err)
		h.Error(c, 400, "获取设备信息失败")
		return
	}
	if len(ids) != 0 {
		err = dal.Getdb().Raw("select sum(play_times) from ad_devices where device_id in ?", ids).First(&static.ADPlayTimes).Error
		if err != gorm.ErrRecordNotFound && err != nil {
			fmt.Println(err)
			h.Error(c, 400, "获取车主端首页信息失败")
			return
		}
	}
	c.JSON(200, response.HomePageResponse{response.Response{200, "获取首页信息成功"}, static})
}
