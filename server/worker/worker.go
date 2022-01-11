package main

import (
	"context"
	"fmt"
	api "mapred/server/master/api/worker"
	"mapred/server/worker/config"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// stop signal
	stop := make(chan struct{})
	defer close(stop)

	// worker id
	workerID := fmt.Sprintf("%s:%d", config.IP, config.Port)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.MasterIP, config.MasterPort), grpc.WithInsecure())
	if err != nil {
		fmt.Println("无法请求Master服务器")
		return
	}
	defer conn.Close()
	client := api.NewWorkerServiceClient(conn)

	// send heartbeat in backend
	go func(stop chan struct{}) {
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Println("发送心跳")
				client.HeartBeat(context.Background(), &api.HeartBeatReq{Id: workerID})
				time.Sleep(5 * time.Second)
			}
		}
	}(stop)

	for {
		resp, _ := client.ApplyForJob(context.Background(), &api.ApplyForJobReq{Id: workerID})
		if resp.Success {
			// TODO 从master节点拉取job文件和输入文件
			// job文件为 job.so
			// mapper的输入文件为 input.txt
			// reducer的输入文件为 temp.txt
		} else {
			fmt.Println(resp.Message)
			time.Sleep(10 * time.Second)
		}
	}
}
