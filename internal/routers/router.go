package routers

import (
	"gin-admin/global"
	"gin-admin/internal/middleware"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(middleware.GinRecovery(true))
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		app.WriteResponse(c, nil, gin.H{"msg": "pong"})
	})

	r.StaticFS("/static", gin.Dir(global.AppSetting.UploadSavePath, true))

	registerAdminRouters(r)
	return r
}
