package dao

import (
	"gorm.io/gorm"
)

type Dao interface {
	AdminUserDao() AdminUserDao
	AdminRoleDao() AdminRoleDao
	AdminMenuDao() AdminMenuDao
	AdminRoleMenuDao() AdminRoleMenuDao
	AdminPermissionDao() AdminPermissionDao
	AdminRolePermissionDao() AdminRolePermissionDao
}

type dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) Dao {
	return &dao{db: db}
}

func (d *dao) AdminUserDao() AdminUserDao {
	return NewAdminUserDao(d.db)
}

func (d *dao) AdminRoleDao() AdminRoleDao {
	return NewAdminRoleDao(d.db)
}

func (d *dao) AdminMenuDao() AdminMenuDao {
	return NewAdminMenuDao(d.db)
}

func (d *dao) AdminRoleMenuDao() AdminRoleMenuDao {
	return NewAdminRoleMenuDao(d.db)
}

func (d *dao) AdminPermissionDao() AdminPermissionDao {
	return NewAdminPermissionDao(d.db)
}

func (d *dao) AdminRolePermissionDao() AdminRolePermissionDao {
	return NewAdminRolePermissionDao(d.db)
}

func GetWhereQuery(db *gorm.DB, query ...interface{}) *gorm.DB {
	if len(query) >= 2 {
		db = db.Where(query[0], query[1:]...)
	}
	return db
}
