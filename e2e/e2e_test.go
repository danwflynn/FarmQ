//go:build e2e

package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

type JobResponse struct {
	ID string `json:"job_id"`
}

type JobStatus struct {
	ID     string `json:"job_id"`
	Status string `json:"status"`
}

// Tests using the api endpoint
func TestSubmitAndProcessJob(t *testing.T) {
	body := []byte(`{"job_type":"test","payload":{"text":"hello"}}`)

	resp, err := http.Post(
		"http://localhost:8080/jobs",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		t.Fatalf("failed to submit job: %v", err)
	}
	defer resp.Body.Close()

	var jobResp JobResponse
	json.NewDecoder(resp.Body).Decode(&jobResp)

	if jobResp.ID == "" {
		t.Fatal("job ID is empty")
	}

	var status JobStatus
	for range 10 {
		time.Sleep(1 * time.Second)

		r, err := http.Get("http://localhost:8080/jobs/" + jobResp.ID)
		if err != nil {
			t.Fatalf("failed to fetch job: %v", err)
		}

		json.NewDecoder(r.Body).Decode(&status)
		r.Body.Close()

		if status.Status == "COMPLETED" {
			break
		}
	}

	if status.Status != "COMPLETED" {
		t.Fatalf("job never completed, status: %s", status.Status)
	}
}
