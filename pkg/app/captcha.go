package app

import (
	"image/color"
	"strings"

	"github.com/mojocn/base64Captcha"
)

type CaptchaOption struct {
	height          int
	width           int
	noiseCount      int
	showLineOptions int
	length          int
	source          string
	bgColor         *color.RGBA
	fonts           []string
}

func GenerateCaptcha(opts *CaptchaOption) (string, string, error) {
	c := base64Captcha.NewCaptcha(getCaptchaDriver(opts), base64Captcha.DefaultMemStore)
	return c.Generate()
}

func getCaptchaDriver(opts *CaptchaOption) base64Captcha.Driver {
	driver := base64Captcha.NewDriverString(
		opts.height,
		opts.width,
		opts.noiseCount,
		opts.showLineOptions,
		opts.length,
		opts.source,
		opts.bgColor,
		opts.fonts,
	)
	return driver
}

func NewCaptchaOption() *CaptchaOption {
	return &CaptchaOption{
		height:          54,
		width:           117,
		noiseCount:      5,
		showLineOptions: base64Captcha.OptionShowSlimeLine,
		length:          5,
		source:          base64Captcha.TxtNumbers + base64Captcha.TxtAlphabet,
		bgColor:         &color.RGBA{0, 0, 0, 0},
		fonts:           []string{"chromohv.ttf", "wqy-microhei.ttc", "DENNEthree-dee.ttf"},
	}
}

func (o *CaptchaOption) WithHeight(h int) *CaptchaOption {
	if h > 0 {
		o.height = h
	}
	return o
}

func (o *CaptchaOption) WithWidth(w int) *CaptchaOption {
	if w > 0 {
		o.width = w
	}
	return o
}
func (o *CaptchaOption) WithLength(l int) *CaptchaOption {
	if l > 0 {
		o.length = l
	}
	return o
}

func (o *CaptchaOption) WithNoiseCount(n int) *CaptchaOption {
	if n > 0 {
		o.noiseCount = n
	}
	return o
}

func ValidCaptcha(captchaId, captchaData string) bool {
	data := base64Captcha.DefaultMemStore.Get(captchaId, true)
	return strings.EqualFold(strings.ToLower(data), strings.ToLower(captchaData))
}
