package middleware

import (
	"buyfree/dal"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

var shortKey = "shortenUrl:%s"
var local = "localhost:9003/"
var remote = "https://bf.shoggothy.xyz/"

func GetShortenKey(short string) string {
	return fmt.Sprintf(shortKey, short)
}

// 使用中间件实现重定向
func RedirectShort() gin.HandlerFunc {
	return func(c *gin.Context) {
		short := c.Request.URL.String()[1:]
		fmt.Println(short)
		rdb := dal.Getrdb()
		ctx := context.Background()
		key := GetShortenKey(remote + short)
		s, err := rdb.Get(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(500, "未能获取到对应短链接信息")
			return
		} else {
			c.Redirect(302, s)
		}
		c.Next()
	}
}
