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

	t.Log(utils.AddSales(ctx, rdb, utils.GetAllKeys("dsm"), "233"))
	t.Log("测试通过")

}
