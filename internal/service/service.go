package service

import (
	"gin-admin/internal/dao"
)

type Service interface {
	AdminUserService() AdminUserService
	AuthService() AuthService
	CaptchaService() CaptchaService
	UploadService() UploadService
	AdminRoleService() AdminRoleService
	AdminMenuService() AdminMenuService
	AdminPermissionService() AdminPermissionService
}

type service struct {
	dao dao.Dao
}

func NewService(dao dao.Dao) Service {
	return &service{dao: dao}
}

func (s service) AdminUserService() AdminUserService {
	return newAdminUserService(s.dao)
}

func (s service) AuthService() AuthService {
	return NewAuthService(s.dao)
}

func (s service) CaptchaService() CaptchaService {
	return newCaptchaService()
}

func (s service) UploadService() UploadService {
	return newUploadService()
}

func (s service) AdminRoleService() AdminRoleService {
	return newAdminRoleService(s.dao)
}

func (s service) AdminMenuService() AdminMenuService {
	return newAdminMenuService(s.dao)
}

func (s service) AdminPermissionService() AdminPermissionService {
	return newAdminPermissionService(s.dao)
}
