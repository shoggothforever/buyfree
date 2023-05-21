package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
)

//key:utils.LOCATION,lgt 经度，lat 纬度 loc 地名
func LocAdd(c context.Context, rdb *redis.ClusterClient, key, lgt, lat, loc string) (interface{}, error) {
	return rdb.Do(c, "geoadd", key, lgt, lat, loc).Result()
}

// source:源点 destination: 终点 unit:默认为m,可选km:千米,mi:英里,ft：英尺)
func LocDist(c context.Context, rdb *redis.ClusterClient, key, source, destination, unit string) (interface{}, error) {
	return rdb.Do(c, "geodist", key, source, destination, unit).Result()
}

// unit:默认为m,可选km:千米,mi:英里,ft：英尺)WithDist作为默认选项已经内置在函数中了，还可以添加其他选项(WithCOORD:返回经纬度信息,WithHash:返回哈希值)
func LocRadiusWithDist(c context.Context, rdb *redis.ClusterClient, key, lgt, lat, radius string, unit string) (interface{}, error) {
	return rdb.Do(c, "georadius", key, lgt, lat, radius, unit, "WITHDIST").Result()
}

// unit:默认为m,可选km:千米,mi:英里,ft：英尺)WithDist作为默认选项已经内置在函数中了，还可以添加其他选项(WithCOORD:返回经纬度信息)
func LocRadiusWithCoord(c context.Context, rdb *redis.ClusterClient, key, lgt, lat, radius string, unit string) (interface{}, error) {
	return rdb.Do(c, "georadius", key, lgt, lat, radius, unit, "WITHCOORD", "ASC").Result()
}
