package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

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
func GetSales(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetCurve(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func GetDevCnt(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func GetLocation(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func AnalyzeAD(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetSaleRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
