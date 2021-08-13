package admin_role

import "gin-admin/internal/service"

type AdminRole struct {
	serv service.Service
}

func NewAdminRole(serv service.Service) *AdminRole {
	return &AdminRole{serv: serv}
}
