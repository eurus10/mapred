package client

import (
	"context"
	"fmt"
	api "mapred/server/master/api/client"
	"mapred/server/master/config"
	"mapred/server/master/scheduler"
	"os"
	"time"
)

type ClientService struct{}

func (c *ClientService) Submit(ctx context.Context, req *api.SubmitReq) (*api.SubmitResp, error) {
	jobName := fmt.Sprintf("%s_%d", req.JobName, time.Now().Unix())
	// create Job WorkSpace
	jobDirPath := fmt.Sprintf("%s/%s", config.APPS, jobName)
	err := os.Mkdir(jobDirPath, os.ModePerm)
	if err != nil {
		return &api.SubmitResp{Success: false, Message: "提交任务失败"}, nil
	}
	// save Job File (.so)
	jobFilePath := fmt.Sprintf("%s/job.so", jobDirPath)
	jobFile, err := os.Create(jobFilePath)
	if err != nil {
		return &api.SubmitResp{Success: false, Message: "上传Job文件失败"}, nil
	}
	fmt.Fprint(jobFile, req.JobFile)
	jobFile.Close()
	// save Input File (.txt)
	inputFilePath := fmt.Sprintf("%s/input.txt", jobDirPath)
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		return &api.SubmitResp{Success: false, Message: "上传Input文件失败"}, nil
	}
	fmt.Fprint(inputFile, req.InputFile)
	inputFile.Close()
	// create Jobs: Mapper & Reducer
	scheduler.Submit(jobName)
	// Result File (.txt)
	resultFilePath := fmt.Sprintf("%s/result.txt", jobDirPath)
	return &api.SubmitResp{Success: true, Message: "提交任务成功", ResultFile: resultFilePath}, nil
}
