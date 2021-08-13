package admin_role

import (
	"gin-admin/global"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (r *AdminRole) List(c *gin.Context) {
	list, err := r.serv.AdminRoleService().GetAdminRoleList()

	if err != nil {
		global.Logger.Errorf("get admin user list err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, list)
}
