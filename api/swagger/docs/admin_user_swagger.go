package docs

import (
	"gin-admin/internal/model"
	"gin-admin/internal/service"
)

// swagger:route GET /administor 管理员账号管理 AdminUserList
// 账号列表
//
// Security:
//   Bearer:
//
// Responses:
//   200: adminListResponse
//   default: errResponse

// swagger:route POST /administor 管理员账号管理 CreateAdminUser
// 添加账号
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route PUT /administor/{id} 管理员账号管理 UpdateAdminUser
// 修改账号
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route DELETE /administor/{id} 管理员账号管理 DeleteAdminUser
// 删除账号
//
// Security:
//   Bearer:
//
// Responses:
//   200: okResponse
//   default: errResponse

// swagger:route GET /administor/{id} 管理员账号管理 AdminUser
// 账号详情
//
// Security:
//   Bearer:
//
// Responses:
//   200: adminUserResponse
//   default: errResponse

// swagger:route GET /administor/info 管理员账号管理 AdminUser
// 登录用户详情
// Security:
//   Bearer:
//
// Responses:
//   200: adminUserResponse
//   default: errResponse

// swagger:parameters AdminUserList
type GetAdminListRequestWrapper struct {
	//in:body
	service.AdminUserListParam
}

// swagger:response adminListResponse
type GetAdminListResponseWrapper struct {
	//in:body
	Body model.AdminUserList
}

// swagger:parameters CreateAdminUser
type CreateAdminUserReuestWrapper struct {
	//in:body
	service.CreateAdminUserParam
}

// swagger:parameters UpdateAdminUser
type UpdateAdminUserRequestWrapper struct {
	//in:body
	service.UpdateAdminUserParam
}

// swagger:response adminUserResponse
type AdminUserResponse struct {
	//in:body
	model.AdminUser
}
