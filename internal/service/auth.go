package service

import (
	"gin-admin/internal/dao"
	"gin-admin/internal/errcode"
	"gin-admin/pkg/app"
	"gin-admin/pkg/auth"
)

type LoginParam struct {
	Username  string `form:"username" json:"username" binding:"min=5,max=18"`
	Password  string `form:"password" json:"password" binding:"min=6,max=18"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required"`
}

type AuthService interface {
	Authorize(param *LoginParam) (string, error)
}

type authService struct {
	dao dao.Dao
}

func NewAuthService(dao dao.Dao) *authService {
	return &authService{dao: dao}
}

func (a authService) Authorize(param *LoginParam) (string, error) {

	adminUser, err := a.dao.AdminUserDao().GetAdminUser("username=?", param.Username)

	if err != nil {
		return "", errcode.UnauthorizedInvalid.WithDetails(err.Error())
	}

	if adminUser.ID == 0 {
		return "", errcode.UnauthorizedInvalid
	}

	err = auth.Compare(adminUser.Password, param.Password)

	if err != nil {
		return "", errcode.UnauthorizedInvalid
	}

	token, err := app.GenerateToken(uint64(adminUser.ID))

	if err != nil {
		return "", errcode.UnauthorizedTokenGenerate.WithDetails(err.Error())
	}

	return token, err
}
