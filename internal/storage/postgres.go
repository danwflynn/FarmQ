package storage

import (
	"context"
	"encoding/json"

	"github.com/danwflynn/FarmQ/internal/jobs"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresStore struct for database
type PostgresStore struct {
	db  *pgxpool.Pool
	ctx context.Context
}

// NewPostgresStore constructor
func NewPostgresStore(connString string) *PostgresStore {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(err)
	}

	return &PostgresStore{
		db:  pool,
		ctx: ctx,
	}
}

// Save function to save job to database
func (p *PostgresStore) Save(job *jobs.Job) {
	data, _ := json.Marshal(job)

	_, err := p.db.Exec(p.ctx,
		`INSERT INTO jobs (id, data)
		 VALUES ($1, $2)
		 ON CONFLICT (id)
		 DO UPDATE SET data = $2`,
		job.ID, data,
	)

	if err != nil {
		panic(err)
	}
}

// Get function to get job from job id
func (p *PostgresStore) Get(id string) (*jobs.Job, bool) {
	row := p.db.QueryRow(p.ctx,
		`SELECT data FROM jobs WHERE id = $1`,
		id,
	)

	var data []byte
	err := row.Scan(&data)
	if err != nil {
		return nil, false
	}

	var job jobs.Job
	json.Unmarshal(data, &job)

	return &job, true
}

// Ping the database
func (p *PostgresStore) Ping() error {
	var tmp int
	return p.db.QueryRow(p.ctx, "SELECT 1").Scan(&tmp)
}
