package admin_user

import (
	"gin-admin/global"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (u *AdminUser) Info(c *gin.Context) {
	adminUserId := c.GetUint64("admin_user_id")

	adminUser, err := u.serv.AdminUserService().GetAdminUserById(adminUserId)
	if err != nil {
		global.Logger.Errorf("get admin user  err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}
	app.WriteResponse(c, nil, adminUser)
}
