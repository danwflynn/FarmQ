package api

import (
	"log"
	"net/http"
	"time"

	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

// watForDB to wait for the database
func waitForDB(store storage.Store) {
	for range 30 {
		if err := store.Ping(); err == nil {
			log.Println("Database ready")
			return
		}
		log.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}
	log.Fatal("Database not ready after timeout")
}

// Start the api server
func Start(q queue.Queue, store storage.Store) {
	waitForDB(store)

	handler := &Handler{
		Queue: q,
		Store: store,
	}

	http.HandleFunc("/jobs", handler.handleCreateJob)
	http.HandleFunc("/jobs/", handler.handleGetJob)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
