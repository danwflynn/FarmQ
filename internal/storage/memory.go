package storage

import (
	"sync"

	"github.com/danwflynn/FarmQ/internal/jobs"
)

// MemoryStore implementation for local storage
type MemoryStore struct {
	mu   sync.RWMutex
	jobs map[string]*jobs.Job
}

// NewMemoryStore function for creation
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		jobs: make(map[string]*jobs.Job),
	}
}

// Save function for saving a job to the store
func (s *MemoryStore) Save(job *jobs.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = job
}

// Get function for getting job
func (s *MemoryStore) Get(id string) (*jobs.Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	job, ok := s.jobs[id]
	return job, ok
}

// Ping always returns no error
func (s *MemoryStore) Ping() error {
	return nil
}
