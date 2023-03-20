package main

import (
	"buyfree/config"
	"buyfree/dal"
	"buyfree/repo/gen"
	"buyfree/service"
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
	//id, _ := uuid.Parse("a870e804-cf1e-3dc3-1190-5726a7d46039")
	//u, _ := gen.Passenger.GetByUUID(id)
	//fmt.Println(u.ID)
	dal.InitRedis()
	//defer Exit()
	service.PlatFormrouter()

}
