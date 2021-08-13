package errcode

import "gin-admin/pkg/err"

var (
	// InvalidParams - 400: Validation failed.
	InvalidParams = err.NewError(100001, 400, "请求参数错误")

	// ErrTokenInvalid - 401: Token invalid.
	TokenInvalid = err.NewError(100002, 401, "Token invalid")

	// ErrPageNotFound - 404: Page not found.
	PageNotFound = err.NewError(100003, 404, "Page not found")

	Forbidden = err.NewError(100004, 403, "暂无权限")
)

var (
	ErrGenerateCaptcha  = err.NewError(100101, 400, "生成图形验证码失败")
	ErrorUploadFileFail = err.NewError(100102, 400, "上传文件失败")
)
