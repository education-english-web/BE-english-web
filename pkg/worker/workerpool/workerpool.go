package workerpool

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"sync"

	"github.com/education-english-web/BE-english-web/pkg/worker"
)

// JobManager will take care jobs sent to workerPool
type JobManager interface {
	Execute(job worker.Job)
}

// workerPool runs the jobs using its workers
type workerPool struct {
	jobsQueue  chan worker.Job
	jobManager JobManager
}

var (
	once     sync.Once
	instance worker.Workers
)

// New creates workerPool to run jobs
func New(workerNumber int, jobManager JobManager) worker.Workers {
	once.Do(func() {
		pool := workerPool{
			jobsQueue:  make(chan worker.Job),
			jobManager: jobManager,
		}

		for i := 0; i < workerNumber; i++ {
			go func() {
				for job := range pool.jobsQueue {
					pool.jobManager.Execute(job)
				}
			}()
		}

		instance = &pool
	})

	return instance
}

func GetWorkerPool() worker.Workers {
	return instance
}

// Run submits task to the worker pool.
func (p *workerPool) Run(w worker.Job) {
	p.jobsQueue <- w
}
