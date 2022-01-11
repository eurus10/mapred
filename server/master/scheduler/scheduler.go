package scheduler

import (
	"fmt"
	"mapred/server/master/mr"
)

var (
	jobID = 1
	jobs  []mr.Job
)

func Submit(jobName string) {
	mapper := mr.Job{
		ID:   jobID,
		Name: jobName,
		Type: mr.MAPPER,
	}
	fmt.Printf("提交Mapper任务, JobID=%d", jobID)
	jobID++
	reducer := mr.Job{
		ID:       jobID,
		Name:     jobName,
		Type:     mr.REDUCER,
		DepJobID: jobID - 1,
	}
	fmt.Printf("提交Reducer任务, JobID=%d", jobID)
	jobID++
	jobs = append(jobs, mapper, reducer)
}

func GetJob(worker string) (bool, mr.Job) {
	found := false
	var job mr.Job
	for i, job := range jobs {
		if job.Worker == "" {
			found = true
			job.Worker = worker
			jobs[i] = job
			break
		}
	}
	return found, job
}

func Done(jobID int) bool {
	target := -1
	for i, job := range jobs {
		if job.ID == jobID {
			target = i
			break
		}
	}
	if target != -1 {
		jobs = append(jobs[:target], jobs[target+1:]...)
		return true
	}
	return false
}
