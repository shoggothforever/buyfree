package dal

import (
	"buyfree/config"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var Ctx = context.Background()
var rdb *redis.Client
var addr, password string

func Getrdb() *redis.Client {
	return rdb
}
func readRedisInfo() {
	info := config.Reader.GetStringMapString("redis")
	addr = info[config.Redisaddr]
	password = info[config.Redispassword]
}
func CloseDB() {
	//DB.Close()
	rdb.Close()
}

var Ptimers sync.Pool

func init() {
	readRedisInfo()
	//单机部署
	rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           10,
		ReadTimeout:  time.Millisecond * time.Duration(500),
		WriteTimeout: time.Millisecond * time.Duration(500),
		PoolSize:     64,
		MinIdleConns: 16,
		PoolFIFO:     true,
		MaxRetries:   3,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		logrus.Info(err)
	} else {
		logrus.Info("成功连接redis")
	}

}
