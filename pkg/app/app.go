package app

import (
	"gin-admin/pkg/err"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	//错误码
	Code int `json:"code"`

	//错误信息
	Msg string `json:"msg"`

	//详细信息
	Details []string `json:"details"`
}

func WriteResponse(c *gin.Context, e error, data interface{}) {
	if e != nil {
		err := err.ParseErr(e)
		c.JSON(err.HttpStatus(), ErrResponse{
			Code:    err.Code(),
			Msg:     err.Msg(),
			Details: err.Details(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
