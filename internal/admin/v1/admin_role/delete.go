package admin_role

import (
	"gin-admin/internal/errcode"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (r *AdminRole) Delete(c *gin.Context) {
	id := convert.StrTo(c.Param("id")).MustUint64()
	err := r.serv.AdminRoleService().DeleteAdminRoleById(id)
	if err != nil {
		app.WriteResponse(c, errcode.ErrDeleteAdminUser.WithDetails(err.Error()), gin.H{})
		return
	}
	app.WriteResponse(c, nil, gin.H{"msg": "删除成功"})
}
