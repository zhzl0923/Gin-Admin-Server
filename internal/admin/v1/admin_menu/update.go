package admin_menu

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (m *AdminMenu) Update(c *gin.Context) {
	var param service.AdminMenuParam
	id := convert.StrTo(c.Param("id")).MustUint64()
	var err error
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("create admin menu BindAndValid err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}
	err = m.serv.AdminMenuService().UpdateAdminMenu(id, &param)
	if err != nil {
		global.Logger.Errorf("create admin menu err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}

	app.WriteResponse(c, nil, gin.H{})
}
