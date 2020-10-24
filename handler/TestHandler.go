package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danborodin/go-chat-server/database"
	"github.com/danborodin/go-chat-server/models"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	database.AddUser(user)
}
