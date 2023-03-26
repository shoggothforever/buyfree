package main

import (
	"buyfree/service"
	"os"
	"sync"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      bf.shoggothy.xyz
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
var once sync.Once

//func Exit() {
//	//意外关闭的时候要注意将管道中的数据做持久化处理
//	close(utils.Refundchannel)
//	close(utils.Orderchannel)
//}
var PWD string

func init() {
	PWD, _ = os.Getwd()
}

func main() {
	//config.Init()
	//once.Do(dal.InitPostgresSQL)
	//gen.SetDefault(dal.DB)
	//once.Do(dal.InitRedis)
	//once.Do(utils.InitIDWorker)
	//defer Exit()

	service.PlatFormrouter()

}
