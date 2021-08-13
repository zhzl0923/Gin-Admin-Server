package service

import (
	"gin-admin/internal/dao"
	"gin-admin/internal/errcode"
	"gin-admin/internal/model"
)

type AdminUserService interface {
	GetAdminUserById(id uint64) (model.AdminUser, error)
	GetAdminUserList(param *AdminUserListParam) (*model.AdminUserList, error)
	CreateAdminUser(param *CreateAdminUserParam) error
	UpdateAdminUser(id uint64, param *UpdateAdminUserParam) error
	DeleteAdminUserById(id uint64) error
}

type adminUserService struct {
	dao dao.Dao
}

type AdminUserListParam struct {
	// 当前页数
	Page int `form:"page"`

	// 每页条数
	PageSize int `form:"page_size"`
}

type CreateAdminUserParam struct {
	// Required: true
	// 用户名
	Username string `form:"username" json:"username" binding:"required,min=5,max=18"`

	//昵称
	Nickname string `form:"nickname" json:"nickname" binding:"max=20"`

	// Required: true
	//密码
	Password string `form:"password" json:"password" binding:"required,min=6,max=18"`

	//头像
	Avatar string `form:"avatar" json:"avatar"`

	//是否为超级管理员，1是，0 否
	IsSuper uint8 `form:"is_super" json:"is_super"`

	//角色ID
	RoleID uint64 `form:"role_id" json:"role_id"`
}

type UpdateAdminUserParam struct {
	Nickname string `form:"nickname" json:"nickname" binding:"max=20"`
	Password string `form:"password" json:"password" binding:"max=18"`
	Avatar   string `form:"avatar" json:"avatar"`
	IsSuper  uint8  `form:"is_super" json:"is_super"`
	RoleID   uint64 `form:"role_id" json:"role_id"`
}

func newAdminUserService(dao dao.Dao) AdminUserService {
	return &adminUserService{dao: dao}
}

func (s adminUserService) GetAdminUserList(param *AdminUserListParam) (*model.AdminUserList, error) {
	adminUserList, err := s.dao.AdminUserDao().GetAdminUserList(param.Page, param.PageSize)
	if err != nil {
		err = errcode.ErrGetAdminUserList.WithDetails(err.Error())
	}
	return adminUserList, err
}

func (s adminUserService) CreateAdminUser(param *CreateAdminUserParam) error {
	adminUserDao := s.dao.AdminUserDao()
	adminUser, err := adminUserDao.GetAdminUser("username=?", param.Username)

	if err == nil && adminUser.ID > 0 {
		return errcode.ErrDuplicateUsername
	}
	err = adminUserDao.CreateAdminUser(&dao.AdminUser{
		Username: param.Username,
		Nickname: param.Nickname,
		Password: param.Password,
		Avatar:   param.Avatar,
		IsSuper:  param.IsSuper,
		RoleID:   param.RoleID,
	})
	if err != nil {
		err = errcode.ErrCreateAdminUser.WithDetails(err.Error())
	}
	return err
}

func (s adminUserService) UpdateAdminUser(id uint64, param *UpdateAdminUserParam) error {
	err := s.dao.AdminUserDao().UpdateAdminUser(id, &dao.AdminUser{
		Nickname: param.Nickname,
		Password: param.Password,
		Avatar:   param.Avatar,
		IsSuper:  param.IsSuper,
		RoleID:   param.RoleID,
	})
	if err != nil {
		err = errcode.ErrUpdateAdminUser.WithDetails(err.Error())
	}
	return err
}

func (s adminUserService) GetAdminUserById(id uint64) (model.AdminUser, error) {
	adminUser, err := s.dao.AdminUserDao().GetAdminUser("id=?", id)
	if err != nil {
		return adminUser, errcode.AdminUserNotFound.WithDetails(err.Error())
	}
	if adminUser.ID == 0 {
		return adminUser, errcode.AdminUserNotFound.WithDetails(err.Error())
	}
	return adminUser, nil
}

func (s adminUserService) DeleteAdminUserById(id uint64) error {
	if id <= 0 {
		return errcode.AdminUserNotFound
	}
	return s.dao.AdminUserDao().DeleteAdminUserById(id)
}
