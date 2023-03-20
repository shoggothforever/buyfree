package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type SalesController struct {
	BaseController
}

//	@Summary		获取营销额
//	@Tags			Screen
//	@Accept			json
//	@Description	get platform's sales
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"待填"
//	@Param			id				formData	int		true	"待填"
//	@Failure		403				{object}	model.Response
//	@Success		200				{object}	model.Response
//	@Router			/url/pause [Post]
func (s *SalesController) GetSales(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func (s *SalesController) GetCurve(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func (s *SalesController) GetDevCnt(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func (s *SalesController) GetLocation(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func (s *SalesController) AnalyzeAD(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func (s *SalesController) GetSaleRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
