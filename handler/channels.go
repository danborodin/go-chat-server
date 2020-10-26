package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/danborodin/go-chat-server/models"
	"github.com/dgrijalva/jwt-go"
)

// GetChannels return all users channels
func GetChannels(w http.ResponseWriter, r *http.Request) {

	bearerToken := r.Header.Get("Authorization")

	token, err := ValidateToken(bearerToken)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := token.Claims.(*models.Token)

	json.NewEncoder(w).Encode(fmt.Sprintf("%s channels", user.Username))
}

// ValidateToken validates the token with the secret key and return the object
func ValidateToken(bearerToken string) (*jwt.Token, error) {

	tokenString := strings.Split(bearerToken, " ")[0]

	token, err := jwt.ParseWithClaims(tokenString, &models.Token{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	return token, err
}
