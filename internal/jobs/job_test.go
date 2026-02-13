package jobs

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewJobDefaults(t *testing.T) {
	payload := json.RawMessage(`{"hello":"world"}`)

	job := NewJob("email", payload)

	// ID should not be empty
	if job.ID == "" {
		t.Fatal("expected ID to be generated")
	}

	// ID should be a valid UUID
	if _, err := uuid.Parse(job.ID); err != nil {
		t.Fatalf("expected valid UUID, got %v", err)
	}

	// Type should match input
	if job.Type != "email" {
		t.Errorf("expected Type 'email', got %s", job.Type)
	}

	// Status should default to pending
	if job.Status != StatusPending {
		t.Errorf("expected Status %s, got %s", StatusPending, job.Status)
	}

	// Payload should match
	if string(job.Payload) != string(payload) {
		t.Errorf("expected payload %s, got %s", payload, job.Payload)
	}

	// Result should be nil
	if job.Result != nil {
		t.Errorf("expected Result to be nil, got %v", job.Result)
	}

	// Retries should default to 0
	if job.Retries != 0 {
		t.Errorf("expected Retries to be 0, got %d", job.Retries)
	}

	// Timestamps should be set
	if job.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}

	if job.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set")
	}

	// UpdatedAt should be close to CreatedAt
	if job.UpdatedAt.Sub(job.CreatedAt) > time.Second {
		t.Error("expected UpdatedAt and CreatedAt to be close")
	}
}
