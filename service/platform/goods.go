package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

func (g *GoodsController) GetRank(c *gin.Context) {
	//mode := c.Param("mode")

	c.JSON(200, response.Response{
		200,
		"ok"})
}
