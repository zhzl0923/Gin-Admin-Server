package admin_permission

import (
	"gin-admin/global"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (p *AdminPermission) List(c *gin.Context) {
	param := service.AdminPermissionListParam{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	list, err := p.serv.AdminPermissionService().GetAdminPermissionList(&param)

	if err != nil {
		global.Logger.Errorf("get admin permission list err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, list)
}
