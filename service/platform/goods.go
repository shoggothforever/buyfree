package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func GetDailyRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func GetMonthlyRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func GetAnnuallyRank(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
