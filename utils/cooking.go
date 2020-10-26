package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/danborodin/go-chat-server/config"
)

var pepper string = config.GetEnvVar("PEPPER")

// Magic ...
func Magic(password string) (string, string) {

	length := 10
	arr := make([]byte, length)
	_, err := rand.Read(arr)
	if err != nil {
		log.Println(err)
	}

	salt := sha256.Sum256([]byte(arr))
	condiments := fmt.Sprintf("%x", salt) + password + pepper
	passwordHash := sha256.Sum256([]byte(condiments))

	return fmt.Sprintf("%x", passwordHash), fmt.Sprintf("%x", salt)
}

// IsMagicEqual ...
func IsMagicEqual(password string, passwordHash string, salt string) bool {

	condiments := salt + password + pepper
	_passHash := sha256.Sum256([]byte(condiments))
	if fmt.Sprintf("%x", _passHash) != passwordHash {
		return false
	}

	return true
}

//jwt.sign(user, secretKey)
