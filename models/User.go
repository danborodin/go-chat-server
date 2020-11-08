package models

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model structure
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Salt     string             `bson:"salt,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`

	Room *Room
	Conn *websocket.Conn
	Send chan []byte
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

func (u *User) readPump() {
	defer func() {
		u.Room.Unregister <- u
		u.Conn.Close()
	}()
	u.Conn.SetReadLimit(maxMessageSize)
	u.Conn.SetReadDeadline(time.Now().Add(pongWait))
	u.Conn.SetPongHandler(func(string) error {
		u.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	var msg Message
	for {
		err := u.Conn.ReadJSON(msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error while reading Pump: %v", err)
			}
			break
		}
		u.Room.Brodcast <- []byte(fmt.Sprintf("%v", msg))
	}
}

func (u *User) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		u.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-u.Send:
			if !ok {
				// The room closed the channel
				u.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.Conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(u.Send)
			for i := 0; i < n; i++ {
				w.Write(<-u.Send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			u.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := u.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
