package models

// Message struct
type Message struct {
	ID     int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Sender string `json:"sender,omitempty" bson:"sender,omitempty"`
	Text   string `json:"text,omitempty" bson:"text,omitempty"`
	Date   string `json:"date,omitempty" bson:"date,omitempty"`
}
