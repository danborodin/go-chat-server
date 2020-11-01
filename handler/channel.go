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

// GetChannelByID ...
func GetChannelByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
