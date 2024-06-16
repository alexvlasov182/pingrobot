package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alexvlasov182/http/pingrobot/backend/backend/workerpool"
)

func StartHandler(wp *workerpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job workerpool.Job
		if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		wp.Push(job)
		w.WriteHeader(http.StatusOK)
	}
}
