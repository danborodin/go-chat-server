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

// GetRooms return all users rooms
func GetRooms(w http.ResponseWriter, r *http.Request) {

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

	rooms, err := database.GetRooms()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

// AddNewRoom ...
func AddNewRoom(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var room models.Room
	err := dec.Decode(&room)
	if err != nil {
		log.Println("Error while decoding room from request", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Json decoding failed"))
		return
	}

	err = database.AddRoom(room)
	if err != nil {
		log.Println("Error while adding room to database", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Room added successfully"))

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

// need room id , not pointer to a room

func ConnectToRoom(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error in ConnectToRoom handler: ", err)
		return
	}

	_, roomID, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error while getting room id from user")
	}
	log.Println(roomID)

	//get room from db
	//un slice cu room-urile active..?
	// user := &models.User{
	// 	Room: room,
	// }
}
