package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alexvlasov182/http/pingrobot/backend/workerpool"
)

func ResultsHandler(results chan workerpool.Result) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res []workerpool.Result
		for result := range results {
			res = append(res, result)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
