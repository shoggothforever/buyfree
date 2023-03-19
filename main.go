package main

import (
	"buyfree/config"
	"buyfree/dal"
	"buyfree/service"
	"sync"
)

var once sync.Once

func main() {
	config.Init()
	once.Do(dal.InitPostgresSQL)

	//gen.SetDefault(dal.DB)
	//id, _ := uuid.Parse("a870e804-cf1e-3dc3-1190-5726a7d46039")
	//u, _ := gen.Passenger.GetByUUID(id)
	//fmt.Println(u.ID)
	//dal.InitRedis()
	service.PlatFormrouter()
}
