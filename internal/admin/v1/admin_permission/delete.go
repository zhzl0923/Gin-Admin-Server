package admin_permission

import (
	"gin-admin/internal/errcode"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (p *AdminPermission) Delete(c *gin.Context) {
	id := convert.StrTo(c.Param("id")).MustUint64()
	err := p.serv.AdminPermissionService().DeleteAdminPermissionById(id)
	if err != nil {
		app.WriteResponse(c, errcode.ErrDeleteAdminPermission.WithDetails(err.Error()), gin.H{})
		return
	}
	app.WriteResponse(c, nil, gin.H{"msg": "删除成功"})
}
