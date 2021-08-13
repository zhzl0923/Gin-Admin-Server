package admin_user

import (
	"gin-admin/global"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (u *AdminUser) List(c *gin.Context) {
	param := service.AdminUserListParam{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	list, err := u.serv.AdminUserService().GetAdminUserList(&param)

	if err != nil {
		global.Logger.Errorf("get admin user list err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, list)
}
