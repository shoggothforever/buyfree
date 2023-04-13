package service

import (
	"buyfree/config"
	"buyfree/middleware"
	"buyfree/service/passenger"
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

func Passengerrouter() {
	var r *gin.Engine
	if *config.D == false {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		r = gin.Default()
	}
	//r.Static("/static", "./public")
	r.Use(middleware.Cors())
	srv := http.Server{
		Addr:    ":9005",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/", func(c *gin.Context) {
		w := c.Writer
		w.Write([]byte("welecome to passenger.buyfree.com"))
	})
	//与微信服务器交互
	var wat passenger.WeiXinAuthController
	{
		r.POST("/login", wat.Login)
	}
	//r.Use(middleware.AuthJwt())
	var ht passenger.HomePageController
	home := r.Group("home")
	{
		home.GET("/:id", ht.GetStatic)
		home.POST("/pay", ht.Pay)

	}
	mir := r.Group("infos")
	{
		mir.GET("/orders", ht.GetOrders)
		mir.GET("/:id/orders", ht.GetOrders)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Passenger Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Passenger Server Shutdown:", err)
	}
	log.Println("Passenger Server exiting")
}
