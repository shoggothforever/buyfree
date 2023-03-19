package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func GetADList(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func GetAddAD(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func GetADContent(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func GetADEfficient(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
