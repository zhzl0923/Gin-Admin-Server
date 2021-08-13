package errcode

import "gin-admin/pkg/err"

var (
	ErrGetAdminMenuList = err.NewError(100501, 400, "获取菜单列表失败")
	ErrCreateAdminMenu  = err.NewError(100502, 400, "创建菜单失败")
	ErrUpdateAdminMenu  = err.NewError(100503, 400, "更新菜单失败")
	ErrDuplicateMenu    = err.NewError(100504, 400, "菜单已存在")
	AdminMenuNotFound   = err.NewError(100505, 400, "菜单不存在")
	ErrDeleteAdminMenu  = err.NewError(100506, 400, "删除菜单失败")
	ErrMenuHasSubMenu   = err.NewError(100507, 400, "请先删除子菜单")
)
