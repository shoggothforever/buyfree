package passenger

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
)

type BasePaController struct {
	ctx *gin.Context
}

func (b *BasePaController) Ping(c *gin.Context) {
	w := c.Writer
	w.Write([]byte("welecome to passenger.buyfree.com"))
	c.Set("hello", "How are you?")
	c.Next()
}
func (con BasePaController) Success(c *gin.Context, code int64, msg string, Res ...interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": Res,
	})
}

func (con BasePaController) Error(c *gin.Context, code int64, msg string) {
	c.JSON(200, response.Response{Code: code, Msg: msg})
}
