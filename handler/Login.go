package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danborodin/go-chat-server/config"
	"github.com/danborodin/go-chat-server/database"
	"github.com/danborodin/go-chat-server/models"
	"github.com/danborodin/go-chat-server/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

var key = []byte(config.GetEnvVar("TOKEN_SECRET_KEY"))

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	//user data from request
	var user models.User
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)
	if err != nil {
		log.Println(err)
	}

	//user data from database
	_user, err := database.GetUserByUsername(user.UserName)
	if err != nil {
		log.Println(err)
	}
	if _user.UserName == "" {
		log.Println(_user.UserName)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Login failed, username not found"))
		return
	}

	//compare user data send in request with user data in database
	res := utils.IsMagicEqual(user.Password, _user.Password, _user.Salt)
	if res {
		log.Println(fmt.Sprintf("User with id %s logged in", _user.ID))

		var tokenClaim = models.Token{
			Username: _user.UserName,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(100 * time.Minute).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)

		tokenString, err := token.SignedString(key)

		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(tokenString)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Login failed, username or password are incorrect"))
	return
}
