package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//热销榜，排行榜data struct ZSET  key:data    val: productname  score sale func:zrankbyscore,
//清理过去数据使用 ZREMRANGEBYLEX
//统计七天销售数据 data struct List key:date	val:sales 策略：每天0点 使用lpush操作添加
const DailySalesKey string = "Sales:Daily"          //val:product
const Constantly7aysSalesKey string = "Sales:7days" //要求连续 val:
const WeeklySalesKey string = "Sales:Weekly"        //不要求连续 val:
const MonthSalesKey string = "Sales:Monthly"        //+month
const AnnuallySalesKey string = "Sales:Annually"    //+year
//加锁lua脚本
func luaLock() *redis.Script {
	return redis.NewScript(`
	local lockKey = tostring(KEYS[1])
	local val = tostring(ARGV[1])
	local ok= redis.call("set",lockKey,val)
	if ok~=0 then
		return 1
	end
	return 0
`)
}

//解锁lua脚本
func luaUnlock() *redis.Script {
	return redis.NewScript(`
	local key = tostring(KEYS[1])
	local userid =tostring(ARGV[1])
	if redis.call("get",key)==ARGV[1] then
		return redis.call("del",key)
	end
`)
}

//改变销量信息的lua脚本
func luaAddSales(goodsID int64) *redis.Script {

	return redis.NewScript(`
	local

`)

}

//获得一天开头的确切时间
func GetBeginningOfTheDay() string {
	y, m, d := time.Now().In(time.Local).Date()
	return fmt.Sprintf("%d-%d-%d 00:00:00", y, m, d)
}

//获取每周第一天（周一）的日期
func GetFirstDayOfWeek() (int, int, int) {
	now := time.Now().In(time.Local)

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	y, m, d := time.Now().In(time.Local).AddDate(0, 0, offset).Date()
	return y, int(m), d
}

//获取每月第一天的日期
func GetFirstDayOfMonth() (int, int, int) {
	now := time.Now().In(time.Local)
	y, m, _ := now.Date()
	return y, int(m), 1
}

//获取每年第一天的日期
func GetFirstDayOfYear() (int, int, int) {
	now := time.Now().In(time.Local)
	y, _, _ := now.Date()
	return y, 1, 1
}

//根据模式获取相应的时间 0一天的开始，1：一周的开始，2：当月第一天，3：当年第一天.4:连续七天
func GetTimeKeyByMode(uname string, mode int) string {
	now := time.Now().In(time.Local)
	y, timem, d := now.Date()
	m := int(timem)
	if mode == 0 {
		return fmt.Sprintf("%s:%s:%d-%d-%d 00:00:00:", uname, DailySalesKey, y, m, d)
	} else if mode == 1 {
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6
		}
		y, m, d := time.Now().In(time.Local).AddDate(0, 0, offset).Date()
		return fmt.Sprintf("%s:%s:%d-%d-%d 00:00:00:", uname, WeeklySalesKey, y, m, d)
	} else if mode == 2 {
		return fmt.Sprintf("%s:%s:%d-%d-%d 00:00:00:", uname, MonthSalesKey, y, m, 1)
	} else if mode == 3 {
		return fmt.Sprintf("%s:%s:%d-%d-%d 00:00:00:", uname, AnnuallySalesKey, y, 1, 1)
	} else if mode == 4 {
		return fmt.Sprintf("%s:%s:", uname, Constantly7aysSalesKey)
	}
	return "root:root:2006-1-2 15:-4:-5"
}

var key = []string{"mlock"}
var val = "42"

func Lualock(c context.Context, rdb *redis.Client, key []string, val string) {
	script := luaLock()
	sha, _ := script.Load(c, rdb).Result()
	ret := rdb.EvalSha(c, sha, key, val)
	res, _ := ret.Result()
	fmt.Println("加锁结果", res)
}
func Luaunlock(c context.Context, rdb *redis.Client, key []string, val string) {
	script := luaUnlock()
	sha, _ := script.Load(c, rdb).Result()
	res, _ := rdb.EvalSha(c, sha, key, val).Result()
	fmt.Println("解锁结果", res)
}
