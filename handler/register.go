package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/danborodin/go-chat-server/database"
	"github.com/danborodin/go-chat-server/models"
	"github.com/danborodin/go-chat-server/utils"
	"github.com/go-playground/validator"
)

var validate *validator.Validate = validator.New()

// Register is the resigter handler
func Register(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var user models.User
	err := dec.Decode(&user)
	if err != nil {
		log.Println("Error while decoding json in register handler", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Json decoding failed"))
		return
	}

	res, err := database.CheckUsernameExist(user.UserName)
	if err != nil {
		log.Println(err)
	}
	if res {
		log.Println("Register failed, username taken")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Register failed, username taken"))
		return
	}

	user.Password, user.Salt = utils.Magic(user.Password)

	err = database.AddUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account created successfully"))

	return
}

// func validateUser(user models.User) bool {
// 	err := validate.Struct(user)
// 	if err != nil {
// 		if _, ok := err.(*validator.InvalidValidationError); ok {
// 			log.Println(err)
// 			//return false
// 		}
// 		for _, err := range err.(validator.ValidationErrors) {
// 			log.Println(err.Field())
// 		}
// 	}

// 	return true
// }

// func validateNickname(nickname string) bool {
// 	errs := validate.Var(nickname, "required")
// 	if errs != nil {
// 		log.Println(errs)
// 		return false
// 	}

// 	return true
// }
