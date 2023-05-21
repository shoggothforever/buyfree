package test

import (
	"buyfree/repo/model"
	"buyfree/utils"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestRedis(t *testing.T) {
	//var ukey string = "usercount"
	//GetUserCounterKey := func(uid int64) string {
	//	return fmt.Sprintf("%s_%d", ukey, uid)
	//}
	ctx := context.TODO()
	RC := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       10,
	})
	//pipe := RC.Pipeline()
	//userCounters := []map[string]interface{}{
	//	{"user_id": "155612", "got_digg_count": 10693, "got_view_count": 2238438, "followee_count": 176, "follower_count": 9895, "follow_collect_set_count": 0, "subscribe_tag_count": 95},
	//	{"user_id": "1111", "got_digg_count": 19, "got_view_count": 4},
	//	{"user_id": "2222", "got_digg_count": 1238, "follower_count": 379},
	//}
	//for _, counter := range userCounters {
	//	uid, err := strconv.ParseInt(counter["user_id"].(string), 10, 64)
	//	key := GetUserCounterKey(uid)
	//	fmt.Println(key)
	//	rw, err := pipe.Del(ctx, key).Result()
	//	if err != nil {
	//		fmt.Printf("del %s, rw=%d\n", key, rw)
	//	}
	//	_, err = pipe.HMSet(ctx, key, counter).Result()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("设置 uid=%d, key=%s\n", uid, key)
	//}
	//// 批量执行上面for循环设置好的hmset命令
	//_, err := pipe.Exec(ctx)
	//if err != nil { // 报错后进行一次额外尝试
	//	_, err = pipe.Exec(ctx)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	//for _, counter := range userCounters {
	//	uid, _ := strconv.ParseInt(counter["user_id"].(string), 10, 64)
	//	key := GetUserCounterKey(uid)
	//	fmt.Println(key)
	//	rs, _ := RC.HGet(ctx, key, "user_id").Result()
	//	fmt.Println(rs)
	//}
	pt := model.User{}
	pt.ID = 97111
	pt.PasswordSalt = "123"
	pt.Password = "123"
	pt.Role = 3
	pt.Name = "dsm"
	pt.Balance = 0
	pt.Pic = "233"
	pt.Level = 0
	//l := model.LoginInfo{123, "123", "123", "123"}
	var loginInfo model.LoginInfo
	var err error
	loginInfo.UserID = pt.ID
	loginInfo.Salt = pt.PasswordSalt
	loginInfo.Password = utils.Messagedigest5(pt.Password, pt.PasswordSalt)
	loginInfo.Jwt, err = utils.GeneraterJwt(pt.ID, pt.Name, pt.PasswordSalt)
	if err != nil {
		logrus.Info("JWT created fail")
	}
	fmt.Println(loginInfo.Jwt)
	RC.Set(ctx, loginInfo.Jwt, 1, utils.EXPIRE)
	au, _ := RC.Get(ctx, loginInfo.Jwt).Result()
	fmt.Println(au) // 打印出1 表示存在jwt ，打印出空行 表示不存在

}
