package docs

import (
	"gin-admin/internal/service"
	"gin-admin/pkg/app"
)

// swagger:route GET /captcha 基础接口 Captcha
// 获取验证码
// Security:
//   Bearer:
//
//
// Responses:
//   200: getCaptchaResponse
//   default: errResponse

// swagger:operation POST /upload/file 基础接口 UploadFile
// 上传文件
// ---
// Security:
// - Bearer:
//
// consumes:
// - multipart/form-data
// Produces:
// - application/json
//
// parameters:
// - name: file
//   type: file
//   require: true
//   description: 文件
//   in: formData
//
// responses:
//   '200':
//     "$ref": "#/responses/uploadFileResponse"
//   default:
//     "$ref": "#/responses/errResponse"

// 获取验证码请求
// swagger:parameters Captcha
type GetCaptchaRequestWrapper struct {
	// in:query
	service.GenerateCaptchaParam
}

// 获取验证码响应
// swagger:response getCaptchaResponse
type GetCaptchaResponseWrapper struct {
	// in:body
	service.Captcha
}

// 发生错误响应
// swagger:response errResponse
type ErrResponseWrapper struct {
	// in:body
	Body app.ErrResponse
}

// swagger:response okResponse
type OkResponseWrapper struct {
	Body struct{}
}

// 上传文件响应
// swagger:response uploadFileResponse
type UploadFileResponseWrapper struct {
	// in:body

	// 文件路径
	Path string `json:"path"`
	// 访问host
	Host string `json:"host"`
}
