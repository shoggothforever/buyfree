package middleware

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	PAADMIN string = "paadmin"
	PTADMIN string = "ptadmin"
	DRADMIN string = "dradmin"
	FADMIN  string = "fadmin"
)

// 使用中间件实现鉴权
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("Authorization")
		if len(jwt) > 7 { //为Bearer Token去除前七位数据
			jwt = jwt[7:]
		} else {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 401, "msg": "验证信息失败1",
			})
			return
		} //如果在PostMan中使用 Bearer Token 会在jwt前加上bearer: 前缀
		rdb := dal.Getrdb()
		ctx := context.Background()
		s, err := rdb.Get(ctx, jwt).Result()
		authjwt, _ := strconv.ParseInt(s, 10, 64)
		//logger.Loger.Info("从redis中获取用户信息", authjwt)
		//logger.Loger.Info("redis操作返回err", err)
		if err != nil {
			//logger.Loger.Info(err)
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 401, "msg": "验证信息失败2",
			})
			return
		} else {
			c.Set("AuthInfo", "Success!")
			c.Set("Jwt", jwt)
			var id int64
			err = dal.Getdb().Raw("select user_id from login_infos where jwt=?", jwt).First(&id).Error
			if err != nil {
				c.Set("AUthInfo", "Failed!")
				c.AbortWithStatusJSON(200, gin.H{
					"code": 401, "msg": "验证信息失败4",
				})
				return
			}
			if authjwt == model.PASSENGER {
				//logger.Loger.Info("验证乘客信息")
				var passenger []model.Passenger
				dal.Getdb().Model(&model.Passenger{}).Where("id=?", id).First(&passenger)
				if len(passenger) != 0 {
					c.Set(PAADMIN, passenger[0])
					c.Next()
				}
			} else if authjwt == model.PLATFORMADMIN {
				//logger.Loger.Info("验证平台信息")
				var ptadmin []model.Platform
				dal.Getdb().Model(&model.Platform{}).Where("id=?", id).First(&ptadmin)
				if len(ptadmin) != 0 {
					c.Set(PTADMIN, ptadmin[0])
					c.Next()
				}
			} else if authjwt == model.DRIVER {
				//logger.Loger.Info("验证车主信息")
				var dradmin []model.Driver
				dal.Getdb().Model(&model.Driver{}).Where("id=?", id).First(&dradmin)
				if len(dradmin) != 0 {
					c.Set(DRADMIN, dradmin[0])
					c.Next()
				}
			} else if authjwt == model.FACTORYADMIN {
				//logger.Loger.Info("验证场站信息")
				var fadmin []model.Factory
				dal.Getdb().Model(&model.Factory{}).Where("id=?", id).First(&fadmin)
				if len(fadmin) != 0 {
					c.Set(FADMIN, fadmin[0])
					c.Next()
				}
			}
		}
		c.Next()
	}
}
