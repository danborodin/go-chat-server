package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var username = "test123"
var password = "test321"

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var cred credentials
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&cred)
	if err != nil {
		fmt.Println(err)
	}

	if (cred.Username != username) || (cred.Password != password) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Error")
		return
	}
	//w.Write([]byte("Succes\n"))
	fmt.Println("Succes")

	//fmt.Println(cred)
}
