package main

import (
	"buyfree/config"
	"buyfree/dal"
)

func main() {
	config.Init()
	dal.InitPostgresSQL()
	//dal.InitRedis()
	//initrouter()
}
