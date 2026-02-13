package main

import (
	"github.com/danwflynn/FarmQ/internal/api"
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

func main() {
	q := queue.NewMemoryQueue(100)
	store := storage.NewMemoryStore()

	api.Start(q, store)
}
