package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetOnShelf(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

//从数据库获取相关信息
func Getsoldout(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func Getdownshelf(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetGoodinfo(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func TakeOn(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
