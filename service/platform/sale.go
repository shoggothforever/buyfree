package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type SalesController struct {
	BaseController
}

// @Summary	展示销售数据
// @Description	在数据大屏上展示管理场站的销售数据
// @Tags Platform
// @Accept json
// @Accept mpfd
// @Produce json
// @Success 200 {object} response.ScreenInfoResponse
// @Failure 400 {object} response.Response
// @Router /pt/screen [get]
func (s *SalesController) GetScreenData(c *gin.Context) {
	var si response.ScreenInfo
	rdb := dal.Getrdb()
	iadmin, ok := c.Get("admin")
	if ok != true {
		s.Error(c, 400, "获取用户信息失败")
		return
	}
	admin := iadmin.(model.Platform)
	name := admin.Name
	curve := utils.SalesOf7Days(c, rdb, name)
	err := dal.Getdb().Raw("select count(*) from devices").First(&si.DevNums).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		s.Error(c, 400, "无法获取设备数量")
		return
	}
	err = dal.Getdb().Raw("select count(*) from devices where is_online = ?", true).First(&si.OnlineDevNums).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		s.Error(c, 400, "无法获取上线设备信息")
		return
	}
	err = dal.Getdb().Raw("select * from advertisements  where platform_id= ? order by profit desc limit 10", admin.ID).Find(&si.ADList).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		s.Error(c, 400, "无法获取广告信息")
		return
	}
	si.OfflineDevNums = si.DevNums - si.OnlineDevNums

	info, err := utils.GetSalesInfo(c, rdb, name)
	if err != nil {
		s.Error(c, 400, "获取用户信息失败")
		return
	}
	//fmt.Println(info)
	var salesinfo model.SalesData
	salesinfo.DailySales = info[0]
	salesinfo.WeeklySales = info[1]
	salesinfo.MonthlySales = info[2]
	salesinfo.AnnuallySales = info[3]
	salesinfo.TotalSales = info[4]
	ranklist, err := utils.GetRankList(c, rdb, utils.Ranktype1, name, 1)
	if err != nil {
		s.Error(c, 400, "获取排名信息失败")
	}
	si.SalesData = salesinfo
	si.SalesCurve = curve
	si.ProductRankList = ranklist
	c.JSON(200, response.ScreenInfoResponse{
		response.Response{200, "获取统计数据成功"},
		si})
}

func getsalesmessage(mode int64) string {
	switch mode {
	case 1:
		return "获取本周商品销售额排行数据成功"

	case 2:
		return "获取本月商品销售额排行数据成功"

	case 3:
		return "获取本年商品销售额排行数据成功"
	case 4:
		return "获取总商品销售额排行数据成功"
	default:
		return "获取本日商品销售额排行数据成功"
	}

}

//TODO: 统计数据补全计划
// @Summary	销售数据统计
// @Description	展示管理场站的销售数据，获取详细的销售排行信息
// @Tags Platform
// @Accept json
// @Accept mpfd
// @Produce json
// @Param mode path int true "根据模式获取平台商品的排行信息，mode=0 今日排行，mode=1 本周排行 ,mode=2 本月排行,mode=3 本年排行,mode=4 总排行"
// @Success 200 {object} response.SaleStaticResponse
// @Failure 400 {object} response.Response
// @Router /pt/static/{mode} [get]
func (s *SalesController) GetSales(c *gin.Context) {
	iadmin, ok := c.Get("admin")
	if ok != true {
		s.Error(c, 400, "获取用户信息失败")
		return
	}
	mode, err := strconv.ParseInt(c.Param("mode"), 10, 64)
	if err != nil {
		s.Error(c, 400, "请输入正确的模式信息")
		return
	}
	name := iadmin.(model.Platform).Name
	rdb := dal.Getrdb()
	info, err := utils.GetSalesInfo(c, rdb, name)
	if err != nil {
		s.Error(c, 400, "无法获取销量信息")
		return
	}
	var salesinfo model.SalesData
	salesinfo.DailySales = info[0]
	salesinfo.WeeklySales = info[1]
	salesinfo.MonthlySales = info[2]
	salesinfo.AnnuallySales = info[3]
	salesinfo.TotalSales = info[4]
	ranklist, err := utils.GetRankList(c, rdb, utils.Ranktype1, name, int(mode))
	if err != nil {
		s.Error(c, 400, "获取排名信息失败")
	} else {
		c.JSON(200, response.SaleStaticResponse{
			response.Response{
				200,
				getsalesmessage(mode)},
			salesinfo,
			ranklist,
		})
	}
}

////从数据库获取相关信息
//func (s *SalesController) GetDevCnt(c *gin.Context) {
//	c.JSON(200, response.Response{
//		200,
//		"ok"})
//}
//func (s *SalesController) GetLocation(c *gin.Context) {
//	c.JSON(200, response.Response{
//		200,
//		"ok"})
//}
//
//func (s *SalesController) AnalyzeAD(c *gin.Context) {
//	c.JSON(200, response.Response{
//		200,
//		"ok"})
//}
//
//func (s *SalesController) GetSaleRank(c *gin.Context) {
//	c.JSON(200, response.Response{
//		200,
//		"ok"})
//}
