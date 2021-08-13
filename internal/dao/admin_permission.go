package dao

import (
	"errors"
	"gin-admin/internal/errcode"
	"gin-admin/internal/model"
	"gin-admin/pkg/app"
	"math"

	"gorm.io/gorm"
)

type AdminPermissionDao interface {
	GetAllAdminPermission() ([]*model.AdminPermission, error)
	GetAdminPermission(query ...interface{}) (model.AdminPermission, error)
	GetAdminPermissionList(page, pageSize int, query ...interface{}) (*model.AdminPermissionList, error)
	CreateAdminPermission(param *AdminPermission) error
	UpdateAdminPermission(id uint64, param *AdminPermission) error
	DeleteAdminPermissionById(id uint64) error
}

type AdminPermission struct {
	Name       string
	HttpMethod string
	HttpPath   string
}

type adminPermissionDao struct {
	db *gorm.DB
}

func NewAdminPermissionDao(db *gorm.DB) AdminPermissionDao {
	return &adminPermissionDao{db: db}
}

func (p *adminPermissionDao) GetAllAdminPermission() ([]*model.AdminPermission, error) {
	var list []*model.AdminPermission
	err := p.db.Order("id desc").Find(&list).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return list, nil
}

func (p *adminPermissionDao) GetAdminPermission(query ...interface{}) (model.AdminPermission, error) {
	var adminPermission model.AdminPermission
	err := GetWhereQuery(p.db, query...).First(&adminPermission).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return adminPermission, err
	}
	return adminPermission, nil
}

func (p *adminPermissionDao) GetAdminPermissionList(page, pageSize int, query ...interface{}) (*model.AdminPermissionList, error) {
	offset := app.GetPageOffset(page, pageSize)
	list := &model.AdminPermissionList{Page: page, PageSize: pageSize}
	err := GetWhereQuery(p.db, query...).
		Offset(offset).
		Limit(pageSize).
		Order("id desc").
		Find(&list.Items).
		Offset(-1).
		Limit(-1).
		Count(&list.Total).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	list.LastPage = int(math.Ceil(float64(list.Total) / float64(pageSize)))
	return list, nil
}

func (p *adminPermissionDao) CreateAdminPermission(param *AdminPermission) error {
	adminPermission := model.AdminPermission{
		Name:       param.Name,
		HttpPath:   param.HttpPath,
		HttpMethod: param.HttpMethod,
	}

	return p.db.Create(&adminPermission).Error
}

func (p *adminPermissionDao) UpdateAdminPermission(id uint64, param *AdminPermission) error {

	if id == 0 {
		return nil
	}
	perm, _ := p.GetAdminPermission("http_path=? and http_method=?", param.HttpPath, param.HttpMethod)
	if perm.ID > 0 && perm.ID != id {
		return errcode.ErrDuplicatePermission
	}
	values := map[string]interface{}{
		"name":        param.Name,
		"http_path":   param.HttpPath,
		"http_method": param.HttpMethod,
	}
	err := p.db.Model(&model.AdminPermission{Model: model.Model{ID: id}}).Updates(values).Error

	if err != nil {
		return errcode.ErrUpdateAdminPermission.WithDetails(err.Error())
	}
	return nil
}

func (p *adminPermissionDao) DeleteAdminPermissionById(id uint64) error {
	return p.db.Delete(&model.AdminPermission{}, id).Error
}
