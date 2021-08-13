package base

import (
	"gin-admin/global"
	"gin-admin/internal/errcode"
	"gin-admin/internal/service"
	"gin-admin/pkg/app"
	"gin-admin/pkg/convert"
	"gin-admin/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (b Base) UploadFile(c *gin.Context) {

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams.WithDetails(err.Error()), gin.H{})
		return
	}
	param := service.UploadFileParam{
		File:       file,
		FileHeader: fileHeader,
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		app.WriteResponse(c, errcode.InvalidParams, gin.H{})
		return
	}
	param.FileType = utils.FileType(fileType)

	fileInfo, err := b.serv.UploadService().UploadFile(&param)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		app.WriteResponse(c, errcode.ErrorUploadFileFail.WithDetails(err.Error()), gin.H{})
		return
	}

	app.WriteResponse(c, nil, gin.H{
		"path": fileInfo.Path,
		"host": fileInfo.Host,
	})
}
