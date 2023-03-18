package service

import (
	"buyfree/handler"
	"buyfree/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initrouter() {
	//r := gin.Default()
	r := gin.New()
	r.Static("/static", "./public")
	r.Use(middleware.Cors())
	srv := http.Server{
		Addr:    ":9999",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.GET("/", func(c *gin.Context) {
		w := c.Writer
		w.Write([]byte("welecome to buyfree.com"))
	})
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	dr := r.Group("/driver")
	{
		dr.POST("/reple", func(c *gin.Context) {

		})
	}
	pr := r.Group("/passenger")
	{
		pr.GET("/buy", func(c *gin.Context) {

		})
	}
	fr := r.Group("/factory")
	{
		fr.POST("/supply", func(c *gin.Context) {

		})
	}
	//plr := r.Group("/platform")
	//{
	//	pds := plr.Group("/pds")
	//	{
	//
	//	}
	//	gct := plr.Group("/product")
	//	{
	//
	//	}
	//	dct := plr.Group("/device")
	//	adt := plr.Group("/ad")
	//	{
	//
	//	}
	//	ana := plr.Group("/ala")
	//	{
	//
	//	}
	//}

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
