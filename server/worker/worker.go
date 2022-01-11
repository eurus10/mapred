package main

import (
	"context"
	"fmt"
	"mapred/server/master/api/worker"
	"mapred/server/worker/config"
	"time"

	"google.golang.org/grpc"
)

func main() {
	stop := make(chan struct{})
	defer close(stop)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.MasterIP, config.MasterPort), grpc.WithInsecure())
	if err != nil {
		fmt.Println("无法请求Master服务器")
		return
	}
	defer conn.Close()
	client := worker.NewWorkerServiceClient(conn)
	go func(stop chan struct{}) {
		defer fmt.Println("测试，Interrupt是否可以终止协程")
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Println("发送心跳")
				client.HeartBeat(context.Background(), &worker.HeartBeatReq{Id: fmt.Sprintf("%s:%d", config.IP, config.Port)})
				time.Sleep(5 * time.Second)
			}
		}
	}(stop)

	for {
		fmt.Println("Worker运行中")
		time.Sleep(10 * time.Second)
	}
}
