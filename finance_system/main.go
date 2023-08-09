package main

import (
	"mentor_app/finance_system/controller"
	"mentor_app/finance_system/middleware/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title 登入系統
func main() {
	bindAddress := "0.0.0.0:2307"

	r := gin.New()

	r.Use(log.LoggerMiddleware())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	controller.Route(r)

	r.Run(bindAddress)
}
