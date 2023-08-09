package main

import (
	"context"
	"fmt"
	"mentor/classroom/config"
	"mentor/classroom/mentor_redis"
	"mentor/classroom/middleware/log"
	"mentor/classroom/route"
	"mentor/classroom/service/class_channel_service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// bindAddress := "0.0.0.0:2304"
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(log.LoggerMiddleware())
	r.Use(gin.RecoveryWithWriter(log.Logger.Out))
	r.Use(cors.New(config.CorsConfig()))
	
	route.Route(r)

	srv := &http.Server{
		Addr:    ":2304",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("here")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutdown Server ...")

	cleanClassroomList()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Info("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Logger.Info("timeout of 5 seconds.")
	}
	log.Logger.Info("Server exiting")
}

func cleanClassroomList() {
	for _, classroom := range class_channel_service.GetClassroomList() {
		mentor_redis.Client.Del(classroom)
	}
}
