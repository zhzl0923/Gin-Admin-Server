package routers

import (
	"gin-admin/global"
	"gin-admin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(middleware.GinRecovery(true))
	r.Use(middleware.Cors())
	r.Use(middleware.Tracing())
	r.StaticFS("/static", gin.Dir(global.AppSetting.UploadSavePath, true))
	registerAdminRouters(r)
	return r
}
