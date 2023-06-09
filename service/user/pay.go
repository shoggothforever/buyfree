package user

import (
	"buyfree/service/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pay(c *gin.Context) {
	if c.GetBool("PAYCHECK") == true {
		//TODO 数据库操作
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusOK,
			Msg:  "支付成功",
		})
	} else {
		c.JSON(http.StatusOK, response.Response{
			Code: http.StatusBadRequest,
			Msg:  "支付失败",
		})
	}
}
