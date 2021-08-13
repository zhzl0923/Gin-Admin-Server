package admin_permission

import (
	"gin-admin/global"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (p *AdminPermission) Get(c *gin.Context) {
	id := c.Param("id")
	permission, err := p.serv.AdminPermissionService().GetAdminPermissionById(convert.StrTo(id).MustUint64())
	if err != nil {
		global.Logger.Errorf("get admin permission  err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}
	app.WriteResponse(c, nil, permission)
}
