package admin_menu

import (
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (m *AdminMenu) List(c *gin.Context) {
	app.WriteResponse(c, nil, m.serv.AdminMenuService().GetAdminMenuList())
}

func (m *AdminMenu) UserMenu(c *gin.Context) {
	userId := c.GetUint64("admin_user_id")
	list, _ := m.serv.AdminMenuService().GetAdminMenuByUserId(userId)
	app.WriteResponse(c, nil, list)
}
