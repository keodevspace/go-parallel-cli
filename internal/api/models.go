package api

// Post represents the structure of the data from the API
// In Go, fields must start with a Capital letter to be "Public" (Exported)
type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// APIInterface defines the contract for fetching data
// Similar to an Interface in C# (but implemented implicitly)
type APIInterface interface {
	FetchPost(id int) (*Post, error)
}