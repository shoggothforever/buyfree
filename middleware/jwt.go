package middleware

import (
	"buyfree/dal"
	"buyfree/repo/model"
	"github.com/gin-gonic/gin"
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
		if len(jwt) == 0 {
			ijwt, ok := c.Get("Authorization")
			if ok != true {
				c.Set("AUthInfo", "Failed!")
				return
			}
			jwt = ijwt.(string)
		}
		if len(jwt) > 7 { //为Bearer Token去除前七位数据
			jwt = jwt[7:]
		} else {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 401, "msg": "验证信息失败",
			})
			return
		} //如果在PostMan中使用 Bearer Token 会在jwt前加上bearer: 前缀
		authjwt, err := dal.Getrdb().Get(c, jwt).Result()
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
			var id int64
			dal.Getdb().Raw("select user_id from login_infos where jwt=?", jwt).First(&id)
			var ptadmin []model.Platform

			var passenger []model.Passenger
			dal.Getdb().Model(&model.Passenger{}).Where("id=?", id).First(&passenger)
			if len(passenger) != 0 {
				c.Set(DRADMIN, passenger[0])
				return
			}
			dal.Getdb().Model(&model.Platform{}).Where("id=?", id).First(&ptadmin)
			if len(ptadmin) != 0 {
				c.Set(PTADMIN, ptadmin[0])
				return
			}
			var dradmin []model.Driver
			dal.Getdb().Model(&model.Driver{}).Where("id=?", id).First(&dradmin)
			if len(dradmin) != 0 {
				c.Set(DRADMIN, dradmin[0])
				return
			}
			var fadmin []model.Factory
			dal.Getdb().Model(&model.Factory{}).Where("id=?", id).First(&fadmin)
			if len(fadmin) != 0 {
				c.Set(FADMIN, fadmin[0])
				return
			}

			c.Next()
		}
	}
}
