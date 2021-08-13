package base

import (
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

// @Summary 获取验证码
// @Produce json
// @Param username body string true "用户名" minlength(6) maxlength(18)
// @Param password body string true "密码" minlength(6) maxlength(18)
// @Success 200 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
func (b *Base) Captcha(c *gin.Context) {

	var param service.GenerateCaptchaParam
	var err error
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}
	captcha, err := b.serv.CaptchaService().GenetateCaptcha(&param)

	if err != nil {
		app.WriteResponse(c, errcode.ErrGenerateCaptcha, nil)
		return
	}

	app.WriteResponse(c, nil, captcha)
}
