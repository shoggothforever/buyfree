package service

import (
	"buyfree/middleware"
	"buyfree/service/auth"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var d = flag.Bool("d", false, "默认为release，true为debug")
var QuitDriverChan chan os.Signal
var DriverSrv http.Server

func Driverrouter() {
	flag.Parse()
	//if *d == false {
	//	gin.SetMode(gin.ReleaseMode)
	//} else {
	//	gin.SetMode(gin.DebugMode)
	//}
	//r := gin.New()
	r := gin.Default()
	//r.Static("/static", "./public")
	r.Use(middleware.Cors())
	DriverSrv = http.Server{
		Addr:    ":9001",
		Handler: r,
	}
	go func() {
		if err := DriverSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.GET("/", func(c *gin.Context) {
		w := c.Writer
		w.Write([]byte("welecome to driver.buyfree.com"))
	})
	dr := r.Group("/dr")
	{
		dr.POST("/register", auth.DriverRegister)
		dr.POST("/login", auth.DriverLogin)
	}
	{
		dr.POST("/reple", func(c *gin.Context) {

		})
	}
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
