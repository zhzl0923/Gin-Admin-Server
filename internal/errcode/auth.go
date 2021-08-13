package errcode

import "gin-admin/pkg/err"

var (
	Unauthorized               = err.NewError(100300, 403, "未登录")
	UnauthorizedInvalid        = err.NewError(100301, 403, "用户名或密码错误")
	UnauthorizedTokenGenerate  = err.NewError(100303, 403, "token生成失败")
	UnauthorizedInvalidCaptcha = err.NewError(100304, 403, "验证码错误")
	UnauthorizedTokenTimeout   = err.NewError(100305, 403, "token失效")
	UnauthorizedTokenError     = err.NewError(100306, 403, "token错误")
)
