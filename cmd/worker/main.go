package main

import (
	"log"
	"time"

	"github.com/danwflynn/FarmQ/internal/jobs"
	"github.com/danwflynn/FarmQ/internal/queue"
)

var jobQueue = queue.NewJobQueue(100)

func main() {
	log.Println("Worker started")

	for {
		job := jobQueue.Dequeue()

		log.Println("Processing job:", job.ID)

		job.Status = jobs.StatusRunning
		time.Sleep(2 * time.Second)

		job.Status = jobs.StatusCompleted
		log.Println("Completed job:", job.ID)
	}
}
