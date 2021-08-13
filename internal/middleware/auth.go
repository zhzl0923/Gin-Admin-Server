package middleware

import (
	"gin-admin/global"
	"gin-admin/internal/dao"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminUserID := c.GetUint64("admin_user_id")
		httpPath := c.FullPath()
		service := service.NewService(dao.NewDao(global.DB))
		if !service.AdminPermissionService().UserHasPermission(adminUserID, httpPath, c.Request.Method) {
			app.WriteResponse(c, errcode.Forbidden, gin.H{})
			c.Abort()
			return
		}
		c.Next()
	}

}
