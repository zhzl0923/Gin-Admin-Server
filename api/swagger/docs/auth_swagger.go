package docs

import (
	"gin-admin/internal/service"
)

// swagger:route POST /login 基础接口 Login
// 登录
//
// Security:
// -
// Responses:
//   200: loginResponse
//   default: errResponse

// swagger:parameters Login
type LoginRequestWrapper struct {
	//in:body
	service.LoginParam
}

// 登录成功返回
// swagger:response loginResponse
type LoginResponseWrapper struct {
	Body struct {
		Token string `json:"token"`
	}
}
