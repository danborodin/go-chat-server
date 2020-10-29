package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Message struct
type Message struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Sender primitive.ObjectID `json:"sender,omitempty" bson:"sender,omitempty"`
	Text   string             `json:"text,omitempty" bson:"text,omitempty"`
	Date   string             `json:"date,omitempty" bson:"date,omitempty"`
}
