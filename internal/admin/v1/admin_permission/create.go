package admin_permission

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (p *AdminPermission) Create(c *gin.Context) {
	var param service.AdminPermissionParam
	var err error
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("create admin permission BindAndValid err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}
	err = p.serv.AdminPermissionService().CreateAdminPermission(&param)

	if err != nil {
		global.Logger.Errorf("create admin permission err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, gin.H{})
}
