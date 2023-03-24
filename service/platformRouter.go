package service

import (
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

func PlatFormrouter() {
	r := gin.Default()
	//r.Static("/static", "./public")
	r.Use(middleware.Cors())
	srv := http.Server{
		Addr:    ":9003",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, response.PingResponse{"welecome to platform.buy-free.com"})
	})
	//注册与登录
	pt := r.Group("/pt")
	{
		pt.POST("/register", auth.Register)
		pt.POST("/login", auth.Login)

	}
	//鉴权
	pt.Use(middleware.AuthJwt())

	//数据大屏
	var gdc platform.GoodsController
	var salect platform.SalesController
	psc := pt.Group("/screen")
	{
		//TODO:
		psc.GET("", salect.GetScreenData)
	}
	//设备管理
	var devct platform.DevadminController
	rdv := pt.Group("/dev-admin")
	{
		rdv.GET("/list/:mode", devct.GetdevBystate)
		rdv.POST("/devs", devct.AddDev)
		var devinfoct platform.DevinfoController
		//设备详情
		rdv.GET("/infos/:id", devinfoct.LsInfo)

		////TODO:下架功能取消
		//rdv.PUT("/shelf", devinfoct.TakeDown)
	}

	//商品管理
	//var orderct platform.OrderController
	ord := pt.Group("/products")
	{

		//默认展示全部
		ord.GET("/factory/:mode/:factory_name/", gdc.GetAllProducts)
		ord.GET("/infos/:sku", gdc.GetGoodsInfo)

		//TODO:上下架操作整合
		ord.PUT("/on/:id", gdc.OnShelfGoods)
		ord.PUT("/down/:id", gdc.DownShelfGoods)
	}
	//销售统计
	//TODO:默认显示
	pt.GET("/static/:mode", salect.GetSales)
	//广告管理
	var adct platform.ADController
	ads := pt.Group("/ads")
	{
		ads.POST("", adct.AddAD)
		ads.GET("/list/:page", adct.GetADList)
		ads.GET("/infos/:id", adct.GetADContent)
		ads.GET("/efficient/:id", adct.GetADEfficient)

	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
