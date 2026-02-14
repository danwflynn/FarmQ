package api

import (
	"log"
	"net/http"

	"github.com/danwflynn/FarmQ/internal/queue"
	"github.com/danwflynn/FarmQ/internal/storage"
)

// Start the api server
func Start(q queue.Queue, store storage.Store) {
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
