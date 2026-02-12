package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danwflynn/FarmQ/internal/jobs"
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

var (
	jobQueue = queue.NewJobQueue(100)
	store    = storage.NewMemoryStore()
)

func main() {
	http.HandleFunc("/jobs", handleCreateJob)

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCreateJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		JobType string          `json:"job_type"`
		Payload json.RawMessage `json:"payload"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	job := jobs.NewJob(req.JobType, req.Payload)

	store.Save(job)
	jobQueue.Enqueue(job)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"job_id": job.ID,
		"status": string(job.Status),
	})
}
