package main

import (
	"buyfree/config"
	"buyfree/dal"
	"buyfree/mrpc"
	"buyfree/service"
	"buyfree/utils"
	"context"
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
	//cmd := exec.Command("swag", "init")
	//err := cmd.Run()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("swagger.json生成成功")
	path, _ := os.Getwd()
	os.Setenv("cfgPATH", path)
}
func SalesCounter() {
	var uname string = utils.PTNAME
	rdb := dal.Getrdb()
	ctx := context.TODO()
	utils.ModifySales(ctx, rdb, utils.Ranktype1, uname, "1")
	utils.ModifySales(ctx, rdb, utils.Ranktype2, uname, "1")
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "yith", 123)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "9e", 12)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "2a", 123)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "3s", 1234)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype1, uname, "4b", 12345)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "yith", 123)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "9e", 12)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "2a", 123)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "3s", 1234)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype2, uname, "4b", 12345)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "yith", 123)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "9e", 12)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "2a", 123)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "3s", 1234)
	utils.ModifyTypeRanks(ctx, rdb, utils.Ranktype3, uname, "4b", 12345)
}
func main() {
	if *config.D {
		//go func() {
		//	logger.Loger.Info(http.ListenAndServe(":6060", nil))
		//}()
	}
	mrpc.PlatFormService.Run()
	mrpc.DriverService.Run()
	go service.Passengerrouter()
	//go service.Factoryrouter()
	go service.Driverrouter()
	//SalesCounter()
	service.PlatFormrouter()

}
