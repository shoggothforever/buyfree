package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func AnaSales(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func LsDev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func LsDriver(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func LsDevProduct(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func TakeDown(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
