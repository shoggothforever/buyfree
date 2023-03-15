package dal

import (
	"buyfree/config"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var Ctx = context.Background()
var Redisclient *redis.Client
var addr, password string

func readRedisInfo() {
	info := config.Reader.GetStringMapString("redis")
	addr = info[config.Redisaddr]
	password = info[config.Redispassword]
}
func CloseDB() {
	DB.Close()
	Redisclient.Close()
}
func InitRedis() {
	readRedisInfo()
	Redisclient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       10,
	})
	pong, err := Redisclient.Ping(Ctx).Result()
	if err != nil {
		logrus.Info(err)

	} else {
		logrus.Info(pong, "成功连接redis")
	}

}
