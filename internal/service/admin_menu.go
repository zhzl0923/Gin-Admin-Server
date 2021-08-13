package service

import (
	"gin-admin/internal/dao"
	"gin-admin/internal/errcode"
	"gin-admin/internal/model"
)

type AdminMenuService interface {
	GetAdminMenuList() model.AdminMenuList
	GetAdminMenuById(id uint64) (model.AdminMenu, error)
	CreateAdminMenu(param *AdminMenuParam) error
	UpdateAdminMenu(id uint64, param *AdminMenuParam) error
	DeleteAdminMenuById(id uint64) error
	GetAdminMenuByUserId(id uint64) (model.AdminMenuList, error)
}

type adminMenuService struct {
	dao dao.Dao
}

type AdminMenuParam struct {
	ParentID   uint64 `form:"parent_id" json:"parent_id"`
	Name       string `form:"name" json:"name" binding:"required,min=1,max=20"`
	Permission string `form:"permission" json:"permission"`
	Path       string `form:"path" json:"path"`
	Type       uint8  `form:"type" json:"type"`
	Icon       string `form:"icon" json:"icon"`
	Component  string `form:"component" json:"component"`
	Sort       uint64 `formL:"sort" json:"sort"`
	IsDisabled uint8  `form:"is_disabled" json:"is_disabled"`
}

func newAdminMenuService(dao dao.Dao) AdminMenuService {
	return &adminMenuService{dao: dao}
}

func (s *adminMenuService) GetAdminMenuList() model.AdminMenuList {
	list := s.dao.AdminMenuDao().GetAdminMenuList()
	return list.ToTree()
}

func (s *adminMenuService) CreateAdminMenu(param *AdminMenuParam) error {
	adminMenuDao := s.dao.AdminMenuDao()
	adminMenu, _ := adminMenuDao.GetAdminMenu("name=?", param.Name)
	if adminMenu.ID > 0 {
		return errcode.ErrDuplicateMenu
	}
	err := adminMenuDao.CreateAdminMenu(&dao.AdminMenu{
		ParentID:   param.ParentID,
		Name:       param.Name,
		Permission: param.Permission,
		Path:       param.Path,
		Type:       param.Type,
		Icon:       param.Icon,
		Component:  param.Component,
		IsDisabled: param.IsDisabled,
		Sort:       param.Sort,
	})
	if err != nil {
		return errcode.ErrCreateAdminMenu.WithDetails(err.Error())
	}
	return nil
}

func (s *adminMenuService) UpdateAdminMenu(id uint64, param *AdminMenuParam) error {
	adminMenuDao := s.dao.AdminMenuDao()
	adminMenu, _ := adminMenuDao.GetAdminMenu("name=?", param.Name)
	if adminMenu.ID > 0 && adminMenu.ID != id {
		return errcode.ErrDuplicateMenu
	}
	err := adminMenuDao.UpdateAdminMenu(id, &dao.AdminMenu{
		ParentID:   param.ParentID,
		Name:       param.Name,
		Permission: param.Permission,
		Path:       param.Path,
		Type:       param.Type,
		Icon:       param.Icon,
		Component:  param.Component,
		IsDisabled: param.IsDisabled,
		Sort:       param.Sort,
	})
	if err != nil {
		return errcode.ErrUpdateAdminMenu
	}
	return nil
}

func (s *adminMenuService) GetAdminMenuById(id uint64) (model.AdminMenu, error) {
	adminMenu, err := s.dao.AdminMenuDao().GetAdminMenu("id=?", id)
	if err != nil || adminMenu.ID == 0 {
		return adminMenu, errcode.AdminMenuNotFound.WithDetails(err.Error())
	}
	return adminMenu, nil
}

func (s *adminMenuService) DeleteAdminMenuById(id uint64) error {
	adminMenuDao := s.dao.AdminMenuDao()
	menu, _ := adminMenuDao.GetAdminMenu("parent_id=?", id)
	if menu.ID > 0 {
		return errcode.ErrMenuHasSubMenu
	}
	err := adminMenuDao.DeleteAdminMenuById(id)
	if err != nil {
		return errcode.ErrDeleteAdminMenu.WithDetails(err.Error())
	}
	s.dao.AdminRoleMenuDao().DeleteAdminRoleMenu("menu_id=?", id)
	return nil
}

func (s *adminMenuService) GetAdminMenuByUserId(id uint64) (model.AdminMenuList, error) {
	var menu model.AdminMenuList
	adminUser, err := s.dao.AdminUserDao().GetAdminUser("id=?", id)
	if err != nil {
		return menu, err
	}
	if adminUser.IsSuper == 1 {
		menu = s.dao.AdminMenuDao().GetAdminMenuList()
	} else {
		menuIds := s.dao.AdminRoleMenuDao().GetAdminMenuIds("role_id=?", adminUser.RoleID)
		menu = s.dao.AdminMenuDao().GetAdminMenuList("id in ?", menuIds)
	}
	return menu.ToTree(), nil
}
