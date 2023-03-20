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
		psc.GET("/",
			salect.GetSales, salect.GetCurve,
			salect.GetDevCnt, salect.AnalyzeAD,
			salect.GetSaleRank, salect.GetLocation,
		)

	}
	//设备管理
	var devct platform.DevadminController
	pdv := pt.Group("/dev")
	{

		//默认显示
		//pdv.GET("/alldev", platform.GetAlldev)
		pdv.GET("/", devct.GetAlldev)

		pdv.GET("/activated", devct.GetActivated)
		pdv.GET("/nonactivated", devct.GetNotActivated)
		pdv.GET("/online", devct.GetOndev)
		pdv.GET("/offline", devct.GetOffdev)
		pdv.POST("/adddev", devct.AddDev)
		//设备详情
		pdc := pdv.Group("/info") //传入设备ID
		{
			var devinfoct platform.DevinfoController
			//下架
			pdc.PUT("/down", devinfoct.TakeDown)
			//当日营销额(还缺少需要的信息，暂时不合并）
			pdc.GET("/salesinfo", devinfoct.AnaSales)
			//可以合并的操作
			//pdc.GET("/devinfo", devinfoct.LsDev)
			//pdc.GET("/driverinfo", devinfoct.LsDriver)
			//pdc.GET("/productinfo", devinfoct.LsDevProduct)
			pdc.GET("/", devinfoct.LsInfo)
		}
	}

	//商品管理
	var orderct platform.OrderController
	ssc := pt.Group("/order")
	{

		//默认展示全部
		ssc.GET("/allorder", orderct.GetOrders)
		ssc.GET("/onshelf", orderct.GetOnShelf)
		ssc.GET("/soldout", orderct.Getsoldout)
		ssc.GET("/downshelf", orderct.Getdownshelf)
		ssc.GET("/info", orderct.GetGoodinfo)
		//ssc.PUT("/on", orderct.TakeOn)

	}
	//销售统计
	var goodsct platform.GoodsController
	sst := pt.Group("/anasales")
	{

		//默认显示
		sst.GET("/static", salect.GetSales)
		sst.GET("/daily", goodsct.GetDailyRank)

		sst.GET("/monthly", goodsct.GetMonthlyRank)
		sst.GET("/annually", goodsct.GetAnnuallyRank)
	}
	//广告管理
	var adct platform.ADController
	ana := pt.Group("/ads")
	{
		ana.GET("/list", adct.GetADList)
		ana.GET("/add", adct.GetAddAD)
		adinfo := ana.Group("/info")
		{
			adinfo.GET("/content", adct.GetADContent)
			adinfo.GET("/efficent", adct.GetADEfficient)
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
