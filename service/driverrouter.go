package service

import (
	"buyfree/config"
	"buyfree/middleware"
	"buyfree/service/auth"
	"buyfree/service/driverapp"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var QuitDriverChan chan os.Signal
var DriverSrv http.Server

func Driverrouter() {
	flag.Parse()
	var r *gin.Engine
	if *config.D == false {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		r = gin.Default()
	}
	//r := gin.Default()
	//r.Static("/static", "../public")
	r.Use(middleware.Cors())
	DriverSrv = http.Server{
		Addr:    ":9004",
		Handler: r,
	}
	go func() {
		if err := DriverSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	var base driverapp.BaseDrController
	var ht driverapp.HomePageController
	var it driverapp.InventoryController
	var ft driverapp.FactoryController
	var dft driverapp.InfoController
	var cat driverapp.CartController
	var det driverapp.DeviceController
	r.GET("/", base.Ping)
	r.GET("/shorten/*any", middleware.RedirectShort())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	dr := r.Group("/dr")
	{
		dr.POST("/register", auth.DriverRegister)
		dr.POST("/login", auth.DriverLogin)
		dr.POST("/userinfo", auth.DriverUserInfo)
	}
	dr.GET("/inventory/:device_id", it.GetDeviceByScan)

	dr.Use(middleware.AuthJwt())
	dr.GET("/home", ht.GetStatic)
	dr.POST("/ping", ht.Ping)
	dr.GET("/inventory", it.GetInventory)
	fa := dr.Group("/factory", middleware.AuthJwt())
	{
		fa.POST("", ft.FactoryOverview)
		fa.GET("/cart", cat.OpenCart)
		fa.POST("/infos", ft.Detail)
	}
	od := dr.Group("/order")
	{
		od.POST("/replenish", ft.Modify)
		od.PATCH("/choose", ft.Choose)
		od.POST("/submit", ft.Submit)
		od.POST("/submit2", ft.SubmitMany)
		od.POST("/pay", ft.Pay)
		od.GET("/:id/load", ft.Load)

	}
	devr := dr.Group("/devices")
	{
		devr.GET("/QR", det.QR)
		devr.GET("/scan", det.Scan)
		devr.POST("/bind", det.BindDevice)
	}
	pr := dr.Group("/infos")
	{
		pr.GET("/devices", dft.Getdevice)
		pr.GET("/orderform/:mode", dft.GetOrders)
		pr.GET("/orderdetail/:id", dft.GetOrder)
		//pr.GET("/balance", dft.GetBalance)
		pr.GET("/withdraw", dft.Withdraw)

	}
	//for _, route := range r.Routes() {
	//	dal.GetCasbinModel().Enforcer.AddPolicy(strconv.Itoa(int(model.DRIVER)), route.Path, route.Method)
	//}
	QuitDriverChan = make(chan os.Signal)
	signal.Notify(QuitDriverChan, os.Interrupt)
	<-QuitDriverChan
	log.Println("Shutdown Driver Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := DriverSrv.Shutdown(ctx); err != nil {
		log.Fatal("Driver Server Shutdown:", err)
	}
	log.Println("Driver Server exiting")
}
