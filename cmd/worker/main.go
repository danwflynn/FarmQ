package main

import (
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
	"github.com/danwflynn/FarmQ/internal/worker"
)

func main() {
	q := queue.NewJobQueue(100)
	store := storage.NewMemoryStore()

	go worker.Start(q, store)
}
