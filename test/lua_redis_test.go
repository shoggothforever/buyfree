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
	//t.Log(utils.GetAllTypeRankKeys(utils.Ranktype1, "dsm"))
	//{
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, "dsm", "0", 123)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, "dsm", "1", 12)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, "dsm", "2", 123)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, "dsm", "3", 1234)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, "dsm", "4", 12345)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, "dsm", "0", 123)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, "dsm", "1", 12)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, "dsm", "2", 123)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, "dsm", "3", 1234)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, "dsm", "4", 12345)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, "dsm", "0", 123)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, "dsm", "1", 12)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, "dsm", "2", 123)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, "dsm", "3", 1234)
	//	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, "dsm", "4", 12345)
	//}
	//t.Log(utils.GetRankList(ctx, rdb, "dsm", 0))
	//t.Log(utils.SalesOf7Days(ctx, rdb, "dsm"))
	t.Log(utils.GetSalesInfo(ctx, rdb, "dsm"))
	t.Log("测试通过")

}
