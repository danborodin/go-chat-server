package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/danborodin/go-chat-server/database"
	"github.com/danborodin/go-chat-server/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

// GetChannels return all users channels
func GetChannels(w http.ResponseWriter, r *http.Request) {

	bearerToken := r.Header.Get("Authorization")

	token, err := ValidateToken(bearerToken)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	channels, err := database.GetChannels()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}

// AddNewChannel ...
func AddNewChannel(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var channel models.Channel
	err := dec.Decode(&channel)
	if err != nil {
		log.Println("Error while decoding channel from request", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Json decoding failed"))
		return
	}

	err = database.AddChannel(channel)
	if err != nil {
		log.Println("Error while adding channel to database", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Channel added successfully"))

	return
}

// ValidateToken validates the token with the secret key and return the object
func ValidateToken(bearerToken string) (*jwt.Token, error) {

	if bearerToken == "" {
		return nil, errors.New("Bearer token mising, unauthorized")
	}

	if len(strings.Split(bearerToken, " ")) <= 1 {
		return nil, errors.New("Bearer token mising, unauthorized")
	}

	tokenString := strings.Split(bearerToken, " ")[1]

	token, err := jwt.ParseWithClaims(tokenString, &models.Token{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	return token, err
}

//-----------------------------------------------------

var connections []*websocket.Conn
var m = 0

var msgData = models.Message{
	ID:     1,
	Sender: "Pidor",
	Text:   "Huynea text",
	Date:   "huynea date",
}

func ConnectToChannel(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	connections = append(connections, conn)
	log.Println(&conn)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client connected")
	err = conn.WriteJSON(msgData)
	if err != nil {
		log.Println(err)
	}

	for {
		_, msg2, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		for n := 0; n < len(connections); n++ {
			if err := connections[n].WriteJSON(msg2); err != nil {
				log.Println(err)
			}
		}
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			m--
			log.Println("Client futut")
		}
		//connections[0].WriteMessage(msgType, msgData)
		log.Println(msgData)
	}
}
