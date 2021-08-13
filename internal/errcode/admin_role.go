package errcode

import "gin-admin/pkg/err"

var (
	ErrGetAdminRoleList = err.NewError(100401, 400, "获取角色列表失败")
	ErrCreateAdminRole  = err.NewError(100402, 400, "创建角色失败")
	ErrUpdateAdminRole  = err.NewError(100403, 400, "更新角色失败")
	ErrDuplicateRole    = err.NewError(100404, 400, "角色已存在")
	AdminRoleNotFound   = err.NewError(100405, 400, "角色不存在")
	ErrDeleteAdminRole  = err.NewError(100406, 400, "删除角色失败")
)
