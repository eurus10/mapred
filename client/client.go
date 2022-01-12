package client

import (
	"context"
	"fmt"
	"io/ioutil"
	api "mapred/server/master/api/client"

	"google.golang.org/grpc"
)

type MapRedClient struct {
	Master string
}

func (m *MapRedClient) SubmitJob(jobName string, jobFilePath string, inputFilePath string) (bool, string) {
	jobFile, err := ioutil.ReadFile(jobFilePath)
	if err != nil {
		fmt.Printf("无法打开任务文件[%s]\n", jobFilePath)
		return false, ""
	}
	inputFile, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Printf("无法打开输入文件[%s]\n", inputFilePath)
		return false, ""
	}
	conn, err := grpc.Dial(m.Master, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("集群@%s无法访问\n", m.Master)
		return false, ""
	}
	defer conn.Close()
	client := api.NewClientServiceClient(conn)
	resp, _ := client.Submit(context.Background(), &api.SubmitReq{
		JobName:   jobName,
		JobFile:   jobFile,
		InputFile: inputFile,
	})
	fmt.Println(resp.Message)
	return resp.Success, resp.OutputFile
}

func (m *MapRedClient) PullOutput(outputFilePath string, savePath string) bool {
	return false
}
