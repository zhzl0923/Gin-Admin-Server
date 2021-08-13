package docs

import (
	"gin-admin/internal/dao"
	"gin-admin/internal/model"
	"gin-admin/internal/service"
)

// swagger:route GET /role 角色管理 AdminRoleList
// 角色列表
//
// Security:
//   Bearer:
//
// Responses:
//   200: roleListResponse
//   default: errResponse

// swagger:route POST /role 角色管理 CreateAdminRole
// 添加角色
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route PUT /role/{id} 角色管理 UpdateAdminRole
// 修改角色
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route DELETE /role/{id} 角色管理 DeleteAdminRole
// 删除角色
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route GET /role/{id} 角色管理 AdminRole
// 角色详情
//
// Security:
//   Bearer:
//
// Responses:
//   200: adminRoleResponse
//   default: errResponse

// swagger:response roleListResponse
type GetRoleListResponseWrapper struct {
	//in:body
	Body []*model.AdminRole
}

// swagger:parameters CreateAdminRole UpdateAdminRole
type CreateRoleReuestWrapper struct {
	//in:body
	service.AdminRoleParam
}

// swagger:response adminRoleResponse
type AdminRoleResponse struct {
	//in:body
	dao.AdminRole
}
