package worker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/danwflynn/FarmQ/internal/jobs"
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

// Start function for starting worker server
func Start(q queue.Queue, store storage.Store) {
	log.Println("Worker started")

	for {
		// Wait for the next job
		job, err := q.Dequeue()
		if err != nil {
			log.Println("Error getting job from queue:", err)
			continue
		}
		log.Println("Processing job:", job.ID)

		// Mark as running
		job.Status = jobs.StatusRunning
		store.Save(job)

		// Simulate work
		time.Sleep(2 * time.Second)

		// Generate a result
		result := map[string]string{"message": "Work done successfully!"}
		resultJSON, err := json.Marshal(result)
		if err != nil {
			log.Printf("Failed to marshal result for job %s: %v", job.ID, err)
			job.Status = jobs.StatusFailed
			store.Save(job)
			continue
		}
		job.Result = resultJSON

		// Mark as completed
		job.Status = jobs.StatusCompleted
		store.Save(job)

		log.Println("Completed job:", job.ID)
	}
}
