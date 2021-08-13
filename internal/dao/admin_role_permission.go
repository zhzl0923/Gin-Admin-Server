package dao

import (
	"gin-admin/internal/model"

	"gorm.io/gorm"
)

type AdminRolePermissionDao interface {
	GetRolePermission(query ...interface{}) model.AdminRolePermission
	CreateAdminRolePermission(roleID uint64, permissions []uint64)
	DeleteAdminRolePermission(query ...interface{})
}

type adminRolePermissionDao struct {
	db *gorm.DB
}

func NewAdminRolePermissionDao(db *gorm.DB) AdminRolePermissionDao {
	return &adminRolePermissionDao{db: db}
}

func (d *adminRolePermissionDao) GetRolePermission(query ...interface{}) model.AdminRolePermission {
	var rolePermission model.AdminRolePermission
	GetWhereQuery(d.db, query...).First(&rolePermission)
	return rolePermission
}

func (d *adminRolePermissionDao) CreateAdminRolePermission(roleID uint64, permissions []uint64) {
	if len(permissions) == 0 {
		return
	}
	rolePermission := make([]model.AdminRolePermission, len(permissions))
	for key, id := range permissions {
		rolePermission[key] = model.AdminRolePermission{RoleID: roleID, PermissionID: id}
	}
	if len(rolePermission) > 0 {
		d.db.Create(&rolePermission)
	}
}

func (d *adminRolePermissionDao) DeleteAdminRolePermission(query ...interface{}) {
	GetWhereQuery(d.db, query...).Delete(&model.AdminRolePermission{})
}
