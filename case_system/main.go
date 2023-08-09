package main

import (
	"case_system/controller"
	"case_system/middleware/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	bindAddress := "0.0.0.0:2305"
	r := gin.New()

	r.Use(log.LoggerMiddleware())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	controller.Route(r)

	r.Run(bindAddress)
}
