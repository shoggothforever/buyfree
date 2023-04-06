package main

import (
	"buyfree/mrpc"
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
func init() {
	//cmd := exec.Command("swag", "init")
	//err := cmd.Run()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("swagger.json生成成功")
	path, _ := os.Getwd()
	os.Setenv("cfgPATH", path)
}
func main() {
	//defer Exit()
	//go service.Factoryrouter()
	mrpc.PlatFormService.Run()
	go service.Driverrouter()
	service.PlatFormrouter()

}
