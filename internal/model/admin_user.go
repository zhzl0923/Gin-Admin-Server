package model

import (
	"gin-admin/pkg/auth"

	"gorm.io/gorm"
)

type AdminUser struct {
	Model

	// 用户名
	Username string `json:"username"`

	// 密码
	Password string `json:"-"`

	// 昵称
	Nickname string `json:"nickname"`

	// 头像
	Avatar string `json:"avatar"`

	// 是否为超级管理员，1是，0否
	IsSuper uint8 `json:"is_super"`

	//角色id
	RoleID uint64 `json:"role_id"`
}

type AdminUserList struct {
	// 总条数
	Total int64 `json:"total"`

	// 当前页数
	Page int `json:"page"`

	// 没页条数
	PageSize int `json:"page_size"`

	// 总页数
	LastPage int `json:"last_page"`

	// 当前页数据
	Items []*AdminUser `json:"items"`
}

func (m AdminUser) TableName() string {
	return "admin_user"
}

func (m *AdminUser) BeforeCreate(tx *gorm.DB) (err error) {
	m.Password, err = auth.Encrypt(m.Password)
	return
}
