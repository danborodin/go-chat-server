package models

import (
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model structure
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Salt     string             `bson:"salt,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`

	Conn *websocket.Conn
}
