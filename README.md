# mapred
MIT6.824 Lab1：简单的分布式MapReduce框架，Golang实现

## 一、项目结构

* client：客户端API，主要用于提交任务和拉取输出结果（目前还没写配套的分布式文件系统）

* server：服务端应用

  * master：主节点应用，包含子节点管理、任务调度、任务资源管理等功能
  
  * worker：工作节点应用，启动后会定时发送心跳加入集群，同时会在闲置时不断向主节点申请任务

## 二、配置说明

* master节点的配置在`mapred/server/master/config/config.go`中，需要配置本机IP、占用端口以及任务资源持久化路径

* worker节点的配置在`mapred/server/worker/config/config.go`中，需要配置本机IP、占用端口、Master节点IP、Master节点通信端口以及任务资源持久化路径

## 三、其他说明

* 提交的任务必须为Golang编译生成的动态库，即.so文件，项目中包含一个WordCount程序的示例

* 任务代码中必须包含Map和Reduce两个函数，声明如下：

```golang
func Map(name string, context string) []mr.KV
func Reduce(key string, values []string) string
```