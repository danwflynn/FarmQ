package main

import (
	"github.com/danwflynn/FarmQ/internal/api"
	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

func main() {
	q := queue.NewRedisQueue("redis:6379")
	store := storage.NewPostgresStore("postgres://farmq:farmq@db:5432/farmq?sslmode=disable")

	api.Start(q, store)
}
