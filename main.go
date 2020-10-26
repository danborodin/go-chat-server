package main

import (
	"fmt"

	"github.com/danborodin/go-chat-server/database"
)

func main() {
	// router.SetupRoutes()
	// log.Printf("Server running at address %s", router.Host)
	// log.Fatal(http.ListenAndServe(router.Host, router.Router))

	res, err := database.GetChannels()
	if err != nil {
		//log.Println(err)
	}
	fmt.Println(res)
}
