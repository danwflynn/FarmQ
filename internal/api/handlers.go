package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/danwflynn/FarmQ/internal/jobs"
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

// Handler for queues and storage
type Handler struct {
	Queue queue.Queue
	Store storage.Store
}

// JobResponse for json responses to requests
type JobResponse struct {
	ID        string          `json:"job_id"`
	Type      string          `json:"job_type"`
	Status    string          `json:"status"`
	Payload   json.RawMessage `json:"payload"`
	Result    json.RawMessage `json:"result,omitempty"`
	Retries   int             `json:"retries"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *Handler) handleCreateJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req struct {
		JobType string          `json:"job_type"`
		Payload json.RawMessage `json:"payload"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON request")
		return
	}

	if req.JobType == "" {
		writeJSONError(w, http.StatusBadRequest, "job_type is required")
		return
	}

	job := jobs.NewJob(req.JobType, req.Payload)
	h.Store.Save(job)

	if err := h.Queue.Enqueue(job); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to enqueue job")
		return
	}

	log.Printf("Created job: %s type: %s", job.ID, job.Type)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JobResponse{
		ID:        job.ID,
		Type:      job.Type,
		Status:    string(job.Status),
		Payload:   job.Payload,
		Result:    job.Result,
		Retries:   job.Retries,
		CreatedAt: job.CreatedAt.Format(time.RFC3339),
		UpdatedAt: job.UpdatedAt.Format(time.RFC3339),
	})
}

func (h *Handler) handleGetJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.URL.Path[len("/jobs/"):]
	if id == "" {
		writeJSONError(w, http.StatusBadRequest, "missing job id")
		return
	}

	job, ok := h.Store.Get(id)
	if !ok {
		writeJSONError(w, http.StatusNotFound, "job not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JobResponse{
		ID:        job.ID,
		Type:      job.Type,
		Status:    string(job.Status),
		Payload:   job.Payload,
		Result:    job.Result,
		Retries:   job.Retries,
		CreatedAt: job.CreatedAt.Format(time.RFC3339),
		UpdatedAt: job.UpdatedAt.Format(time.RFC3339),
	})
}
