package worker

import (
	"context"
	"fmt"
	api "mapred/server/master/api/worker"
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
			fmt.Println("开始扫描...")
			for k, w := range workers {
				fmt.Printf("检查节点@%s, lasthb=%d\n", k, w.LastHeartBeat)
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
	// 从scheduler获取任务信息
	job := api.Job{}
	return &api.ApplyForJobResp{Success: true, Message: "申请任务成功", Job: &job}, nil
}

func (w *WorkerService) DoneJob(ctx context.Context, req *api.DoneJobReq) (*api.GenericResp, error) {
	// 从scheduler中删除已完成的任务
	return &api.GenericResp{Success: true, Message: "任务结果提交成功"}, nil
}
