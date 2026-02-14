package storage

import "github.com/danwflynn/FarmQ/internal/jobs"

// Store interface for using different types of storage
type Store interface {
	Save(job *jobs.Job)
	Get(id string) (*jobs.Job, bool)
	Ping() error
}
