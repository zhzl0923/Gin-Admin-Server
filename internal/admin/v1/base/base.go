package base

import (
	"gin-admin/internal/service"
)

type Base struct {
	serv service.Service
}

func NewBase(serv service.Service) *Base {
	return &Base{serv: serv}
}
