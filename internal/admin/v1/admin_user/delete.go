package admin_user

import (
	"gin-admin/internal/errcode"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (u *AdminUser) Delete(c *gin.Context) {
	id := convert.StrTo(c.Param("id")).MustUint64()
	if id == c.GetUint64("admin_user_id") {
		app.WriteResponse(c, errcode.ErrDeleteAdminUser, gin.H{})
		return
	}
	err := u.serv.AdminUserService().DeleteAdminUserById(id)
	if err != nil {
		app.WriteResponse(c, errcode.ErrDeleteAdminUser.WithDetails(err.Error()), gin.H{})
		return
	}
	app.WriteResponse(c, nil, gin.H{"msg": "删除成功"})
}
