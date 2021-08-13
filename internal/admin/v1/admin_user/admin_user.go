package admin_user

import (
	"gin-admin/internal/service"
)

type AdminUser struct {
	serv service.Service
}

func NewAdminUser(serv service.Service) *AdminUser {
	return &AdminUser{serv: serv}
}
