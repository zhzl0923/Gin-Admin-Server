package model

type AdminRole struct {
	Model
	RoleName    string             `json:"role_name"` //角色名称
	Menus       []*AdminMenu       `json:"menus" gorm:"many2many:admin_role_menu;foreignKey:ID;joinForeignKey:RoleID;References:ID;JoinReferences:MenuID"`
	Permissions []*AdminPermission `json:"permissions" gorm:"many2many:admin_role_permission;foreignKey:ID;joinForeignKey:RoleID;References:ID;JoinReferences:PermissionID"`
}

func (model AdminRole) TableName() string {
	return "admin_role"
}
