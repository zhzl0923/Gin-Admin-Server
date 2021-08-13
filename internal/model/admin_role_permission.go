package model

type AdminRolePermission struct {
	ID           uint64 `json:"id"`            //
	RoleID       uint64 `json:"role_id"`       // 角色id
	PermissionID uint64 `json:"permission_id"` // 权限id

}

func (model AdminRolePermission) TableName() string {
	return "admin_role_permission"
}
