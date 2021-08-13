package main

import (
	_ "gin-admin/api/swagger/docs"
	"gin-admin/global"
	"gin-admin/internal"
	"gin-admin/internal/routers"
	"gin-admin/internal/setup"

	"github.com/gin-gonic/gin"
)

func init() {
	setup.Setup()
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	internal.Run(
		global.ServerSetting.HttpPort,
		router,
		global.ServerSetting.ReadTimeout,
		global.ServerSetting.WriteTimeout,
	)
}
