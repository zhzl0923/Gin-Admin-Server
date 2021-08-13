package admin_user

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (u *AdminUser) Update(c *gin.Context) {
	var param service.UpdateAdminUserParam
	var err error
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("update admin user BindAndValid err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}
	id := convert.StrTo(c.Param("id")).MustUint64()
	if param.Password != "" && len(param.Password) < 6 {
		app.WriteResponse(c, errcode.InvalidParams.WithDetails("Password长度必须至少为6个字符"), nil)
		return
	}
	err = u.serv.AdminUserService().UpdateAdminUser(id, &param)
	if err != nil {
		global.Logger.Errorf("update admin user err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}
	app.WriteResponse(c, nil, gin.H{})
}
