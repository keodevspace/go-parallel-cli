package processor

import (
	"fmt"
	"sync"
	"github.com/keodevspce/go-parallel-cli/internal/api"
)

// Worker consumes IDs from the jobs channel and processes them
// It uses a pointer to sync.WaitGroup to notify the main thread
func Worker(id int, client *api.Client, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Equivalent to a finally block that calls Done()

	for jobID := range jobs {
		// Fetching data from API
		post, err := client.FetchPost(jobID)
		if err != nil {
			fmt.Printf("[Worker %d] Failed to fetch post %d: %v\n", id, jobID, err)
			continue
		}

		// Success log
		fmt.Printf("[Worker %d] Processed Post ID: %d | Title: %.30s...\n", id, post.ID, post.Title)
	}
}