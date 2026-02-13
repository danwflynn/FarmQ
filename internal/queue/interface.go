package queue

import "github.com/danwflynn/FarmQ/internal/jobs"

type Queue interface {
	Enqueue(job *jobs.Job) error
	Dequeue() (*jobs.Job, error)
}
