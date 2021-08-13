package admin_menu

import "gin-admin/internal/service"

type AdminMenu struct {
	serv service.Service
}

func NewAdminMenu(serv service.Service) *AdminMenu {
	return &AdminMenu{serv: serv}
}
