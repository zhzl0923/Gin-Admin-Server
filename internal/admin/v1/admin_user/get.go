package admin_user

import (
	"gin-admin/global"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (u *AdminUser) Get(c *gin.Context) {
	adminUserId := c.Param("id")
	adminUser, err := u.serv.AdminUserService().GetAdminUserById(convert.StrTo(adminUserId).MustUint64())
	if err != nil {
		global.Logger.Errorf("get admin user  err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}
	app.WriteResponse(c, nil, adminUser)
}
