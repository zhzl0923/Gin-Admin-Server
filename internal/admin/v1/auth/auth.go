package auth

import (
	"gin-admin/internal/service"
)

type Auth struct {
	serv service.Service
}

func NewAuth(serv service.Service) *Auth {
	return &Auth{serv: serv}
}
