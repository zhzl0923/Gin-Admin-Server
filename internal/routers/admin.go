package routers

import (
	"gin-admin/global"
	"gin-admin/internal/admin/v1/admin_menu"
	"gin-admin/internal/admin/v1/admin_permission"
	"gin-admin/internal/admin/v1/admin_role"
	"gin-admin/internal/admin/v1/admin_user"
	"gin-admin/internal/admin/v1/auth"
	"gin-admin/internal/admin/v1/base"
	"gin-admin/internal/dao"
	"gin-admin/internal/middleware"
	"gin-admin/internal/service"

	"github.com/gin-gonic/gin"
)

func registerAdminRouters(r *gin.Engine) {
	serv := service.NewService(dao.NewDao(global.DB))
	adminBase := base.NewBase(serv)
	r.GET("admin/v1/captcha", adminBase.Captcha)
	adminAuth := auth.NewAuth(serv)
	r.POST("admin/v1/login", adminAuth.Login)
	v1 := r.Group("admin/v1")
	v1.Use(middleware.JWTAuth())
	v1.Use(middleware.Auth())
	v1.POST("/upload/file", adminBase.UploadFile)

	adminUserV1 := v1.Group("/administor")
	{
		adminUser := admin_user.NewAdminUser(serv)
		adminUserV1.POST("", adminUser.Create)
		adminUserV1.DELETE("/:id", adminUser.Delete)
		adminUserV1.PUT("/:id", adminUser.Update)
		adminUserV1.GET("", adminUser.List)
		adminUserV1.GET("/info", adminUser.Info)
		adminUserV1.GET("/:id", adminUser.Get)
	}

	adminRoleV1 := v1.Group("/role")
	{
		adminRole := admin_role.NewAdminRole(serv)
		adminRoleV1.POST("", adminRole.Create)
		adminRoleV1.DELETE("/:id", adminRole.Delete)
		adminRoleV1.PUT("/:id", adminRole.Update)
		adminRoleV1.GET("", adminRole.List)
		adminRoleV1.GET("/:id", adminRole.Get)
	}

	adminMenuV1 := v1.Group("/menu")
	{
		adminMenu := admin_menu.NewAdminMenu(serv)
		adminMenuV1.POST("", adminMenu.Create)
		adminMenuV1.DELETE("/:id", adminMenu.Delete)
		adminMenuV1.PUT("/:id", adminMenu.Update)
		adminMenuV1.GET("", adminMenu.List)
		adminMenuV1.GET("/:id", adminMenu.Get)
		adminMenuV1.GET("user_menus", adminMenu.UserMenu)
	}

	adminPermissionV1 := v1.Group("/permission")
	{
		adminPermission := admin_permission.NewAdminPermission(serv)
		adminPermissionV1.POST("", adminPermission.Create)
		adminPermissionV1.DELETE("/:id", adminPermission.Delete)
		adminPermissionV1.PUT("/:id", adminPermission.Update)
		adminPermissionV1.GET("", adminPermission.List)
		adminPermissionV1.GET("/:id", adminPermission.Get)
		adminPermissionV1.GET("/all", adminPermission.All)
	}

}
