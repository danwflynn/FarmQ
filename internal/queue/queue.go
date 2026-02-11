package queue

import "github.com/danwflynn/FarmQ/internal/jobs"

type JobQueue struct {
	Jobs chan *jobs.Job
}

func NewJobQueue(bufferSize int) *JobQueue {
	return &JobQueue{
		Jobs: make(chan *jobs.Job, bufferSize),
	}
}

func (q *JobQueue) Enqueue(job *jobs.Job) {
	q.Jobs <- job
}

func (q *JobQueue) Dequeue() *jobs.Job {
	return <-q.Jobs
}
