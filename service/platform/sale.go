package platform

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"buyfree/service/response"
	"buyfree/utils"
	"github.com/gin-gonic/gin"
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
	dal.Getdb().Raw("select count(*) from advertisements").First(&si.DevNums)
	c.JSON(200, response.ScreenInfoResponse{
		response.Response{200, "获取统计数据成功"},
		si})
}

//TODO: 统计数据补全计划
// @Summary	销售数据统计
// @Description	展示管理场站的销售数据，获取详细的销售排行信息
// @Tags Platform
// @Accept json
// @Accept mpfd
// @Produce json
// @Param mode path int true "mode=0 今日排行，mode=1 本周排行 ,mode=2 本月排行,mode=3 本年排行,mode=4 总排行"
// @Success 200 {object} response.SaleStaticResponse
// @Failure 400 {object} response.Response
// @Router /pt/static/{mode} [get]
func (s *SalesController) GetSales(c *gin.Context) {
	iadmin, ok := c.Get("admin")
	mode, _ := strconv.ParseInt(c.Param("mode"), 10, 64)
	name := iadmin.(model.Platform).Name
	if ok != true {
		s.Error(c, 400, "获取用户信息失败")
	}
	rdb := dal.Getrdb()
	info := utils.GetSalesInfo(c, rdb, name)
	var salesinfo model.SalesData
	salesinfo.DailySales = float64(info[0])
	salesinfo.DailySales = float64(info[1])
	salesinfo.DailySales = float64(info[2])
	salesinfo.DailySales = float64(info[3])
	salesinfo.DailySales = float64(info[4])
	ranklist, err := utils.GetRankList(c, rdb, name, int(mode))
	if err != nil {
		s.Error(c, 400, "获取排名信息失败")
	}
	c.JSON(200, response.SaleStaticResponse{
		response.Response{
			200,
			"获取销售数据成功"},
		salesinfo,
		ranklist,
	})
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
