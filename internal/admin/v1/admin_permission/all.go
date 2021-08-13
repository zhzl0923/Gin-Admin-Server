package admin_permission

import (
	"gin-admin/global"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (p *AdminPermission) All(c *gin.Context) {

	list, err := p.serv.AdminPermissionService().GetAllAdminPermission()

	if err != nil {
		global.Logger.Errorf("get admin permission list err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, list)
}
