package queue

import "github.com/danwflynn/FarmQ/internal/jobs"

type MemoryQueue struct {
	Jobs chan *jobs.Job
}

func NewMemoryQueue(bufferSize int) *MemoryQueue {
	return &MemoryQueue{
		Jobs: make(chan *jobs.Job, bufferSize),
	}
}

func (q *MemoryQueue) Enqueue(job *jobs.Job) error {
	q.Jobs <- job
	return nil
}

func (q *MemoryQueue) Dequeue() (*jobs.Job, error) {
	job := <-q.Jobs
	return job, nil
}
