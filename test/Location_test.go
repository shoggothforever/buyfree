package test

import (
	"buyfree/dal"
	"buyfree/utils"
	"context"
	"testing"
)

func TestLocation(t *testing.T) {
	rdb := dal.Getrdb()
	ctx := context.Background()
	//res, err := rdb.Do(ctx, "geoadd", utils.LOCATION, 122.2222, 30.123, "hdu").Result()
	//t.Log(res, err)
	//t.Log(utils.LocAdd(ctx, rdb, utils.LOCATION, "122.11111", "30.111111", "silence1"))
	//t.Log(utils.LocAdd(ctx, rdb, utils.LOCATION, "122.12", "30.12", "cat"))
	//t.Log(utils.LocAdd(ctx, rdb, utils.LOCATION, "122.2", "30.2", "loc2"))
	//t.Log(utils.LocDist(ctx, rdb, utils.LOCATION, "hdu", "silence", "km"))
	//t.Log(utils.LocRadiusWithDist(ctx, rdb, utils.LOCATION, "122.222", "30.123", "10000", "m"))
	//res, _ := utils.LocRadiusWithDist(ctx, rdb, utils.LOCATION, "122.222", "30.123", "50000", "m")
	//t.Log(res, len(res.([]interface{})))
	//ires := res.([]interface{})
	//for _, v := range ires {
	//	t.Log(v.([]interface{})[0].(string))
	//	t.Log(v.([]interface{})[1].(string))
	//}
	res, _ := utils.LocRadiusWithCoord(ctx, rdb, utils.LOCATION, "122.222", "30.123", "50000", "m")
	t.Log(res.([]interface{}), len(res.([]interface{})))
	for _, iiv := range res.([]interface{}) {
		iv := iiv.([]interface{})
		t.Log(iv[0])
		v := iv[1].([]interface{})
		t.Log(v[0])
		t.Log(v[1])
	}
}
