package service

import (
	"buyfree/config"
	_ "buyfree/docs"
	"buyfree/middleware"
	"buyfree/service/auth"
	"buyfree/service/platform"
	"buyfree/service/response"
	"context"
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

// var b = flag.Bool("b", false, "默认为release，true为debug")
var QuitPlatformChan chan os.Signal
var PlatFormSrv http.Server

func PlatFormrouter() {
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
	PlatFormSrv = http.Server{
		Addr:    ":9003",
		Handler: r,
	}
	go func() {
		if err := PlatFormSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/shorten/*any", middleware.RedirectShort())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, response.PingResponse{"welecome to platform.buy-free.com"})
	})
	//注册与登录
	pt := r.Group("/pt")
	{
		pt.POST("/register", auth.PlatformRegister)
		pt.POST("/login", auth.PlatformLogin)
		pt.POST("/userinfo", auth.PlatformUserInfo)

	}
	var fat platform.FactoryadminController
	fa := r.Group("/fa")
	{
		fa.POST("/register", auth.FactoryRegister)
		fa.POST("/login", auth.FactoryLogin)
		fa.POST("/userinfo", auth.FactoryUserInfo)
		fa.POST("/inventory", middleware.AuthJwt(), fat.Add)
		fa.PATCH("/inventory/add/:product_name/:inv", middleware.AuthJwt(), fat.AddInv)
		fa.GET("/infos/all/:mode", middleware.AuthJwt(), fat.GetAllProducts)
		fa.GET("/infos/detail/:product_name", middleware.AuthJwt(), fat.GetGoodsInfo)
		fa.GET("/infos/orders/:mode/:page", middleware.AuthJwt(), fat.GetDriverOrders)
		//fa.GET("/infos/orders/detail/", middleware.AuthJwt(), fat.GetDriverOrders)
	}
	//鉴权
	pt.Use(middleware.AuthJwt())

	//数据大屏
	var gdc platform.GoodsController
	var salect platform.SalesController

	pt.GET("/static/:mode", salect.GetSales)
	psc := pt.Group("/screen")
	{
		psc.GET("", salect.GetScreenData)
		psc.GET("/location", salect.GetNearbyDriver)
	}

	fdr := pt.Group("/factory-admin")
	{
		fdr.POST("/register", fat.PRegister)
		fdr.POST("/:factory_name/products", fat.PAdd)
		fdr.GET("/:factory_name/products/:product_name/:inv", fat.PAddInv)
		//fdr.GET("/:factory_name/orders/:mode", fat.PGetDriverOrders)
	}

	//设备管理
	var devct platform.DevadminController
	rdv := pt.Group("/dev-admin")
	{
		rdv.GET("/list/:mode/:page", devct.GetdevBystate)
		rdv.POST("/devs", devct.AddDev)
		//rdv.GET("/owner/:id")
		//设备详情
		rdv.GET("/infos/:id", devct.LsInfo)
		rdv.POST("/launch/:dev_id", devct.Launch)
	}

	ord := pt.Group("/products")
	{

		//默认展示全部
		ord.GET("/:mode/factory/:factory_name", gdc.PGetAllProducts)
		ord.GET("/infos/:factory_name/:product_name", gdc.PGetGoodsInfo)
		ord.PATCH("/turn", gdc.TurnOver)
	}
	//销售统计
	//TODO:默认显示

	//广告管理
	var adct platform.ADController
	ads := pt.Group("/ads")
	{
		ads.POST("", adct.AddAD)
		ads.GET("/list/:page", adct.GetADList)
		ads.GET("/infos/:id", adct.GetADContent)
		ads.GET("/efficient/:id", adct.GetADEfficient)

		ads.PATCH("/modify/:ad_id", adct.Shelf)

	}
	QuitPlatformChan = make(chan os.Signal)
	time.Sleep(5 * time.Second)
	signal.Notify(QuitPlatformChan, os.Interrupt)
	<-QuitPlatformChan
	log.Println("Shutdown  PlatForm Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := PlatFormSrv.Shutdown(ctx); err != nil {
		log.Fatal("PlatForm Server Shutdown:", err)
	}
	log.Println("Plat Form Server exiting")
}
