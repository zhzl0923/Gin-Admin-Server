package dao

import (
	"gin-admin/internal/model"

	"gorm.io/gorm"
)

type AdminRoleMenuDao interface {
	GetAdminMenuIds(query ...interface{}) []uint64
	CreateAdminRoleMenu(roleID uint64, menus []uint64)
	DeleteAdminRoleMenu(query ...interface{})
}

type adminRoleMenuDao struct {
	db *gorm.DB
}

func NewAdminRoleMenuDao(db *gorm.DB) AdminRoleMenuDao {
	return &adminRoleMenuDao{db: db}
}

func (d *adminRoleMenuDao) GetAdminMenuIds(query ...interface{}) []uint64 {
	var adminRoleMenus []model.AdminRoleMenu
	GetWhereQuery(d.db, query...).Select("menu_id").Find(&adminRoleMenus)
	menuIds := make([]uint64, len(adminRoleMenus))
	for k, rm := range adminRoleMenus {
		menuIds[k] = rm.MenuID
	}
	return menuIds
}

func (d *adminRoleMenuDao) CreateAdminRoleMenu(roleID uint64, menus []uint64) {
	if len(menus) == 0 {
		return
	}
	roleMenu := make([]model.AdminRoleMenu, len(menus))
	for key, menuID := range menus {
		roleMenu[key] = model.AdminRoleMenu{RoleID: roleID, MenuID: menuID}
	}
	if len(roleMenu) > 0 {
		d.db.Create(&roleMenu)
	}
}

func (d *adminRoleMenuDao) DeleteAdminRoleMenu(query ...interface{}) {
	GetWhereQuery(d.db, query...).Delete(&model.AdminRoleMenu{})
}
