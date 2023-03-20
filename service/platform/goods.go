package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

func (g *GoodsController) GetDailyRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func (g *GoodsController) GetMonthlyRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func (g *GoodsController) GetAnnuallyRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
