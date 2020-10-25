package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/danborodin/go-chat-server/config"
)

// Magic ...
func Magic(password string) (string, string) {

	c := 10
	salt := make([]byte, c)
	_, err := rand.Read(salt)
	if err != nil {
		log.Println(err)
	}

	saltHash := sha256.Sum256([]byte(salt))
	pepper := config.GetEnvVar("PEPPER")
	condiments := fmt.Sprintf("%x", saltHash) + password + pepper
	passwordHash := sha256.Sum256([]byte(condiments))

	return fmt.Sprintf("%x", passwordHash), fmt.Sprintf("%x", saltHash)
}
