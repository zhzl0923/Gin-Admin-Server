package service

import (
	"gin-admin/internal/dao"
	"gin-admin/internal/errcode"
	"gin-admin/internal/model"
)

type AdminRoleService interface {
	GetAdminRoleById(id uint64) (*dao.AdminRole, error)
	GetAdminRoleList() ([]*model.AdminRole, error)
	CreateAdminRole(param *AdminRoleParam) error
	UpdateAdminRole(id uint64, param *AdminRoleParam) error
	DeleteAdminRoleById(id uint64) error
}

type adminRoleService struct {
	dao dao.Dao
}

type AdminRoleParam struct {

	//Required:true
	//角色名称
	RoleName string `json:"role_name" form:"role_name" binding:"required,min=1"`

	//菜单id
	Menus []uint64 `json:"menus" form:"menus"`

	//权限id
	Permissions []uint64 `json:"permissions" form:"permissions"`
}

func newAdminRoleService(dao dao.Dao) AdminRoleService {
	return &adminRoleService{dao: dao}
}

func (s adminRoleService) GetAdminRoleList() ([]*model.AdminRole, error) {
	adminRoleList, err := s.dao.AdminRoleDao().GetAdminRoleList()
	if err != nil {
		err = errcode.ErrGetAdminRoleList.WithDetails(err.Error())
	}
	return adminRoleList, err
}

func (s adminRoleService) CreateAdminRole(param *AdminRoleParam) error {
	adminRoleDao := s.dao.AdminRoleDao()
	adminRole, err := adminRoleDao.GetAdminRole("role_name=?", param.RoleName)
	if err == nil && adminRole.ID > 0 {
		return errcode.ErrDuplicateRole
	}
	roleId, err := adminRoleDao.CreateAdminRole(&dao.CreateAdminRoleParam{RoleName: param.RoleName})
	if err != nil {
		err = errcode.ErrCreateAdminRole.WithDetails(err.Error())
	}
	s.dao.AdminRoleMenuDao().CreateAdminRoleMenu(roleId, param.Menus)
	s.dao.AdminRolePermissionDao().CreateAdminRolePermission(roleId, param.Permissions)
	return err
}

func (s adminRoleService) UpdateAdminRole(id uint64, param *AdminRoleParam) error {
	adminRole, _ := s.dao.AdminRoleDao().GetAdminRole("role_name=?", param.RoleName)
	if adminRole.ID > 0 && adminRole.ID != id {
		return errcode.ErrDuplicateRole
	}
	err := s.dao.AdminRoleDao().UpdateAdminRole(id, param.RoleName)
	if err != nil {
		err = errcode.ErrUpdateAdminRole.WithDetails(err.Error())
	}
	adminRoleMenuDao := s.dao.AdminRoleMenuDao()
	adminRoleMenuDao.DeleteAdminRoleMenu("role_id=?", id)
	adminRoleMenuDao.CreateAdminRoleMenu(id, param.Menus)
	adminRolePermission := s.dao.AdminRolePermissionDao()
	adminRolePermission.DeleteAdminRolePermission("role_id=?", id)
	adminRolePermission.CreateAdminRolePermission(id, param.Permissions)
	return err
}

func (s adminRoleService) GetAdminRoleById(id uint64) (*dao.AdminRole, error) {
	var adminRole *dao.AdminRole
	var err error
	adminRole, err = s.dao.AdminRoleDao().GetAdminRole("id=?", id)
	if err != nil || adminRole.ID == 0 {
		return adminRole, errcode.AdminRoleNotFound.WithDetails(err.Error())
	}
	return adminRole, nil
}

func (s adminRoleService) DeleteAdminRoleById(id uint64) error {
	if id <= 0 {
		return errcode.AdminRoleNotFound
	}
	err := s.dao.AdminRoleDao().DeleteAdminRoleById(id)
	if err != nil {
		return err
	}
	s.dao.AdminRoleMenuDao().DeleteAdminRoleMenu("role_id=?", id)
	return nil
}
