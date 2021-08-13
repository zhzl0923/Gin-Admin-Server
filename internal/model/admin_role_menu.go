package model

type AdminRoleMenu struct {
	ID     uint64 `gorm:"primary_key" json:"id"` //自增ID
	RoleID uint64 `json:"role_id"`               //角色ID
	MenuID uint64 `json:"menu_id"`               //菜单ID
}

func (model AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}
