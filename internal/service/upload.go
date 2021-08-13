package service

import (
	"errors"
	"gin-admin/global"
	"gin-admin/pkg/utils"
	"mime/multipart"
	"os"
)

type UploadService interface {
	UploadFile(param *UploadFileParam) (*FileInfo, error)
}

type uploadService struct {
}

type FileInfo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Host string `json:"host"`
}

type UploadFileParam struct {
	FileType   utils.FileType
	File       multipart.File
	FileHeader *multipart.FileHeader
}

func newUploadService() UploadService {
	return &uploadService{}
}

func (u uploadService) UploadFile(param *UploadFileParam) (*FileInfo, error) {
	fileName := utils.GetFileName(param.FileHeader.Filename)
	uploadSavePath := utils.GetSavePath()
	dst := global.AppSetting.UploadSavePath + "/" + uploadSavePath
	if !utils.CheckContainExt(param.FileType, fileName) {
		return nil, errors.New("不支持的文件后缀")
	}
	if utils.CheckSavePath(dst) {
		err := utils.CreateSavePath(dst, os.ModePerm)
		if err != nil {
			return nil, errors.New("创建上传目录失败")
		}
	}
	if utils.CheckMaxSize(param.FileType, param.File) {
		return nil, errors.New("上传文件超出最大文件限制")
	}
	if utils.CheckPermission(uploadSavePath) {
		return nil, errors.New("文件权限不足")
	}
	if err := utils.SaveFile(param.FileHeader, dst+"/"+fileName); err != nil {
		return nil, err
	}
	return &FileInfo{Name: fileName, Path: "/" + uploadSavePath + "/" + fileName, Host: global.AppSetting.UploadServerUrl}, nil
}
