package middleware

import (
	"buyfree/dal"
	"github.com/gin-gonic/gin"
)

//使用中间件实现鉴权
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("Authorization")
		if len(jwt) > 7 { //为Bearer Token去除前七位数据
			jwt = jwt[7:]
		} else {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 401, "msg": "验证信息失败",
			})
			return
		} //如果在PostMan中使用 Bearer Token 会在jwt前加上bearer: 前缀
		authjwt, err := dal.Getrd().Get(c, jwt).Result()
		//fmt.Println(authjwt)
		if err != nil {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 401, "msg": "验证信息失败",
			})
			return
		} else if authjwt != "1" {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 401, "msg": "验证信息失败",
			})
			return
		} else {
			c.Set("AuthInfo", "Success!")
			c.Set("Jwt", jwt)
			c.Next()
		}
	}
}
