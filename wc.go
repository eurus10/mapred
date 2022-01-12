package main

import (
	"fmt"
	"mapred/client"
	"mapred/server/master/mr"
	"strconv"
	"strings"
)

func Map(name string, context string) []mr.KV {
	kvs := []mr.KV{}
	lines := strings.Split(context, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			if word != "" {
				kvs = append(kvs, mr.KV{Key: word, Value: "1"})
			}
		}
	}
	return kvs
}

func Reduce(key string, values []string) string {
	total := len(values)
	return strconv.Itoa(total)
}

func main() {
	client := client.MapRedClient{Master: "127.0.0.1:8888"}
	fmt.Println(client.SubmitJob("WordCount", "wc.so", "wc.txt"))
}
