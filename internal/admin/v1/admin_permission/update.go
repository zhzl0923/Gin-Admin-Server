package admin_permission

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (p *AdminPermission) Update(c *gin.Context) {
	var param service.AdminPermissionParam
	var err error
	id := convert.StrTo(c.Param("id")).MustUint64()
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("update admin permission BindAndValid err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}
	err = p.serv.AdminPermissionService().UpdateAdminPermission(id, &param)
	if err != nil {
		global.Logger.Errorf("update admin permission err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}
	app.WriteResponse(c, nil, gin.H{})
}
