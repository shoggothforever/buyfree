package driverapp

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type BaseDrController struct {
	ctx *gin.Context
}

func (b *BaseDrController) Ping(c *gin.Context) {
	w := c.Writer
	w.Write([]byte("welecome to driver.buyfree.com"))
	c.Set("hello", "How are you?")
	c.Next()
}
func (b BaseDrController) Success(c *gin.Context, code int64, msg string, Res ...interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": Res,
	})
}

func (b BaseDrController) Error(c *gin.Context, code int64, msg string) {
	c.JSON(200, response.Response{Code: code, Msg: msg})
}
