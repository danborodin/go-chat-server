package models

// Channel struct
type Channel struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Users    []User    `json:"users"`
	Messages []Message `json:"messages"`
}
