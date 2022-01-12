package scheduler

import (
	"fmt"
	"mapred/server/master/mr"
)

var (
	globalID = 1
	jobs     map[int]mr.Job
)

func Submit(jobName string) {
	if jobs == nil {
		jobs = make(map[int]mr.Job)
	}
	mapper := mr.Job{
		ID:   globalID,
		Name: jobName,
		Type: mr.MAPPER,
	}
	globalID++
	jobs[mapper.ID] = mapper
	fmt.Printf("[Scheduler]提交名为%s的Mapper任务, JobID=%d\n", jobName, mapper.ID)
	reducer := mr.Job{
		ID:       globalID,
		Name:     jobName,
		Type:     mr.REDUCER,
		DepJobID: mapper.ID,
	}
	globalID++
	jobs[reducer.ID] = reducer
	fmt.Printf("[Scheduler]提交名为%s的Reducer任务, JobID=%d\n", jobName, reducer.ID)
}

func GetJob(worker string) (bool, mr.Job) {
	found := false
	var target mr.Job
	for jobID, job := range jobs {
		if job.Worker == "" {
			if job.DepJobID != 0 {
				_, unfinished := jobs[job.DepJobID]
				if unfinished {
					continue
				}
			}
			found = true
			target = job
			target.Worker = worker
			jobs[jobID] = target
			break
		}
	}
	return found, target
}

func Done(jobID int) {
	job, unfinished := jobs[jobID]
	if unfinished {
		delete(jobs, jobID)
		fmt.Printf("[Scheduler]%+v已完成\n", job)
	}
}
