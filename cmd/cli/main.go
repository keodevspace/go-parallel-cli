package main

import (
	"fmt"
	"sync"

	"github.com/keodevspce/go-parallel-cli/internal/api"
	"github.com/keodevspce/go-parallel-cli/internal/processor"
)

func main() {
	fmt.Println("Starting Parallel CLI in Go...")

	client := api.NewClient()

	// Create a channel for IDs (Jobs)
	jobs := make(chan int, 100)

	// WaitGroup to synchronize the 10 goroutines
	var wg sync.WaitGroup

	// 1. Launch 10 Goroutines (The Worker Pool)
	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go processor.Worker(w, client, jobs, &wg)
	}

	// 2. Feed the jobs channel with 50 tasks
	for i := 1; i <= 50; i++ {
		jobs <- i
	}

	// 3. Close the channel (tells workers: "no more work coming")
	close(jobs)

	// 4. Block until all workers finish
	wg.Wait()

	fmt.Println("All 50 tasks completed successfully.")
}
