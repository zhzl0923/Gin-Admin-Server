package admin_role

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (r *AdminRole) Update(c *gin.Context) {
	var param service.AdminRoleParam
	id := convert.StrTo(c.Param("id")).MustUint64()
	var err error
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("create admin role BindAndValid err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}
	err = r.serv.AdminRoleService().UpdateAdminRole(id, &param)
	if err != nil {
		global.Logger.Errorf("create admin role err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, gin.H{})
}
