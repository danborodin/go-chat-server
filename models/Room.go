package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Room struct
type Room struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Owner    string             `json:"owner,omitempty" bson:"owner,omitempty"`
	Messages []Message          `json:"messages,omitempty" bson:"messages,omitempty"`

	Users      map[*User]bool
	Brodcast   chan []byte
	Register   chan *User
	Unregister chan *User
}

// NewRoom is a constructor
func NewRoom(name string, messages []Message, owner string) *Room {
	room := Room{
		ID:         primitive.NewObjectID(),
		Name:       name,
		Messages:   messages,
		Owner:      owner,
		Users:      make(map[*User]bool),
		Brodcast:   make(chan []byte),
		Register:   make(chan *User),
		Unregister: make(chan *User),
	}

	return &room
}

// func (room *Room) Run() {
// 	for {
// 		select {
// 		case user := <-room.Register:
// 			room.Users[user] = true
// 		case user := <-room.Unregister:
// 			if _, ok := room.Users[user]; ok {
// 				delete(room.Users, user)
// 				close(user.Send)
// 			}
// 		case message := <-room.Brodcast:
// 			for user := range room.Users {
// 				select {
// 				case user.Send <- message:
// 				default:
// 					close(user.Send)
// 					delete(room.Users, user)
// 				}
// 			}
// 		}
// 	}
// }
