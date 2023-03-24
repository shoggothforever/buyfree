package test

import (
	"buyfree/utils"
	"github.com/go-redis/redis/v8"
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
	ctx := rdb.Context()
	//utils.Lualock(ctx, rdb, key, val)
	//utils.Luaunlock(ctx, rdb, key, val)

	//t.Log(utils.ModifySales(ctx, rdb, "dsm", "2333"))
	//t.Log(utils.GetAllProductRankKeys("dsm"))
	{
		utils.ModifyProductRanks(ctx, rdb, "dsm", "sku", 123)
		utils.ModifyProductRanks(ctx, rdb, "dsm", "sku1", 12)
		utils.ModifyProductRanks(ctx, rdb, "dsm", "sku2", 123)
		utils.ModifyProductRanks(ctx, rdb, "dsm", "sku3", 1234)
		utils.ModifyProductRanks(ctx, rdb, "dsm", "sku4", 12345)
	}
	t.Log(utils.GetRankList(ctx, rdb, "dsm", 0))
	//t.Log(utils.SalesOf7Days(ctx, rdb, "dsm"))
	//t.Log(utils.GetSalesInfo(ctx, rdb, "dsm"))
	t.Log("测试通过")

}
