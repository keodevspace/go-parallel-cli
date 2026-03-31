package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client holds the HTTP client configuration
type Client struct {
	HTTPClient *http.Client
}

// NewClient acts like a Constructor in C#
func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second, // Crucial for production apps
		},
	}
}

// FetchPost retrieves a single post from the public API
func (c *Client) FetchPost(id int) (*Post, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Like 'using' block in C#, ensures resource cleanup

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error: status %d", resp.StatusCode)
	}

	var post Post
	// Decoding the JSON directly into our struct
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}
