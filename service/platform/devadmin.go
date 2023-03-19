package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func GetAlldev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetOndev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetOffdev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func GetActivated(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func GetNotActivated(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func AddDev(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
