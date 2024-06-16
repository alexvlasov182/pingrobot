package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type URLRequest struct {
	URLs []string `json:"urls"`
}

type PingResult struct {
	URL          string `json:"url"`
	Status       string `json:"status"`
	ResponseTime string `json:"response_time"`
}

var results []PingResult

func startPing(w http.ResponseWriter, r *http.Request) {
	var req URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results = []PingResult{} // Очищаем предыдущие результаты

	for _, url := range req.URLs {
		start := time.Now()
		resp, err := http.Get(url)
		elapsed := time.Since(start)

		status := "OK"
		if err != nil || resp.StatusCode != http.StatusOK {
			status = "Failed"
		}

		results = append(results, PingResult{
			URL:          url,
			Status:       status,
			ResponseTime: elapsed.String(),
		})
	}

	w.WriteHeader(http.StatusAccepted)
}

func getResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/start", startPing).Methods("POST")
	router.HandleFunc("/api/results", getResults).Methods("GET")

	// Setup CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}
