package queue

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/danwflynn/FarmQ/internal/jobs"
)

func TestMemoryQueueEnqueueDequeue(t *testing.T) {
	q := NewMemoryQueue(10)

	payload := json.RawMessage(`{"foo":"bar"}`)
	job := jobs.NewJob("test", payload)

	// Enqueue the job
	if err := q.Enqueue(job); err != nil {
		t.Fatalf("enqueue failed: %v", err)
	}

	// Dequeue the job
	dequeuedJob, err := q.Dequeue()
	if err != nil {
		t.Fatalf("dequeue failed: %v", err)
	}

	// Make sure it's the same job
	if dequeuedJob.ID != job.ID {
		t.Errorf("expected job ID %s, got %s", job.ID, dequeuedJob.ID)
	}

	if string(dequeuedJob.Payload) != string(job.Payload) {
		t.Errorf("expected payload %s, got %s", job.Payload, dequeuedJob.Payload)
	}

	if dequeuedJob.Type != job.Type {
		t.Errorf("expected type %s, got %s", job.Type, dequeuedJob.Type)
	}

	if dequeuedJob.Status != job.Status {
		t.Errorf("expected status %s, got %s", job.Status, dequeuedJob.Status)
	}

	// Test channel blocking behavior with timeout
	done := make(chan struct{})
	go func() {
		_, _ = q.Dequeue()
		close(done)
	}()

	select {
	case <-done:
		t.Errorf("expected dequeue to block on empty queue")
	case <-time.After(100 * time.Millisecond):
		// Expected: channel is empty, so Dequeue should block
	}
}
