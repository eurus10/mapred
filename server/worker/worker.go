package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	api "mapred/server/master/api/worker"
	"mapred/server/master/mr"
	"mapred/server/worker/config"
	"os"
	"plugin"
	"sort"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type ByKey []mr.KV

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

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
		time.Sleep(2 * time.Second)
		resp, _ := client.ApplyForJob(context.Background(), &api.ApplyForJobReq{Id: workerID})
		if resp.Success {
			jobInfo := resp.Job

			// create Job WorkSpace
			jobDirPath := fmt.Sprintf("%s/%s", config.APPS, jobInfo.Name)
			err := os.Mkdir(jobDirPath, os.ModePerm)
			if err != nil {
				// TODO 向master报告失败 ReportFailure
				fmt.Println("创建工作空间失败")
				continue
			}

			// pull Job File (.so)
			resp, _ := client.PullFile(context.Background(), &api.PullFileReq{
				JobName:  jobInfo.Name,
				FileName: "job.so",
			})
			if !resp.Success {
				// TODO 向master报告失败 ReportFailure
				fmt.Println(resp.Message)
				continue
			}
			jobFile, err := os.Create(fmt.Sprintf("%s/job.so", jobDirPath))
			if err != nil {
				// TODO 向master报告失败 ReportFailure
				fmt.Println("保存任务文件失败")
				continue
			}
			fmt.Fprint(jobFile, resp.Data)
			jobFile.Close()

			// pull Input File (.txt)
			resp, _ = client.PullFile(context.Background(), &api.PullFileReq{JobName: jobInfo.Name, FileName: "input.txt"})
			if !resp.Success {
				// TODO 向master报告失败 ReportFailure
				fmt.Println(resp.Message)
				continue
			}
			var inputType string
			if jobInfo.Type == "Map" {
				inputType = "input"
			} else {
				inputType = "temp"
			}
			inputFile, err := os.Create(fmt.Sprintf("%s/%s.txt", jobDirPath, inputType))
			if err != nil {
				// TODO 向master报告失败 ReportFailure
				fmt.Println("保存输入文件失败")
				continue
			}
			fmt.Fprint(inputFile, resp.Data)

			// load Job
			p, err := plugin.Open(fmt.Sprintf("%s/job.so", jobDirPath))
			if err != nil {
				// TODO 向master报告失败 ReportFailure
				fmt.Println("读取任务文件失败")
				continue
			}
			doJob, err := p.Lookup(jobInfo.Type)
			if err != nil {
				// TODO 向master报告失败 ReportFailure
				fmt.Println("任务文件不包含指定类型任务")
				continue
			}
			if jobInfo.Type == "Map" {
				// stage Map
				doMap := doJob.(func(string, string) []mr.KV)
				input, _ := ioutil.ReadAll(inputFile)
				kvs := doMap(inputType, string(input))
				inputFile.Close()

				// stage Shuffle
				sort.Sort(ByKey(kvs))
				outputFile, err := os.Create(fmt.Sprintf("%s/temp.txt", jobDirPath))
				if err != nil {
					// TODO 向master报告失败 ReportFailure
					fmt.Println("存储计算结果失败")
					continue
				}
				for _, kv := range kvs {
					fmt.Fprintf(outputFile, "%s %s\n", kv.Key, kv.Value)
				}
				output, _ := ioutil.ReadAll(outputFile)
				resp, _ := client.WriteFile(context.Background(), &api.WriteFileReq{
					JobName:  jobInfo.Name,
					FileName: "temp.txt",
					Data:     output,
				})
				outputFile.Close()
				if !resp.Success {
					// TODO 向master报告失败 ReportFailure
					fmt.Println(resp.Message)
					continue
				}
			} else {
				doReduce := doJob.(func(string, []string) string)
				scanner := bufio.NewReader(inputFile)
				outputFile, err := os.Create(fmt.Sprintf("%s/output.txt", jobDirPath))
				if err != nil {
					inputFile.Close()
					// TODO 向master报告失败 ReportFailure
					fmt.Println("存储计算结果失败")
					continue
				}
				var key string
				var values []string
				for {
					line, _, err := scanner.ReadLine()
					if err == io.EOF {
						if len(values) != 0 {
							fmt.Fprintf(outputFile, "%s %s\n", key, doReduce(key, values))
						}
						break
					}
					tokens := strings.Split(string(line), " ")
					if len(tokens) != 2 {
						inputFile.Close()
						outputFile.Close()
						// TODO 向master报告失败 ReportFailure
						fmt.Println("输入文件数据格式错误")
						continue
					}
					if tokens[0] != key {
						// old Partition
						if len(values) != 0 {
							fmt.Fprintf(outputFile, "%s %s\n", key, doReduce(key, values))
							values = values[:0]
						}
						// new Partition
						key = tokens[0]
					}
					values = append(values, tokens[1])
				}
				inputFile.Close()
				output, _ := ioutil.ReadAll(outputFile)
				resp, _ := client.WriteFile(context.Background(), &api.WriteFileReq{
					JobName:  jobInfo.Name,
					FileName: "output.txt",
					Data:     output,
				})
				outputFile.Close()
				if !resp.Success {
					// TODO 向master报告失败 ReportFailure
					fmt.Println(resp.Message)
					continue
				}
			}
		} else {
			fmt.Println(resp.Message)
		}
	}
}
