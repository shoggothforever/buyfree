package service

import (
	"buyfree/middleware"
	"buyfree/service/auth"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Factoryrouter() {
	//r := gin.Default()
	r := gin.New()
	//r.Static("/static", "./public")
	r.Use(middleware.Cors())
	srv := http.Server{
		Addr:    ":9002",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.GET("/", func(c *gin.Context) {
		w := c.Writer
		w.Write([]byte("welecome to factory.buyfree.com"))
	})
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	fr := r.Group("/factory")
	{
		fr.POST("/supply", func(c *gin.Context) {

		})
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
