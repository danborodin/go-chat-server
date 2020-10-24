package models

// Message struct
type Message struct {
	ID     int    `json:"id"`
	Sender User   `json:"sender"`
	Text   string `json:"text"`
	Date   string `json:"date"`
}
