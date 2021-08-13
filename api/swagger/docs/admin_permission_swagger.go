package docs

import (
	"gin-admin/internal/model"
	"gin-admin/internal/service"
)

// swagger:route GET /permission 权限管理 AdminPermissionList
// 权限列表
//
// Security:
//   Bearer:
//
// Responses:
//   200: roleListResponse
//   default: errResponse

// swagger:route POST /permission 权限管理 CreateAdminPermission
// 添加权限
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route PUT /permission/{id} 权限管理 UpdateAdminPermission
// 修改权限
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route DELETE /permission/{id} 权限管理 DeleteAdminPermission
// 删除权限
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route GET /permission/{id} 权限管理 AdminPermission
// 权限详情
//
// Security:
//   Bearer:
//
// Responses:
//   200: adminPermissionResponse
//   default: errResponse

// swagger:route GET /permission/all 权限管理 AllPermission
// 所有权限列表
//
// Security:
//   Bearer:
//
// Responses:
//   200: allPermissionResponse
//   default: errResponse

// swagger:parameters AdminPermissionList
type GetPermissionListRequestWrapper struct {
	service.AdminPermissionListParam
}

// swagger:response roleListResponse
type GetPermissionListResponseWrapper struct {
	//in:body
	Body model.AdminPermissionList
}

// swagger:parameters CreateAdminPermission UpdateAdminPermission
type PermissionReuestWrapper struct {
	//in:body
	service.AdminPermissionParam
}

// swagger:response adminPermissionResponse
type AdminPermissionResponse struct {
	//in:body
	model.AdminPermission
}

// swagger:response allPermissionResponse
type AllAdminPermissionResponse struct {
	Body []*model.AdminPermission
}
