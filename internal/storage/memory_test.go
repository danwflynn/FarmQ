package storage

import (
	"encoding/json"
	"testing"

	"github.com/danwflynn/FarmQ/internal/jobs"
)

func TestMemoryStoreSaveAndGet(t *testing.T) {
	store := NewMemoryStore()

	payload := json.RawMessage(`{"foo":"bar"}`)
	job := jobs.NewJob("test", payload)

	// Save the job
	store.Save(job)

	// Retrieve the job
	gotJob, ok := store.Get(job.ID)
	if !ok {
		t.Fatalf("expected to find job with ID %s", job.ID)
	}

	if gotJob.ID != job.ID {
		t.Errorf("expected ID %s, got %s", job.ID, gotJob.ID)
	}

	if gotJob.Type != job.Type {
		t.Errorf("expected type %s, got %s", job.Type, gotJob.Type)
	}

	if string(gotJob.Payload) != string(job.Payload) {
		t.Errorf("expected payload %s, got %s", job.Payload, gotJob.Payload)
	}

	// Try to get a non-existent job
	_, ok = store.Get("non-existent-id")
	if ok {
		t.Errorf("expected not to find a job with ID non-existent-id")
	}
}
