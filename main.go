package main

import (
	"Backend/config"
	"Backend/handler"
	"Backend/middleware"
	"Backend/model"
	
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	load()

	gin.SetMode(config.Val.Mode)
	r := gin.Default()
	r.Use(middleware.CROSS())

	r.GET("/ping", handler.Ping)

	api := r.Group("/api")
	{
		api.POST("/task", handler.CreateTask)
		api.GET("return", handler.GetReturn)
	}
	
	r.Run(":" + config.Val.Port)

	log.Infof("serve port: %v \n", config.Val.Port)
}

func load() {
	config.Init()
	model.Init()
}
