package app

import (
	"gin-admin/global"
	"strings"

	"github.com/gin-gonic/gin"

	valdate "github.com/go-playground/validator/v10"
)

type ValidateError struct {
	Key string
	Msg string
}

type ValidateErrors []*ValidateError

func (v *ValidateError) Error() string {
	return v.Msg
}

func (v ValidateErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidateErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidateErrors) {
	var errs ValidateErrors
	err := c.ShouldBind(v)
	if err != nil {
		verrs, ok := err.(valdate.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(global.Trans) {
			errs = append(errs, &ValidateError{
				Key: key,
				Msg: value,
			})
		}
		return false, errs
	}
	return true, nil
}
