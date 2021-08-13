package errcode

import "gin-admin/pkg/err"

var (
	ErrGetAdminPermissionList = err.NewError(100601, 400, "获取权限列表失败")
	ErrCreateAdminPermission  = err.NewError(100602, 400, "创建权限失败")
	ErrUpdateAdminPermission  = err.NewError(100603, 400, "更新权限失败")
	ErrDuplicatePermission    = err.NewError(100604, 400, "权限已存在")
	AdminPermissionNotFound   = err.NewError(100605, 400, "权限不存在")
	ErrDeleteAdminPermission  = err.NewError(100606, 400, "删除权限失败")
)
