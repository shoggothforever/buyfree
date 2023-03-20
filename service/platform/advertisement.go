package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type ADController struct {
	BaseController
}

func (a *ADController) GetADList(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func (a *ADController) GetAddAD(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
func (a *ADController) GetADContent(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func (a *ADController) GetADEfficient(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
