package admin_menu

import (
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"

	"github.com/gin-gonic/gin"
)

func (m *AdminMenu) Delete(c *gin.Context) {
	id := convert.StrTo(c.Param("id")).MustUint64()
	err := m.serv.AdminMenuService().DeleteAdminMenuById(id)
	if err != nil {
		app.WriteResponse(c, err, gin.H{})
		return
	}
	app.WriteResponse(c, nil, gin.H{"msg": "删除成功"})
}
