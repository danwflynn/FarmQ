package queue

import "github.com/danwflynn/FarmQ/internal/jobs"

// MemoryQueue implementation for queue
type MemoryQueue struct {
	Jobs chan *jobs.Job
}

// NewMemoryQueue function for creating new memory queue
func NewMemoryQueue(bufferSize int) *MemoryQueue {
	return &MemoryQueue{
		Jobs: make(chan *jobs.Job, bufferSize),
	}
}

// Enqueue function for enqueue
func (q *MemoryQueue) Enqueue(job *jobs.Job) error {
	q.Jobs <- job
	return nil
}

// Dequeue function for dequeue
func (q *MemoryQueue) Dequeue() (*jobs.Job, error) {
	job := <-q.Jobs
	return job, nil
}
