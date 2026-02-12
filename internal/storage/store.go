package storage

import "github.com/danwflynn/FarmQ/internal/jobs"

type Store interface {
	Save(job *jobs.Job)
	Get(id string) (*jobs.Job, bool)
}
