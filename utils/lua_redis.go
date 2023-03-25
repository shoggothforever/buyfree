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
const TOTALSalesKey string = "Sales:Totally"
const ExLock time.Duration = 30

//加锁lua脚本,设置过期时间 ExLock
func luaLock() *redis.Script {
	return redis.NewScript(`
	local lockKey = tostring(KEYS[1])
	local val,dur = tostring(ARGV[1]),tonumber(ARGV[2])
	local ok= redis.call("set",lockKey,val)
	if ok~=nil then
		redis.call("expire",lockKey,dur)
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

//向特定的榜单中添加数据
func listPopPush() *redis.Script {
	return redis.NewScript(`
		local key =KEYS[1]
		local val =tonumber(ARGV[1])
		print(tonumber(val))
		local lval =tonumber(redis.call("lpop",key))
		if lval ==nil then 
		return 0
		end
		print(tonumber(lval))
		local ll =tonumber(redis.call("lpush",key,lval+val))
		if ll ==nil then
		return 0
		end
		return ll
`)
}

//获取连续七天销量信息
func salesOF7days() *redis.Script {
	return redis.NewScript(`
	local key=tostring(KEYS[1])
	local st,ed=tonumber(ARGV[1]),tonumber(ARGV[2])
	local array={}
	local exec=redis.call("ltrim",key,0,6)
	if exec==0 then 
	return exec
	end
	for i=1,7,1 do
	local e=tonumber(redis.call("lpop",key))
	array[i]=tonumber(e)
	redis.call("rpush",key,e)
	end
	return array

`)
}

//改变销量信息的lua脚本
func addSales() *redis.Script {
	//KEYS[1]今日开始时间,KEYS[2]昨日开始时间,KEYS[3]:日榜，KEYS[4]周榜，KEYS[5]月榜，KEYS[6]年榜，KEYS[7]7天连榜，KEYS[8]总榜
	//日榜keys[9]-0day,KEYS[10]-1day,keys[11]-2day,KEYS[12]-3day,KEYS[13]-4day,KEYS[14]-5day,KEYS[15]-6day
	//ARGV[1]订单金额
	return redis.NewScript(`
	local keys={}
	local val=tonumber(ARGV[1])
	local vals={}
	for i=1,15,1 do
	keys[i]=tostring(KEYS[i])
	end
	local bit =tonumber(redis.call("getbit",keys[1],0))
    if bit == 0 then
		redis.call("setbit",keys[1],0,1)
		redis.call("del",keys[2])
		redis.call("lpush",keys[3],0)
		vals[2]=tonumber(redis.call("lpop",keys[3])) or 0
		redis.call("lpush",keys[3],val+vals[2])
	else
		vals[3]=tonumber(redis.call("lpop",keys[3])) or 0
		redis.call("lpush",keys[3],vals[3]+val)

		vals[4]=tonumber(redis.call("lpop",keys[4])) or 0
		redis.call("rpop",keys[4])
		redis.call("lpush",keys[4],vals[4]+val)

		vals[5]=tonumber(redis.call("lpop",keys[5])) or 0
		redis.call("rpop",keys[5])
		redis.call("lpush",keys[5],vals[5]+val)

		vals[6]=tonumber(redis.call("lpop",keys[6])) or 0
		redis.call("rpop",keys[6])
		redis.call("lpush",keys[6],vals[6]+val)

		vals[7]=tonumber(redis.call("lpop",keys[7])) or 0
		redis.call("lpush",keys[7],vals[7]+val)

		vals[8]=tonumber(redis.call("lpop",keys[8])) or 0
		redis.call("lpush",keys[8],vals[8]+val)

	end
	local array={}
	for i=7,1,-1 do
		local e=tonumber(redis.call("lindex",keys[16-i],0)) or 0
		array[8-i]=e
		redis.call("lpush",keys[7],e)
	end
	redis.call("ltrim",keys[7],0,6)
	redis.call("ltrim",keys[6],0,6)
	redis.call("ltrim",keys[5],0,6)
	redis.call("ltrim",keys[4],0,6)
	redis.call("ltrim",keys[3],0,6)
	return array
`)
}

func Lualock(c context.Context, rdb *redis.Client, key []string, val ...string) {
	script := luaLock()
	sha, _ := script.Load(c, rdb).Result()
	ret := rdb.EvalSha(c, sha, key, val, ExLock)
	res, _ := ret.Result()
	fmt.Println("加锁结果", res)
}
func Luaunlock(c context.Context, rdb *redis.Client, key []string, val ...string) {
	script := luaUnlock()
	sha, _ := script.Load(c, rdb).Result()
	res, _ := rdb.EvalSha(c, sha, key, val).Result()
	fmt.Println("解锁结果", res)
}
func ChangeTodaySales(c context.Context, rdb *redis.Client, key []string, val ...string) {
	script := listPopPush()
	sha, _ := script.Load(c, rdb).Result()
	ret := rdb.EvalSha(c, sha, []string{"testlist"}, val)
	res, _ := ret.Result()
	fmt.Println("列表长度", res)
}
func SalesOf7Days(c context.Context, rdb *redis.Client, key []string, val ...string) []int64 {
	script := salesOF7days()
	sha, _ := script.Load(c, rdb).Result()
	ret := rdb.EvalSha(c, sha, []string{"testlist"}, val)
	res, _ := ret.Result()
	return res.([]int64)
}
func ChangeAnalySalesList(c context.Context, rdb *redis.Client, keys []string, val ...string) {
	script := listPopPush()
	sha, _ := script.Load(c, rdb).Result()
	ret := rdb.EvalSha(c, sha, []string{"testlist"}, val)
	res, _ := ret.Result()
	fmt.Println("列表长度", res)
}

//改变销量信息的lua脚本
func AddSales(c context.Context, rdb *redis.Client, key []string, val ...string) []int64 {
	script := addSales()
	sha, _ := script.Load(c, rdb).Result()
	//sha := "bec071e3ab167970eee28b4152388cda1af7148c" //lua脚本在redis中的缓存
	ret := rdb.EvalSha(c, sha, GetAllKeys("dsm"), val)
	res, err := ret.Result()
	if err != nil {
		fmt.Println("ERROR HAPPENS", err)
	}
	var array []int64
	for _, v := range res.([]interface{}) {
		array = append(array, v.(int64))
	}
	return array
}
