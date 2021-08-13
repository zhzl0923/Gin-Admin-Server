package admin_menu

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (m *AdminMenu) Create(c *gin.Context) {
	var param service.AdminMenuParam
	var err error
	valid, err := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("create admin menu BindAndValid err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}

	err = m.serv.AdminMenuService().CreateAdminMenu(&param)

	if err != nil {
		global.Logger.Errorf("create admin menu err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, gin.H{})
}
