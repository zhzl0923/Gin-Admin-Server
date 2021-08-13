package errcode

import "gin-admin/pkg/err"

var (
	ErrGetAdminUserList  = err.NewError(100201, 400, "获取管理员列表失败")
	ErrCreateAdminUser   = err.NewError(100202, 400, "创建管理员失败")
	ErrUpdateAdminUser   = err.NewError(100203, 400, "更新管理员失败")
	ErrDuplicateUsername = err.NewError(100204, 400, "用户名重复")
	AdminUserNotFound    = err.NewError(100205, 400, "管理员不存在")
	ErrDeleteAdminUser   = err.NewError(100206, 400, "删除管理员失败")
)
