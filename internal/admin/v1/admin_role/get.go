package admin_role

import (
	"gin-admin/global"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (r *AdminRole) Get(c *gin.Context) {
	adminRoleId := c.Param("id")
	adminRole, err := r.serv.AdminRoleService().GetAdminRoleById(convert.StrTo(adminRoleId).MustUint64())
	if err != nil {
		global.Logger.Errorf("get admin role  err: %v", err)
		app.WriteResponse(c, err, gin.H{})
		return
	}
	app.WriteResponse(c, nil, adminRole)
}
