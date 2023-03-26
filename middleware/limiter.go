package middleware

import (
	"github.com/gin-gonic/gin"
)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		//rdb := dal.Getrdb()
		//limiter, err := redis_rate.NewLimiter()
	}
}
