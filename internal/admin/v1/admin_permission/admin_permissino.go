package admin_permission

import (
	"gin-admin/internal/service"
)

type AdminPermission struct {
	serv service.Service
}

func NewAdminPermission(serv service.Service) *AdminPermission {
	return &AdminPermission{serv: serv}
}
