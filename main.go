package main

import (
	"buyfree/logger"
	"buyfree/mrpc"
	"buyfree/service"
	"net/http"
	_ "net/http/pprof"
	"os"
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

func init() {
	path, _ := os.Getwd()
	os.Setenv("cfgPATH", path)
}
func main() {
	//defer Exit()
	//go service.Factoryrouter()
	go func() {
		logger.Loger.Info(http.ListenAndServe(":6060", nil))
	}()
	mrpc.PlatFormService.Run()
	mrpc.DriverService.Run()
	go service.Passengerrouter()
	go service.Driverrouter()
	service.PlatFormrouter()

}
