package utils

import (
	"buyfree/repo/model"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
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

const DailyRanksKey string = "Ranks:Daily"       //val:product
const WeeklyRanksKey string = "Ranks:Weekly"     //不要求连续 val:
const MonthRanksKey string = "Ranks:Monthly"     //+month
const AnnuallyRanksKey string = "Ranks:Annually" //+year
const TOTALRanksKey string = "Ranks:Totally"

func GetProductKey(sku int64) []string {
	return []string{}
}

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
	local keys={}
	for i=1,15,1 do
		keys[i]=KEYS[i]
	end
	local array={}
	for i=7,1,-1 do
		local e=tonumber(redis.call("lindex",keys[16-i],0)) or 0
		array[i]=e
		redis.call("lpush",keys[7],e)
	end
	redis.call("ltrim",keys[7],0,6)
	return array

`)
}

/*
改变销量信息的lua脚本,暂时不支持浮点数
KEYS[1]今日开始时间,KEYS[2]昨日开始时间,KEYS[3]:日榜，KEYS[4]周榜，KEYS[5]月榜，KEYS[6]年榜，KEYS[7]7天连榜，KEYS[8]总榜
日榜keys[9]-0day,KEYS[10]-1day,keys[11]-2day,KEYS[12]-3day,KEYS[13]-4day,KEYS[14]-5day,KEYS[15]-6day
ARGV[1]订单金额
*/
func modifySales() *redis.Script {

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
		if val~=0 then
			for i=3,8,1 do
				vals[i]=tonumber(redis.call("lpop",keys[i])) or 0
				redis.call("lpush",keys[i],vals[i]+val)
			end
		end
	end
	local array={}
	for i=7,1,-1 do
		local e=tonumber(redis.call("lindex",keys[16-i],0)) or 0
		array[i]=e
		redis.call("lpush",keys[7],e)
	end
	for i=3,7,1 do
	redis.call("ltrim",keys[i],0,6)
	end
	return array
`)
}

//改变商品排行信息
func modifyranks() *redis.Script {
	return redis.NewScript(`
	local keys={}
	for i=1,10,1 do
		keys[i]=tostring(KEYS[i])
	end
	local field =tostring(ARGV[1])
	local sales =tonumber(ARGV[2])
	print(field,sales)
	for i=1,10,1 do
		redis.call("zincrby",keys[i],sales,field)
	end
	return 1
	

`)
}

//商品销量信息
func getSalesInfo() *redis.Script {
	return redis.NewScript(`
	local array={}
	for i=1,5,1 do
		array[i]=tonumber(redis.call("lindex",KEYS[i+2],0))
	end
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
func SalesOf7Days(c context.Context, rdb *redis.Client, uname string, val ...string) []int64 {
	script := salesOF7days()
	sha, _ := script.Load(c, rdb).Result()
	ret := rdb.EvalSha(c, sha, GetAllTimeKeys(uname), val)
	res, err := ret.Result()
	//fmt.Println(res, err)
	if err == nil {
		logrus.Info("获取七天销量数据失败")
	}
	var sales []int64
	for _, v := range res.([]interface{}) {
		sales = append(sales, v.(int64))
	}
	return sales
}
func ChangeAnalySalesList(c context.Context, rdb *redis.Client, keys []string, val ...string) {
	script := listPopPush()
	sha, _ := script.Load(c, rdb).Result()
	ret := rdb.EvalSha(c, sha, []string{"testlist"}, val)
	res, _ := ret.Result()
	fmt.Println("列表长度", res)
}

//改变销量信息的lua脚本
func ModifySales(c context.Context, rdb *redis.Client, uname string, val ...string) []int64 {
	script := modifySales()
	sha, _ := script.Load(c, rdb).Result()
	//sha := "bec071e3ab167970eee28b4152388cda1af7148c" //lua脚本在redis中的缓存
	ret := rdb.EvalSha(c, sha, GetAllTimeKeys(uname), val)
	res, err := ret.Result()
	if err != nil {
		fmt.Println("ERROR HAPPENS while modifying Sales", err)
	}
	var array []int64
	for _, v := range res.([]interface{}) {
		array = append(array, v.(int64))
	}
	return array
}
func ModifyProductRanks(c context.Context, rdb *redis.Client, uname, sku string, sales int64) {
	script := modifyranks()
	sha, _ := script.Load(c, rdb).Result()
	//fmt.Println(sha)
	//sha := "4de03aadfd76a083106f6183ce602e36e32fc0cb" //lua脚本在redis中的缓存
	ret := rdb.EvalSha(c, sha, GetAllProductRankKeys(uname), sku, sales) //KEYS,SKU(FIELD),SALES(SCORE)
	_, err := ret.Result()
	if err != nil {
		fmt.Println("ERROR HAPPENS ", err)
	}
	//fmt.Println(res)
}
func ModifyADRanks(c context.Context, rdb *redis.Client, adname, sku string, sales int64) {
	script := modifyranks()
	sha, _ := script.Load(c, rdb).Result()
	//sha := "bec071e3ab167970eee28b4152388cda1af7148c" //lua脚本在redis中的缓存
	ret := rdb.EvalSha(c, sha, GetAllADRankKeys(adname), sku, sales) //KEYS,SKU(FIELD),SALES(SCORE)
	_, err := ret.Result()
	if err != nil {
		fmt.Println("ERROR HAPPENS ", err)
	}
	//fmt.Println(res)
}

//获取广告或者商品的排行
func GetRankList(c context.Context, rdb *redis.Client, queryname string, mode int) ([10]model.ProductRank, error) {
	if mode < 0 || mode > 5 {
		mode = 0
	}
	ret := rdb.ZRangeWithScores(c, GetRankKeyByMode(queryname, mode), 0, 9)
	res, err := ret.Result()
	if err != nil {
		fmt.Println("get ranklist error while getting ranklist", err)
		return [10]model.ProductRank{}, err
	}
	//fmt.Println(res)
	var ranklist = [10]model.ProductRank{}
	for i, v := range res {
		ranklist[i] = model.ProductRank(v)
	}
	return ranklist, nil
}

//获取销量信息
func GetSalesInfo(c context.Context, rdb *redis.Client, uname string) []int64 {
	script := getSalesInfo()
	sha, _ := script.Load(c, rdb).Result()
	//fmt.Println(sha)
	//sha := "4de03aadfd76a083106f6183ce602e36e32fc0cb" //lua脚本在redis中的缓存
	ret := rdb.EvalSha(c, sha, GetAllTimeKeys(uname))
	res, err := ret.Result()
	fmt.Println(res)
	if err != nil {
		fmt.Println("ERROR HAPPENS while getting sales info", err)
	}
	var array []int64
	for _, v := range res.([]interface{}) {
		array = append(array, v.(int64))
	}
	return array
}
