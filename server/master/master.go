package main

import (
	"fmt"
	capi "mapred/server/master/api/client"
	wapi "mapred/server/master/api/worker"
	"mapred/server/master/config"
	cservice "mapred/server/master/service/client"
	wservice "mapred/server/master/service/worker"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// stop signal
	stop := make(chan struct{})
	defer close(stop)

	// ResourceManager
	wservice.Init()
	go wservice.Scan(stop)

	// RPC-Server
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		fmt.Println("无法监听指定端口")
		return
	}
	defer conn.Close()
	server := grpc.NewServer()
	wapi.RegisterWorkerServiceServer(server, &wservice.WorkerService{})
	capi.RegisterClientServiceServer(server, &cservice.ClientService{})
	server.Serve(conn)
}
