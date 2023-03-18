package test

import (
	"buyfree/config"
	"buyfree/dal"
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestRedis(t *testing.T) {
	var ukey string = "usercount"
	GetUserCounterKey := func(uid int64) string {
		return fmt.Sprintf("%s_%d", ukey, uid)
	}
	ctx := context.TODO()
	config.Init()
	dal.InitRedis()
	RC := dal.Getrd()
	pipe := RC.Pipeline()
	userCounters := []map[string]interface{}{
		{"user_id": "155612", "got_digg_count": 10693, "got_view_count": 2238438, "followee_count": 176, "follower_count": 9895, "follow_collect_set_count": 0, "subscribe_tag_count": 95},
		{"user_id": "1111", "got_digg_count": 19, "got_view_count": 4},
		{"user_id": "2222", "got_digg_count": 1238, "follower_count": 379},
	}
	for _, counter := range userCounters {
		uid, err := strconv.ParseInt(counter["user_id"].(string), 10, 64)
		key := GetUserCounterKey(uid)
		fmt.Println(key)
		rw, err := pipe.Del(ctx, key).Result()
		if err != nil {
			fmt.Printf("del %s, rw=%d\n", key, rw)
		}
		_, err = pipe.HMSet(ctx, key, counter).Result()
		if err != nil {
			panic(err)
		}
		fmt.Printf("设置 uid=%d, key=%s\n", uid, key)
	}
	// 批量执行上面for循环设置好的hmset命令
	_, err := pipe.Exec(ctx)
	if err != nil { // 报错后进行一次额外尝试
		_, err = pipe.Exec(ctx)
		if err != nil {
			panic(err)
		}
	}
	for _, counter := range userCounters {
		uid, _ := strconv.ParseInt(counter["user_id"].(string), 10, 64)
		key := GetUserCounterKey(uid)
		fmt.Println(key)
		rs, _ := RC.HGet(ctx, key, "user_id").Result()
		fmt.Println(rs)
	}

}
