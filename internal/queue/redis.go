package queue

import (
	"context"
	"encoding/json"

	"github.com/danwflynn/FarmQ/internal/jobs"
	"github.com/redis/go-redis/v9"
)

// RedisQueue struct
type RedisQueue struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisQueue constructor
func NewRedisQueue(addr string) *RedisQueue {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisQueue{
		client: rdb,
		ctx:    context.Background(),
	}
}

// Enqueue function
func (r *RedisQueue) Enqueue(job *jobs.Job) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return r.client.LPush(r.ctx, "jobs", data).Err()
}

// Dequeue function
func (r *RedisQueue) Dequeue() (*jobs.Job, error) {
	result, err := r.client.BRPop(r.ctx, 0, "jobs").Result()
	if err != nil {
		return nil, err
	}

	var job jobs.Job
	if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
		return nil, err
	}

	return &job, nil
}
