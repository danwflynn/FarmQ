package jobs

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// JobStatus for status of a job
type JobStatus string

// enum for JobStatus
const (
	StatusPending   JobStatus = "PENDING"
	StatusRunning   JobStatus = "RUNNING"
	StatusCompleted JobStatus = "COMPLETED"
	StatusFailed    JobStatus = "FAILED"
	StatusRetrying  JobStatus = "RETRYING"
)

// Job struct for necessary fields
type Job struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Status    JobStatus       `json:"status"`
	Payload   json.RawMessage `json:"payload"`
	Result    json.RawMessage `json:"result"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Retries   int             `json:"retries"`
}

// NewJob function for creating new job
func NewJob(jobType string, payload json.RawMessage) *Job {
	return &Job{
		ID:        uuid.New().String(),
		Type:      jobType,
		Status:    StatusPending,
		Payload:   payload,
		Result:    nil,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Retries:   0,
	}
}
