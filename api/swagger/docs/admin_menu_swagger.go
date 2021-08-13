package docs

import (
	"gin-admin/internal/dao"
	"gin-admin/internal/model"
	"gin-admin/internal/service"
)

// swagger:route GET /menu 菜单管理 AdminMenuList
// 菜单列表
//
// Security:
//   Bearer:
//
// Responses:
//   200: menuListResponse
//   default: errResponse

// swagger:route POST /menu 菜单管理 CreateAdminMenu
// 添加菜单
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route PUT /menu/{id} 菜单管理 UpdateAdminMenu
// 修改菜单
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route DELETE /menu/{id} 菜单管理 DeleteAdminMenu
// 删除菜单
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route GET /menu/{id} 菜单管理 AdminMenu
// 菜单详情
//
// Security:
//   Bearer:
//
// Responses:
//   200: adminMenuResponse
//   default: errResponse

// swagger:route GET /menu/user_menus 菜单管理 AdminUserMenu
// 获取用户菜单
//
// Security:
//   Bearer:
//
// Responses:
//   200: menuListResponse
//   default: errResponse

// swagger:response menuListResponse
type GetMenuListResponseWrapper struct {
	//in:body
	Body model.AdminMenuList
}

// swagger:parameters CreateAdminMenu UpdateAdminMenu
type CreateMenuReuestWrapper struct {
	//in:body
	service.AdminMenuParam
}

// swagger:response adminMenuResponse
type AdminMenuResponse struct {
	//in:body
	dao.AdminMenu
}
