package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexvlasov182/http/pingrobot/backend/handlers"
	"github.com/alexvlasov182/http/pingrobot/backend/workerpool"
)

const (
	INTERVAL        = time.Second * 10
	REQUEST_TIMEOUT = time.Second * 2
	WORKERS_COUNT   = 3
)

var urls = []string{
	"https://google.com",
	"https://facebook.com",
	"https://some-fake.com",
}

func main() {
	results := make(chan workerpool.Result)
	workerPool := workerpool.New(WORKERS_COUNT, REQUEST_TIMEOUT, results)

	workerPool.Init()
	go generateJobs(workerPool)
	go processResults(results)

	http.HandleFunc("/start", handlers.StartHandler(workerPool))
	http.HandleFunc("/results", handlers.ResultsHandler(results))

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("server failed: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	workerPool.Stop()
}

func processResults(results chan workerpool.Result) {
	for result := range results {
		fmt.Println(result.Info())
	}
}

func generateJobs(wp *workerpool.Pool) {
	for {
		for _, url := range urls {
			wp.Push(workerpool.Job{URL: url})
		}
		time.Sleep(INTERVAL)
	}
}
