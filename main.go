package main

import (
	"buyfree/config"
	"buyfree/dal"
	"buyfree/repo/gen"
	"buyfree/service"
	"buyfree/utils"
	"sync"
)

var once sync.Once

//func Exit() {
//	//意外关闭的时候要注意将管道中的数据做持久化处理
//	close(utils.Refundchannel)
//	close(utils.Orderchannel)
//}
func main() {
	config.Init()
	once.Do(dal.InitPostgresSQL)
	gen.SetDefault(dal.DB)
	once.Do(dal.InitRedis)
	once.Do(utils.InitIDWorker)
	//defer Exit()
	service.PlatFormrouter()

}
