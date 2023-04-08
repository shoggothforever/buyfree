package test

import (
	"buyfree/logger"
	"buyfree/utils"
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"testing"
)

var key = []string{"mlock"}
var val = "42"

func TestLua(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       10,
	})
	ctx := context.TODO()
	var uname string = "dsm"
	sales, err := rdb.Get(ctx, utils.GetSalesKeyByMode(utils.Ranktype1, uname, 0)).Result()
	if sales == "" {
		logger.Loger.Info(sales)
		logger.Loger.Info(err)
	}
	t.Log(strconv.ParseFloat(sales, 64))
	//t.Log(utils.GetSalesKeyByMode(utils.Ranktype1, uname, 0))
	//utils.Lualock(ctx, rdb, []string{"mylock"}, val, "30")
	//utils.Luaunlock(ctx, rdb, key, val)
	//utils.ChangeTodaySales(ctx, rdb, key, "123")
	//更新销售额榜单信息
	//t.Log(utils.ModifySales(ctx, rdb, utils.Ranktype1, uname, "2333.123"))

	//t.Log("测试根据类型和用户名生成所有的排行榜redis键名")
	//t.Log(utils.GetAllTypeRankKeys(utils.Ranktype1, "dsm"))
	//更改商品排行信息
	{
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "yith", 123)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "9e", 12)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "2a", 123)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "3s", 1234)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "4b", 12345)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "yith", 123)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "9a", 12)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "2e", 123)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "3s", 1234)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "4b", 12345)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "yith", 123)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "9e", 12)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "2a", 123)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "3s", 1234)
		//utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "4b", 12345)
	}

	//测试根据ranktype获得的KEYS,返回值为{[]{username:score}}
	//t.Log(utils.GetRankList(ctx, rdb, utils.Ranktype1, uname, 0))
	//测试根据Ranktype和用户名获得七日销售数据
	//t.Log(utils.SalesOf7Days(ctx, rdb, utils.Ranktype1, uname))
	//t.Log(utils.GetSalesInfo(ctx, rdb, utils.Ranktype1, uname))
	//t.Log(utils.GetDriverSalesKeys(uname))
	//t.Log(utils.GetHomeStatic(ctx, rdb, uname))
	//获取司机首页需要显示的信息
	//t.Log(utils.GetDriverSalesKeys(uname))

	//t.Log("测试通过")
}
