package admin_menu

import (
	"gin-admin/global"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (m *AdminMenu) Get(c *gin.Context) {
	adminMenuId := c.Param("id")
	adminMenu, err := m.serv.AdminMenuService().GetAdminMenuById(convert.StrTo(adminMenuId).MustUint64())
	if err != nil {
		global.Logger.Errorf("get admin user  err: %v", err)
		app.WriteResponse(c, err, nil)
		return
	}
	app.WriteResponse(c, nil, adminMenu)
}
