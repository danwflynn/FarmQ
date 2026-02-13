package worker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/danwflynn/FarmQ/internal/jobs"
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

func Start(q *queue.JobQueue, store storage.Store) {
	log.Println("Worker started")

	for {
		// Wait for the next job
		job := q.Dequeue()
		log.Println("Processing job:", job.ID)

		// Mark as running
		job.Status = jobs.StatusRunning
		store.Save(job)

		// Simulate work
		time.Sleep(2 * time.Second)

		// Generate a result
		result := map[string]string{"message": "Work done successfully!"}
		resultJSON, _ := json.Marshal(result)
		job.Result = resultJSON

		// Mark as completed
		job.Status = jobs.StatusCompleted
		store.Save(job)

		log.Println("Completed job:", job.ID)
	}
}
