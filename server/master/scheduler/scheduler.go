package scheduler

import (
	"fmt"
	"mapred/server/master/mr"
)

var (
	jobID = 1
	jobs  map[int]mr.Job
)

func Submit(jobName string) {
	mapper := mr.Job{
		ID:   jobID,
		Name: jobName,
		Type: mr.MAPPER,
	}
	jobs[jobID] = mapper
	fmt.Printf("提交Mapper任务, JobID=%d", jobID)
	jobID++
	reducer := mr.Job{
		ID:       jobID,
		Name:     jobName,
		Type:     mr.REDUCER,
		DepJobID: jobID - 1,
	}
	jobs[jobID] = reducer
	fmt.Printf("提交Reducer任务, JobID=%d", jobID)
	jobID++
}

func GetJob(worker string) (bool, mr.Job) {
	found := false
	var job mr.Job
	for jobID, job := range jobs {
		if job.Worker == "" {
			if job.DepJobID != 0 {
				_, unfinished := jobs[job.DepJobID]
				if unfinished {
					continue
				}
			}
			found = true
			job.Worker = worker
			jobs[jobID] = job
			break
		}
	}
	return found, job
}

func Done(jobID int) {
	_, unfinished := jobs[jobID]
	if unfinished {
		delete(jobs, jobID)
	}
}
