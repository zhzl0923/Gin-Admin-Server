package service

import (
	"gin-admin/internal/dao"
	"gin-admin/internal/errcode"
	"gin-admin/internal/model"
)

type AdminPermissionService interface {
	GetAllAdminPermission() ([]*model.AdminPermission, error)
	GetAdminPermissionById(id uint64) (model.AdminPermission, error)
	GetAdminPermissionList(param *AdminPermissionListParam) (*model.AdminPermissionList, error)
	CreateAdminPermission(param *AdminPermissionParam) error
	UpdateAdminPermission(id uint64, param *AdminPermissionParam) error
	DeleteAdminPermissionById(id uint64) error
	UserHasPermission(adminUserID uint64, httpPath string, HttpMethod string) bool
}

type adminPermissionService struct {
	dao dao.Dao
}

type AdminPermissionListParam struct {
	//当前页数
	Page int `form:"page"`
	//每页条数
	PageSize int `form:"page_size"`
}

type AdminPermissionParam struct {
	// 权限名称
	Name string `form:"name" json:"name" binding:"required"`
	// 请求路径
	HttpPath string `form:"http_path" json:"http_path" binding:"required"`
	// 请求方法
	HttpMethod string `form:"http_method" json:"http_method" binding:"required"`
}

func newAdminPermissionService(dao dao.Dao) AdminPermissionService {
	return &adminPermissionService{dao: dao}
}

func (s adminPermissionService) GetAllAdminPermission() ([]*model.AdminPermission, error) {
	adminPermissionList, err := s.dao.AdminPermissionDao().GetAllAdminPermission()
	if err != nil {
		err = errcode.ErrGetAdminPermissionList.WithDetails(err.Error())
	}
	return adminPermissionList, err
}

func (s adminPermissionService) GetAdminPermissionList(param *AdminPermissionListParam) (*model.AdminPermissionList, error) {
	adminPermissionList, err := s.dao.AdminPermissionDao().GetAdminPermissionList(param.Page, param.PageSize)
	if err != nil {
		err = errcode.ErrGetAdminPermissionList.WithDetails(err.Error())
	}
	return adminPermissionList, err
}

func (s adminPermissionService) CreateAdminPermission(param *AdminPermissionParam) error {
	adminPermissionDao := s.dao.AdminPermissionDao()
	adminPermission, err := adminPermissionDao.GetAdminPermission("name=?", param.Name)

	if err == nil && adminPermission.ID > 0 {
		return errcode.ErrDuplicatePermission
	}
	err = adminPermissionDao.CreateAdminPermission(&dao.AdminPermission{
		Name:       param.Name,
		HttpMethod: param.HttpMethod,
		HttpPath:   param.HttpPath,
	})
	if err != nil {
		err = errcode.ErrCreateAdminPermission.WithDetails(err.Error())
	}
	return err
}

func (s adminPermissionService) UpdateAdminPermission(id uint64, param *AdminPermissionParam) error {
	err := s.dao.AdminPermissionDao().UpdateAdminPermission(id, &dao.AdminPermission{
		Name:       param.Name,
		HttpMethod: param.HttpMethod,
		HttpPath:   param.HttpPath,
	})
	if err != nil {
		err = errcode.ErrUpdateAdminPermission.WithDetails(err.Error())
	}
	return err
}

func (s adminPermissionService) GetAdminPermissionById(id uint64) (model.AdminPermission, error) {
	adminPermission, err := s.dao.AdminPermissionDao().GetAdminPermission("id=?", id)
	if err != nil {
		return adminPermission, errcode.AdminPermissionNotFound.WithDetails(err.Error())
	}
	if adminPermission.ID == 0 {
		return adminPermission, errcode.AdminPermissionNotFound.WithDetails(err.Error())
	}
	return adminPermission, nil
}

func (s adminPermissionService) DeleteAdminPermissionById(id uint64) error {
	if id == 0 {
		return errcode.AdminPermissionNotFound
	}
	err := s.dao.AdminPermissionDao().DeleteAdminPermissionById(id)
	if err != nil {
		return err
	}
	s.dao.AdminRolePermissionDao().DeleteAdminRolePermission("permission_id=?", id)
	return nil
}

func (s adminPermissionService) UserHasPermission(adminUserID uint64, httpPath string, HttpMethod string) bool {
	perm, _ := s.dao.AdminPermissionDao().GetAdminPermission("http_path=? AND http_method=?", httpPath, HttpMethod)
	if perm.ID > 0 {
		adminUser, _ := s.dao.AdminUserDao().GetAdminUser("id=?", adminUserID)
		if adminUser.IsSuper == 1 {
			return true
		}
		rolePerm := s.dao.AdminRolePermissionDao().GetRolePermission("role_id=? AND permission_id=?", adminUser.RoleID, perm.ID)
		if rolePerm.ID == 0 {
			return false
		}
	}
	return true
}
