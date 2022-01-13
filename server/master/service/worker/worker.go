package worker

import (
	"context"
	"fmt"
	"io/ioutil"
	api "mapred/server/master/api/worker"
	"mapred/server/master/config"
	"mapred/server/master/scheduler"
	"os"
	"time"
)

type Worker struct {
	ID            string
	LastHeartBeat int64
}

var workers map[string]Worker

func Init() {
	if workers == nil {
		workers = make(map[string]Worker)
	}
	for k := range workers {
		delete(workers, k)
	}
	fmt.Println("[ResourceManager]已初始化Worker队列")
}

func Scan(stop chan struct{}) {
	for {
		now := time.Now().Unix()
		select {
		case <-stop:
			fmt.Println("[ResourceManager]Worker队列监听已关闭")
		default:
			for k, w := range workers {
				if now-w.LastHeartBeat > 10 {
					delete(workers, k)
					fmt.Printf("[ResourceManager]节点Worker@%s心跳已过期\n", k)
				}
			}
			time.Sleep(3 * time.Second)
		}
	}
}

type WorkerService struct{}

func (w *WorkerService) HeartBeat(ctx context.Context, req *api.HeartBeatReq) (*api.GenericResp, error) {
	worker, ok := workers[req.Id]
	if ok {
		worker.LastHeartBeat = time.Now().Unix()
		workers[req.Id] = worker
	} else {
		workers[req.Id] = Worker{
			ID:            req.Id,
			LastHeartBeat: time.Now().Unix(),
		}
		fmt.Printf("[ResourceManager]节点Worker@%s心跳已加入集群\n", req.Id)
	}
	return &api.GenericResp{Success: true, Message: "发送心跳成功"}, nil
}

func (w *WorkerService) ApplyForJob(ctx context.Context, req *api.ApplyForJobReq) (*api.ApplyForJobResp, error) {
	ok, job := scheduler.GetJob(req.Id)
	if !ok {
		return &api.ApplyForJobResp{Success: false, Message: "暂无可分配任务"}, nil
	}
	fmt.Printf("Worker@%s申请任务%+v\n", req.Id, job)
	return &api.ApplyForJobResp{Success: true, Message: "申请任务成功", Job: &api.Job{
		Id:   int32(job.ID),
		Name: job.Name,
		Type: job.Type,
	}}, nil
}

func (w *WorkerService) DoneJob(ctx context.Context, req *api.DoneJobReq) (*api.GenericResp, error) {
	scheduler.Done(int(req.JobId))
	return &api.GenericResp{Success: true, Message: "任务结果提交成功"}, nil
}

func (w *WorkerService) ReportFailure(ctx context.Context, req *api.ReportFailureReq) (*api.GenericResp, error) {
	scheduler.GiveUpJob(int(req.JobId))
	fmt.Printf("Worker@%s放弃任务[ID=%d], 原因: %s\n", req.WorkerId, req.JobId, req.Message)
	return &api.GenericResp{Success: true, Message: "报告成功"}, nil
}

func (w *WorkerService) PullFile(ctx context.Context, req *api.PullFileReq) (*api.PullFileResp, error) {
	filePath := fmt.Sprintf("%s/%s/%s", config.APPS, req.JobName, req.FileName)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &api.PullFileResp{Success: false, Message: "文件不存在"}, nil
	}
	return &api.PullFileResp{Success: true, Message: "拉取成功", Data: data}, nil
}

func (w *WorkerService) WriteFile(ctx context.Context, req *api.WriteFileReq) (*api.GenericResp, error) {
	filePath := fmt.Sprintf("%s/%s/%s", config.APPS, req.JobName, req.FileName)
	file, err := os.Create(filePath)
	fmt.Fprint(file, string(req.Data))
	file.Close()
	if err != nil {
		return &api.GenericResp{Success: false, Message: "写入文件失败"}, nil
	}
	return &api.GenericResp{Success: true, Message: "写入文件成功"}, nil
}
