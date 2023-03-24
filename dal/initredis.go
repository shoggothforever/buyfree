package dal

import (
	"buyfree/config"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var Ctx = context.Background()
var RDB *redis.Client
var addr, password string

func Getrdb() *redis.Client {
	return RDB
}
func readRedisInfo() {
	info := config.Reader.GetStringMapString("redis")
	addr = info[config.Redisaddr]
	password = info[config.Redispassword]
}
func CloseDB() {
	//DB.Close()
	RDB.Close()
}
func init() {
	readRedisInfo()
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       10,
	})
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		logrus.Info(err)

	} else {
		logrus.Info("成功连接redis")
	}

}
