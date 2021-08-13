package dao

import (
	"errors"
	"gin-admin/internal/model"

	"gorm.io/gorm"
)

type AdminRoleDao interface {
	GetAdminRole(query ...interface{}) (*AdminRole, error)
	GetAdminRoleList(query ...interface{}) ([]*model.AdminRole, error)
	CreateAdminRole(param *CreateAdminRoleParam) (uint64, error)
	UpdateAdminRole(id uint64, roleName string) error
	DeleteAdminRoleById(id uint64) error
}

type adminRoleDao struct {
	db *gorm.DB
}

type AdminRole struct {
	// 角色id
	ID uint64 `json:"id"`

	//角色名称
	RoleName string `json:"role_name"`

	//菜单id
	Menus []uint64 `json:"menus"`

	//权限id
	Permissions []uint64 `json:"permissions"`
}

type CreateAdminRoleParam struct {
	RoleName string `json:"role_name" form:"role_name"`
}

type UpdateAdminRoleParam struct {
	ID       uint64 `json:"id" form:"id"`
	RoleName string `json:"role_name" form:"role_name"`
}

func NewAdminRoleDao(db *gorm.DB) AdminRoleDao {
	return &adminRoleDao{db: db}
}

func (r *adminRoleDao) GetAdminRole(query ...interface{}) (*AdminRole, error) {
	var adminRole model.AdminRole
	err := GetWhereQuery(r.db, query...).Preload("Menus").Preload("Permissions").First(&adminRole).Error
	var role *AdminRole
	if err != nil {
		return role, err
	}
	role = &AdminRole{
		ID:          adminRole.ID,
		RoleName:    adminRole.RoleName,
		Menus:       make([]uint64, len(adminRole.Menus)),
		Permissions: make([]uint64, len(adminRole.Permissions)),
	}
	for key, roleMenu := range adminRole.Menus {
		role.Menus[key] = roleMenu.ID
	}
	for key, permission := range adminRole.Permissions {
		role.Permissions[key] = permission.ID
	}
	return role, nil
}

func (r *adminRoleDao) GetAdminRoleList(query ...interface{}) ([]*model.AdminRole, error) {
	var list []*model.AdminRole
	err := GetWhereQuery(r.db, query...).Order("id desc").Find(&list).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return list, nil
}

func (r *adminRoleDao) CreateAdminRole(param *CreateAdminRoleParam) (uint64, error) {
	adminRole := model.AdminRole{
		RoleName: param.RoleName,
	}

	err := r.db.Create(&adminRole).Error
	if err != nil {
		return 0, err
	}
	return adminRole.ID, nil
}

func (r *adminRoleDao) UpdateAdminRole(id uint64, roleName string) error {
	if id == 0 {
		return nil
	}

	values := map[string]interface{}{
		"role_name": roleName,
	}

	return r.db.Model(&model.AdminRole{Model: model.Model{ID: id}}).Updates(values).Error
}

func (r *adminRoleDao) DeleteAdminRoleById(id uint64) error {
	return r.db.Delete(&model.AdminRole{}, id).Error
}
