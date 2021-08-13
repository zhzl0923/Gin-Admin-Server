package service

import (
	"gin-admin/pkg/app"
)

type GenerateCaptchaParam struct {
	//图片宽度
	Width int `form:"width" json:"width"`
	//图片高度
	Height int `form:"height" json:"height"`
	//验证码长度
	Length int `form:"length" json:"length"`
}

type CaptchaService interface {
	GenetateCaptcha(param *GenerateCaptchaParam) (*Captcha, error)
}

type captchaService struct {
}

type Captcha struct {
	//验证码id
	Id string `json:"captcha_id"`
	//验证码
	Data string `json:"captcha"`
}

func newCaptchaService() *captchaService {
	return &captchaService{}
}

func (s captchaService) GenetateCaptcha(param *GenerateCaptchaParam) (*Captcha, error) {
	opt := app.NewCaptchaOption()
	opt = opt.WithWidth(param.Width).
		WithHeight(param.Height).
		WithLength(param.Length)
	id, b64s, err := app.GenerateCaptcha(opt)
	if err != nil {
		return nil, err
	}
	captcha := &Captcha{
		Id:   id,
		Data: b64s,
	}
	return captcha, nil
}
