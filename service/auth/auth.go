package auth

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}

func Login(c *gin.Context) {
	c.JSON(200, response.Response{
		200,
		"ok"})
}
