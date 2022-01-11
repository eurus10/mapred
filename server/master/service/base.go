package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"mapred/server/master/api"
	"mapred/server/master/config"
	"os"
	"time"
)

type BaseService struct{}

func (b *BaseService) UploadFile(ctx *context.Context, req *api.UploadFileReq) (*api.UploadFileResp, error) {
	filePath := fmt.Sprintf("%s/%s_%d", config.APPS, req.FileName, time.Now().Unix())
	file, err := os.Create(filePath)
	if err != nil {
		return &api.UploadFileResp{Success: false, Message: "服务端创建持久化文件失败"}, nil
	}
	fmt.Fprint(file, req.Data)
	return &api.UploadFileResp{Success: true, Message: "上传文件成功"}, nil
}

func (b *BaseService) DownloadFile(ctx *context.Context, req *api.DownloadFileReq) (*api.DownloadFileResp, error) {
	data, err := ioutil.ReadFile(req.FileName)
	if err != nil {
		return &api.DownloadFileResp{Success: false, Message: "文件不存在"}, nil
	}
	return &api.DownloadFileResp{Success: true, Message: "下载文件成功", FileName: req.FileName, Data: data}, nil
}
