package auth

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"

	"github.com/gin-gonic/gin"
)

func (a *Auth) Login(c *gin.Context) {
	var param service.LoginParam
	var err error
	valid, err := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("login BindAndValid err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.(app.ValidateErrors).Errors()...), nil)
		return
	}
	if !app.ValidCaptcha(param.CaptchaId, param.Captcha) {
		app.WriteResponse(c, errcode.UnauthorizedInvalidCaptcha, nil)
		return
	}
	token, err := a.serv.AuthService().Authorize(&param)
	if err != nil {
		app.WriteResponse(c, err, nil)
		return
	}
	app.WriteResponse(c, nil, gin.H{"token": token})
}
