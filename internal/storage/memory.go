package storage

import (
	"sync"

	"github.com/danwflynn/FarmQ/internal/jobs"
)

type MemoryStore struct {
	mu   sync.RWMutex
	jobs map[string]*jobs.Job
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		jobs: make(map[string]*jobs.Job),
	}
}

func (s *MemoryStore) Save(job *jobs.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = job
}

func (s *MemoryStore) Get(id string) (*jobs.Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[id]
	return job, ok
}
