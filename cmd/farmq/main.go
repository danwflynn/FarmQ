package main

import (
	"github.com/danwflynn/FarmQ/internal/api"
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
	"github.com/danwflynn/FarmQ/internal/worker"
)

func main() {
	q := queue.NewMemoryQueue(100)
	store := storage.NewMemoryStore()

	go worker.Start(q, store)

	api.Start(q, store)
}
