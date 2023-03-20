package platform

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (con BaseController) Success(c *gin.Context, code int64, msg string, Res ...interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": Res,
	})
}

func (con BaseController) Error(c *gin.Context, code int64, msg string) {
	c.JSON(200, response.Response{Code: code, Msg: msg})
}
