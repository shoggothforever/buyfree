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
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	//pt.Use(middleware.AuthJwt())

	//数据大屏
	var salect platform.SalesController
	psc := pt.Group("/screen")
	{

		//可能不需要路由啦，直接加载
		//psc.GET("/sales",platform.GetSales)
		//psc.GET("/curve",platform.GetCurve)
		//psc.GET("/cntdev",platform.GetDevCnt)
		//psc.GET("/adstatic",platform.AnalyzeAD)
		//psc.GET("/hotdata",platform.GetSaleRank)
		//psc.GET("/locate",platform.GetLocation)
		//TODO:
		psc.GET("",
			salect.GetSales, salect.GetCurve,
			salect.GetDevCnt, salect.AnalyzeAD,
			salect.GetSaleRank, salect.GetLocation,
		)

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

		//TODO:下架
		rdv.PUT("/shelf", devinfoct.TakeDown)
		//TODO:当日营销额(还缺少需要的信息，暂时不合并）
		rdv.GET("/infos/sale", devinfoct.AnaSales)
	}

	//商品管理
	var orderct platform.OrderController
	ord := pt.Group("/orders")
	{

		//默认展示全部
		ord.GET("/factory/:mode", orderct.GetFactoryOrders)
		ord.GET("/infos/:sku", orderct.GetGoodinfo)
		//TODO:上下架操作整合
		ord.GET("/shelf/:mode", orderct.ModifyGoods)

	}
	//销售统计
	var goodsct platform.GoodsController
	sst := pt.Group("/ana-sales")
	{

		//TODO:默认显示
		sst.GET("/static", salect.GetSales)
		sst.GET("/daily", goodsct.GetDailyRank)
		sst.GET("/monthly", goodsct.GetMonthlyRank)
		sst.GET("/annually", goodsct.GetAnnuallyRank)
	}
	//广告管理
	var adct platform.ADController
	ads := pt.Group("/ads")
	{
		ads.POST("", adct.AddAD)
		ads.GET("/list", adct.GetADList)
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
