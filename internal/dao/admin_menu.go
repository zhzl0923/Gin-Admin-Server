package dao

import (
	"errors"
	"gin-admin/internal/model"

	"gorm.io/gorm"
)

type AdminMenuDao interface {
	GetAdminMenu(query ...interface{}) (model.AdminMenu, error)
	GetAdminMenuList(query ...interface{}) model.AdminMenuList
	CreateAdminMenu(param *AdminMenu) error
	UpdateAdminMenu(id uint64, param *AdminMenu) error
	DeleteAdminMenuById(id uint64) error
}

type adminMenuDao struct {
	db *gorm.DB
}

type AdminMenu struct {
	ParentID   uint64
	Name       string
	Permission string
	Path       string
	Type       uint8
	Icon       string
	Component  string
	Sort       uint64
	IsDisabled uint8
}

func NewAdminMenuDao(db *gorm.DB) AdminMenuDao {
	return &adminMenuDao{db: db}
}

func (r *adminMenuDao) GetAdminMenuList(query ...interface{}) model.AdminMenuList {
	var adminMenuList model.AdminMenuList
	GetWhereQuery(r.db, query...).Order("sort ASC").Order("id ASC").Find(&adminMenuList)
	return adminMenuList
}

func (r *adminMenuDao) GetAdminMenu(query ...interface{}) (model.AdminMenu, error) {
	var adminMenu model.AdminMenu
	err := GetWhereQuery(r.db, query...).First(&adminMenu).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return adminMenu, err
	}
	return adminMenu, nil
}

func (r *adminMenuDao) CreateAdminMenu(param *AdminMenu) error {
	adminMenu := model.AdminMenu{
		ParentID:   param.ParentID,
		Name:       param.Name,
		Permission: param.Permission,
		Path:       param.Path,
		Type:       param.Type,
		Icon:       param.Icon,
		Component:  param.Component,
		IsDisabled: param.IsDisabled,
		Sort:       param.Sort,
	}
	return r.db.Create(&adminMenu).Error
}

func (r *adminMenuDao) UpdateAdminMenu(id uint64, param *AdminMenu) error {
	if id == 0 {
		return nil
	}

	values := map[string]interface{}{
		"parent_id":   param.ParentID,
		"type":        param.Type,
		"is_disabled": param.IsDisabled,
		"icon":        param.Icon,
		"component":   param.Component,
		"path":        param.Path,
		"sort":        param.Sort,
		"permission":  param.Permission,
	}

	if param.Name != "" {
		values["name"] = param.Name
	}

	return r.db.Model(&model.AdminMenu{Model: model.Model{ID: id}}).Updates(values).Error
}

func (r *adminMenuDao) DeleteAdminMenuById(id uint64) error {
	return r.db.Delete(&model.AdminMenu{}, id).Error
}
