package dao

import (
	"errors"
	"gin-admin/internal/model"
	"gin-admin/pkg/app"
	"gin-admin/pkg/auth"
	"math"

	"gorm.io/gorm"
)

type AdminUserDao interface {
	GetAdminUser(query ...interface{}) (model.AdminUser, error)
	GetAdminUserList(page, pageSize int, query ...interface{}) (*model.AdminUserList, error)
	CreateAdminUser(param *AdminUser) error
	UpdateAdminUser(id uint64, param *AdminUser) error
	DeleteAdminUserById(id uint64) error
}

type AdminUser struct {
	Username string
	Password string
	Nickname string
	Avatar   string
	IsSuper  uint8
	RoleID   uint64
}

type adminUserDao struct {
	db *gorm.DB
}

func NewAdminUserDao(db *gorm.DB) AdminUserDao {
	return &adminUserDao{db: db}
}

func (u *adminUserDao) GetAdminUser(query ...interface{}) (model.AdminUser, error) {
	var adminUser model.AdminUser
	err := GetWhereQuery(u.db, query...).First(&adminUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return adminUser, err
	}
	return adminUser, nil
}

func (u *adminUserDao) GetAdminUserList(page, pageSize int, query ...interface{}) (*model.AdminUserList, error) {
	offset := app.GetPageOffset(page, pageSize)
	list := &model.AdminUserList{Page: page, PageSize: pageSize}
	err := GetWhereQuery(u.db, query...).
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

func (u *adminUserDao) CreateAdminUser(param *AdminUser) error {
	adminUser := model.AdminUser{
		Username: param.Username,
		Nickname: param.Nickname,
		Password: param.Password,
		Avatar:   param.Avatar,
		IsSuper:  param.IsSuper,
		RoleID:   param.RoleID,
	}

	return u.db.Create(&adminUser).Error
}

func (u *adminUserDao) UpdateAdminUser(id uint64, param *AdminUser) error {

	if id == 0 {
		return nil
	}
	values := map[string]interface{}{
		"nickname": param.Nickname,
		"avatar":   param.Avatar,
		"role_id":  param.RoleID,
	}

	if param.Password != "" {
		values["password"], _ = auth.Encrypt(param.Password)
	}

	if param.IsSuper != 1 {
		values["is_super"] = 0
	}

	return u.db.Model(&model.AdminUser{Model: model.Model{ID: id}}).Updates(values).Error
}

func (u *adminUserDao) DeleteAdminUserById(id uint64) error {
	return u.db.Delete(&model.AdminUser{}, id).Error
}
