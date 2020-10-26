package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Channel struct
type Channel struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Users    []User             `json:"users,omitempty" bson:"users,omitempty"`
	Messages []Message          `json:"messages,omitempty" bson:"messages,omitempty"`
}

//un user are maim ulte chaneluri , nu un chanel mai muti useri... de regandit asta
