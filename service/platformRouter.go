package service

import (
	"buyfree/middleware"
	"buyfree/service/auth"
	"buyfree/service/platform"
	"buyfree/service/response"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func PlatFormrouter() {
	r := gin.Default()
	r.Static("/static", "./public")
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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, response.PingResponse{"welecome to platform.buyfree.com"})
	})

	pt := r.Group("/pt")
	{
		pt.POST("/register", auth.Register)
		pt.POST("/login", auth.Login)

	}
	//鉴权
	pt.Use(middleware.AuthJwt())
	//数据大屏
	psc := pt.Group("/screen")
	{
		//可能不需要路由啦，直接加载
		//psc.GET("/sales",platform.GetSales)
		//psc.GET("/curve",platform.GetCurve)
		//psc.GET("/cntdev",platform.GetDevCnt)
		//psc.GET("/adstatic",platform.AnalyzeAD)
		//psc.GET("/hotdata",platform.GetSaleRank)
		//psc.GET("/locate",platform.GetLocation)
		psc.GET("/",
			platform.GetSales, platform.GetCurve,
			platform.GetDevCnt, platform.AnalyzeAD,
			platform.GetSaleRank, platform.GetLocation,
		)

	}
	//设备管理
	pdv := pt.Group("/dev")
	{
		//默认显示
		//pdv.GET("/alldev", platform.GetAlldev)
		pdv.GET("/", platform.GetAlldev)

		pdv.GET("/activated", platform.GetActivated)
		pdv.GET("/nonactivated", platform.GetNotActivated)
		pdv.GET("/online", platform.GetOndev)
		pdv.GET("/offline", platform.GetOffdev)
		pdv.POST("/adddev", platform.AddDev)
		//设备详情
		pdc := pdv.Group("/info")
		{
			pdc.PUT("/down", platform.TakeDown)
			pdc.GET("/salesinfo", platform.AnaSales)
			pdc.GET("/devinfo", platform.LsDev)
			pdc.GET("/driverinfo", platform.LsDriver)
			pdc.GET("/productinfo", platform.LsDevProduct)
		}
	}

	//商品管理
	ssc := pt.Group("/goods")
	{
		ssc.GET("/allorder", platform.GetOrders)
		ssc.GET("/onshelf", platform.GetOnShelf)
		ssc.GET("/soldout", platform.Getsoldout)
		ssc.GET("/downshelf", platform.Getdownshelf)
		ssc.GET("/info", platform.GetGoodinfo)
		ssc.PUT("/on", platform.TakeOn)

	}
	//销售统计
	sst := pt.Group("/anasales")
	{
		//默认显示
		sst.GET("/static", platform.GetSales)
		sst.GET("/daily", platform.GetDailyRank)

		sst.GET("/monthly", platform.GetMonthlyRank)
		sst.GET("/annually", platform.GetAnnuallyRank)
	}
	//广告管理
	ana := pt.Group("/ads")
	{
		ana.GET("/list", platform.GetADList)
		ana.GET("/add", platform.GetAddAD)
		adinfo := ana.Group("/info")
		{
			adinfo.GET("/content", platform.GetADContent)
			adinfo.GET("/efficent", platform.GetADEfficient)
		}
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
