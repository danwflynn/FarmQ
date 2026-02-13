package queue

import "github.com/danwflynn/FarmQ/internal/jobs"

// Queue interface for interacting with different queues
type Queue interface {
	Enqueue(job *jobs.Job) error
	Dequeue() (*jobs.Job, error)
}
