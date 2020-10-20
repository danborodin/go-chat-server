package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/danborodin/go-chat-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to database
func Connect() {

	pass := config.Config("DATABASE_PASSWORD")
	dbName := config.Config("DATABASE_NAME")
	connString := fmt.Sprintf("mongodb+srv://admin:%s@chatdb.ngveg.mongodb.net/%s?retryWrites=true&w=majority", pass, dbName)

	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
}
